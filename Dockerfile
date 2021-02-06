FROM golang:alpine as build
COPY . .
RUN cd cmd/vigenere-jorin && GOPATH=$(pwd) go build

FROM scratch
COPY --from=build /go/cmd/vigenere-jorin/vigenere-jorin /
ENTRYPOINT ["/vigenere-jorin"]
CMD ["/vigenere-jorin"]
