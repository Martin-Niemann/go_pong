## Raylib PONG clone ported to Go with [raylib-go](https://github.com/gen2brain/raylib-go). 

#### Based on [this](https://www.youtube.com/watch?v=VLJlTaFvHo4) and [this](https://www.youtube.com/watch?v=LvpS3ILwQNA) video tutorial.

Special attention has been paid to make the code as readable as possible through separation of concerns and avoidance of code reuse.

### How to play
Move the left paddle with the `up` and `down` keyboard arrows.

### Build instructions

Build on Windows 

`go build -o "go_pong.exe" -ldflags "-H=windowsgui" .`

Build on Arch Linux
1. Install the [`raylib`](https://archlinux.org/packages/extra/x86_64/raylib/) package.
2. Run `go build -o go_pong`

### Issues
It is possible to control the CPU paddle together with the player paddle by calling `p2.paddle.update()` instead of `p2.update(&ball)` in the `main.go` `update()` function. 

There seems to be no good way in Go to override methods OOP style. See [Golang Method Override](https://stackoverflow.com/questions/38123911/golang-method-override) on Stack Overflow.
