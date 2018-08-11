IMAGE ?= $(shell basename `pwd`)
TAG ?= latest
PORT ?= 10000

help: ## List targets & descriptions
	@cat Makefile* | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: bin/routeguide ## Build binary

bin/routeguide:
	go build -o ./bin/routeguide

clean: ## Remove build output
	rm -rf bin/

docker-build: ## Build Docker image
	docker build -t $(IMAGE):$(TAG) .

docker-run-server: docker-build ## Build Docker image and run in server mode with defaults
	docker run -it --rm -p$(PORT):$(PORT) $(IMAGE):$(TAG)
