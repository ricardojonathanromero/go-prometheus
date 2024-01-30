FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk add --no-cache ca-certificates
COPY go.mod .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o health main.go


## deploy stage
FROM alpine as final

MAINTAINER "ricardo.jonathan.romero@gmail.com"
LABEL service="health"
LABEL owner="ricardojonathanromero"

# update ca-certificates
RUN apk --no-cache add ca-certificates wget
# Create a group and user
RUN addgroup -S nonroot && adduser -S nonroot -G nonroot

# defining workdir
WORKDIR /home/nonroot/app

# copy binary
COPY --from=builder /app/health .

# change owner
RUN chmod +x ./health && chown -R nonroot:nonroot ./health

# configure user
USER nonroot

ENTRYPOINT ["/home/nonroot/app/health"]
