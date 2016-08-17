package middlewares

import (
    "log"
    "net/url"
    "net/http"
    "github.com/vulcand/oxy/forward"
    "github.com/vulcand/oxy/roundrobin"
    "github.com/vulcand/oxy/stream"
    "github.com/xaviergodart/gydro/models"
)

type ReverseProxy struct {
    EntryPoints map[string]*stream.Streamer
}

func NewReverseProxy(apis []*models.Api) *ReverseProxy {
    conf := make(map[string]*stream.Streamer)

    // Loading routing and backend configuration from database
    for _, api := range apis {
        fwd, _ := forward.New()
        lb, _ := roundrobin.New(fwd)
        for _, backend := range api.Backends {
            target, _ := url.Parse(backend)
            lb.UpsertServer(target)
        }
        stream, _ := stream.New(lb, stream.Retry(`(IsNetworkError() || ResponseCode() >= 500) && Attempts() < 2`))
        conf[api.Route] = stream
    }

    return &ReverseProxy{EntryPoints: conf}
}

func (rp *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("X-GydroProxy", "GydroProxy")
    log.Println(r.URL.Path)
    if backends, ok := rp.EntryPoints[r.URL.Path]; ok {
        log.Println("proxy: custom")
        log.Println(ok)
        backends.ServeHTTP(w, r)
        return
    }

    log.Println("proxy: default")
    http.NotFoundHandler().ServeHTTP(w, r)
}
