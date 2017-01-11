# Installing Git

Git is a version control system. We need it for two reasons: (1) many components 
of Go are distributed via Git and (2) the Windows installer includes a terminal
that makes it behave more like Unix operating systems (like OSX or Linux).

Download and install msysgit from [here](https://git-for-windows.github.io/). 
When the installer asks make sure to select "Use Git from the Windows Command Prompt".

# Installing Atom

Atom is a text editor that we will use to write code. If you are comfortable
using a different editor feel free to do so.

Go to [atom.io](https://atom.io) and click download. Installation should
be fairly straightforward.

We are also going to install improved support for Go in Atom.

Navigate to `Packages -> Settings View -> Open`. Then click on the `Install`
tab and search for `go-plus`. It should find a package called `go-plus` with
the description: `Makes working with go in Atom awesome`. Click `Install`.

# Installing Go

You can install the Go Programming Language from [here](https://golang.org/dl/).
Choose windows-amd64.msi for 64 bit or windows-386.msi for 32 bit (you probably 
want the 64 bit version).

We need to configure two environment variables: PATH and GOPATH. PATH is used by
the operating system to locate executables. GOPATH is used by the Go 
compiler to find go source code.

Open Git Bash and execute the following:

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
