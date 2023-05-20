## 1. Build Image
FROM golang:1.19 AS build-stage
WORKDIR /app
COPY src/go.mod src/go.sum ./
RUN go mod download
COPY src/*.go ./ 
RUN CGO_ENABLED=0 GOOS=linux go build -o /dynamitedb-coordinator

## 2. Test image
FROM build-stage as run-test-stage
RUN go test -v ./...

## Build minimal image without entire go toolchain
FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build-stage /dynamitedb-coordinator /dynamitedb-coordinator
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT [ "/dynamitedb-coordinator" ]
