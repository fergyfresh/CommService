FROM golang:1.12.4-stretch

ENV GOOS=linux
ENV GOARCH=amd64
ENV PATH=/home/jenkins/firefox:/usr/lib64/qt-3.3/bin:/bin:/usr/local/bin:/usr/bin:/usr/bin:/usr/local/go/bin:${GOBIN}


RUN mkdir /opt/code && \
mkdir /opt/code/src && \
mkdir /opt/code/bin && \
mkdir /opt/commservice/

#Add the code to the container
ADD . /opt/code

WORKDIR /opt/code/

RUN go mod init

WORKDIR /opt/code/cmd/main/

RUN go build -o commserv ./... && \
ls

RUN mv commserv /opt/commservice/
WORKDIR /opt/commservice/
RUN ls -la

EXPOSE 8080

ENTRYPOINT ["./commserv"]