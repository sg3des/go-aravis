FROM golang:1.15-buster AS build
WORKDIR /src

# build essential, Gig-EV dependencies from manual, plus some util packages (vim, less)
RUN apt-get update && apt-get install -y \
    build-essential sudo \
    vim-tiny less nano git ninja-build python3-pip cmake libglib2.0-dev \
    libxml2-dev libusb-1.0-0-dev 

RUN python3 -m pip install meson

RUN git clone https://github.com/AravisProject/aravis.git /opt/aravis && \
    cd /opt/aravis && \
    git checkout main && \
    ls -a && \
    meson -Dviewer=disabled -Dintrospection=disabled -Dgst-plugin=disabled -Ddocumentation=disabled build && \
    cd build && \
    ninja && \
    ninja install && \
    ldconfig

RUN arv-tool-0.8

COPY . /src/github.com/hybridgroup/go-aravis
WORKDIR /src/github.com/hybridgroup/go-aravis

RUN go build -mod=vendor -o /src/listdevices ./examples/list_devices.go

ENTRYPOINT ["/src/listdevices"]
