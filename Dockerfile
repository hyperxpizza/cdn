from golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

ENV grpc=true
ENV rest=false
ENV secure=false


WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main ./cmd/cdn/main.go

WORKDIR /dist

RUN cp /build/main .
RUN cp /build/config.json .

EXPOSE 8888
EXPOSE 8887

CMD ["/dist/main", "--config=./config.json", "--grpc=true", "--rest=true", "--secure=false"]