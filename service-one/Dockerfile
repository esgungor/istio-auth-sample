### Build
FROM golang:1.17-buster AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /user-service

### Deploy
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /user-service /user-service
EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/user-service"]