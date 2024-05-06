.PHONY: all test clean build

build:
	- rm build/dao-generator
	go build -o build/dao-generator .