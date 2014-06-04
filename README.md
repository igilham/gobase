gobase
======

Overview
--------

gobase is intended as a base package of simple unix-like utilities,
similar to [GNU coreutils](http://www.gnu.org/software/coreutils/) and
[sbase from suckless.org](http://suckless.org/). These tools are generally
inspired by sbase and aim to be as simple as possible, not to provide a
replacement for GNU coreutils or any other production-quality set of utilities.

The applications are not thoroughly tested and may or may not work. They are
not expected to be useful when compared to the many more mature and complete
implementations available.

License
-------

All code is licensed under the 2-clause FreeBSD license (see file: LICENSE).

Prerequisites
-------------

* Go compiler and tools
* Ruby
* rubygems
* gem install bundler

You will need an installation of the [Go compiler](http://golang.org/) and
tools. The website has some documentation on how to get that set up. If you can
build a hello world program, you should be able to build gobase.

Ruby and cucumber are used for integration testing of the binary applications.

Dependencies
------------

gobase has been built against go 1.0 and 1.1. No major problems have been found
so far.

There are no dependencies outside of the go standard library.

Building
--------

All the build steps are automated in the Makefile (Unix-like only). A simple
build script is provided for Windows users.

The Go convention is to designate the repository root as the *workspace*. In
order to interact with the `go` program and drive the compiler and other tools,
the environment variable `GOPATH` should be set to the workspace root. The
Windows build script and Makefile show an example of this.

In order to build gobase, just run `make install` or `build.cmd` from the
workspace root.

The build scripts pass an argument to the `go` tool and run it for all libs and
applications in the workspace. They will work with all the usual arguments:
`build`, `install`, `test`, `fmt`, `fix`, `vet`, etc.

To build an individual library or application (THING), you can run `go install
THING`.
To run any associated unit tests, run `go test THING`. You can also use the
`fmt`, `vet`, `fix` etc. commands in the same way.

The gobase package will be placed in `pkg`. Application binaries will be placed
in `bin`.

To run the cucumber tests, just run `cucumber` from the checkout root directory.
