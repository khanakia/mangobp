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


## Nats Test
nats request foo 'test'


## Strucutre
`services` - this directory will contain the modules which can only be used via nats only. al the services will have the possibilty to move into separate microservices out of monolithing code if needed at some point.
* No service can import other service directly.
* One service should communicate with other services using NATS only
* Apigql should call services using NATS only.
* Sometimes two services are dependent each other and they may need to call each other then find a way to make them loosely coupled as much as possibile

`mango` - this directory will contain the source code for whole mango framework. later on we will convert this to go module and can be downloaded separately

`pkg` - any modules which we cannot decide where to put can go to pkg for now. 

`apigql` - will be app specific code. Devs are free to modify any code inside that directory.

`tools` - any module which should to be available in vendor direct even if not used should declare inside `tools.go.`

`wireapp` - this is where we include all the modules and works as entrypoint. devs are free to change the code inside this directory.