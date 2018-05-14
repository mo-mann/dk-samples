FROM golang:1.8.3-alpine AS builder
COPY src /go/src/github.com/rajch/dk-samples/inside-container
WORKDIR /go/src/github.com/rajch/dk-samples/inside-container
ENV CGO_ENABLED="0" GOARCH="amd64" GOOS="linux"
RUN go build

FROM scratch AS final
LABEL maintainer="Raj Chaudhuri <rajch@hotmail.com>"
COPY --from=builder /go/src/github.com/rajch/dk-samples/inside-container/inside-container /inside-container
ENV PORT_NUMBER=8080
ENTRYPOINT [ "/inside-container" ]
