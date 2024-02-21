FROM golang:1.21-alpine AS build-stage

WORKDIR /app
RUN cd /app

ARG PORT=8082
ENV PORT=${PORT}
ADD . /app/

#RUN go install github.com/codegangsta/gin@latest && \
#    go install github.com/onsi/ginkgo/ginkgo@latest && \
#    go install github.com/go-delve/delve/cmd/dlv@latest && \
#    go install github.com/onsi/gomega/...@latest \

COPY go.mod go.sum ./

RUN go mod download

RUN go build -o product-service

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM alpine

WORKDIR /app
COPY --from=build-stage /app/product-service /app
COPY .env /app/.env

EXPOSE $PORT

CMD ["/app/product-service"]