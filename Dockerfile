# Build stage
FROM golang:1.12.0-alpine as build

RUN apk add --no-cache git make
RUN apk --no-cache add ca-certificates

# Don't run as root
RUN adduser -D -u 2019 -g 2019 frequencify

ENV GO111MODULE=on

RUN mkdir /app 
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o frequencify

# Final stage
FROM scratch

COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /app/frequencify /app/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER frequencify
ENTRYPOINT ["/app/frequencify"]