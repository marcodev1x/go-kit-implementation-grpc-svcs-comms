package transports

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/marco-kit/kit-home-service/pkg/pb/protocols/testing/testing"
	"github.com/testing-2/pkg/endpoint"
	"go.elastic.co/apm/module/apmgorilla/v2"
)

func NewHTTPServer(endpoint endpoint.EndpointSetup, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	apmgorilla.Instrument(r)

	r.Methods("GET").
		Path("/test").
		Handler(httptransport.NewServer(
			endpoint.Test,
			getTestDecodeHTTPRequest,
			encodeHttpResponse,
		))

	return r
}

func getTestDecodeHTTPRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req testing.TestRequest
	vars := mux.Vars(r)

	req.Name = vars["test"]
	return &req, nil
}

func encodeHttpResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		encodeHttpError(ctx, e, w)
		return nil
	}

	// encode.HttpHeaders(ctx, w)
	return json.NewEncoder(w).Encode(response)
}

func encodeHttpError(ctx context.Context, err error, w http.ResponseWriter) {
	// encode.HttpHeaders(ctx, w)

	e := strings.ReplaceAll(err.Error(), "rpc error: code = Unknown desc = ", "")
	switch e {
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":    e,
		"messages": "msg",
	})
}
