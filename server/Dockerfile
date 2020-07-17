########################################
## Build Stage
########################################
FROM golang:1.14.3 as builder

ENV GO111MODULE="on"

# Copy the code from the host and compile it
WORKDIR /go/src/github.com/redestro/zeronet/server

# install dependencies
ADD ./go.sum ./go.sum
ADD ./go.mod ./go.mod
RUN go mod download

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -o zero main.go

########################################
## Production Stage
########################################
FROM scratch

# Set timezone
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Kolkata

# set working directory
WORKDIR /root

# add required files from host
COPY --from=builder /go/src/github.com/redestro/zeronet/server/zero ./zero

EXPOSE 8000

ENTRYPOINT ["./zero"]