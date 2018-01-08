### YulMails Application
[![Build Status](https://travis-ci.org/yulPa/yulmails.svg?branch=master)](https://travis-ci.org/yulPa/yulmails)

The goal of this `go` application is to check mails between sender and recipients. (#TODO: Add a better description)

# Getting started

Please make sure that [Docker](https://www.docker.com/) is installed on your machine. Then you can clone this repo.

```shell
$ docker version
$ git clone https://github.com/yulpa/yulmails
$ cd yulmails/
```

Now, you can `build` locally your `docker` image, for testing purpose:

```shell
$ docker-compose build --no-cache
$ docker images
```

# Configuration

Please update `yulmails.yaml` configuration file following your installation. (#TODO: We could provide a MySql Database with Compose ?)

# Run your application

Set environment variable then run

```shell
$ export YMAILS_IP_LISTENING=127.0.0.1
$ export YMAILS_PORT_LISTENING=80
$ export YMAILS_VOLUMES_LOGS=/var/log/ymails

$ docker-compose up -d
$ docker-compose logs
```

# Contributing

If you want to contribute to this project (thanks !), please fork this repo and commit following commit [style](https://github.com/angular/angular.js/blob/master/DEVELOPERS.md#-git-commit-guidelines) by AngularJS.

__Todo__

- [ ] Add a NoSql test database
- [ ] Provides data to populate NoSql test database
- [ ] Add Mock to test server API
- [ ] Add Mock to test Mongo DB dial
- [ ] Create a `API.md` in order to document API
- [ ] Add a Makefile to build project from sources

Thanks and happy coding !

# API

This is the list of available API:

* Create an Entity
```golang
Route{
  Method:      "POST",
  Pattern:     "/api/v1/entity",
}
```
Parameter:

```json
{
  "name": "An entity",
  "abuse": "abuse@domain.tld",
  "conservation":{
    "sent": 5,
    "unsent": 2,
    "keep": true
  }
}
```

* Get list of Entitys
```golang
Route{
  Method:      "GET",
  Pattern:     "/api/v1/entity",
}
```
Return:

```json

{
  "entitys":[
    {
      "name": "entity_name",
      "abuse": "abuse@domain.tld",
      "conservation":{
        "sent": 5,
        "unsent": 2,
        "keep": true
      }
    },    
    {
      "name": "Another entity",
      "abuse": "another-abuse@domain.tld",
      "conservation":{
        "sent": 1,
        "unsent": 2,
      }
    }   
  ]
}
```

* Create an environment associated to an entity
```golang
Route{
  Method:      "POST",
  Pattern:     "/api/v1/entity/{entity_name}/environment",
}
```

Parameter:
```Json
{
  "ips": [
    "192.168.0.1",
    "192.168.0.2",
    "192.168.0.3"
  ],
  "abuse": "abuse@domain.tld",
  "open": false,
  "quota": {
    "tenlastminutes": 150,
    "sixtylastminutes": 200,
    "lastday": 1000,
    "lastweek": 3000,
    "lastmonth": 10000
  }
}
```
