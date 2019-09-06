FROM golang:1.12.1 as bd
WORKDIR /github.com/layer5io/meshery-octarine
ADD . .
RUN go build -ldflags="-w -s" -a -o /meshery-octarine .
RUN find . -name "*.go" -type f -delete; mv octarine /

FROM octarinesec/octactl-container:0.13.1 as oc

FROM alpine:latest
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN adduser -D appuser
ADD ./.octactl.yaml /home/appuser
COPY --from=oc /usr/local/bin/octactl /usr/local/bin/
COPY --from=bd /meshery-octarine /app/
RUN chown -R appuser:appuser /home/appuser
USER appuser
WORKDIR /app
CMD ./meshery-octarine
