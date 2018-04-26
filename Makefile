.PHONY: test clean

test: test_parser

test_parser:
	@( go test ./ )

clean: 
	rm callgraph.dot
