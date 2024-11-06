
.PHONY: up
up:
	docker compose up --build --watch

test:
	docker compose -f ./docker-compose.yml exec -T go sh -c "cd /app && go test -v -cover ./tests/integration"


db-bash:
	docker compose exec -it db bash



.PHONY: model_gen
model_gen:
	sqlboiler mysql --wipe --pkgname datamodel --output api/internal/infra/datamodel --templates ${GOPATH}/pkg/mod/github.com/volatiletech/sqlboiler/v4@v4.14.2/templates --templates api/internal/infra/templates
	

.PHONY: docker_up_db
docker_up_db:
	docker compose -f docker-compose.yml up --build -d db


# マイグレーション名のデフォルト値
MIGRATION_NAME?=default_migration

#make create-migration MIGRATION_NAME=create_users_table
create-migration:
	migrate create -ext sql -dir api/migrations -seq $(MIGRATION_NAME)
