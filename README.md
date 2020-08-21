# go-utils

[![Build Status](https://travis-ci.org/030/go-utils.svg?branch=master)](https://travis-ci.org/030/go-utils)

```bash
go get github.com/030/go-utils
```

and add the following to the go file:

```bash
import (
    "github.com/pkg/errors"

    "github.com/030/go-utils"
)
```

## CheckErrorPrintStackTraceAndExit

```bash
utils.CheckErrorPrintStackTraceAndExit(errors.WithStack(err))
```
