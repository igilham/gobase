gocoreutils
===========

gocoreutils is intended as a base package of simple unix-like utilities,
similar to [GNU coreutils](http://www.gnu.org/software/coreutils/) and
[sbase from suckless.org](http://suckless.org). These tools are generally
inspired by sbase and aim to be as simple as possible.

All code is licensed under the 2-clause FreeBSD license (see file: LICENSE).

Building
--------

In order to build **gocoreutils**, just run `./all.bash install` or `all.cmd install` from the
workspace root.

The build scripts pass an argument to the `go` tool and run it for all libs and applications in
the workspace. They will work with all the usual arguments: `build`, `install`, `test`, `fmt`,
`fix`, `vet`, etc.

