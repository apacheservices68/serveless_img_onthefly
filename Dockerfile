FROM ubuntu:18.04

LABEL mantainer="Dam Viet <vietdien2005@gmail.com>"

ENV LIBVIPS_VERSION 8.7.0
# Go version to use
ENV GOLANG_VERSION 1.11.2
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN \
    # Install dependencies
    apt-get update && apt-get install -y \
    ca-certificates automake build-essential curl \
    gobject-introspection gtk-doc-tools libglib2.0-dev \
    libjpeg-turbo8-dev libpng-dev libwebp-dev libtiff5-dev \
    libgif-dev libexif-dev libxml2-dev libpoppler-glib-dev \
    swig libmagickwand-dev libpango1.0-dev libmatio-dev \
    libopenslide-dev libcfitsio-dev nasm libgsf-1-dev \
    fftw3-dev liborc-0.4-dev librsvg2-dev libglib2.0-0 \
    libjpeg-turbo8 libpng16-16 libopenexr22 libwebp6 \
    libtiff5 libgif7 libexif12 libxml2 libpoppler-glib8 \
    libmagickwand-6.q16-dev libpango1.0-0 libmatio4 libopenslide0 \
    libgsf-1-114 fftw3 liborc-0.4 librsvg2-2 libcfitsio5 \
    libmagickwand-6.q16-3 gcc git libc6-dev make cmake

RUN \
    # Build libvips
    cd /tmp && \
    curl -OL https://github.com/libvips/libvips/releases/download/v${LIBVIPS_VERSION}/vips-${LIBVIPS_VERSION}.tar.gz && \
    tar zvxf vips-${LIBVIPS_VERSION}.tar.gz && \
    cd /tmp/vips-${LIBVIPS_VERSION} && \
    ./configure --enable-debug=no --without-python $1 && \
    make && \
    make install && \
    ldconfig

# Clean up
RUN apt-get autoremove -y && \
    apt-get autoclean && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN curl -fsSL --insecure "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
    && tar -C /usr/local -xzf golang.tar.gz \
    && rm golang.tar.gz

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

# Fetch the latest version of the package
RUN go get golang.org/x/net/context
RUN go get github.com/golang/dep/cmd/dep

COPY . /go/src/tto-resize

RUN cd /go/src/tto-resize/lilliput/deps && \
    ./build-deps-linux.sh