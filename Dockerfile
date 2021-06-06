FROM alpine 

COPY hello $GOPATH/src/hello
WORKDIR $GOPATH/src

RUN chmod +x hello
