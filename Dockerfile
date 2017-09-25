FROM alpine:latest

MAINTAINER Edward Muller <edward@heroku.com>

WORKDIR "/opt"

ADD .docker_build/cisco-spark-bot /opt/bin/cisco-spark-bot
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/cisco-spark-bot"]

