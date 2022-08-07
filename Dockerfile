FROM ubuntu:20.04
RUN rm /bin/sh && ln -s /bin/bash /bin/sh
RUN mkdir /app
WORKDIR /app

RUN apt-get update && apt-get dist-upgrade -y
RUN apt-get install -y make gcc gcc-arm-linux-gnueabi
COPY --from=golang:1.17.13-bullseye /usr/local/go/ /usr/local/go/
RUN echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
#RUN go version