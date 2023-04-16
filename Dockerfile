FROM alpine:3.17

COPY /nsbouncer /usr/local/bin/nsbouncer
RUN chmod +x /usr/local/bin/nsbouncer

ENTRYPOINT ["nsbouncer"]