#!/bin/bash
# driver script to build/install/test/fix/vet all projects

GOPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
LIBS="gobase"
APPS="basename cat cksum dirname echo false head ls mkdir pwd rm seq sleep sort tee touch true uniq wc whoami yes"

if [[ $# > 0 ]]; then
  COMMAND=$1
else
  echo "usage: $0 [build|test|fmts| vet]"
  exit 1
fi

go ${COMMAND} ${LIBS} ${APPS}

