package bapi

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

type API struct {
	URL string
}

func NewAPI(url string) *API {
	return &API{URL: url}
}

func (a *API) httpGET(c context.Context, path string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", a.URL, path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

func (a *API) GetQueList(c context.Context, page string, amount string) ([]byte, error) {
	body, err := a.httpGET(c, fmt.Sprintf("%s/%s/%s", "/api/question/get_question_list", page, amount))
	if err != nil {
		return nil, err
	}

	return body, nil
}
