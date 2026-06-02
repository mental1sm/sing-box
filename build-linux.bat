@echo off
echo Building sing-box for Linux (amd64)...

set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0

go build -v -trimpath -ldflags "-s -w -buildid=" -tags "with_quic,with_grpc,with_dhcp,with_wireguard,with_utls,with_clash_api,with_gvisor" ./cmd/sing-box

echo.
echo Build complete!
pause