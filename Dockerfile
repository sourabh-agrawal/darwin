FROM golang:1.19
WORKDIR /darwin
COPY . .
RUN make go-build
EXPOSE 8080
ENTRYPOINT ["./bin/darwin"]
