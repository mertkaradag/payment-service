# Payment Service

This project is a **Payment Service** application developed to manage routing to different banks based on specific commission rates.

## Project Structure

```
payment-service/
├── cmd/
│   └── payment/
│       └── main.go                 # Main entry point of the application
├── internal/
│   ├── bank/
│   │   ├── isbank.go               # Operations related to İş Bankası
│   │   ├── akbank.go               # Operations related to Akbank
│   │   ├── halkbank.go             # Operations related to Halkbank
│   │   └── bank.go                 # Core structure managing bank operations
│   └── strategy/
│       ├── strategy.go             # Strategy interface definitions
│       ├── highest_commision.go    # Strategy for the highest commission rate
│       └── lowest_commision.go     # Strategy for the lowest commission rate
├── go.mod                          # Dependencies for the Go module
├── go.sum                          # Summary of dependency versions
└── Makefile                        # Makefile for build and other tasks
```

## Installation and Execution

### Requirements
- **Go 1.23.2+** must be installed.

### Installation Steps
1. Navigate to the project directory:
    ```bash
    cd payment-service
    ```

2. Install necessary dependencies:
    ```bash
    go mod tidy
    ```

3. Run the application:
    ```bash
    go run cmd/payment/main.go
    ```

### Makefile Commands
You can use `Makefile` for project-related tasks:
- **`make run`**: Runs the application
- **`make unit-test`**: Runs the tests
- **`make code-coverage-test`**: Runs the tests and generates a code coverage report

## Usage

- This service routes transactions to three different banks (İş Bankası, Akbank, Halkbank) based on commission rates. The **strategy** package implements strategies to determine the bank with the highest or lowest commission rate.

### Ideas

- Since this project uses the Strategy Pattern, I didn’t set it up as a full microservice. I could have built a structure more suited to business needs by applying Hexagonal Architecture.
