# fdate
Go package similar to the native `time` package except implemented on the
[French Republican Calendar](https://en.wikipedia.org/wiki/French_Republican_Calendar). 

Unlike the `time` package, this one measures dates relative to how many days have passed since
22 September 1792 and doesn't measure anything smaller than days.

## Install

To install, have Go installed and run
```
go get github.com/rfaulhaber/fdate
```