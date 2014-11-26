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

To build the software you need to run the following command:

```bash
rake app:build
```

## License

Please see the [LICENSE.md](LICENSE.md).
