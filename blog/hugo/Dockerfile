FROM alpine:3.1

ENV HUGO_VERSION=0.35 \
    HUGO_USER=hugo \
    HUGO_SITE=/srv/hugo

RUN set -x && \
    apk add --update wget ca-certificates && \
    wget https://github.com/spf13/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-64bit.tar.gz && \
    tar xzf hugo_${HUGO_VERSION}_Linux-64bit.tar.gz && \
    rm -r hugo_${HUGO_VERSION}_Linux-64bit.tar.gz && \
    mv hugo /usr/bin/hugo && \
    apk del wget ca-certificates && \
    addgroup ${HUGO_USER} && \
    adduser -G ${HUGO_USER} -D ${HUGO_USER} && \
    rm /var/cache/apk/*

# https://github.com/gohugoio/hugo/releases

# Install HUGO
# RUN set -x && \
#     apk add --update wget ca-certificates && \
#     wget https://github.com/spf13/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-64bit.tar.gz && \
#     tar xzf hugo_${HUGO_VERSION}_Linux-64bit.tar.gz && \
#     rm -r hugo_${HUGO_VERSION}_Linux-64bit.tar.gz && \
#     mv hugo /usr/bin/hugo && \
#     apk del wget ca-certificates && \
#     rm /var/cache/apk/*


# RUN apk --no-cache add \
#     curl \
#     git \
#     && curl -SL https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-64bit.tar.gz \
#     -o /tmp/hugo.tar.gz \
#     && tar -xzf /tmp/hugo.tar.gz -C /tmp \
#     && apk del curl \
#     && mv /tmp/hugo /usr/local/bin/ \
#     && mkdir -p ${HUGO_SITE} \
#     && addgroup ${HUGO_USER} \
#     && adduser -G ${HUGO_USER} -D ${HUGO_USER} \
#     && chown -R ${HUGO_USER}:${HUGO_USER} ${HUGO_SITE} \
#     && rm -rf /tmp/*



USER ${HUGO_USER}

WORKDIR ${HUGO_SITE}

# VOLUME ${HUGO_SITE}

EXPOSE 1313

ENTRYPOINT ["/usr/bin/hugo"]


# FROM alpine:3.6

# ENV HUGO_VERSION=0.35 \
#     HUGO_USER=hugo \
#     HUGO_SITE=/srv/hugo

# RUN apk --no-cache add \
#     curl \
#     git \
#     && curl -SL https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-64bit.tar.gz \
#     -o /tmp/hugo.tar.gz \
#     && tar -xzf /tmp/hugo.tar.gz -C /tmp \
#     && apk del curl \
#     && mv /tmp/hugo /usr/local/bin/ \
#     && mkdir -p ${HUGO_SITE} \
#     && addgroup ${HUGO_USER} \
#     && adduser -G ${HUGO_USER} -D ${HUGO_USER} \
#     && chown -R ${HUGO_USER}:${HUGO_USER} ${HUGO_SITE} \
#     && rm -rf /tmp/*

# USER ${HUGO_USER}

# WORKDIR ${HUGO_SITE}

# VOLUME ${HUGO_SITE}

# EXPOSE 1313

# CMD hugo server --bind 0.0.0.0 \
#     --navigateToChanged \
#     --templateMetrics \
#     --buildDrafts