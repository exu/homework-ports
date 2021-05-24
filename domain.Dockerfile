FROM golang:alpine as builder
WORKDIR /app 
COPY . .
RUN cd cmd/domain;CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" .

FROM scratch
WORKDIR /app
COPY --from=builder /app/cmd/domain/domain /usr/bin/
ENTRYPOINT ["domain"]
