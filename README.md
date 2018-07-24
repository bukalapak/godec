[![Build Status](https://travis-ci.org/bukalapak/godec.svg?branch=master)](https://travis-ci.org/bukalapak/godec)
[![Go Report Card](https://goreportcard.com/badge/github.com/bukalapak/godec)](https://goreportcard.com/report/github.com/bukalapak/godec)
[![Documentation](https://godoc.org/github.com/bukalapak/godec?status.svg)](http://godoc.org/github.com/bukalapak/godec)

# Godec

## Description

Godec is decorator fabricator. Godec will automatically create a decorator for desired interface using desired template. The desired template must be placed in folder [template/](https://github.com/bukalapak/godec/tree/master/template).

## Usage

1. Install godec

    ```sh
    go install github.com/bukalapak/godec/app/godec
    ```

2. Add `go:generate` hook in desired file

    ```go
    //go:generate godec MyInterface canceler
    type MyInterface interface {
      Sample(int, string) (float64, error)
    }
    ```

    godec needs two or more arguments to run correctly.

    ```sh
    godec [interface name] [decorator name]...
    ```

    [interface name] is the name of interface to be decorated. In the example above, interface's name is `MyInterface`.
    [decorator name] is the name of decorator. In the example above, decorator's name is `canceler`. Please, note that desired decorator must be placed in [template/](https://github.com/bukalapak/godec/tree/master/template) folder. Godec will find the desired template there. The name of the template decorator must be `<name>.go.tmpl`. Please, look at [template/](https://github.com/bukalapak/godec/tree/master/template) for clarity.

3. Run `go generate`

    ```sh
    go generate
    ```

    Decorator will be created in folder `decorator/`.
