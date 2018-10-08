# errors [![Travis-CI](https://travis-ci.org/theodesp/bisect.svg)](https://travis-ci.org/theodesp/bisect)  
[![GoDoc](https://godoc.org/github.com/theodesp/bisect?status.svg)](http://godoc.org/github.com/theodesp/bisect) 
[![Report card](https://goreportcard.com/badge/github.com/theodesp/bisect)](https://goreportcard.com/report/github.com/theodesp/bisect) 

Package bisect implements common bisection algorithms.

`go get -u github.com/theodesp/bisect`

## Examples

```go
package main

import (
	"github.com/theodesp/bisect"
	"fmt"
)

func main()  {
	var ints bisect.IntSlice
	for i:=0;i<10;i += 2 {
		ints = append(ints, i)
	}

	ints.InsortLeft(-1)
	ints.InsortRight(16)

	fmt.Println(ints) // [16 -1 0 2 4 6 8]
}
```

[Read the package documentation for more information](https://godoc.org/github.com/theodesp/bisect).

## Contributing

We welcome pull requests, bug fixes and issue reports.
Before proposing a change, please discuss your change by raising an issue.

## License

Copyright © 2017 Theo Despoudis MIT license