FROM alpine:3.21
RUN addgroup -S nexus && adduser -S nexus -G nexus
# Remove package manager to prevent runtime exploits
RUN rm -rf /sbin/apk /usr/share/apk
WORKDIR /app
USER nexus
ENTRYPOINT ["/usr/local/bin/api"]