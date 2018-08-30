# image with development tools
FROM xena/go:1.10 AS build
ENV GOPATH /root/go
RUN apk --no-cache add git protobuf
COPY . /root/go/src/github.com/Xe/printerfacts
WORKDIR /root/go/src/github.com/Xe/printerfacts
ENV CGO_ENABLED 0
RUN ./scripts/generate.sh && ./scripts/build.sh

# runner image
FROM xena/alpine
COPY --from=build /root/go/src/github.com/Xe/printerfacts/bin/printerfacts /printerfacts
CMD /printerfacts
