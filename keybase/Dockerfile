############################
# STEP 1 build executable binary
############################
# FROM: https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
# Build: docker build -f Dockerfile -t smallest-secured-golang .

# TODO: figure out if i really need timezone and/or certs


# golang alpine 1.12.6
FROM golang@sha256:cee6f4b901543e8e3f20da3a4f7caac6ea643fd5a46201c3c2387183a332d989 as builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git

# Create appuser
RUN adduser -D -g '' appuser

# TODO: rename!!!
WORKDIR $GOPATH/src/jlwarren1.com/keybase/
COPY . .

# Fetch dependencies.
RUN go get -d -v

# Build the binary
# TODO: rename!!!
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/keybase .

############################
# STEP 2 build a small image
############################
FROM scratch

# Import from builder.
# COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# COPY --from=builder /etc/passwd /etc/passwd

# TODO: copy template and static files?

ARG LISTEN_PORT
ENV LISTEN_PORT=${LISTEN_PORT}

# Copy our static executable
# TODO: rename!!!
COPY --from=builder /go/bin/keybase /go/bin/keybase

# Use an unprivileged user.
USER appuser

ARG LISTEN_PORT
ENV LISTEN_PORT=${LISTEN_PORT}

EXPOSE ${LISTEN_PORT}

# Run the keybase binary.
ENTRYPOINT ["/go/bin/keybase"]