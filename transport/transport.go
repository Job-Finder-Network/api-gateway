package transport

import (
	"context"
	"net/http"

	"github.com/Job-Finder-Network/api-gateway/endpoint"
	"github.com/Job-Finder-Network/api-gateway/middleware"
	"github.com/Job-Finder-Network/api-gateway/reqresp"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http/httputil"
	"net/url"
)

func NewHTTPServer(ctx context.Context, endpoints endpoint.UserEndpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(middleware.JsonMiddleware)
	target, _ := url.Parse("http://localhost:8000")
	proxy := httputil.NewSingleHostReverseProxy(target)
	r.HandleFunc("/login1", handler(proxy))
	r.Methods("POST").Path("/register").Handler(httptransport.NewServer(
		endpoints.CreateUser,
		reqresp.DecodeCreateUserReq,
		reqresp.EncodeResponse,
	))
	r.Methods("POST").Path("/login").Handler(httptransport.NewServer(
		endpoints.Login,
		reqresp.DecodeLoginReq,
		reqresp.EncodeResponse,
	))

	return r

}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p.ServeHTTP(w, r)
	}
}

func checkToken(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p.ServeHTTP(w, r)
	}
}
