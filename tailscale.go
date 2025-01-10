package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"tailscale.com/tsnet"
)

var (
	hostname = flag.String("hostname", "hello", "hostname for the tailnet")
)

// NewProxy takes target host and creates a reverse proxy
func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	return httputil.NewSingleHostReverseProxy(url), nil
}

// ProxyRequestHandler handles the http request using proxy
func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

func main() {
	flag.Parse()

	// initialize a reverse proxy and pass the actual backend server url here
	proxy, err := NewProxy("http://127.0.0.1:4000")
	if err != nil {
		panic(err)
	}

	s := &tsnet.Server{
		Hostname: *hostname,
	}

	defer s.Close()

	ln, err := s.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	// lc, err := s.LocalClient()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	log.Fatal(http.Serve(ln, http.HandlerFunc(ProxyRequestHandler(proxy))))

	// log.Fatal(http.Serve(ln, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	who, err := lc.WhoIs(r.Context(), r.RemoteAddr)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), 500)
	// 		return
	// 	}

	// 	fmt.Fprintf(w, "<html><body><h1>Hello, world!</h1>\n")
	// 	fmt.Fprintf(w, "<p>You are <b>%s</b> from <b>%s</b> (%s)</p>",
	// 		html.EscapeString(who.UserProfile.LoginName),
	// 		html.EscapeString(firstLabel(who.Node.ComputedName)),
	// 		r.RemoteAddr)
	// })))
}

func firstLabel(s string) string {
	if hostname, _, ok := strings.Cut(s, "."); ok {
		return hostname
	}

	return s
}
