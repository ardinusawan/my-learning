# Pub Sub

## How to run
### Consumer
1. Open two terminal
1. Terminal 1: 
    ```go 
    go run receive_logs.go > logs_from_rabbit.log
    ```
1. Terminal 2: 
    ```go 
    go run receive_logs.go
    ```

### Publisher
1. go run emit_logs.go [OPTIONAL]message

Message will be consumed by all consumer
