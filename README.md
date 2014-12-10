go-tle
======

just what is it?
----------------

just a quick and dirty implementation of two-line element sets. essentially it takes in a stream of
data, converts it to the TLE format, and then outputs to another stream the JSON.

how do i get this awesome piece of software?
--------------------------------------------

`go get github.com/bmallred/go-tle`

how do i use it?
----------------

```
package main

import (
	"bufio"
	"os"

	"github.com/bmallred/go-tle"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	tle.Scan(os.Stdin, out)
}
```

then build it

`go build`

which should create your executable (will assume it is called `bodacious`) and then do something
like this

```
curl "http://www.celestrak.com/NORAD/elements/tle-new.txt" | ./bodacious
```
