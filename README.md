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

[API list](https://github.com/yulPa/yulmails/blob/master/API.md)

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

- [ ] Provides data to populate NoSql test database
- [x] Add Mock to test Mongo DB dial
- [x] Create a `API.md` in order to document API
- [ ] Add a Makefile to build project from sources

Thanks and happy coding !
