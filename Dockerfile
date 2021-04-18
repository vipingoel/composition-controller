FROM golang

ENV GO111MODULE=on

WORKDIR /

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

ENTRYPOINT ["/sample-controller"]