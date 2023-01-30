package test_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/sheason2019/spoved/libs/router"
)

func HttpTestWithRecorder[T any](method, urlPath string, params any) (T, *httptest.ResponseRecorder, error) {
	var resp T
	var body *bytes.Reader = bytes.NewReader([]byte{})
	if params != nil {
		if method != "GET" && method != "DELETE" {
			buf, err := json.Marshal(params)
			if err != nil {
				return resp, nil, err
			}
			body = bytes.NewReader(buf)
		} else {
			m := map[string]any{}
			buf, err := json.Marshal(params)
			if err != nil {
				return resp, nil, err
			}
			err = json.Unmarshal(buf, &m)
			if err != nil {
				return resp, nil, err
			}
			query := url.Values{}
			for k, v := range m {
				if query[k] == nil {
					query[k] = []string{}
				}
				query[k] = append(query[k], fmt.Sprint(v))
			}
			urlPath = urlPath + "?" + query.Encode()
		}
	}

	r := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, urlPath, body)
	r.ServeHTTP(w, req)

	err := json.Unmarshal(w.Body.Bytes(), &resp)
	return resp, w, err
}

func HttpTest[T any](method, url string, params any) (T, error) {
	t, _, e := HttpTestWithRecorder[T](method, url, params)
	return t, e
}
