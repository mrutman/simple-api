# SimpleAPI

SimpleAPI is a simple and small REST application for test purpose. It uses mariadb as database.

## Build
To build, use the make command:
```sh
$ make
```

## Run
To run, execute ./target/simple -api:
```sh
$ ./target/simple-api
```
or run script:
```sh
$ ./run-simple-api
```

## Configure
SimpleAPI is confirugured through environment variables:
- SIMPLEAPI_SERVER_PORT="8080"
- SIMPLEAPI_DB_HOST="172.17.0.2"
- SIMPLEAPI_DB_PORT="3306"
- SIMPLEAPI_DB_USER="root"
- SIMPLEAPI_DB_PASSWORD="mypass"

## Endpoints:
- /api/config - returns current config (GET)
- /api/db - returns all record from DB (GET)
- /api/db/<endpoint> - returns record from DB for endpoint. If record does not exists it wiil be added and returned (GET, POST work the same way)
- /healthz - returns Code 200 if service is up and running
- /api/healthz - the same as previous
- /livez - the same as previous
- /api/livez - the same as previous
- /readyz - returns Code 200 if service is up and running and database is accessible
- /api/readyz - the same as previous
