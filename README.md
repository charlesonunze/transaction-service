# Transaction Service

A microservice that exposes credit and debit operations.

## Running the application

Make sure you have [Git](https://git-scm.com/downloads) and [Docker](https://docs.docker.com/get-docker/) and [Golang](https://go.dev/doc/install) installed locally.

Also make sure nothing is running on ports **8080**, **7070**, **9090**. Or you can change the ports on in the **.env** file.

Make sure the **WALLET_CLIENT_PORT** value corresponds to the port of the wallet service.

#### Clone the project locally

```bash
git clone https://github.com/charlesonunze/transaction-service.git && cd transaction-service
```

#### Install packages

```bash
go get -u ./...
```

#### Running the database

```bash
make db
```

#### Running the microservice

```bash
make run
```

#### Running the tests

```bash
make test
```

## Possible Improvements

#### Tests

I did not add tests for the GRPC services. This is bad and should never make it to prod.
