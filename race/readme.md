# Go Race Detector

## Race condition

<https://www.wikiwand.com/en/Race_condition>

## Deadlock

<https://www.wikiwand.com/zh/%E6%AD%BB%E9%94%81>
<https://www.wikiwand.com/en/Deadlock>

## Blog

<https://blog.golang.org/race-detector>

## Memory model

<https://golang.org/ref/mem>

## Sanitizer

<https://github.com/google/sanitizers/wiki/ThreadSanitizerGoManual>
<https://github.com/google/sanitizers/wiki/ThreadSanitizerAlgorithm>

## Using the race detector

``` sh
# -race 选项可以用在多个命令中
$ go test -race mypkg          // test the package
$ go run -race mysrc.go        // compile and run the program
$ go build -race mycmd         // build the command
$ go install -race mypkg       // install the package
$ go get -race path/to/package // get the package
```