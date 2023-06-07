#!/usr/bin/env bash

printf " -\033[35m setting up\033[0m hayashi root\n"
mkdir -p ~/.hayashi/repo
cd ~/.hayashi/repo

printf " -\033[35m cloning\033[0m hayashi from\033[33m https://github.com/crispybaccoon/hayashi\033[0m\n"
git clone --filter=blob:none https://github.com/crispybaccoon/hayashi && cd hayashi

printf " -\033[35m building\033[0m hayashi\n"
go build -o ./hayashi .

printf " -\033[35m setting up\033[0m environment\n"
./hayashi config init

printf " -\033[35m finishing\033[0m installation\n"
./hayashi task pack hayashi

