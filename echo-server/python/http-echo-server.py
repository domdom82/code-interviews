#!/usr/bin/env python

from http.server import BaseHTTPRequestHandler, HTTPServer

hostName = "localhost"
serverPort = 8080

class MyServer(BaseHTTPRequestHandler):
    def do_POST(self):
        self.send_response(200)
        self.send_header("Content-type", "text/plain")
        self.end_headers()

        content_length = int(self.headers["Content-Length"])
        data = self.rfile.read(content_length)
        self.wfile.write(data)
          
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
