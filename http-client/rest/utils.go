package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"

	"github.com/bygui86/go-traces/http-client/commons"
	"github.com/bygui86/go-traces/http-client/logging"
)

const (
	// urls
	rootProductsEndpoint     = "/products"
	productsIdEndpoint       = rootProductsEndpoint + "/{id:[0-9]+}" // used by mux router
	productsIdServerEndpoint = rootProductsEndpoint + "/%d"          // used to reach http-server

	// headers
	// keys
	headerContentType  = "Content-Type"
	headerAccept       = "Accept"
	headerUserAgent    = "User-Agent"
	headerCustomSource = "Custom-Source"
	// values
	headerUserAgentClient = "GoTracesHttpClient/1.0"
	headerApplicationJson = "application/json"
)

// SERVER

func (s *Server) setupRestClient() error {
	logging.Log.Debug("Setup base URL")

	var urlErr error
	s.baseURL, urlErr = url.Parse(fmt.Sprintf("http://%s:%d", s.config.restServerHost, s.config.restServerPort))
	if urlErr != nil {
		return urlErr
	}

	s.restClient = &http.Client{
		Timeout: 3 * time.Second,
	}

	return nil
}

func (s *Server) setupRouter() {
	logging.Log.Debug("Create new router")

	s.router = mux.NewRouter().StrictSlash(true)

	s.router.Use(requestInfoPrintingMiddleware)

	s.router.HandleFunc(rootProductsEndpoint, s.getProducts).Methods(http.MethodGet)
	s.router.HandleFunc(productsIdEndpoint, s.getProduct).Methods(http.MethodGet)
	s.router.HandleFunc(rootProductsEndpoint, s.createProduct).Methods(http.MethodPost)
	s.router.HandleFunc(productsIdEndpoint, s.updateProduct).Methods(http.MethodPut)
	s.router.HandleFunc(productsIdEndpoint, s.deleteProduct).Methods(http.MethodDelete)
}

func (s *Server) setupHTTPServer() {
	logging.SugaredLog.Debugf("Create new HTTP server on port %d", s.config.restPort)

	if s.config != nil {
		s.httpServer = &http.Server{
			Addr:    fmt.Sprintf(commons.HttpServerHostFormat, s.config.restHost, s.config.restPort),
			Handler: s.router,
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: commons.HttpServerWriteTimeoutDefault,
			ReadTimeout:  commons.HttpServerReadTimeoutDefault,
			IdleTimeout:  commons.HttpServerIdelTimeoutDefault,
		}
		return
	}

	logging.Log.Error("HTTP server creation failed: REST server configurations not loaded")
}

// HANDLERS

func sendJsonResponse(writer http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	writer.Header().Set(headerContentType, headerApplicationJson)
	writer.WriteHeader(code)
	_, writeErr := writer.Write(response)
	if writeErr != nil {
		logging.SugaredLog.Errorf("Error sending JSON response: %s", writeErr.Error())
	}
}

func sendErrorResponse(writer http.ResponseWriter, code int, message string) {
	sendJsonResponse(writer, code, map[string]string{"error": message})
}

func closeBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		logging.SugaredLog.Errorf("Closing request body failed: %s", err.Error())
	}
}

// OTHERS

func (s *Server) setRequestHeaders(restRequest *http.Request) {
	restRequest.Header.Set(headerAccept, headerApplicationJson)
	restRequest.Header.Set(headerUserAgent, headerUserAgentClient)
	ip, ipErr := getPublicIp()
	if ipErr != nil {
		logging.SugaredLog.Errorf("Get public IP address failed: %s", ipErr.Error())
	} else {
		restRequest.Header.Set(headerCustomSource, ip)
	}
}

func getPublicIp() (string, error) {
	url := "https://api.ipify.org?format=text"
	logging.SugaredLog.Infof("Get public IP address from %s", url)

	response, respErr := http.Get(url)
	if respErr != nil {
		return "", respErr
	}
	defer closeBody(response.Body)

	ip, bodyErr := ioutil.ReadAll(response.Body)
	if bodyErr != nil {
		return "", bodyErr
	}

	logging.SugaredLog.Infof("Public IP '%s' will be added as 'Custom-Source' header", ip)

	return string(ip), nil
}
