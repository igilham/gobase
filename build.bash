#!/bin/bash
# driver script to build/install/test/fix/vet all projects

GOPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
LIBS="gobase"
APPS="basename cat cksum dirname echo false head ls mkdir pwd rm seq sleep sort tee touch true uniq wc whoami yes"
COMMAND=$1

go ${COMMAND} ${LIBS} ${APPS}

