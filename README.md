### YulMails Application
[![Build Status](https://travis-ci.org/yulPa/yulmails.svg?branch=master)](https://travis-ci.org/yulPa/yulmails)

The goal of this `go` application is to check mails between sender and recipients. (#TODO: Add a better description)
[API list](https://github.com/yulPa/yulmails/blob/master/API.md)

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

Please update `yulmails.yaml` configuration file following your installation:

`version` is the current running version of `yulmails` configuration.
`services` is an array of 4 services:

  * `entrypoint`: This is the interface between `yulmails` and your system
  * `compute`: This service will check if a mail is considered as a spam
  * `sender`: This service will send mails to recipients
  * `archiving_db`: Archiving database is provided by `yulmails` but you can use your own, you just need to use a mongo DB database.

A default configuration file is already provided in order to run `yulmails` on a single machine.

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

- [ ] Find a good way to document APIs
- [ ] Add details to logger (env, ent, etc.): not only "not found"
- [ ] Wrap mongo call with a function in order to get args function (assert.CalledWith something in this A.K.A)

Thanks and happy coding !
