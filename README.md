### Simple Go Application

The goal of this `go` application is to check mails between sender and recipients. (#TODO: Add a better description)

# Getting started

Please make sure that [Docker](https://www.docker.com/) is installed on your machine. Then you can clone this repo.

```shell
$ docker version
$ git clone https://github.com/yulpa/check_mails
$ cd check_mails/
```

Now, you can `build` locally your `docker` image, for testing purpose:

```shell
$ docker-compose build --no-cache
$ docker images
```

# Configuration

Please update `check_mail.yaml` configuration file following your installation. (#TODO: We could provide a MySql Database with Compose ?)

# Run your application

```shell
$ docker-compose up -d
$ docker-compose logs
```
