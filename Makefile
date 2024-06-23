run:
	go mod tidy
	go run main.go

build:
	go mod tidy
	go build -o EMPLOYEE_CRUD_MUX

