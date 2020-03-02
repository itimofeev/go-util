package util

import (
	"net/http"
	"net/http/pprof"
)

func GetDebugHandler(prefix string) http.Handler {
	var debugMux = http.NewServeMux()
	debugMux.HandleFunc(prefix+"/", pprof.Index)
	debugMux.HandleFunc(prefix+"/cmdline", pprof.Cmdline)
	debugMux.HandleFunc(prefix+"/profile", pprof.Profile)
	debugMux.HandleFunc(prefix+"/symbol", pprof.Symbol)
	debugMux.HandleFunc(prefix+"/trace", pprof.Trace)

	debugMux.Handle(prefix+"/allocs", pprof.Handler("allocs"))
	debugMux.Handle(prefix+"/goroutine", pprof.Handler("goroutine"))
	debugMux.Handle(prefix+"/heap", pprof.Handler("heap"))
	debugMux.Handle(prefix+"/threadcreate", pprof.Handler("threadcreate"))
	debugMux.Handle(prefix+"/block", pprof.Handler("block"))
	debugMux.Handle(prefix+"/mutex", pprof.Handler("mutex"))

	return debugMux
}

func RunDebugHandlerIfNeeded(addr string) {
	if len(addr) == 0 {
		return
	}

	go func() {
		_ = http.ListenAndServe(addr, GetDebugHandler("/debug/pprof"))
	}()
}
