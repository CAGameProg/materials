# Setting the GOBIN correctly

Run this command to set your GOBIN environment variable:

```
$ penv set GOBIN $GOPATH/bin
```

```
$ penv append PATH $GOBIN
```

# Windows SFML Installation

Head over to this page to download CSFML: http://www.sfml-dev.org/download/csfml/

Download the `Visual C++ / GCC - 32-bit` version of `CSFML 2.3`.

Unzip the file and place it in your home directory.

Now run the following commands to set everything up correctly.

```
$ penv set C_INCLUDE_PATH ~/CSFML/include
```

```
$ penv set LIBRARY_PATH ~/CSFML/lib
```

```
$ penv append PATH ~/CSFML/bin
```

# Installing GoSFML

Now you can install the Go version of SFML:

```
$ go get github.com/manyminds/gosfml
```

You should be finished now.

# Testing GoSFML

Try running the [`rect.go`](./rect.go) program. It should open a window and
show a blue rectangle.
