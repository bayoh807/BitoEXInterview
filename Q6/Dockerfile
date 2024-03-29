FROM golang:alpine3.16 AS builder

RUN apk add make; \
    apk add openssh; \
    apk add openssh-client; \
    apk add openrc; \
    apk add zip; \
    apk add tzdata;
RUN zip -r -0 /zoneinfo.zip /usr/share/zoneinfo
RUN apk add --no-cache bash

ENV WORK_HOME="app"
ENV CMD_NAME="tinder-server"
WORKDIR /$WORK_HOME
COPY . /$WORK_HOME

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOGC=50

RUN cd /$WORK_HOME && \
    make init build

WORKDIR /dist

RUN cp /$WORK_HOME/build/main main



# # Build a small image ===================================
FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /dist /

ENV ZONEINFO /zoneinfo.zip
COPY --from=builder /zoneinfo.zip /

ARG DB_USERNAME
ARG DB_PASSWORD
ARG DB_HOST
ARG DB_NAME
ARG ENV
ARG COMMIT_ID
ARG BUILD_VERSION

ENV DB_USERNAME="${DB_USERNAME}"
ENV DB_PASSWORD="${DB_PASSWORD}"
ENV DB_HOST="${DB_HOST}"
ENV DB_NAME="${DB_NAME}"
ENV ENV="${ENV}"
ENV VERSION="${BUILD_VERSION}"
ENV COMMIT_ID="${COMMIT_ID}"


ENTRYPOINT ["/main"]