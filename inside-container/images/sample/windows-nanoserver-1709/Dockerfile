FROM golang:1.10.2-windowsservercore-1709 AS builder
COPY src C:/go/src/github.com/rajch/dk-samples/inside-container
WORKDIR C:/go/src/github.com/rajch/dk-samples/inside-container
ENV CGO_ENABLED="0" GOARCH="amd64" GOOS="windows"
RUN go build

FROM microsoft/nanoserver:1709 AS final
LABEL maintainer="Raj Chaudhuri <rajch@hotmail.com>"
COPY --from=builder C:/go/src/github.com/rajch/dk-samples/inside-container/inside-container.exe C:/inside-container.exe
ENV PORT_NUMBER=8080
ENTRYPOINT [ "C:/inside-container.exe" ]
