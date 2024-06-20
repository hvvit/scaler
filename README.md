
# Scaler

Project to scale in and scaleout replicas based on the current number of replicas present.

## Project Structure

```
LICENSE                         - license file
Makefile                        - make file
README.md                       - README file
cmd                             - cmd package
-/root.go                       - handles cobra init for root command that is if i start scaler command
-/watch.go                      - handles the subcommand with scaler watch, has all the default parameters needed to be triggered.
go.mod                          - go mod file
go.sum                          - go sum file
main.go                         - main file that calls the cmd package
watcher                         - watcher package that handles watch for changes
-/replica.go                    - file that handles replica logic for changes, like if there is need for scaling.
-/replica_test.go               - testfile for checking desiredreplica logic, can be expanded
-/watcher.go                    - has the watcher logic for scraping details, this file calls functions from replica.go
webhandler                      - webhandler package that handles web requests to app
-/structs.go                    - has all the structs required for json payload requests.
-/webhandler.go                 - has all the functions for making http requests and url formatting. only handles web requests, app logic should not be kept here
```

## Building the Project

for default use `make build` 
```sh
➜  scaler make build
CGO_ENABLED=0
go build -o scaler
```

## Testing the Project 

for triggering test run `make test`

```
go test -v ./...
?       github.com/hvvit/scaler [no test files]
?       github.com/hvvit/scaler/cmd     [no test files]
=== RUN   TestIsChangeRequired
--- PASS: TestIsChangeRequired (0.00s)
=== RUN   TestCalculateDesiredReplica
--- PASS: TestCalculateDesiredReplica (0.00s)
PASS
ok      github.com/hvvit/scaler/watcher (cached)
=== RUN   TestUrlFormatter
--- PASS: TestUrlFormatter (0.00s)
PASS
ok      github.com/hvvit/scaler/webhandler      (cached)
```

## How to run the Project

Once the binary is generated

there are multiple configurable options like which can be explored

```sh
➜  scaler ./scaler watch --help
takes parameters like threshold value, scrape time and url to watch the server load and change replica if required

Usage:
  scaler watch [flags]

Flags:
  -h, --help                help for watch
  -s, --scrape-time float   Time to wait before scraping the server load (default 1)
  -t, --threshold float     Threshold value for the server load to monitor (default 0.8)
  -u, --url string          URL of the server to monitor (default "http://localhost:8123")
```

Triggering with default parameters
```
➜  scaler ./scaler watch
2024/06/21 00:43:41 watching the server load with threshold value:  0.8  scrape time:  1  url:  http://localhost:8123
2024/06/21 00:43:41 Watcher is running
2024/06/21 00:43:41 will scrape url: http://localhost:8123 after 1 seconds
2024/06/21 00:43:42 current cpu utilization for host: http://localhost:8123 is 0.84
2024/06/21 00:43:42 current replicas are: 13 ,current metrics value is: 0.84
2024/06/21 00:43:42 desired replicas are: 14
2024/06/21 00:43:42 error in updating the replica Error in getting the response from the server with status code: 500, with message: error updating replicas

2024/06/21 00:43:42 will scrape url: http://localhost:8123 after 1 seconds
2024/06/21 00:43:43 current cpu utilization for host: http://localhost:8123 is 0.82
2024/06/21 00:43:43 current replicas are: 13 ,current metrics value is: 0.82
2024/06/21 00:43:43 desired replicas are: 14
2024/06/21 00:43:43 replicas are updated to: 14
2024/06/21 00:43:43 will scrape url: http://localhost:8123 after 1 seconds
2024/06/21 00:43:44 error in getting the response from the server Error in getting the response from the server with status code: 500
2024/06/21 00:43:44 will scrape url: http://localhost:8123 after 1 seconds
2024/06/21 00:43:45 current cpu utilization for host: http://localhost:8123 is 0.8
2024/06/21 00:43:45 current replicas are: 14 ,current metrics value is: 0.8
2024/06/21 00:43:45 will scrape url: http://localhost:8123 after 1 seconds
2024/06/21 00:43:46 error in getting the response from the server Error in getting the response from the server with status code: 500
2024/06/21 00:43:46 will scrape url: http://localhost:8123 after 1 seconds
2024/06/21 00:43:47 current cpu utilization for host: http://localhost:8123 is 0.7
2024/06/21 00:43:47 current replicas are: 14 ,current metrics value is: 0.7
2024/06/21 00:43:47 desired replicas are: 13
2024/06/21 00:43:47 replicas are updated to: 13
2024/06/21 00:43:47 will scrape url: http://localhost:8123 after 1 seconds
2024/06/21 00:43:48 current cpu utilization for host: http://localhost:8123 is 0.74
2024/06/21 00:43:48 current replicas are: 13 ,current metrics value is: 0.74
2024/06/21 00:43:48 desired replicas are: 13
2024/06/21 00:43:48 replicas are updated to: 13
2024/06/21 00:43:48 will scrape url: http://localhost:8123 after 1 seconds
2024/06/21 00:43:49 current cpu utilization for host: http://localhost:8123 is 0.67
2024/06/21 00:43:49 current replicas are: 13 ,current metrics value is: 0.67
2024/06/21 00:43:49 desired replicas are: 11
2024/06/21 00:43:49 replicas are updated to: 11
```
