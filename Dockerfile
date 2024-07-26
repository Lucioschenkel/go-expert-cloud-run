FROM golang:1.22.2-alpine3.19 as build

WORKDIR /app

COPY go.* .

RUN go mod download && go mod verify

COPY . .

RUN go build -o /bin/server ./cmd/server/main.go

FROM scratch

WORKDIR /app

COPY --from=build /bin/server .

EXPOSE 8080

ENTRYPOINT [ "./server" ]