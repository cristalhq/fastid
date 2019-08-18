# fastid

[![Build Status][travis-img]][travis-url]
[![GoDoc][doc-img]][doc-url]
[![Go Report Card][reportcard-img]][reportcard-url]
[![Go Report Card][coverage-img]][coverage-url]

Fast ID generator based on timestamp, sequence number and worker id.

## Features

* Concurrent safe.
* Time based.
* Minimal.

## Install

```
go get github.com/cristalhq/fastid
```

## Example

```go
gen, err := NewGenerator(100, 200)

id := gen.Next()

ts, seq, w := id.Parts()

ts = id.Timestamp()
seq = id.Sequence()
w = id.WorkerID()
```

## Documentation

See [these docs](https://godoc.org/github.com/cristalhq/fastid).

## License

[MIT License](LICENSE).

[travis-img]: https://travis-ci.org/cristalhq/fastid.svg?branch=master
[travis-url]: https://travis-ci.org/cristalhq/fastid
[doc-img]: https://godoc.org/github.com/cristalhq/fastid?status.svg
[doc-url]: https://godoc.org/github.com/cristalhq/fastid
[reportcard-img]: https://goreportcard.com/badge/cristalhq/fastid
[reportcard-url]: https://goreportcard.com/report/cristalhq/fastid
[coverage-img]: https://coveralls.io/repos/github/cristalhq/fastid/badge.svg?branch=master
[coverage-url]: https://coveralls.io/github/cristalhq/fastid?branch=master
