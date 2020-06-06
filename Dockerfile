# First stage: build the executable.
FROM golang:1.13 as builder
WORKDIR /go/src/github.com/rekzi/elasticsearch-provisioner/
COPY . .
# CGO_ENABLED=0 to build a statically-linked executable
# GOFLAGS=-mod=vendor to force `go build` to look into the `/vendor` folder.
ENV CGO_ENABLED=0 GOFLAGS=-mod=vendor
RUN go build -installsuffix 'static' -o elasticsearch-provisioner .

# Final stage: the running container.
FROM alpine:3.9 AS final
RUN apk add --update --no-cache ca-certificates
WORKDIR /bin/
# Import the compiled executable from the first stage.
COPY --from=builder /go/src/github.com/rekzi/elasticsearch-provisioner/elasticsearch-provisioner .
EXPOSE 9000
ENTRYPOINT [ "/bin/elasticsearch-provisioner" ]
