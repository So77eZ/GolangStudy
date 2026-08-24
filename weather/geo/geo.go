package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Анонимные запросы по https ipapi.co режет по 429 (эвристики free-tier),
// по http отвечает. Канал открытый — ответ можно подменить, доверяем только городу.
const apiURL = "http://ipapi.co/json/"
const apiCityResURL = "https://countriesnow.space/api/v0.1/countries/population/cities"

var client = &http.Client{Timeout: 5 * time.Second}

// GeoData хранит данные о местоположении.
type GeoData struct {
	City string
}

type apiResponse struct {
	City   string `json:"city"`
	Error  bool   `json:"error"`
	Reason string `json:"reason"`
}

type CityPopulationRes struct {
	Error bool `json:"error"`
}

func GetMyLocation(city string) (GeoData, error) {
	if city != "" {
		if checkCity(city) {
			return GeoData{
				City: city,
			}, nil
		}
		panic("Передаваемый город не найден")
	}

	res, err := client.Get(apiURL)
	if err != nil {
		return GeoData{}, fmt.Errorf("geo: запрос местоположения: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(io.LimitReader(res.Body, 1<<20)) // потолок 1 МБ
	if err != nil {
		return GeoData{}, fmt.Errorf("geo: чтение тела: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return GeoData{}, fmt.Errorf("geo: статус %d: %s",
			res.StatusCode, bytes.TrimSpace(body[:min(len(body), 512)]))
	}

	var r apiResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return GeoData{}, fmt.Errorf("geo: перекодировка json: %w", err)
	}
	if r.Error {
		return GeoData{}, fmt.Errorf("geo: api вернул ошибку: %s", r.Reason)
	}
	if r.City == "" {
		return GeoData{}, errors.New("geo: ответ без города")
	}

	return GeoData{City: r.City}, nil
}

func checkCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})

	res, err := client.Post(apiCityResURL, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return false //GeoData{}, fmt.Errorf("geo: проверка города: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(io.LimitReader(res.Body, 1<<20)) // потолок 1 МБ
	if err != nil {
		fmt.Errorf("geo: чтение тела: %w", err)
		return false
	}

	if res.StatusCode != http.StatusOK {
		fmt.Errorf("geo: статус %d: %s",
			res.StatusCode, bytes.TrimSpace(body[:min(len(body), 512)]))
		return false
	}

	var populationRes CityPopulationRes
	json.Unmarshal(body, &populationRes)
	return !populationRes.Error
}
