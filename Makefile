.PHONY: build test run pkger build-alpine


generate-models:
	sh contrib/db/generate.sh

install-pkger:
	go install github.com/markbates/pkger/cmd/pkger

pkger:
	pkger -o cmd/stargazer

ci-sign:
	drone sign public-awesome/stargazer --save
build: pkger
	go build -o build/stargazer github.com/public-awesome/stargazer/cmd/stargazer

build-linux: pkger
	GOOS=linux GOARCH=amd64 go build -o build/stargazer github.com/public-awesome/stargazer/cmd/stargazer

build-alpine: pkger
	GOOS=linux GOARCH=amd64 go build -tags muslc -o build/stargazer github.com/public-awesome/stargazer/cmd/stargazer
 
build-docker:
	docker build -t publicawesome/stargazer .

run-auto-migrate: pkger
	go run github.com/public-awesome/stargazer/cmd/stargazer --auto-migrate
run:
	go run github.com/public-awesome/stargazer/cmd/stargazer

start:
	./build/stargazer
	
fake-post:
	stakecli tx curating post  1 $(POST_ID) "post body"  --from validator --keyring-backend test --trust-node --chain-id $(shell stakecli status | jq '.node_info.network') -b block -y

	
fake-upvote:
	stakecli tx curating upvote 1 $(POST_ID) 1  --from validator --keyring-backend test --trust-node --chain-id $(shell stakecli status | jq '.node_info.network') -b block -y


install-tools: install-sqlboiler install-sqlmigrate


install-sqlboiler:
	VERSION=v4.4.0 ./contrib/dev/install-sqlboiler.sh

install-sqlmigrate:
	go install github.com/rubenv/sql-migrate/sql-migrate

sql-migrate:
	sql-migrate up
sql-migrate-down:
	sql-migrate down 
sql-reset:
	sql-migrate down --limit=0
	sql-migrate up

docker-up:
	docker-compose up -d
docker-down:
	docker-compose down
docker-cleanup:
	docker-compose down --volumes
