hello:
	@echo "Hello, Mango"

watch:
	@echo `air`

migratedb:
	@echo `go run . plugin migratedb`

wire:
	@echo `cd wireapp && wire`	

gql:
	@echo `cd apigql && go run github.com/99designs/gqlgen generate`