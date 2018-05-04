# FROM golang:1.7.4-alpine
# ENV SOURCES /go/src/github.com/LeeTrent/cloud-native-go/
# COPY . ${SOURCES}
# RUN cd ${SOURCES} && CGO_ENABLED=0 go install -a
# ENV PORT 8080
# EXPOSE 8080 
# ENTRYPOINT cloud-native-go

FROM alpine:3.5
LABEL MAINTAINER Lee Trent
COPY ./cloud-native-go /app/cloud-native-go
RUN chmod +x /app/cloud-native-go
ENV PORT 8080 
EXPOSE 8080
ENTRYPOINT /app/cloud-native-go