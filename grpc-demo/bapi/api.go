package bapi

import (
	"bytes"
	"context"
	"encoding/json"
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

func (a *API) httpPOST(c context.Context, path string, ch int64, location string) ([]byte, error) {
	type Res struct {
		Ch       int64
		Location string
	}
	re := Res{
		Ch:       ch,
		Location: location,
	}
	jre, _ := json.Marshal(re)

	resp, err := http.Post(fmt.Sprintf("%s/%s", a.URL, path), "application/x-www-form-urlencoded", bytes.NewBuffer(jre))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

func (a *API) GetQueList(c context.Context, page string, amount string, ch int64, location string) ([]byte, error) {
	body, err := a.httpPOST(c, fmt.Sprintf("%s/%s/%s", "api/question/get_question_list", page, amount), ch, location)
	if err != nil {
		return nil, err
	}

	return body, nil
}
