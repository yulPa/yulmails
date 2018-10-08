# api

in order to be able to customize yulmails, we setted up a few apis. Each api is in REST format, following [CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) architecture.

api will be reachable on http://localhost:8080/ but behind a proxy (example provided nginx), you'll need to get the domain name and the SSL certificates, the authentication is made through SSL certificates.

`server.go` is the entrypoint to start the HTTP server and create the database context, to pass to the router.

`router.go` is the file where each `resource` is defined.

`handlers.go` is the file where the `handlers` belonging to the `resources` are defined.

`api_test.go` is the main file to test global behaviors.

# workflow

in order to add a new API, you'll need to add a `resource`, bind it to the new `handler`, add the associated `mock`, finally add your test.

# example 

to create an entity:

we add the new [resource](https://github.com/yulPa/yulmails/blob/develop/api/router.go#L44)

Then we create the associated [handler](https://github.com/yulPa/yulmails/blob/develop/api/handlers.go#L13)

we add the mock to perform the unit test, and finally, we create the [test](https://github.com/yulPa/yulmails/blob/develop/api/api_test.go#L85)

