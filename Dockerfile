FROM golang:onbuild
ENV PORT 5000
VOLUME ["/go/src/app"]
EXPOSE 5000
