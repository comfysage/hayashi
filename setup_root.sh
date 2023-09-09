#!/usr/bin/env bash

chr=/mnt/root
mkdir -p $chr
mkdir -p $chr/{bin,lib,lib64}
cd $chr

items=(
  /bin/bash
  /bin/touch
  /bin/ls
  /bin/rm
)

for item in $items; do
  cp -v "$item" "$chr/$(dirname $item)"
  list=$(ldd $item | egrep -o '/lib.*\.[0-9]')
  for i in $list; do
    cp -v --parents "$i" "${chr}"
  done
done

