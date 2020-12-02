FROM golang:alpine AS base
WORKDIR $GOPATH/indiependente/aws-lambda-container
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app

FROM gcr.io/distroless/static
COPY --from=base /app /app
ENTRYPOINT [ "/app" ]
