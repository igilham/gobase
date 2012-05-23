gocoreutils
===========

gocoreutils is intended as a base package of simple unix-like utilities,
similar to [GNU coreutils](http://www.gnu.org/software/coreutils/) and
[sbase from suckless.org](http://suckless.org).

All code is licensed under the 2-clause FreeBSD license (see file: LICENSE).

Building
--------

In order to build **gocoreutils**, just run `./all.bash` or `all.cmd` from the
project's root folder, then copy the **_bin** folder wherever you want.

The build script is built using Go-GB, so if you have it installed, simply run
`gb -pi /path/to/gocoreutils/src` or just `gb` is it is already set up nicely.
