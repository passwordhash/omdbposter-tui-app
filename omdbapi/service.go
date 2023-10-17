package omdbapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseUrl = "https://www.omdbapi.com/"

func SearchByTitle(title, apiKey string) (SearchResult, error) {
	var res SearchResult
	url := getFormattedQuery(map[string]string{
		"apikey": apiKey,
		"s":      title,
	})
	resp, err := http.Get(url)
	if err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 400, 401, 404:
		return res, fmt.Errorf("запрос вернул ошибку: %v", resp.Status)
	}

	if parseErr := json.NewDecoder(resp.Body).Decode(&res); parseErr != nil {
		resp.Body.Close()
		return res, err
	}
	return res, nil
}

func GetById(id string, apiKey string) (Movie, error) {
	var res Movie
	url := getFormattedQuery(map[string]string{
		"apikey": apiKey,
		"i":      id,
	})
	resp, err := http.Get(url)

	switch resp.StatusCode {
	case 400, 401, 404:
		return res, fmt.Errorf("запрос вернул ошибку: %v", resp.Status)
	}

	if parseErr := json.NewDecoder(resp.Body).Decode(&res); parseErr != nil {
		resp.Body.Close()
		return res, err
	}
	return res, nil
}

func getFormattedQuery(params map[string]string) string {
	res := baseUrl + "?"
	for k, v := range params {
		res += fmt.Sprintf("%v=%v&", k, v)
	}
	return res[:len(res)-1]
}
