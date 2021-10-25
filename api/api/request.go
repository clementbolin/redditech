package api

import "net/http"

type Request struct {
	typeRequest string
	body        map[string]interface{}
	token       string
	url         string
	req         *http.Request
	err         error
}

func (data *Request) createRequest(req *Request) {
	data.token = req.token
	data.body = req.body
	data.typeRequest = req.token
	data.req, data.err = http.NewRequest(data.typeRequest, data.url, nil)
}
