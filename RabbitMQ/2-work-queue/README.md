# Work Queue

## How to run
### Worker
1. go run worker.go
You can run multiple worker

### Publisher
1. go run new_task [OPTIONAL]message

If message contains dot, 1 dot = 1 second worker sleep. It for simulate slow task
