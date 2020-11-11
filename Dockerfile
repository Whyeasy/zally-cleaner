FROM alpine

COPY zally-cleaner /usr/bin/
ENTRYPOINT ["/usr/bin/zally-cleaner"]