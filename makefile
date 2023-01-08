BINARY_NAME=main.out
MIGRATE=docker exec -it api_container migrate -path=migration   -database "postgres://spuser:SPuser96@postgres_db:5432/project?sslmode=disable" -verbose
migrate-up:
		$(MIGRATE) up
migrate-down:
		$(MIGRATE) down
force:
		@read -p  "Which version do you want to force?" VERSION; \
		$(MIGRATE) force $$VERSION

goto:
		@read -p  "Which version do you want to migrate?" VERSION; \
		$(MIGRATE) goto $$VERSION

drop:
		$(MIGRATE) drop

create:
		@read -p  "What is the name of migration?" NAME; \
		migrate create -ext sql -seq -dir migration  $$NAME

.PHONY: migrate-up migrate-down force goto drop create

build:
	go build -o ${BINARY_NAME} main.go

test:
	go test -v main.go

run:
	go build -o ${BINARY_NAME} main.go
	 ./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go
