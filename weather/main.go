package main

import (
	"GolangCourse/weather/geo"
	"GolangCourse/weather/weather"
	"flag"
	"fmt"
	"os"
)

// утилита для парсинга погоды по передаваемым параметрам и флагам
func main() {
	// Баннер и ошибки — в stderr, в stdout только результат:
	// иначе вывод нельзя передать по конвейеру.
	fmt.Fprintln(os.Stderr, "Погодный парсер")

	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "формат вывода погоды: 1-4")
	flag.Parse()
	if *format < 1 || *format > 4 {
		fmt.Fprintf(os.Stderr, "формат %d не поддерживается, допустимы 1-4\n", *format)
		os.Exit(2)
	}

	// город можно задать флагом или первым позиционным аргументом: weather Moscow
	name := *city
	if name == "" && flag.NArg() > 0 {
		name = flag.Arg(0)
	}

	geoData, err := geo.GetMyLocation(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "не удалось определить город: %v\n", err)
		os.Exit(1)
	}

	weatherData, err := weather.GetWeather(geoData, *format)
	if err != nil {
		fmt.Fprintf(os.Stderr, "не удалось получить погоду: %v\n", err)
		os.Exit(1)
	}

	// форматы 3-4 уже содержат название города в ответе wttr.in
	if *format <= 2 {
		fmt.Printf("%s: %s\n", geoData.City, weatherData)
	} else {
		fmt.Println(weatherData)
	}
}
