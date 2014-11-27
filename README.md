# local webserver

This repository provides a local webserver which serves the local directory in
which it was stored. If the server ist stored at
`/home/user/tmp/directory.d/server` it will serve all files in
`/home/user/tmp/directory.d/`.

## Usage

```bash
# Start server with random port (1023 < port < 65535) and localhost as interface and be verbose
server

# Start server with port and interface defined
server --port 1234 --interface 127.0.0.2

# Do not output anything
server --silent
```

## Example

```bash
% ./server
Open browser with http://127.0.0.1:7574/index.html
Server listens on 127.0.0.1:7574

Requests:
127.0.0.1 - - [26/Nov/2014:14:34:10 +0100] "GET / HTTP/1.1" 200 41 "" "Mozilla/5.0 (X11; Linux x86_64; rv:33.0) Gecko/20100101 Firefox/33.0"
```

## Development

See [this page](http://docs.drone.io/golang.html) for a good introduction for a
`Go`-installation which can compile `Go`-binaries for multiple operating
systems - tested on Archlinux (64bit) only.

### Configure shell

This is optional and highly opionated. Just make sure the go path is set
correctly.

**.bashrc** / **.zshrc**

```bash
export GOPATH=~/.local/share/go
export PATH=~/.local/share/go/bin:$PATH
export PATH=~/.local/share/golang/bin:$PATH
```

### Clone repository

```bash
hg clone -u release https://code.google.com/p/go <golang_path>
hg update default
```

### Build go

```bash
cd <golang_path>/src
GOOS=windows GOARCH=amd64 ./make.bash --no-clean 2> /dev/null 1> /dev/null
GOOS=darwin  GOARCH=amd64 ./make.bash --no-clean 2> /dev/null 1> /dev/null
GOOS=linux  GOARCH=amd64 ./make.bash --no-clean 2> /dev/null 1> /dev/null
```

### Bootstrap application

```bash
./bootstrap.sh
```

### Build application

To build the software you need to run the following command. This will create
three files.

```bash
./build.sh
```

## License

Please see the [LICENSE.md](LICENSE.md).
