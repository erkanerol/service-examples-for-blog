FROM golang:1.15.2 AS builder

WORKDIR /go/src/github.com/erkanerol/service-examples-for-blog/middle
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o middle main.go


FROM scratch
COPY --from=builder /go/src/github.com/erkanerol/service-examples-for-blog/middle/middle /middle
CMD ["/middle"]
