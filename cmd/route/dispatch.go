package route

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"test/pkg/action"
)

type Dispatch struct {
	http.Handler
}

func (d *Dispatch) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var actionStr = strings.ReplaceAll(req.URL.Path, "/", "")
	var value, found = route.Load(actionStr)
	if !found {
		fmt.Printf("Can't found handler for action: %v\n", actionStr)
	}
	var a, ok = value.(action.Action)
	if !ok {
		fmt.Printf("Can't cast to a handler, value: %v\n", value)
	}
	var buf, _ = ioutil.ReadAll(req.Body)
	_ = json.Unmarshal(buf, &a)
	values := map[string]interface{}{
		"http.request":      req,
		"http.request.body": string(buf)}
	ctx := context.WithValue(req.Context(), "http", values)
	var response = a.Process(ctx)
	var bytes, _ = json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}