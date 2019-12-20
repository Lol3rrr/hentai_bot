FROM golang:1.13 AS build

ENV buildType linux

WORKDIR /go/src/hentai_bot
COPY . .

RUN go get -d -v ./...
RUN make build_$buildType

FROM alpine:latest

WORKDIR /app

COPY --from=build /go/src/hentai_bot/app .

CMD ["./app"]
