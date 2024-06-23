# Employee Management API Documentation

## Build

To build the application, run:

```bash
make build
```

## Run

To run the application locally, use:

```bash
make run
```

### How to use?

To use this application, first start the service locally by running:

```bash
make run
```

Then, use the following endpoints:

#### Create Employee

```bash
Endpoint: http://localhost:8000/api/createEmployee
```

Make a POST request to the above endpoint with the following payload:

```bash
{
    "name": "John Doe",
    "age": 35,
    "position": "Senior Developer"
}
```

The request will return a Employee ID  like this:"6678269a735ff8996176dfc6"

#### Get Employees

```bash
Endpoint: http://localhost:8000/api/getAllEmployees/
```

Make a GET request to the above endpoint:

The request will return a Employees Details"

#### Get EmployeeByID

```bash
Endpoint: http://localhost:8000/api/getAllEmployeeById/{id}
```

Make a GET request to the above endpoint:

The request will return a Employee Detail"


#### Update Employee

```bash
Endpoint: http://localhost:8000/api/updateEmployee/{id}
```

Make a PUT request to the above endpoint with the following payload:

```bash
{
    "name": "John Doe",
    "age": 35,
    "position": "Senior Developer"
}
```

The request will return a Message Employee Updated successfully"


#### Patch Employee

```bash
Endpoint: http://localhost:8000/api/patchEmployee/{id}
```

Make a PATCH request to the above endpoint with the following payload:

```bash
{
    "name": "John Doe",
    "age": 35,
    "position": "Senior Developer"
}
```

The request will return a Message Employee patched successfully "

#### Delete Employee

```bash
Endpoint: http://localhost:8000/api/deleteEmployee/{id}
```

Make a Delete request to the above endpoint:

The request will return a Employee deleted successfully"
