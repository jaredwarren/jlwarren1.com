FROM alpine:3.1

ENV HUGO_VERSION 0.35

# https://github.com/gohugoio/hugo/releases

# Install HUGO
RUN set -x && \
    apk add --update wget ca-certificates && \
    wget https://github.com/spf13/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-64bit.tar.gz && \
    tar xzf hugo_${HUGO_VERSION}_Linux-64bit.tar.gz && \
    rm -r hugo_${HUGO_VERSION}_Linux-64bit.tar.gz && \
    mv hugo /usr/bin/hugo && \
    apk del wget ca-certificates && \
    rm /var/cache/apk/*

ENTRYPOINT ["/usr/bin/hugo"]
