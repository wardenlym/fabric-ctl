.PHONY: all run-test-network down-test-network

all:
	go build

run-test-network:
	cd test-env/fabric-samples/test-network && ./network.sh up -ca

down-test-network:
	cd test-env/fabric-samples/test-network && ./network.sh down