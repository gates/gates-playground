package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gates/gates"
)

type runReq struct {
	Code string `json:"code"`
}

type runResp struct {
	Result string `json:"result"`
}

type errResp struct {
	Message string `json:"message"`
}

func runString(s string) (gates.Value, error) {
	r := gates.New()
	program, err := gates.Compile(s)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return r.RunProgram(ctx, program)
}

func main() {
	http.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		switch r.Method {
		case http.MethodOptions:
		case http.MethodPost:
			dec := json.NewDecoder(r.Body)
			payload := &runReq{}
			if err := dec.Decode(payload); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			enc := json.NewEncoder(w)
			code := payload.Code
			v, err := runString(code)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				enc.Encode(&errResp{Message: err.Error()})
				return
			}
			enc.Encode(&runResp{Result: v.ToString()})
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe("127.0.0.1:3001", nil)
}
