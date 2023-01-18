
CC   = go build
MAIN = ./main.go
NAME = hayashi

BIN_WIN   = ./${NAME}.exe
BIN_LINUX = ./${NAME}

ifeq ($(OS),Windows_NT)
	## Windows
	BIN = ${BIN_WIN}
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
	## Linux
	BIN = ${BIN_LINUX}
endif
ifeq ($(UNAME_S),Darwin)
	## Macos
	BIN = ${BIN_LINUX}
endif

all: ${BIN}

${BIN_LINUX}: ${MAIN}
	env GOOS=linux $(CC) -o ${BIN_LINUX} .

${BIN_WIN}: ${MAIN}
	env GOOS=windows $(CC) -o ${BIN_WIN} .

clean:
	rm ${BIN_LINUX} -r
	rm ${BIN_WIN} -r

help_test: ${BIN}
	@echo "\033[01;35m$$\033[0m ${NAME} --help"
	@${BIN} --help

test: ${BIN}
	@echo "\033[01;35m$$\033[0m ${NAME} show some_pkg"
	@${BIN} show some_pkg
	@echo "\033[01;35m$$\033[0m ${NAME} add some_pkg"
	@${BIN} --force add some_pkg
	@echo "\033[01;35m$$\033[0m ${NAME} update some_pkg"
	@${BIN} update some_pkg

pkg_test: ${BIN}
	@echo "\033[01;35m$$\033[0m ${NAME} pkg add pkg_name"
	@${BIN} --force pkg add pkg_name
	@echo "\033[01;35m$$\033[0m ${NAME} show pkg_name"
	@${BIN} show pkg_name
	@echo "\033[01;35m$$\033[0m ${NAME} pkg remove pkg_name"
	@${BIN} pkg remove pkg_name

install: ${BIN}
	go install .
stats:
	wc --lines **/*.go

.PHONY: all test clean install stats
