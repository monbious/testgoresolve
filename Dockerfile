FROM ubuntu:16.04

ADD conf /conf
ADD testserver /

EXPOSE 8008

CMD ["./testserver"]
