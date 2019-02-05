# clean-arch-example

## How to run this project

### Install libraries on local machine

```bash
make init
```

### Get Project with go get

```bash
go get github.com/hiromisuzuki/clean-arch-example
```

### Go to directory

```bash
cd $(go env GOPATH)/src/github.com/hiromisuzuki/clean-arch-example
```

### Download dependencies

```bash
dep ensure
```

### Build and run

Local machine build command.

* create DB image.
* run DB container. (on localhost:3306)
* create API image.
* run API container
* execute API service. (on localhost:8080)
* provide swagger documentation (on localhost:8080/swagger/index.html)
* create phpMyAdmin image.
* run phpMyAdmin container.
* execute phpMyAdmin service. (on localhost:8888)

```bash
cd $GOPATH/src/github.com/hiromisuzuki/clean-arch-example
make build
make migrate
```

### Generate swagger files

```bash
cd $GOPATH/src/github.com/hiromisuzuki/clean-arch-example
make swag
```