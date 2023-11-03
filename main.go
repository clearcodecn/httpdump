package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
)

var (
	port string
)

var (
	nginxConf = `server {
        listen 80;

        server_name changeme;

        location / {
			  proxy_set_header Host $host;
			  proxy_set_header X-Real-IP $remote_addr;
			  proxy_set_header X-Forwarded-For $remote_addr;
			  proxy_pass http://127.0.0.1%s;
        }
}`
)

func init() {
	flag.StringVar(&port, "p", ":9090", "listen port")
}

func main() {
	flag.Parse()

	fmt.Println("nginx config: ")
	fmt.Printf(nginxConf, port)

	s := new(miniServer)

	http.ListenAndServe(port, s)
}

type miniServer struct{}

func (s *miniServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	out, _ := httputil.DumpRequest(r, true)
	fmt.Println(string("<<< -----"))
	fmt.Println(string(out))
	fmt.Println("")
	fmt.Println("")

	w.Write([]byte("s"))
}
