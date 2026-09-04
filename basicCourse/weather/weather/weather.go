package weather

import (
	"GolangCourse/weather/geo"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Свой клиент, а не http.Get: у http.DefaultClient таймаут нулевой,
// то есть ожидание ответа ничем не ограничено.
var client = &http.Client{Timeout: 10 * time.Second}

const baseURL = "http://wttr.in/"

func GetWeather(location geo.GeoData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", fmt.Errorf("weather: формат %d не поддерживается, допустимы 1-4", format)
	}
	base, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("weather: разбор базового url: %w", err)
	}
	// JoinPath экранирует содержимое сегмента. При конкатенации '?' и '#'
	// из названия города меняли бы структуру URL, а не попадали в путь.
	u := base.JoinPath(location.City)
	params := url.Values{}
	params.Add("format", strconv.Itoa(format))
	u.RawQuery = params.Encode()

	res, err := client.Get(u.String())
	if err != nil {
		return "", fmt.Errorf("weather: запрос погоды: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(io.LimitReader(res.Body, 1<<20))
	if err != nil {
		return "", fmt.Errorf("weather: чтение тела: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("weather: статус %d: %s",
			res.StatusCode, bytes.TrimSpace(body[:min(len(body), 512)]))
	}

	return strings.TrimSpace(string(body)), nil
}
