############################
# STEP 0 build a image
############################
FROM golang:1.12-alpine as builder

RUN apk update && apk add --no-cache git make

ENV GO111MODULE=on

WORKDIR /go/src
COPY . .
# Fetch dependencies.
RUN make install
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/juno


############################
# STEP 1 build a small image
############################

FROM alpine
# Copy our static executable.
WORKDIR /juno
COPY --from=builder /go/bin/juno .
# Run the contract binary.
CMD ["./juno"]