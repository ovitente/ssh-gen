FROM golang:alpine AS builder

ARG CI_PROJECT_NAME
ARG CI_PROJECT_NAMESPACE

ADD . $GOPATH/src/gitlab.com/${CI_PROJECT_NAMESPACE}/${CI_PROJECT_NAME}
WORKDIR $GOPATH/src/gitlab.com/${CI_PROJECT_NAMESPACE}/${CI_PROJECT_NAME}
RUN go build -o /go/bin/app .

FROM alpine
COPY --from=builder /go/bin/app /app/sshmount
WORKDIR /sshmount

