# wakeo
Tinygo, rpi pico and some LEDs to let the kids know when it's ok to wake up.

# Usage

````sh
# flash examples in cmd dir
$ tinygo flash -target=pico <pkg_path>
# flash wakeo
$ tinygo flash -target=pico -ldflags="-X main.buildUnix=$(date +%s)" ./cmd/wakeo
````

Run unit tests

````sh
$ go test ./pkg/**
````

# Todos
- [x] pkg/clock
- [x] cmd/wakeo
- [ ] external LED green on day and red otherwise
- [ ] maybe add the pico OLED for debugging 
