FROM golang

COPY bin/ go/bin/

WORKDIR go/src/github.com/gofunct/stencil

COPY . .
ENV GO111MODULE=on
RUN make init
ENTRYPOINT [ "stencil" ]

