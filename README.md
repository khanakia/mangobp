## Mango Framework Boilerplate Quickstart

### How to start
* Install Deps - `yarn install`
* Install go deps - `go mod vendor`

### Start in development watch mode
```
yarn start
```

### Start in production mode
```
go run . server
```

### GQL Playground
http://localhost:7001/gql?key=1234

### Migrate DB
```
yarn cmd:migrate
```

### Graphql GQLGEN build schema
```
yarn gql
```

### Wire build packages
```
yarn wire
```
