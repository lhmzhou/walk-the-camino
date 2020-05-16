FROM localhost:8444/latest
ENV SRC github.com/lhmzhou/walk-the-camino

# Install GCC before code is added to build cache
USER 0
RUN yum install -y gcc
USER 1001

RUN mkdir -p $GOPATH/{SRC}
COPY . $GOPATH/{SRC}
WORKDIR $GOPATH/{SRC}
USER 0
RUN chmod -R 750 $GOPATH/go/ $GOPATH/src/ && \
    chown -R 1001:0 $GOPATH/go/ $GOPATH/src/
USER 1001
ENTRYPOINT go run project.go
