# Echo-Server
Write a simple HTTP server that reads a request body of arbitrary size and echoes it back to the client.

## Possible Challenges:
- Arbitrary size (memory usage)
- Respond as you read (streaming)
- Chunked encoding (vs. content-length)

## How to use

Echo a simple string
```
curl -X POST -d "this is a test" https://localhost:8080/
this is a test%
```

You should also be able to echo larger files
```
curl -v -X POST --data-binary @largefile.bin http://localhost:8080/ > out.bin
```
