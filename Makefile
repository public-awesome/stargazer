.PHONY: build test run


generate-models:
	sqlboiler psql  --struct-tag-casing camel  --wipe
	
build:
	go build -o build/stakewatcher github.com/public-awesome/stakewatcher/cmd/stakewatcher

run:
	go run github.com/public-awesome/stakewatcher/cmd/stakewatcher
