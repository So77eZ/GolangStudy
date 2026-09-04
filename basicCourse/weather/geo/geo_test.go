package geo_test

import (
	"GolangCourse/weather/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange - подготовка теста (expected result)
	city := "Moscow"
	expected := geo.GeoData{
		City: "Moscow",
	}
	// Act - выполняем ф-цию с заданными в моменте Arrange данными
	got, err := geo.GetMyLocation(city)
	// Assert - проверка результата с expected
	if err != nil {
		t.Error("Ошибка получения города:", err)
	}

	if got.City != expected.City {
		t.Errorf("Ожидаемый город - %v, не равен полученному - %v", expected.City, got.City)
	}
}
