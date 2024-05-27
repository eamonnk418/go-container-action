FROM golang:1.22.3-alpine3.20 AS builder

WORKDIR /app

COPY go.mod *go.sum ./
RUN go mod download && go mod verify

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w -extldflags '-static'" -o go-container-action .

FROM scratch

COPY --from=builder /app/go-container-action ./go-container-action

ENTRYPOINT [ "/go-container-action" ]