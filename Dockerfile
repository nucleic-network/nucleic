ARG GO_VERSION="1.22.1"
ARG RUNNER_IMAGE="gcr.io/distroless/static-debian11"
ARG BUILD_TAGS="netgo,ledger,muslc"

# --------------------------------------------------------
# Builder
# --------------------------------------------------------

FROM golang:${GO_VERSION}-alpine3.18 as builder

ARG GIT_VERSION
ARG GIT_COMMIT
ARG BUILD_TAGS

RUN apk add --no-cache \
    ca-certificates \
    build-base \
    linux-headers 

# Download go dependencies
WORKDIR /eve
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    go mod download 

# Grab the static library and copy it to location that will be found by the linker flag `-lwasmvm_muslc`.
RUN ARCH=$(uname -m) && WASMVM_VERSION=$(go list -m all | grep github.com/CosmWasm/wasmvm | awk '{print $2}' | tail -n 1) && \ 
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/libwasmvm_muslc.$ARCH.a \
    -O /lib/libwasmvm_muslc.$ARCH.a && \
    # verify checksum
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/checksums.txt -O /tmp/checksums.txt && \
    sha256sum /lib/libwasmvm_muslc.$ARCH.a | grep $(cat /tmp/checksums.txt | grep libwasmvm_muslc.$ARCH | cut -d ' ' -f 1)

# Copy the remaining files
COPY . .      

# Build eved binary
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    GOWORK=off go build \
    -mod=readonly \
    -tags "netgo,ledger,muslc" \
    -ldflags \
    "-X github.com/cosmos/cosmos-sdk/version.Name="eve" \
    -X github.com/cosmos/cosmos-sdk/version.AppName="eved" \
    -X github.com/cosmos/cosmos-sdk/version.Version=${GIT_VERSION} \
    -X github.com/cosmos/cosmos-sdk/version.Commit=${GIT_COMMIT} \
    -X github.com/cosmos/cosmos-sdk/version.BuildTags=${BUILD_TAGS} \
    -w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'" \
    -trimpath \
    -o /eve/build/eved \
    /eve/cmd/eved

# --------------------------------------------------------
# toolkit
# --------------------------------------------------------

FROM busybox:1.35.0-uclibc as busybox
RUN addgroup --gid 1025 -S eve && adduser --uid 1025 -S eve -G eve    

# --------------------------------------------------------
# Runner
# --------------------------------------------------------

FROM ${RUNNER_IMAGE}

COPY --from=busybox:1.35.0-uclibc /bin/sh /bin/sh

COPY --from=builder /eve/build/eved /bin/eved

# Install eve user
COPY --from=busybox /etc/passwd /etc/passwd
COPY --from=busybox --chown=1025:1025 /home/eve /home/eve

ENV HOME /eve
WORKDIR $HOME

# rest server
EXPOSE 1317
# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657
# grpc
EXPOSE 9090

ENTRYPOINT ["eved"]