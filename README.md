# fastid

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]

Fast ID generator based on timestamp, sequence number and worker id.

## Rationale

Generating IDs quickly is a common task. Making this quickly is a nice thing to have.

## Features

* Simple.
* Thread-safe.
* Time based.
* Dependency-free.

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

See [these docs][pkg-url].

## License

[MIT License](LICENSE).


[build-img]: https://github.com/cristalhq/fastid/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/fastid/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/fastid
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/fastid
[reportcard-img]: https://goreportcard.com/badge/cristalhq/fastid
[reportcard-url]: https://goreportcard.com/report/cristalhq/fastid
[coverage-img]: https://codecov.io/gh/cristalhq/fastid/branch/master/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/fastid
