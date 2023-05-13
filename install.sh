#!/usr/bin/env bash

echo -e " -\1xb[35m setting up\x1b[0m hayashi root"
mkdir -p ~/.hayashi/repo
cd ~/.hayashi/repo

echo -e " -\1xb[35m cloning\x1b[0m hayashi from\x1b[33m https://github.com/crispybaccoon/hayashi\x1b[0m"
git clone --filter=blob:none https://github.com/crispybaccoon/hayashi && cd hayashi

echo -e " -\1xb[35m building\x1b[0m hayashi"
go build -o ./hayashi .

echo -e " -\1xb[35m setting up\x1b[0m environment"
./hayashi config create
./hayashi config init

echo -e " -\1xb[35m finishing\x1b[0m installation"
./hayashi task pack hayashi

