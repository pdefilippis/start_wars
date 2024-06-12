package main

import (
	"fmt"
	"hop/start_wars/cmd/config"
	"hop/start_wars/cmd/handlers"
	"hop/start_wars/internal/datastore"
	"hop/start_wars/internal/repositories"
	"hop/start_wars/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type App struct {
	Router *gin.Engine
	Log *zap.Logger
	DS     *datastore.IDataStore
}

func main(){
	app := App{}
	app.run()
}

func (app *App) run(){
	config, err := config.LoadConfig("../../")
	if err != nil{
		panic("cant load config")
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Errorf("error initializing logger: %v", err))
	}
	app.Log = logger
	app.Router = gin.Default()

	app.initDB(config)

	starshipRepo := repositories.NewVehicleRepository(config.RootStartWarsApi, "starships")
	vehicleRepo := repositories.NewVehicleRepository(config.RootStartWarsApi, "vehicles")
	vehicleExtendRepo := repositories.NewVehicleExtendRepository(*app.DS)

	vehicleSvc := services.NewVehicleService(app.Log, starshipRepo, vehicleRepo, vehicleExtendRepo)

	vehicleManager := handlers.NewVehicleManager(app.Log, vehicleSvc)
	handlers.CreateVehicleEndpoints(app.Router, vehicleManager)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", config.APIPort),
		Handler: app.Router,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func (app *App) initDB(config *config.AppConfig){
	ds, err := datastore.NewPostgresDataStore(*config)
	if err != nil {
		app.Log.Error("init data base", zap.Error(err))
		panic(fmt.Errorf("error init data base: %v", err))
	}

	err = ds.Migrate()
	if err != nil {
		app.Log.Error("migrate data base", zap.Error(err))
		panic(fmt.Errorf("error migrate data base: %v", err))
	}

	app.DS = &ds
}