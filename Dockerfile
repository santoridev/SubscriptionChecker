    FROM golang:1.24.2-alpine3.21

    # Install dependencies
    WORKDIR /testapp

    COPY . .

    RUN  go build -o app .
    ENTRYPOINT [ "./app" ]