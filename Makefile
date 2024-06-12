SHELL := /bin/bash

migrate-status:
	sql-migrate status -config=db/dbconfig.yml

migrate-up:
	sql-migrate up -config=db/dbconfig.yml

migrate-down:
	sql-migrate down -config=db/dbconfig.yml

migrate-new:
	sql-migrate new -config=db/dbconfig.yml