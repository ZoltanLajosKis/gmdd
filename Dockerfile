FROM golang:1.8-stretch AS build-stage
WORKDIR /go/src/github.com/ZoltanLajosKis/gmdd
COPY / ./
RUN apt-get update \
 && apt-get -y upgrade \
 && apt-get -y --no-install-recommends install xz-utils
RUN make \
 && strip gmdd \
 && curl -fsLO https://github.com/upx/upx/releases/download/v3.94/upx-3.94-amd64_linux.tar.xz \
 && tar xf upx-3.94-amd64_linux.tar.xz \
 && upx-3.94-amd64_linux/upx --ultra-brute gmdd

FROM scratch
WORKDIR /
COPY --from=build-stage /go/src/github.com/ZoltanLajosKis/gmdd/gmdd .
EXPOSE 8000
ENTRYPOINT ["/gmdd", "/www"]
