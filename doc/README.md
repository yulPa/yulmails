# yulmails code documentation

## packages

The code is split in a few independants modules: 

* `pkg/client/`: a simple http client in order to post data on different services (ex: iWal notifications)
* `pkg/configuration/`: read the configuration from a YAML file with different services (archiving, senders, computes, entrypoints) and the different options belonging to these services. This configuration can be save on a database into the schema `global/configuration`
* `api/`: the main configuration package, this api server will expose a bunch of *CRUD* (Create Read Update Delete) api in order to manage entities, environments, etc.[more informations](./modules/api.md) 
* `pkg/domain,entity,environment`: utils structure to serialize data into the database [more informations](./modules/domain-entity-environment.md)
* `pkg/logger`: custom logger configuration to call at each new modules, in order to have a split log.
* `pkg/mail`: mail structure in order to serialize them into the database
* `pkg/mocks`: main package to create mocks and fixtures in order to test our apis call and database manipulation without use them for real
* `pkg/mongo`: main module with DAO implementation in order to use database resource without managing session creation and factory. [more informations](./modules/mongo.md)
* `pkg/options`: saving options 
* `pkg/sender`: module to use when you want to sent an email trough a given smtp server
* `pkg/mta`: entrypoint to get email to analyze, a simple MTA server with a redis behind
* `pkg/processing`: this module will consume redis queue and process the email against anti-spam services

## development workflow

There is one main entrypoint: `main.go`with an embedded CLI: you can start your differents service.
There is an example with the `docker-compose` file.

```shell
$ // to start the entrypoint
$ yulmails entrypoint 
$ // to start the sender service
$ yulmails sender
```

You can decide to fire-up your services with Docker or manually. I highly recommend to use Docker: you'll have your redis and postgresql ready to be used. 

```
$ docker-compose up -d redis postgres
$ docker-compose up -d sender entrypoint api...
```