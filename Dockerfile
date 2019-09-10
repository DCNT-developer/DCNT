FROM golang:1.12

# Get git
RUN apt-get update \
    && apt-get -y install curl git \
    && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Get glide
RUN go get github.com/Masterminds/glide

# Where dcnt sources will live
WORKDIR $GOPATH/src/github.com/DCNT-developer/dcnt

# Get the dependencies
COPY glide.yaml glide.lock ./

# Install dependencies
RUN glide install -v

# Get goveralls for testing/coverage
RUN go get github.com/mattn/goveralls

# Populate the rest of the source
COPY . .

ARG GOOS=linux

# Build and install dcnt
RUN go install -ldflags "-X github.com/DCNT-developer/dcnt/engine.Build=`git rev-parse HEAD` -X github.com/DCNT-developer/dcnt/engine.dcntVersion=`cat VERSION`"

# Setup the cache directory
RUN mkdir -p /root/.factom/m2
COPY dcnt.conf /root/.factom/m2/dcnt.conf

ENTRYPOINT ["/go/bin/dcnt","-sim_stdin=false"]

EXPOSE 8088 8090 8108 8109 8110
