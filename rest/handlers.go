package rest

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"

	"github.com/bygui86/go-traces/database"
	"github.com/bygui86/go-traces/logging"
)

func (s *Server) getProducts(writer http.ResponseWriter, request *http.Request) {
	span := opentracing.StartSpan("get-products")
	startTimer := time.Now()
	defer span.Finish()

	logging.Log.Info("Get products")

	count, _ := strconv.Atoi(request.FormValue("count"))
	start, _ := strconv.Atoi(request.FormValue("start"))
	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}
	products, err := database.GetProducts(s.dbConnection, start, count)
	if err != nil {
		sendErrorResponse(writer, http.StatusInternalServerError, "Get products failed: "+err.Error())
		return
	}

	span.SetTag("products-found", len(products))
	span.LogKV("products-found", len(products))

	sendJsonResponse(writer, http.StatusOK, products)

	IncreaseRestRequests("getProducts")
	ObserveRestRequestsTime("getProducts", float64(time.Now().Sub(startTimer).Milliseconds()))
}

func (s *Server) getProduct(writer http.ResponseWriter, request *http.Request) {
	span := opentracing.StartSpan("get-product")
	startTimer := time.Now()
	defer span.Finish()

	vars := mux.Vars(request)
	id, idErr := strconv.Atoi(vars["id"])
	if idErr != nil {
		sendErrorResponse(writer, http.StatusBadRequest, "Get product failed: Invalid product ID")
		return
	}

	logging.SugaredLog.Infof("Get product by ID: %d", id)

	span.SetTag("requested-product-id", id)

	product := &database.Product{ID: id}
	getErr := database.GetProduct(s.dbConnection, product)
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			sendErrorResponse(writer, http.StatusNotFound, "Get product failed: product not found")
		default:
			sendErrorResponse(writer, http.StatusInternalServerError, getErr.Error())
		}

		span.SetTag("requested-product-found", false)
		span.LogKV("requested-product-id", id, "requested-product-found", false)
		return
	}

	span.SetTag("requested-product-found", true)
	span.LogKV("requested-product-id", id, "requested-product-found", true)

	sendJsonResponse(writer, http.StatusOK, product)

	IncreaseRestRequests("getProduct")
	ObserveRestRequestsTime("getProduct", float64(time.Now().Sub(startTimer).Milliseconds()))
}

func (s *Server) createProduct(writer http.ResponseWriter, request *http.Request) {
	span := opentracing.StartSpan("create-product")
	startTimer := time.Now()
	defer span.Finish()

	var product database.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		sendErrorResponse(writer, http.StatusBadRequest, "Create product failed: invalid request payload")
		return
	}
	defer request.Body.Close()

	logging.SugaredLog.Infof("Create product %s", product.String())

	if err := database.CreateProduct(s.dbConnection, &product); err != nil {
		sendErrorResponse(writer, http.StatusInternalServerError, "Create product failed: "+err.Error())

		span.SetTag("new-product-created", false)
		span.LogKV("new-product-created", false)
		return
	}

	span.SetTag("new-product-id", product.ID)
	span.SetTag("new-product-name", product.Name)
	span.SetTag("new-product-price", product.Price)
	span.SetTag("new-product-created", true)
	span.LogKV("new-product-id", product.ID,
		"new-product-name", product.Name,
		"new-product-price", product.Price,
		"new-product-created", true)

	sendJsonResponse(writer, http.StatusCreated, product)

	IncreaseRestRequests("createProduct")
	ObserveRestRequestsTime("createProduct", float64(time.Now().Sub(startTimer).Milliseconds()))
}

func (s *Server) updateProduct(writer http.ResponseWriter, request *http.Request) {
	span := opentracing.StartSpan("update-product")
	startTimer := time.Now()
	defer span.Finish()

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendErrorResponse(writer, http.StatusBadRequest, "Update product failed: invalid product ID")

		span.SetTag("product-updated", false)
		span.LogKV("product-updated", false)
		return
	}

	var product database.Product
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&product); err != nil {
		sendErrorResponse(writer, http.StatusBadRequest, "Update product failed: invalid request payload")

		span.SetTag("product-updated", false)
		span.LogKV("product-updated", false)
		return
	}
	defer request.Body.Close()

	logging.SugaredLog.Infof("Update product: %s", product.String())

	product.ID = id
	if err := database.UpdateProduct(s.dbConnection, &product); err != nil {
		sendErrorResponse(writer, http.StatusInternalServerError, "Update product failed: "+err.Error())

		span.SetTag("product-updated", false)
		span.LogKV("product-updated", false)
		return
	}

	span.SetTag("updated-product-id", product.ID)
	span.SetTag("updated-product-name", product.Name)
	span.SetTag("updated-product-price", product.Price)
	span.SetTag("product-updated", true)
	span.LogKV("updated-product-id", product.ID,
		"updated-product-name", product.Name,
		"updated-product-price", product.Price,
		"product-updated", true)

	sendJsonResponse(writer, http.StatusOK, product)

	IncreaseRestRequests("updateProduct")
	ObserveRestRequestsTime("updateProduct", float64(time.Now().Sub(startTimer).Milliseconds()))
}

func (s *Server) deleteProduct(writer http.ResponseWriter, request *http.Request) {
	span := opentracing.StartSpan("delete-product")
	startTimer := time.Now()
	defer span.Finish()

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendErrorResponse(writer, http.StatusBadRequest, "Delete product failed: invalid Product ID")

		span.SetTag("product-deleted", false)
		span.LogKV("product-deleted", false)
		return
	}

	logging.SugaredLog.Infof("Delete product by ID: %d", id)

	span.SetTag("product-to-delete", id)

	product := &database.Product{ID: id}
	if err := database.DeleteProduct(s.dbConnection, product); err != nil {
		sendErrorResponse(writer, http.StatusInternalServerError, "Delete product failed: "+err.Error())

		span.SetTag("product-deleted", false)
		span.LogKV("product-deleted", false)
		return
	}

	span.SetTag("product-deleted", true)
	span.LogKV("product-deleted", true)

	sendJsonResponse(writer, http.StatusOK, map[string]string{"result": "success"})

	IncreaseRestRequests("deleteProduct")
	ObserveRestRequestsTime("deleteProduct", float64(time.Now().Sub(startTimer).Milliseconds()))
}
