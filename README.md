## Raylib PONG clone ported to Go with [raylib-go](https://github.com/gen2brain/raylib-go), with in-progress P2P LAN UDP network multiplayer.

#### PONG game based on [this](https://www.youtube.com/watch?v=VLJlTaFvHo4) and [this](https://www.youtube.com/watch?v=LvpS3ILwQNA) video tutorial.

Special attention has been paid to make the code as readable as possible through separation of concerns and avoidance of code reuse.

Network code written by me by reading the Golang documentation and using my google-fu.

### How to play

Move the left paddle with the `up` and `down` keyboard arrows.

### Build instructions

Build on Windows

`go build -o "go_pong.exe" -ldflags "-H=windowsgui" .`

Build on Arch Linux

1. Install the `raylib` package.
2. Run `go build -o go_pong`

### How to run multiplayer 

`./go_pong -lan=host/join`

### Debugging stuff

`netcat -u -v 192.168.0.101 34788`

`go tool pprof -http 192.168.0.102:8080 profile.prof`

### Issues

In singleplayer, it is possible to control the CPU paddle together with the player paddle by calling `p2.paddle.update()` instead of `p2.update(&ball)` in the `main.go` `update()` function.

There seems to be no good way in Go to override methods OOP style. See [Golang Method Override](https://stackoverflow.com/questions/38123911/golang-method-override) on Stack Overflow.