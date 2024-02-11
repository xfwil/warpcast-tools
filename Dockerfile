FROM scratch
COPY warpcast-tools /usr/bin/warpcast-tools
ENTRYPOINT ["/usr/bin/warpcast-tools"]