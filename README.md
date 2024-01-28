# palog

Broadcast messages that someone joined/left the game.

![game screen that shows log of joined:Tsuzu](./docs/example.png)

## Usage
### Enable RCON of PalWorld server
* Open `~/Steam/steamapps/common/PalServer/Pal/Saved/Config/LinuxServer/PalWorldSettings.ini`
* Set `RCONEnabled="True"`
* Set `AdminPassword="your_random_password"`
* Restart the PalWorld server

* Ref: https://tech.palworldgame.com/optimize-game-balance

### Docker
```
docker run --name palog -e RCON_ENDPOINT={PalWorldServerIP}:25575 -e RCON_PASSWORD={AdminPassword you set above} ghcr.io/miscord-dev/palog:v0.0.5
```

### Binary installation
#### Install palog
* Use the following guide of installation
* (optional) Install `uconv`
    * Non-ascii characters are converted via `uconv -x latin`
    * On Debian/Ubuntu; `apt-get install -y icu-devtools`

#### Set environment variables
* RCON_ENDPOINT
    * If you run palog on the same server as the PalWorld server, the value is `127.0.0.1:25575`
* RCON_PASSWORD
    * The password you set in the earlier step
* For other options, please check the [Environemnt Variables](#environment-variables) section

#### Launch palog
```
$ ./palog
```

## Installation
### Docker(Recommended)
```
docker pull ghcr.io/miscord-dev/palog:v0.0.5 # or main
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
* Non-ascii are corrupted
    * https://github.com/miscord-dev/palog/issues/6
