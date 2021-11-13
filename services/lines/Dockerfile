FROM golang:1.17-alpine as builder

LABEL maintainer="Rene Schertel <reneschmitt@gmx.de>"

RUN mkdir -p /app

WORKDIR /app
COPY . /app
RUN go get
RUN go build -o /app/bin/app
RUN ls -la /app/bin

# gcr.io/distroless/base
# golang:1.17-alpine
FROM golang:1.17-alpine
COPY --from=builder /app/bin/app /
ENV DB_HOST=test
ENV DB_PORT=5432
ENV DB_NAME=systems
ENV DB_USER=systems
ENV DB_PASSWORD=''
EXPOSE 80
ENTRYPOINT ["/app"]