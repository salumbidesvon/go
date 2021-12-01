FROM golang:1.16.6-alpine as dev

WORKDIR /work

FROM dev as build

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=readonly -o /backend

FROM alpine as prod

COPY --from=build /backend /backend
RUN addgroup -S go && adduser -S go -G go
RUN chown -R go:go /backend

USER go

ENTRYPOINT ["/backend"]