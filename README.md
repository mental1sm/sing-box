# sing-box

The universal proxy platform.

> This custom build enhances the `naive` inbound protocol by adding **Active Probing Protection**.
>
> **New Features:**
> - `fallback_site`: Transparently redirects unauthorized access attempts, web crawlers, and active probing scans to a local decoy server (e.g., your Next.js site) instead of dropping the connection.

## Naiveproxy Structure

```json
{
  "type": "naive",
  "tag": "naive-in",
  "network": "tcp",
  ...
  // Listen Fields

  "users": [
    {
      "username": "sekai",
      "password": "password"
    }
  ],
  "quic_congestion_control": "",
  "disable_udp": false,
  // New structure
  "fallback_site": {
    "address": "localhost",
    "port": 3000,
    "force_https": false
  },
  "tls": {}
}

## Documentation

https://sing-box.sagernet.org

## License

```

## Build for Linux amd64
```bash
git clone https://github.com/mental1sm/sing-box.git
cd ./sing-box
```
```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
go build -v -trimpath \
  -ldflags "-s -w -buildid=" \
  -tags "with_quic,with_grpc,with_dhcp,with_wireguard,with_utls,with_clash_api,with_gvisor" \
  ./cmd/sing-box
```

### On Windows
Open powershell as admin, then ./build-linux.bat

## Copyright
Copyright (C) 2022 by nekohasekai <contact-sagernet@sekai.icu>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.

In addition, no derivative work may use the name or imply association
with this application without prior consent.
```