# golang-utils

[![Build Status](https://travis-ci.org/030/go-utils.svg?branch=master)](https://travis-ci.org/030/go-utils)

Golang Utils

```
go get github.com/030/go-utils
```

and add the following to the go file:

```
import (
    "github.com/030/go-utils/utils"
)
```

## debug

```
func main() {
    utils.Debug()
    
    logrus.Info("HelloWorld!")
}
```

will result in:

```
time="2019-02-21T18:02:23+01:00" level=info msg=HelloWorld! func=main.main file="C:/path/to/go/src/github.com/030/some-go-project/main.go:16"
```
