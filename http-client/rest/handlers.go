package rest

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/bygui86/go-traces/http-client/commons"
	"github.com/bygui86/go-traces/http-client/logging"
)

// TODO set timeout

func (s *Server) getProducts(writer http.ResponseWriter, request *http.Request) {
	span := opentracing.StartSpan("get-products-handler")
	// span, ctx := opentracing.StartSpanFromContext(context.Background(), "get-products-handler")
	startTimer := time.Now()
	defer span.Finish()

	span.Context()

	logging.Log.Info("Get products")

	endpointUrl := &url.URL{Path: rootProductsEndpoint}
	path := s.baseURL.ResolveReference(endpointUrl)
	restRequest, reqErr := http.NewRequest(http.MethodGet, path.String(), nil)
	if reqErr != nil {
		errMsg := "Get products failed: " + reqErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("products-found", 0)
		span.SetTag("error", errMsg)
		span.LogKV("products-found", 0, "error", errMsg)
		return
	}
	restRequest.Header.Set(headerAccept, headerApplicationJson)
	restRequest.Header.Set(headerUserAgent, headerUserAgentClient)

	// Transmit the span's TraceContext as HTTP headers on our outbound request.
	traceErr := opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(restRequest.Header))
	if traceErr != nil {
		errMsg := "Get products failed: " + traceErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("products-found", 0)
		span.SetTag("error", errMsg)
		span.LogKV("products-found", 0, "error", errMsg)
		return
	}

	response, respErr := s.restClient.Do(restRequest)
	if respErr != nil {
		errMsg := "Get products failed: " + respErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("products-found", 0)
		span.SetTag("error", errMsg)
		span.LogKV("products-found", 0, "error", errMsg)
		return
	}
	defer response.Body.Close()
	var products []*commons.Product
	jsonErr := json.NewDecoder(response.Body).Decode(&products)
	if jsonErr != nil {
		errMsg := "Get products failed: " + jsonErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("products-found", 0)
		span.SetTag("error", errMsg)
		span.LogKV("products-found", 0, "error", errMsg)
		return
	}

	span.SetTag("products-found", len(products))
	span.LogKV("products-found", len(products))

	sendJsonResponse(writer, http.StatusOK, products)

	IncreaseRestRequests("getProducts")
	ObserveRestRequestsTime("getProducts", float64(time.Now().Sub(startTimer).Milliseconds()))
}

// TODO

// func (s *Server) getProduct(writer http.ResponseWriter, request *http.Request) {
// 	span, ctx := opentracing.StartSpanFromContext(context.Background(), "get-product-handler")
// 	startTimer := time.Now()
// 	defer span.Finish()
//
// 	vars := mux.Vars(request)
// 	id, idErr := strconv.Atoi(vars["id"])
// 	if idErr != nil {
// 		errMsg := "Get product failed: Invalid product ID"
// 		sendErrorResponse(writer, http.StatusBadRequest, errMsg)
//
// 		span.SetTag("error", errMsg)
// 		span.LogKV("error", errMsg)
// 		return
// 	}
//
// 	logging.SugaredLog.Infof("Get product by ID: %d", id)
// 	span.SetTag("product-id", id)
//
// 	product := &database.Product{ID: id}
// 	getErr := database.GetProduct(s.dbConnection, product, ctx)
// 	if getErr != nil {
// 		var errMsg string
// 		switch getErr {
// 		case sql.ErrNoRows:
// 			errMsg = "Get product failed: product not found"
// 			sendErrorResponse(writer, http.StatusNotFound, errMsg)
// 		default:
// 			errMsg = getErr.Error()
// 			sendErrorResponse(writer, http.StatusInternalServerError, errMsg)
// 		}
//
// 		span.SetTag("product-found", false)
// 		span.SetTag("error", errMsg)
// 		span.LogKV("product-id", id, "product-found", false, "error", errMsg)
// 		return
// 	}
//
// 	span.SetTag("product-found", true)
// 	span.LogKV("product-id", id, "product-found", true)
//
// 	sendJsonResponse(writer, http.StatusOK, product)
//
// 	IncreaseRestRequests("getProduct")
// 	ObserveRestRequestsTime("getProduct", float64(time.Now().Sub(startTimer).Milliseconds()))
// }

// func (s *Server) createProduct(writer http.ResponseWriter, request *http.Request) {
// 	span, ctx := opentracing.StartSpanFromContext(context.Background(), "create-product-handler")
// 	startTimer := time.Now()
// 	defer span.Finish()
//
// 	var product database.Product
// 	unmarshErr := json.NewDecoder(request.Body).Decode(&product)
// 	if unmarshErr != nil {
// 		errMsg := "Create product failed: invalid request payload"
// 		sendErrorResponse(writer, http.StatusBadRequest, errMsg)
//
// 		span.SetTag("product-created", false)
// 		span.SetTag("error", errMsg)
// 		span.LogKV("product-created", false, "error", errMsg)
// 		return
// 	}
// 	defer request.Body.Close()
//
// 	logging.SugaredLog.Infof("Create product %s", product.String())
//
// 	createErr := database.CreateProduct(s.dbConnection, &product, ctx)
// 	if createErr != nil {
// 		errMsg := "Create product failed: " + createErr.Error()
// 		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)
//
// 		span.SetTag("product-created", false)
// 		span.SetTag("error", errMsg)
// 		span.LogKV("product-created", false, "error", errMsg)
// 		return
// 	}
//
// 	span.SetTag("product", product.String())
// 	span.SetTag("product-created", true)
// 	span.LogKV("product", product.String(), "product-created", true)
//
// 	sendJsonResponse(writer, http.StatusCreated, product)
//
// 	IncreaseRestRequests("createProduct")
// 	ObserveRestRequestsTime("createProduct", float64(time.Now().Sub(startTimer).Milliseconds()))
// }

// func (s *Server) updateProduct(writer http.ResponseWriter, request *http.Request) {
// 	span, ctx := opentracing.StartSpanFromContext(context.Background(), "update-product-handler")
// 	startTimer := time.Now()
// 	defer span.Finish()
//
// 	vars := mux.Vars(request)
// 	id, idErr := strconv.Atoi(vars["id"])
// 	if idErr != nil {
// 		errMsg := "Update product failed: invalid product ID"
// 		sendErrorResponse(writer, http.StatusBadRequest, errMsg)
//
// 		span.SetTag("product-updated", false)
// 		span.SetTag("error", errMsg)
// 		span.LogKV("product-updated", false, "error", errMsg)
// 		return
// 	}
//
// 	var product database.Product
// 	unmarshErr := json.NewDecoder(request.Body).Decode(&product)
// 	if unmarshErr != nil {
// 		errMsg := "Update product failed: invalid request payload"
// 		sendErrorResponse(writer, http.StatusBadRequest, errMsg)
//
// 		span.SetTag("product-updated", false)
// 		span.SetTag("error", errMsg)
// 		span.LogKV("product-updated", false, "error", errMsg)
// 		return
// 	}
// 	defer request.Body.Close()
//
// 	logging.SugaredLog.Infof("Update product: %s", product.String())
//
// 	product.ID = id
// 	updateErr := database.UpdateProduct(s.dbConnection, &product, ctx)
// 	if updateErr != nil {
// 		errMsg := "Update product failed: " + updateErr.Error()
// 		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)
//
// 		span.SetTag("product-updated", false)
// 		span.SetTag("error", errMsg)
// 		span.LogKV("product-updated", false, "error", errMsg)
// 		return
// 	}
//
// 	span.SetTag("product", product.String())
// 	span.SetTag("product-updated", true)
// 	span.LogKV("product", product.String(), "product-updated", true)
//
// 	sendJsonResponse(writer, http.StatusOK, product)
//
// 	IncreaseRestRequests("updateProduct")
// 	ObserveRestRequestsTime("updateProduct", float64(time.Now().Sub(startTimer).Milliseconds()))
// }

// func (s *Server) deleteProduct(writer http.ResponseWriter, request *http.Request) {
// 	span, ctx := opentracing.StartSpanFromContext(context.Background(), "delete-product-handler")
// 	startTimer := time.Now()
// 	defer span.Finish()
//
// 	vars := mux.Vars(request)
// 	id, idErr := strconv.Atoi(vars["id"])
// 	if idErr != nil {
// 		errMsg := "Delete product failed: invalid Product ID"
// 		sendErrorResponse(writer, http.StatusBadRequest, errMsg)
//
// 		span.SetTag("product-deleted", false)
// 		span.SetTag("error", errMsg)
// 		span.LogKV("product-deleted", false, "error", errMsg)
// 		return
// 	}
//
// 	logging.SugaredLog.Infof("Delete product by ID: %d", id)
// 	span.SetTag("product-id", id)
//
// 	deleteErr := database.DeleteProduct(s.dbConnection, id, ctx)
// 	if deleteErr != nil {
// 		errMsg := "Delete product failed: " + deleteErr.Error()
// 		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)
//
// 		span.SetTag("product-deleted", false)
// 		span.SetTag("error", errMsg)
// 		span.LogKV("product-deleted", false, "error", errMsg)
// 		return
// 	}
//
// 	span.SetTag("product-deleted", true)
// 	span.LogKV("product-deleted", true)
//
// 	sendJsonResponse(writer, http.StatusOK, map[string]string{"result": "success"})
//
// 	IncreaseRestRequests("deleteProduct")
// 	ObserveRestRequestsTime("deleteProduct", float64(time.Now().Sub(startTimer).Milliseconds()))
// }
