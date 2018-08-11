FROM golang:1.10-stretch AS builder
RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/src/github.com/joerx/grpc_routeguide
COPY Gopkg.* ./
RUN dep ensure -vendor-only -v
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/routeguide .

FROM debian:stretch
ENV PORT=10000
RUN apt-get update && apt-get -y --force-yes install ca-certificates
WORKDIR /bin
COPY --from=builder /bin/routeguide .
COPY ./testdata /data
ENTRYPOINT ["/bin/routeguide", "-db=/data/route_guide_db.json"]
CMD ["server"]
