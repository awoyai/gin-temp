FROM golang:1.21-alpine

WORKDIR /app
COPY ./ /app/server
ARG goproxy
ENV GOPROXY=$goproxy
RUN cd /app/server && go build -o /app/server/server

FROM mirrors.tencent.com/tlinux/tlinux2.6-minimal:latest

WORKDIR /app

COPY --from=0 /app/server/server /usr/local/server/bin/SpiderManagerServer
COPY --from=0 /app/server/config.yaml /usr/local/server/conf/config.yaml
RUN mkdir -p /usr/local/server/log/ && chmod -R 777 /usr/local/server/log/ 

CMD ["/usr/local/server/bin/GinTempServer","-conf","/usr/local/server/conf/config.yaml"]