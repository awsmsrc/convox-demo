FROM gliderlabs/alpine:3.2

RUN apk-install docker git

RUN apk-install go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

RUN go get github.com/ddollar/init
RUN go get github.com/garyburd/redigo/redis

WORKDIR /go/src/github.com/usr/app
COPY . /go/src/github.com/usr/app
RUN ls -la
RUN go get .

ENV PORT 5000
ENTRYPOINT ["/go/bin/init"]
CMD ["app"]
