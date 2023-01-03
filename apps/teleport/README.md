# teleport

reverse proxying websocket

```
[aaron@i7-6700k sse]$ go run ./cmd/demo/
2022/09/27 02:01:09 main.go:69: listening on http://127.0.0.1:8080


$ go run ./cmd/ufo/ rp https://ufo.k0s.io/rp/rp1/rp2 http://127.0.0.1:8080
2022/09/27 02:01:44 ufo.go:24: 🛸 listening on https://rp.ufo.k0s.io
212.2.242.67 - - [27/Sep/2022:02:01:50 +0800] "GET / HTTP/1.1" 200 666
212.2.242.67 - - [27/Sep/2022:02:01:51 +0800] "GET /favicon.ico HTTP/1.1" 200 1
212.2.242.67 - - [27/Sep/2022:02:01:55 +0800] "GET / HTTP/1.1" 200 666
212.2.242.67 - - [27/Sep/2022:02:02:04 +0800] "GET /ha HTTP/1.1" 200 19
212.2.242.67 - - [27/Sep/2022:02:02:08 +0800] "GET /once HTTP/1.1" 200 19
212.2.242.67 - - [27/Sep/2022:02:02:18 +0800] "GET / HTTP/1.1" 200 0

$ websocat wss://rp.ufo.k0s.io/
2022-09-27 02:02:19
2022-09-27 02:02:20
2022-09-27 02:02:21
2022-09-27 02:02:22
2022-09-27 02:02:23
2022-09-27 02:02:24
```

reverse proxying google

```
$ ufo teleport https://ufo.k0s.io/google https://google.com
INFO[0000] gos module inited
INFO[0000] webteleport module inited
2023/01/03 16:45:34.381 INFO    2023/01/04 00:45:34 serve.go:35: 🛸 listening on https://google.ufo.k0s.io
2023/01/03 16:45:34.381 INFO    2023/01/04 00:45:34 serve.go:40: 🔓 publicly accessible without a password
```
