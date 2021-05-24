FROM golang:alpine as builder
WORKDIR /app 
COPY . .
RUN cd cmd/data;CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" .

FROM alpine
WORKDIR /app
COPY ports.json /app/
COPY --from=builder /app/cmd/data/data /usr/bin/
ENTRYPOINT ["data"]
