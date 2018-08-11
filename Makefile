IMAGE ?= $(shell basename `pwd`)
TAG ?= latest
PORT ?= 10000

echo:
	@echo Image: $(IMAGE):$(TAG)

clean:
	rm -rf bin/

build: bin/routeguide

bin/routeguide:
	go build -o ./bin/routeguide

docker-build:
	docker build -t $(IMAGE):$(TAG) .

docker-run-server: docker-build
	docker run -it --rm -p$(PORT):$(PORT) $(IMAGE):$(TAG)
