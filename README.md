### YulMails Application
[![Build Status](https://travis-ci.org/yulPa/yulmails.svg?branch=master)](https://travis-ci.org/yulPa/yulmails)

The goal of this `go` application is to check mails between sender and recipients. (#TODO: Add a better description)

# Getting started

Please make sure that [Docker](https://www.docker.com/) is installed on your machine. Then you can clone this repo.

```shell
$ docker version
$ git clone https://github.com/yulpa/yulmails
$ cd check_mails/
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

# API

This is the list of available API:

* Create a Pool of authorized environment
```golang
Route{
  Method:      "POST",
  Pattern:     "/api/v1/authorizedpool",
}
```
Parameter:

```json
{
  "environments": [
    {
      "name": "zimbra",
      "domains": [
        "zimbra1.com",
        "zimbra2.com",
        "zimbraN.com"
      ]
    },
    {
      "name": "http",
      "domains": [
        "http1.com",
        "http2.com",
        "httpN.com"
      ]
    }
  ]
}
```
