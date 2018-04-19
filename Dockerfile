FROM golang:1.7.4-alpine

ENV GOPATH /Users/leetrent/Development/Software/go/workspace
ENV SOURCES /Users/leetrent/Development/Software/go/workspace/src/github.com/LeeTrent/cloud-native-go/

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install -a

ENTRYPOINT cloud-native-go
EXPOSE 8080