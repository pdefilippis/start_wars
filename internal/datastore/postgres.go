package datastore

import (
	"fmt"
	"hop/start_wars/cmd/config"
	"net/url"

	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type PostgresDataStore struct {
	Db     *gorm.DB
}

func NewPostgresDataStore(cfg config.AppConfig) (IDataStore, error) {
	db, err := initPostgresDB(cfg)
	if err != nil {
		return nil, err
	}

	return PostgresDataStore{Db: db}, nil
}

func initPostgresDB(cfg config.AppConfig) (*gorm.DB, error) {
	dns := url.URL{
		User:     url.UserPassword(cfg.AppDBUserName, cfg.AppDBPassword),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%s", cfg.AppDBHost, cfg.AppDBPort),
		Path:     cfg.AppDBName,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	return gorm.Open(postgres.Open(dns.String()), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "public.",
			SingularTable: false,
		},
	})
}

func (db PostgresDataStore) Migrate() error {
	migrations := &migrate.FileMigrationSource{Dir: "../../db/migrations"}
	db.Db.Exec("set search_path='public'")
	dataBase, err := db.Db.DB()
	if err != nil {
		return fmt.Errorf("search path error for migration: %v", err)
	}

	_, err = migrate.Exec(dataBase, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	return nil
}

func (db PostgresDataStore) GetDB() *gorm.DB {
	return db.Db
}