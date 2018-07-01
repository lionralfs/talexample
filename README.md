# Example TAL Application (but in Go)

An example TV application using the [TAL](https://www.github.com/bbc/tal) framework from the BBC.

## Getting Started

```
git clone git@github.com:lionralfs/talexample.git
cd talexample
npm install
go run main.go
```

> Note: `npm install` is still necessary to install the js-scripts running in the browser and to get the device configs

Visit http://localhost:8080 in your browser. Use the UP, DOWN, LEFT, RIGHT keys to navigate, use ENTER/RETURN to select.

## Benchmark

> Ran on a 2015 MacBook Air (1,6 GHz Intel Core i5)

### Golang
```
Running 40s test @ http://localhost:8080
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    60.58ms   47.38ms 488.57ms   64.18%
    Req/Sec     1.04k   164.20     1.53k    68.81%
  166340 requests in 40.06s, 563.31MB read
  Socket errors: connect 153, read 0, write 0, timeout 0
Requests/sec:   4151.80
Transfer/sec:     14.06MB
```

### NodeJS
```
Running 40s test @ http://localhost:1337
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   161.34ms   13.51ms 466.34ms   93.70%
    Req/Sec   383.65    136.91   666.00     58.58%
  61125 requests in 40.04s, 210.56MB read
  Socket errors: connect 153, read 68, write 22, timeout 0
Requests/sec:   1526.50
Transfer/sec:      5.26MB
```

## More Information

See [github.com/bbc/tal](https://www.github.com/bbc/tal) or [bbc.github.io/tal](http://bbc.github.io/tal/getting-started/introducing-tal.html) for documentation.
