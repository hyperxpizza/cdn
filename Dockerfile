FROM golang:1.17-alpine AS build

RUN apk update && apk upgrade && apk add --no-cache git
WORKDIR /tmp/app
COPY go.mod .
COPY go.sum .
COPY . .
RUN GOOS=linux go build -o ./out/api .

FROM alpine:latest
RUN apk add ca-certificates
COPY --from=build /tmp/app/out/api /app/api
WORKDIR "/app"
EXPOSE 8888
CMD ["./api"]