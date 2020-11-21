FROM golang:1.15-buster AS build
WORKDIR /src

# build essential, Gig-EV dependencies from manual, plus some util packages (vim, less)
RUN apt-get update && apt-get install -y \
    build-essential sudo \
    vim-tiny less nano git ninja-build python3-pip cmake libglib2.0-dev \
    libxml2-dev libusb-1.0-0-dev 

# libgirepository1.0-dev gobject-introspection libgtk-3-dev \
#    libgstreamer1.0-0 gstreamer1.0-plugins-base gstreamer1.0-plugins-good \ 
#    gstreamer1.0-plugins-bad gstreamer1.0-plugins-ugly \
#    gstreamer1.0-libav gstreamer1.0-doc gstreamer1.0-tools \
#    libgstreamer1.0-dev libgstreamer-plugins-base1.0-dev \
#    libfontconfig1-dev libfreetype6-dev libpng-dev \
#    libcairo2-dev libjpeg-dev libgif-dev \
#    libgstreamer-plugins-base1.0-dev gstreamer1.0-x libnotify-dev gtk-doc-tools

RUN python3 -m pip install meson

RUN git clone https://github.com/AravisProject/aravis.git /opt/aravis && \
    cd /opt/aravis && \
    git checkout master && \
    ls -a && \
#    meson build && \
    meson -Dviewer=disabled -Dintrospection=disabled -Dgst-plugin=disabled -Ddocumentation=disabled build && \
    cd build && \
    ninja && \
    ninja install && \
    ldconfig

RUN arv-tool-0.8

# RUN go get -d github.com/hybridgroup/go-aravis && \
#    cd $GOPATH/src/github.com/hybridgroup/go-aravis && \
#    pwd && \
#    git fetch origin && \
#    ls -l && \
#    git checkout update-aravis-8 && \
#    go run ./examples/list_devices.go
#FROM aravis-base AS vision

COPY . /src/github.com/hybridgroup/go-aravis
WORKDIR /src/github.com/hybridgroup/go-aravis

RUN go build -mod=vendor -o /src/listdevices ./examples/list_devices.go

ENTRYPOINT ["/src/listdevices"]
