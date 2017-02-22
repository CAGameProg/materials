# Setting the GOBIN correctly

Run this command to set your GOBIN environment variable:

```
$ echo 'export GOBIN=$GOPATH/bin' >> ~/.bash_profile
```

# Mac SFML Installation

Installing sfml on Mac is very simple.

Run the following terminal commands:

```
$ brew install sfml
```

```
$ brew install csfml
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
