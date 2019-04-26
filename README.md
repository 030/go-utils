# go-utils

[![Build Status](https://travis-ci.org/030/go-utils.svg?branch=master)](https://travis-ci.org/030/go-utils)

```
go get github.com/030/go-utils
```

and add the following to the go file:

```
import (
    "github.com/030/go-utils"
)
```

## URLExists

```
func main() {
    fmt.Println(utils.URLExists("http://releasesoftwaremoreoften.com"))
}
```

will result in:

```
true
```