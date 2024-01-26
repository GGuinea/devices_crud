FROM golang:1.21.6 AS build
RUN apt-get install -y ca-certificates
WORKDIR /go/src/app
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOPROXY=direct
RUN go build -v -o app .

FROM scratch
COPY --from=build /go/src/app/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]

EXPOSE 8080
EXPOSE 80

