#!/usr/bin/env python

from http.server import BaseHTTPRequestHandler, HTTPServer

hostName = "localhost"
serverPort = 8080

class MyServer(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header("Content-type", "text/plain")
        self.end_headers()
        self.wfile.write(bytes("hello world!", "utf-8"))

def main():
  server = HTTPServer((hostName, serverPort), MyServer)
  print("Listening at http://%s:%s" % (hostName, serverPort))

  try:
      server.serve_forever()
  except KeyboardInterrupt:
      pass

  server.server_close()
  
if __name__ == "__main__":
  main()
