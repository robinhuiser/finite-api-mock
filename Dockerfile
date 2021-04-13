FROM golang:1.16 AS build

WORKDIR /go/src
COPY server ./server
COPY ent ./ent
COPY util ./util
COPY generator ./generator
COPY api ./api
COPY swagger-ui ./swagger-ui

COPY main.go .
COPY go.mod .
COPY go.sum .
COPY banner.txt .

# ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -o fca-emu -a -ldflags "-linkmode external -extldflags '-static' -s -w" .

FROM scratch AS runtime

COPY --from=build /go/src/fca-emu ./
EXPOSE 8080/tcp
ENTRYPOINT ["/fca-emu"]
