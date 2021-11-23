
FROM docker.io/library/golang:1.17.3-alpine AS builder

ARG GO111MODULES=on

COPY . /opt

WORKDIR /opt

# Time to build!
RUN go build -o gocker cmd/main.go

################# Final image! #########################

FROM docker.io/library/alpine:latest

# Copy the final binary from the first stage.
COPY --from=builder /opt/gocker /usr/local/bin/

USER 65534:65534

EXPOSE 8080/tcp

ENTRYPOINT ["/usr/local/bin/gocker"]