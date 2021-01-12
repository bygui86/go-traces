package rest

import (
	"net/http"

	"github.com/bygui86/go-traces/http-server/logging"
)

func requestInfoPrintingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		logging.SugaredLog.Infof("RequestURI: %s", request.RequestURI)
		logging.SugaredLog.Infof("URL: %s", request.URL)
		logging.SugaredLog.Infof("x-forwarded-for: %s", request.Header.Get("X-FORWARDED-FOR"))
		logging.SugaredLog.Infof("RemoteAddr: %s", request.RemoteAddr)

		printHeaders(request)

		next.ServeHTTP(writer, request)
		// all remaining handlers on the pipeline will execute before reaching this line
	})
}

func printHeaders(r *http.Request) {
	logging.Log.Info("Headers:")
	for name, values := range r.Header {
		for _, value := range values {
			logging.SugaredLog.Infof("\t%s: %s", name, value)
		}
	}
}
