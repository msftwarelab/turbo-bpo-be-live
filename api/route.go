package api

import (
	"encoding/json"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/lonmarsDev/bpo-golang-grahpql/internal/middleware"
	"github.com/lonmarsDev/bpo-golang-grahpql/internal/services"
)

func Initialize(router chi.Router) {
	//route declaration
	router.Post("/users/login", login)
	router.Get("/pipelines", searchPipelines)
	router.Get("/iforms/{orderID}", iform)
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type loginBody struct {
		Email    string
		Password string
	}

	var loginVar loginBody
	json.NewDecoder(r.Body).Decode(&loginVar)

	ctxIPAddress := middleware.GetUserIPAddress(r.Context())
	res, err := services.Login(r.Context(), loginVar.Email, loginVar.Password, ctxIPAddress)
	if err != nil {
		log.Debug("err %+v", err)
		render.Render(w, r, &ErrResponse{HTTPStatusCode: 404, StatusText: err.Error()})
		return

	}
	if res.Token == nil {
		render.Render(w, r, errInValidUserNameOrPassword)
		return
	}

	render.Respond(w, r, res)
}

func iform(w http.ResponseWriter, r *http.Request) {
	orderNumber := chi.URLParam(r, "orderID")
	if orderNumber == "" {
		render.Render(w, r, ErrNotFound)
		return
	}

	iform, err := services.IformByOrderNumber(r.Context(), orderNumber)
	if err != nil {
		render.Render(w, r, ErrNotFound)

	}

	render.Respond(w, r, iform)
}

func searchPipelines(w http.ResponseWriter, r *http.Request) {
	//log.Debug("%+v", prettyPrint(r))
	_, err := middleware.GetUserCtx(r.Context())
	if err != nil {
		log.Error("err %+v", err)
		render.Render(w, r, &ErrResponse{HTTPStatusCode: 404, StatusText: err.Error()})
		return
	}
	orderNumber := r.FormValue("order-number")
	orderAddress := r.FormValue("order-address")
	limit := int(100)
	if r.FormValue("limit") != "" {
		var err error
		limit, err = strconv.Atoi(r.FormValue("limit"))
		if err != nil {
			log.Error("err %+v", err)
			limit = 100
		}
	}
	pipelineList, err := services.SearchPipelineByOrderNumberOrORderAddress(r.Context(), orderNumber, orderAddress, limit)
	if err != nil {
		log.Debug("err %+v", err)
		render.Render(w, r, &ErrResponse{HTTPStatusCode: 404, StatusText: err.Error()})
		return

	}
	if len(pipelineList) == 0 {
		log.Debug("err %+v", err)
		render.Render(w, r, &ErrResponse{HTTPStatusCode: 404, StatusText: "no record found"})
		return
	}
	render.Respond(w, r, pipelineList)
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
var errInValidUserNameOrPassword = &ErrResponse{HTTPStatusCode: 404, StatusText: "envalid username or password."}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
