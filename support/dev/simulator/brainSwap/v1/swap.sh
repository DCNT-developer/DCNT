#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd $DIR


function copy() { # and set AcksHeight

  if [[ $2 = 0 ]] ; then
    target="${DIR}/.factom/m2/dcnt.conf"
  else
    target="${DIR}/.factom/m2/simConfig/dcnt00${2}.conf"
  fi

  cat "../../simConfig/dcnt00${1}.conf" \
    | sed 's/ChangeAcksHeight = 0/ChangeAcksHeight = 10/' > $target
}

function main() {
  copy 1 0
  copy 4 1
}

main
