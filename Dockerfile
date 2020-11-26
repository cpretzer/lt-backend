FROM golang:1.15-alpine as build

WORKDIR /lt-build
COPY go.mod go.sum ./

COPY pkg pkg

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /out/lt-backend -mod=readonly -ldflags "-s -w" ./pkg

FROM scratch
ENV PATH=$PATH:/go/bin
COPY --from=build /out/lt-backend /go/bin/lt-backend
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/go/bin/lt-backend", "--logtostderr=true"]