# OMP test task
This is a simple csv and json processor.
It consumes a filepath to csv or json from arguments and outputs products with max price and rating values
```
cd omp-test-task
go run main.go db.csv
```
or
```
go run main.go db.json
```
## Running the tests
```
cd omp-test-task/test
go test ./...
```