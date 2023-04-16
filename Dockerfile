FROM alpine:3.17

COPY /kubebouncer /usr/local/bin/kubebouncer
RUN chmod +x /usr/local/bin/kubebouncer

ENTRYPOINT ["kubebouncer"]