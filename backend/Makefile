MIGRATIONS_DIR := ./database/migrations

make_migration:
	migrate create -ext=sql -dir=$(MIGRATIONS_DIR) -seq $(name)
