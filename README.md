# GRY
Go Recon Yourself. A golang reconnaissance tool based on gopsutil.

## What it does

GRY will gather information about the current system and send them to a server over TCP Socket connection.

The transmission consists of two steps, first the client sends the size of the stringified JSON payload as a BigEndian packed 32bit unsigned int and then the actual payload.

## Try it out.

You can try GRY out using the following Python2 server that turns the data into a JSON object (thus verifying its integrity) and prints the JSON object.

```py
import json
import socket
import struct


def recv_helper(sock, n):
    data = b''
    while len(data) < n:
        packet = sock.recv(n - len(data))
        if not packet:
            return None
        data += packet
    return data


def recv(sock):
    raw_msglen = recv_helper(sock, 4)
    if not raw_msglen:
        return None
    msglen = struct.unpack('>I', raw_msglen)[0]
    data = recv_helper(sock, msglen)
    if not data:
        raise None
    try:
        data = data.decode('UTF-8')
    except UnicodeDecodeError:
        pass
    return data

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

try:
    s.bind(('localhost', 8899))
except socket.error:
    exit(-1)

s.listen(2)

try:
    while True:
        conn, addr = s.accept()
        data = recv(conn)
        print json.loads(data)
except KeyboardInterupt:
    s.close()
```

To execute the client (GRY) either:

```sh
go run gry.go
```

or:

```sh
go build gry.go
./gry
```

## Project Status

The project is in early development, in its most crude form. Interfaces will change and features will be added in the feature.
