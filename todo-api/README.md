## getting started

```sh
$ docker-compose build
$ docker-compose run --rm todo-api make init-db
$ docker-compose run --rm todo-api make migrate
$ docker-compose up -d
```

### build development mode

```sh
$ make start
```

### build production mode

```sh
$ make build
```

### test

```sh
$ make test # run all the tests
$ go test -v ./xxx # run tests under the specified directory
```
