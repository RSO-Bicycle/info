FROM alpine
RUN mkdir -p /opt/app
WORKDIR /opt/app
ADD ./main /opt/app

CMD ["/opt/app/main"]