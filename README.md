### YulMails Application
[![Build Status](https://travis-ci.org/yulPa/yulmails.svg?branch=master)](https://travis-ci.org/yulPa/yulmails)

[API list](https://github.com/yulPa/yulmails/blob/master/API.md)

# Getting started - Docker

Please make sure that [Docker](https://www.docker.com/) is installed on your machine. Then you can clone this repo.

```shell
$ docker version
$ git clone git@github.com:yulpa/yulmails.git
$ cd yulmails/
```

Now, you can `build` locally your `docker` image, for testing purpose:

```shell
$ docker-compose build --no-cache
$ docker images
```
# Getting started - Sources

Please make sure that [Dep](https://github.com/golang/dep) is installed on your machine. Then you can clone this repo.

```shell
$ docker version
$ git clone git@github.com:yulpa/yulmails.git
$ cd yulmails/
$ dep ensure -vendor-only
$ go build -o yulmails main.go
$ chmod +x yulmails
$ sudo mv yulmails /usr/bin/yulmails
```

# Run your application

YulMails is split following four services: an api server for configuration, an entrypoint, a compute node and a sender node. So you need to start at least one of this services to have a ready to production YulMails.

`docker-compose` will start each service with the good command.

Set environment variable then run

```shell
$ export YMAILS_IP_LISTENING=127.0.0.1
$ export YMAILS_PORT_LISTENING=443
$ export YMAILS_VOLUMES_LOGS=/var/log/ymails
$ docker-compose up -d
```

or from sources:

```shell
$ export YMAILS_IP_LISTENING=127.0.0.1
$ export YMAILS_PORT_LISTENING=443
$ export YMAILS_VOLUMES_LOGS=/var/log/ymails
$ #start the api service
$ yulmails api \
-tls-ca-file=/path/to/cert-file.crt \
-tls-key-file /path/to/cert-file.key
$ yulmails api -h
Start the API configuration server

Usage:
  yulmails api  [flags]

Flags:
  -h, --help                  help for api
      --tls-crt-file string   A certificate file (default "domain.tld.crt")
      --tls-key-file string   A key file (default "domain.tld.key")
```

# Kubernetes

Yulmails on Kubernets is currently under development. Documentation will come soon :smiley: :fire: :whale:


# Contributing

If you want to contribute to this project (thanks !), please fork this repo and commit following commit [style](https://github.com/angular/angular.js/blob/master/DEVELOPERS.md#-git-commit-guidelines) by AngularJS.

__Todo__

- [ ] Wrap mongo call with a function in order to get args function (assert.CalledWith something in this A.K.A)

Thanks and happy coding !
