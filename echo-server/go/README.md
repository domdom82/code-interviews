# go-echo
A simple HTTP echo server

## How to use

Echo a simple string
```
curl -X POST -d "this is a test" https://my-app.com/echo
this is a test%
```

You can also echo larger files
```
curl -v -X POST --data-binary @largefile.bin http://localhost:8080/echo > out.bin
```
