#test
.PHONY: build
# build
build:
	go build -o ./bin/gin-temp

.PHONY: run
run:
	make build
	./bin/gin-temp

.PHONY: dev
dev:
	make build
	./bin/gin-temp -conf ./config.dev.yaml