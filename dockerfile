FROM golang:1.16 as builder

LABEL maintainer="Top Chotipat <chotipat.p@log21ruby.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY . .
RUN go test ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api/main.go

FROM alpine:latest  

RUN apk --no-cache add tzdata
RUN apk --no-cache add ca-certificates

ENV TZ=Asia/Bangkok

WORKDIR /app/

COPY --from=builder /app/main .

CMD ["./main"] 