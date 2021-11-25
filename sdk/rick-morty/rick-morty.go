package rick_morty

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var baseURL = "https://rickandmortyapi.com/api"

func GetCaracters() (*CaracterPayloadResponse, error) {
	var caracters CaracterPayloadResponse
	err := request(http.MethodGet, "/character/?page=19", nil, &caracters)
	if err != nil {
		return nil, err
	}
	return &caracters, nil
}

func request(verb string, uri string, body interface{}, bindTo interface{}) error {

	if body != nil {
		buff := new(bytes.Buffer)
		err := json.NewEncoder(buff).Encode(body)
		if err !=nil{
			return nil
		}

		req, err := http.NewRequest(verb, baseURL+uri, buff)
		if err != nil {
			return err
		}
		return execRequest(req, bindTo)
	}

	req, err := http.NewRequest(verb, baseURL+uri, nil)
	if err != nil {
		return err
	}

	return execRequest(req, bindTo)
}

func execRequest(req *http.Request, bindTo interface{}) error {
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if bindTo != nil {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return json.Unmarshal(data, bindTo)
	}
	return nil
}
