# palog

Broadcast messages that someone joined/left the game.

![game screen that shows log of joined:Tsuzu](./docs/example.png)

## Installation
### Docker(Recommended)
```
docker pull ghcr.io/miscord-dev/palog:v0.0.3
```

### Download binary
* Download binary from [releases](https://github.com/miscord-dev/palog/releases/latest)

## Environment Variables
* RCON_ENDPOINT: Endpoint of RCON (IP:Port)
* RCON_PASSWORD: Password of RCON
* INTERVAL: Interval to check current players (default: 5s)
* TIMEOUT: Timeout of RCON calls (default: 1s)
* UCONV_LATIN: Set `false` to disable escape string with `uconv -x latin`

## Known issues
* If the message is split via whitespaces, `Broadcast` sends only the first segment
* CKJ characters are corrupted
    * https://github.com/miscord-dev/palog/issues/6
