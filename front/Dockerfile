FROM golang:1.15.2 AS builder

WORKDIR /go/src/github.com/erkanerol/service-examples-for-blog/front
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o front main.go


FROM scratch
COPY --from=builder /go/src/github.com/erkanerol/service-examples-for-blog/front/front /front
CMD ["/front"]
