# image with development tools
FROM xena/go:1.10
ENV GOPATH /root/go
RUN apk --no-cache add git protobuf retool
COPY . /root/go/src/github.com/Xe/printerfacts
WORKDIR /root/go/src/github.com/Xe/printerfacts
RUN retool build \
 && retool do mage -v generate build

# runner image
FROM xena/alpine
RUN apk --no-cache add route-mesh-runit
COPY --from=0 /root/go/src/github.com/Xe/printerfacts/bin/printerfacts /usr/local/bin/printerfacts
CMD /usr/local/bin/printerfacts
