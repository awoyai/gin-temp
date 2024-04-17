#test
.PHONY: build
# build
build:
	go build -o ./bin/gin-temp

.PHONY: run
run:
	make build
	./bin/gin-temp