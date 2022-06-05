FROM alpine

RUN sed -i 's/https/http/' /etc/apk/repositories
RUN apk add curl
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai