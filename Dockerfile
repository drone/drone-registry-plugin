FROM alpine:3.6 as alpine
RUN apk add -U --no-cache ca-certificates

FROM alpine:3.6
EXPOSE 3000

ENV DRONE_DEBUG=false
ENV DRONE_ADDRESS=:3000
ENV GODEBUG netdns=go

COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ADD drone-registry-plugin /bin/
ENTRYPOINT ["/bin/drone-registry-plugin"]