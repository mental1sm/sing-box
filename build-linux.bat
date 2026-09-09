@echo off
echo Building sing-box for Linux (amd64)...

set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0

go build -v -trimpath -ldflags "-s -w -buildid= -checklinkname=0" -tags "with_quic,with_grpc,with_dhcp,with_utls,with_gvisor,with_naive_outbound,with_purego,badlinkname,tfogo_checklinkname0" ./cmd/sing-box

echo.
echo Build complete!
pause