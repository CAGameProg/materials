# Xcode Command Line Tools

Mac OSX lets you install many tools built for software development through
Xcode, but does not install them by default.

Here is how to install them. Open the Terminal (located in `/Applications/Utilities/Terminal.app`).

Now type `gcc` and press enter. If you see `clang: error: no input files` then
the command line tools are already installed and you can skip to the next step.

Otherwise, you should see a window pop up asking if you'd like to install
the tools. Click `install`.

# Installing Homebrew

Homebrew (or Brew for short) is a package manager for Mac. It lets you easily
install software that Apple does not provide in the command line tools.

To install it, open the Terminal, and paste:

```
ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

And press enter. It should let you know what it is doing as it installs.
You may also have to press the enter key and type in your password.

*When you enter your password, you will not be able to see the characters being typed and your cursor won't move. This is for security reasons. Don't worry, you are still typing in the password.*

# Installing Go

To install the Go programming language, enter `brew install go` into the terminal.

We need to configure two environment variables: PATH and GOPATH. PATH is used by
the operating system to locate executables. It's what makes it possible to execute
go version instead of /usr/local/go/bin/go version. GOPATH is used by the Go 
compiler to find go source code.

Open a terminal and execute the following:

```
mkdir ~/gopath
GOPATH=~/gopath go get github.com/badgerodon/penv/...
```

This installs an executable named penv which we can use to set environment variables. 
To update `PATH` run this:

```
~/bin/penv append PATH ~/gopath/bin
```

To update `GOPATH` run this:

```
~/bin/penv set GOPATH ~/gopath
```

Finally, close and reopen the terminal to apply the changes.

To make sure everything was installed correctly, in the terminal, execute

```
go get github.com/k0kubun/tetris ; tetris
```

A game a tetris should begin playing. Press q to close it.

# Installing Atom

Atom is a text editor that we will use to write code. If you are comfortable
using a different editor feel free to do so.

Go to [atom.io](https://atom.io) and click download. Installation should
be fairly straightforward.

We are also going to install improved support for Go in Atom.

Navigate to `Packages -> Settings View -> Open`. Then click on the `Install`
tab and search for `go-plus`. It should find a package called `go-plus` with
the description: `Makes working with go in Atom awesome`. Click `Install`.

# Installing Micro

Micro is a simple terminal-based text editor (made by me!). To install,
open the Terminal and type `brew install micro`.
