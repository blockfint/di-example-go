## To start the development server
```sh
docker compose up --build
```

## Prerequisite to run test
1. go version 1.18
2. [install ginkgo](https://onsi.github.io/ginkgo/#getting-started)

## To run all test
```sh
docker compose -f docker-compose.test.yml up
ginkgo -v -r
```
