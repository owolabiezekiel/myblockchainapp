## Installation

### Install Go 1.14 or higher
Follow the official docs or use your favorite dependency manager
to install Go: [https://golang.org/doc/install](https://golang.org/doc/install)

Verify your `$GOPATH` is correctly set before continuing!

### Setup this repository

Go is a bit picky about where you store your repositories.

The convention is to store:
- the source code inside the `$GOPATH/src`
- the compiled program binaries inside the `$GOPATH/bin`

#### Using Git
```bash
mkdir -p $GOPATH/src/com.owoez/myblockchainapp
cd $GOPATH/src/com.owoez/myblockchainapp

git clone https://github.com/owolabiezekiel/myblockchainapp.git
```

**PS:** Make sure you actually clone it inside the `src/com.owoez/myblockchainapp` directory, not your own, otherwise it won't compile. Go rules.
