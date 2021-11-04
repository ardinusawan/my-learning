# Routing

## How to run
### Consumer
1. Open three terminal
1. Terminal 1: 
    ```go 
    go run receive_logs_direct.go error
    ```
1. Terminal 2: 
    ```go 
    go run receive_logs_direct.go warning
    ```

### Publisher
1. Terminal 3: 
    ```go 
    go run emit_log_direct.go error "Run. Run. Or it will explode. hhe"
    go run emit_log_direct.go warning "Run. Run. Or it will explode. hhe"
    ```

Message will be consumed by correct routing key
