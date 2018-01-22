# image with development tools
FROM xena/go-mini:1.9.2
ENV GOPATH /root/go
RUN apk --no-cache add git protobuf \
 && go download
COPY . /root/go/src/github.com/Xe/printerfacts
WORKDIR /root/go/src/github.com/Xe/printerfacts
RUN go run ./cmd/mage/main.go -v build

# runner image
FROM xena/alpine
RUN apk --no-cache add route-mesh-runit
COPY --from=0 /root/go/src/github.com/Xe/printerfacts/bin/printerfacts /usr/local/bin/printerfacts
CMD /usr/local/bin/printerfacts
