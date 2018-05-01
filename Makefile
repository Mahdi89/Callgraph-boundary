NAME := can-call
DESC := A clang AST analysis tool 

.PHONY: test clean build run

test: test_parser

test_parser:
	@( go test ./ )

build: 
	$(shell export PATH=$PATH:$PWD)
	cd ./bin/ && go build ../src/parser.go

run: build
	can_call ./test/nested.c main buzz

clean: 
	rm -rf callgraph.dot
	rm -rf ./bin/parser
