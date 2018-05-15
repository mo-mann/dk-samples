FROM alpine:3.5
LABEL maintainer="Raj Chaudhuri <rajch@hotmail.com>"
ENV SERVICE_URL="http://www.google.com"
COPY linux/servicetester.sh /app/servicetester.sh
RUN chmod +x /app/servicetester.sh
ENTRYPOINT [ "/app/servicetester.sh" ]
