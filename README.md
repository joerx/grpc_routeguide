# GRPC Tutorial

Based on [Golang gRPC Tutorial][1]

## Usage

Run `make help`

## Kubernetes

```sh
make docker-publish
helm upgrade --install helm/ --image.repository=<...> --image.tag=<...> routeguide
```

[1]: https://grpc.io/docs/tutorials/basic/go.html