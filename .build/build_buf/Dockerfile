FROM buildpack-deps:bullseye-curl
ARG GO_VERSION
ARG NODE_VERSION

# install go
RUN curl -L https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz | tar xz -C /usr/local
ENV PATH $PATH:/usr/local/go/bin

# install node.js
RUN curl -L https://nodejs.org/dist/v${NODE_VERSION}/node-v${NODE_VERSION}-linux-x64.tar.gz | tar xz -C /usr/local --strip-components=1 && \
    unlink /usr/local/CHANGELOG.md && unlink /usr/local/LICENSE && unlink /usr/local/README.md

# # make user for build
# RUN mkdir /home/builduser && mkdir /project && \
#     groupadd builduser && \
#     useradd -g builduser -d /home/builduser builduser && \
#     chown builduser:builduser /home/builduser && \
#     chown builduser /project
# USER builduser

# install npm packages
RUN npm i -g @bufbuild/buf @bufbuild/protoc-gen-connect-web @bufbuild/protoc-gen-es @bufbuild/connect @bufbuild/connect-web @bufbuild/protobuf
ENV PATH $PATH:/usr/local/bin/

# install go modules
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
# ENV PATH $PATH:/home/builduser/go/bin/
ENV PATH $PATH:/root/go/bin/

ENV LANG C.UTF-8
ENV LANGUAGE C.UTF-8
ENV LC_ALL C.UTF-8
ENV DEBUG_COLORS true
ENV FORCE_COLOR true

WORKDIR /project
