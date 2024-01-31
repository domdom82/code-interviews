# Chat-P2P
- Write a simple TCP chat program.
- The program can run as client or server.
- The client expects an address to connect to.
- The server expects an address to listen on.
- Once connected, either party can send chat messages.
- A chat message is a string terminated by '\n'.
- When a chat message is received, it it displayed.

## Possible Challenges:

- timeouts (kick after idle)
- message size (buffers)
- multiple clients

## How to use

**Run as server**
```
./chat-p2p -s localhost:1234
```


**Run as client**
```
./chat-p2p localhost:1234
```
