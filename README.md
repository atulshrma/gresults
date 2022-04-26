<div align="center">

<img src="https://github.com/atulshrma/gresults/raw/main/assets/images/gresults.svg" width="200" height="200" />
<h1>Go-Results</h1>

</div>

`gresults` is an opinionated development framework that enables developers to build _trustless_ services, by having the caller define default return values and error handlers for all function calls.

## Dependencies

* An installation of Go 1.18 or later. 

## Installation

```
go get github.com/atulshrma/gresults
```

## Usage

To create a new `Result` object, use the `NewResult` factory and pass in the data and/or error the object should wrap

```go
package main

import (
  "fmt"
  results "github.com/atulsharma/gresults"
)

func main() {
  want := 42
  var setInErrorHandler int
  successResult := results.NewResult[int, error](&want, nil)
  errorHandler := func(err error) {
    // Executes in current scope
    setInErrorHandler = -42
    fmt.Printf("encountered err %q, setting var to %q", err, setInErrorHandler)
  }
  got := successResult.OnError(errorHandler).Unwrap(0)
  fmt.Printf("want result %q, got result %q", want, got)
}
```

You can also wrap existing functions using the `Resultify` function wrapper which wraps the return value in a `Result` object

```go
package main

import (
  "fmt"
  "strconv"
  results "github.com/atulsharma/gresults"
)

func main() {
	wrappedAtoiResult := results.Resultify[int, error](strconv.Atoi, "-42")
	errorHandler := func(err error) {
		fmt.Printf("encountered err %q", err)
	}
	got := wrappedAtoiResult.OnError(errorHandler).Unwrap(0)
  fmt.Printf("want result %q, got result %q", -42, got)
}
```

## How to test the software

```sh
$ go test
```

## Getting involved

General instructions on _how_ to contribute can be found in the [CONTRIBUTING](CONTRIBUTING.md) document.

----

## Open source licensing info

This project is licensed under the terms of the MIT license. See [LICENSE](LICENSE) for more details.

## Citation

```
@misc{gresults,
  author = {Atul Sharma},
  title = {Go-Results, an opinionated development framework for trustless services.},
  year = {2022},
  publisher = {GitHub},
  journal = {GitHub repository},
  howpublished = {\url{https://github.com/atulshrma/gresults}}
}
```

----

## Credits and references

1. Inspiration for parts of the project: [Resultify](https://github.com/blackblood/Resultify/tree/master)
