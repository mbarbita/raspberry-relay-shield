# ubuntu-golang
FROM ubuntu:15.10
MAINTAINER mbm <mbarbita@gmail.com>
RUN apt-get update && apt-get -y upgrade && apt-get -y install wget && rm -rf /var/lib/apt/lists/*
RUN wget -O - -nv https://storage.googleapis.com/golang/go1.5.2.linux-amd64.tar.gz | tar zxf - -C /usr/local
RUN mkdir $HOME/gows
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=$HOME/gows
