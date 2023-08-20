APP=backEnd_Coffeshop
BUILD="./build/$(APP)"
DB_DRIVER=postgres
DB_SOURCE="postgresql://ekha:@0okmnji9!@127.0.0.1:5433/coffeshop_database?sslmode=disable&search_path=golang"
MIGRATIONS_DIR=./migrations

install:
	go get -u ./... && go mod tidy

build:
	CGO_ENABLED=0 GOOS=linux go build -o ${BUILD}

test:
	go test -cover -v ./...

migrate-init:
	migrate create -dir ${MIGRATIONS_DIR} -ext sql $(name)

migrate-up:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} -verbose up

migrate-down:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} -verbose down

migrate-fix:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} force 0