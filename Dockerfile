# prebuild
FROM golang:1.17-alpine3.14 AS builder

WORKDIR /go/src/app

# manage dependencies
COPY go.mod go.sum ./
RUN go mod download 
# Build
COPY . .
RUN cd cmd/server && CGO_ENABLED=0 go build -o surl ./... 

# create image
FROM alpine:3.14

COPY --from=builder /go/src/app/cmd/server/surl /usr/local/bin/

CMD ["surl"]
