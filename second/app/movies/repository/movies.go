package repository

import (
	"bibit-test/config"
	"bibit-test/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// HTTP ...
type HTTP struct {
	http   *http.Client
	option *Option
}

// Option ...
type Option struct {
	BaseURL string `envconfig:"BASE_URL"`
	APIKey  string `envconfig:"API_KEY"`
}

// New ...
func NewMoviesRepository(option *Option) *HTTP {
	return &HTTP{
		http: &http.Client{
			Timeout: 30 * time.Second,
		},
		option: option,
	}
}

func (h *HTTP) Search(pagination int64, searchword string) (movie *models.MoviesSearchResponse, err error) {
	var res models.MoviesSearchResponse

	addr := fmt.Sprintf("%s/?apikey=%s&page=%d&s=%s", config.Config.App.Address, config.Config.App.APIKey, pagination, searchword)

	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := h.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		io.Copy(ioutil.Discard, resp.Body)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(body) == 0 {
		err = errors.New("Search not found")
		return nil, err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (h *HTTP) Detail(id string) (movie *models.MovieDetailResponse, err error) {
	var res models.MovieDetailResponse

	addr := fmt.Sprintf("%s/?apikey=%s&i=%s", config.Config.App.Address, config.Config.App.APIKey, id)

	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := h.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		io.Copy(ioutil.Discard, resp.Body)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(body) == 0 {
		err = errors.New("Detail not found")
		return nil, err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
