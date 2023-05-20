## 1. Build Image
FROM golang:1.19 AS build-stage
WORKDIR /app
COPY src/go.mod src/go.sum ./
RUN go mod download
COPY src/*.go ./ 
RUN CGO_ENABLED=0 GOOS=linux go build -o /dynamitedb-gossiper

## 2. Test image
FROM build-stage as run-test-stage
RUN go test -v ./...

## Build minimal image without entire go toolchain
FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build-stage /dynamitedb-gossiper /dynamitedb-gossiper
EXPOSE 8070 
USER nonroot:nonroot
ENTRYPOINT [ "/dynamitedb-gossiper" ]
