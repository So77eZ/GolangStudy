package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Создать приложение, которые сначала выдает меню:
// - 1. Посмотреть закладки,
// - 2. Добавить закладку,
// - 3. Удалить закладку,
// - 4. Редактировать закладку,
// - 5. Выход.
// При выборе пункта 1, выводится список закладок, если их нет, то сообщение "Закладок нет".
// При выборе пункта 2, 2 поля ввода: название закладки и ссылка. После ввода, закладка добавляется в мапу.
// При выборе пункта 3, выводится список закладок с их индексами, пользователь вводит индекс закладки,
// которую хочет удалить. После удаления выводится сообщение "Закладка удалена".
// При выборе пункта 4, выводится список закладок с их индексами, пользователь вводит индекс закладки,
// которую хочет редактировать. После редактирования выводится сообщение "Закладка отредактирована".
// При выборе пункта 5, программа завершает работу.
type bookmarkStringMap = map[string]string

func main() {
	m := make(bookmarkStringMap, 3)

	m["A"] = "1"
	m["B"] = "2"
	m["C"] = "3"

	fmt.Println(len(m))

	bookmarks := bookmarkStringMap{
		"google":        "https://www.google.com",
		"youtube":       "https://www.youtube.com",
		"github":        "https://github.com",
		"stackoverflow": "https://stackoverflow.com",
		"reddit":        "https://www.reddit.com",
		"wikipedia":     "https://www.wikipedia.org",
		"lolkek":        "https://lolkek.ru",
		"yandex":        "https://yandex.ru",
	}

	fmt.Println("Программа управления закладками")
	for {
		fmt.Println("\nВыберите пункт меню: ")
		printMenu()
		userInput := getUserInput()
		choice, err := strconv.Atoi(userInput)
		if err != nil || choice < 1 || choice > 5 {
			fmt.Println("\nОшибка: выберите пункт меню от 1 до 5")
			continue
		}
		if choice == 5 {
			fmt.Println("\nВыход из программы")
			break
		} else {
			menuNavigate(choice, bookmarks)
		}
	}
}

func printMenu() {
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку ")
	fmt.Println("4. Редактировать закладку")
	fmt.Println("5. Выход")
	fmt.Println()
}

func getUserInput() string {
	var userInput string
	fmt.Scan(&userInput)
	return userInput
}

func menuNavigate(choice int, bookmarks bookmarkStringMap) {

	switch choice {
	case 1:
		showBookmarks(bookmarks)
	case 2:
		fmt.Println("\nВведите название закладки: ")

		bookmarkKey := strings.ToLower(getUserInput())
		if _, exists := bookmarks[bookmarkKey]; exists {
			fmt.Printf("\nЗакладка '%s' уже существует\n", bookmarkKey)
			return
		}
		fmt.Println("Введите ссылку закладки: ")
		bookmarkValue := getUserInput()
		bookmarks[bookmarkKey] = bookmarkValue

		fmt.Printf("\nЗакладка '%s' добавлена\n", bookmarkKey)
	case 3:
		bookmarkKey, err := selectBookmarkByIndex(bookmarks, "Удалять нечего, закладок нет")
		if err != nil {
			fmt.Printf("\nОшибка: %s\n", err)
			return
		}
		delete(bookmarks, bookmarkKey)
		fmt.Printf("\nЗакладка '%s' удалена\n", bookmarkKey)
	case 4:
		bookmarkKey, err := selectBookmarkByIndex(bookmarks, "Редактировать нечего, закладок нет")
		if err != nil {
			fmt.Printf("\nОшибка: %s\n", err)
			return
		}
		fmt.Printf("\nВы выбрали закладку '%s' - '%s'\n", bookmarkKey, bookmarks[bookmarkKey])
		fmt.Println("\nЧто вы хотите изменить?")
		fmt.Println("1. Название")
		fmt.Println("2. Ссылка")
		fmt.Println("3. Оба поля")

		editChoice := getUserInput()
		editChoiceInt, err := strconv.Atoi(editChoice)
		if err != nil || editChoiceInt < 1 || editChoiceInt > 3 {
			fmt.Println("\nОшибка: выберите пункты 1, 2 или 3")
			return
		}
		oldValue := bookmarks[bookmarkKey]
		newKey := bookmarkKey
		newValue := bookmarks[bookmarkKey]

		if editChoiceInt == 1 || editChoiceInt == 3 {
			fmt.Println("\nВведите новое название закладки: ")
			newKeyInput := getUserInput()
			if newKeyInput != "" {
				newKey = strings.ToLower(newKeyInput)
				_, exists := bookmarks[newKey]
				if newKey != bookmarkKey && exists {
					fmt.Printf("\nОшибка: закладка '%s' уже существует\n", newKey)
					return
				}
			}
		}
		if editChoiceInt == 2 || editChoiceInt == 3 {
			fmt.Println("\nВведите новую ссылку закладки: ")
			newValueInput := getUserInput()
			if newValueInput != "" {
				newValue = newValueInput
			}
		}

		if newKey == bookmarkKey && newValue == oldValue {
			fmt.Println("\nИзменений не внесено")
			return
		}

		if newKey != bookmarkKey {
			delete(bookmarks, bookmarkKey)
			bookmarks[newKey] = newValue
		} else {
			bookmarks[bookmarkKey] = newValue
		}

		fmt.Printf("\nЗакладка '%s' обновлена на '%s' - '%s'\n", bookmarkKey, newKey, newValue)
	}
}

func showBookmarks(bookmarks bookmarkStringMap) {

	if len(bookmarks) == 0 {
		fmt.Println("\nЗакладок нет")
		return
	}

	keys := getSortedKeys(bookmarks)

	fmt.Println("\nСписок закладок:")
	for i, key := range keys {
		fmt.Printf("%d. %s - %s\n", i+1, key, bookmarks[key])
	}
}

func getSortedKeys(bookmarks bookmarkStringMap) []string {
	keys := make([]string, 0, len(bookmarks))
	for key := range bookmarks {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func getKeyByIndex(bookmarks bookmarkStringMap, index int) (string, error) {
	if index < 1 || index > len(bookmarks) {
		return "", fmt.Errorf("индекс должен быть от 1 до %d", len(bookmarks))
	}
	keys := getSortedKeys(bookmarks)
	return keys[index-1], nil
}

func selectBookmarkByIndex(bookmarks bookmarkStringMap, emptyMsg string) (string, error) {
	if len(bookmarks) == 0 {
		return "", fmt.Errorf("%s", emptyMsg)
	}
	showBookmarks(bookmarks)
	fmt.Println("\nВведите номер закладки: ")
	inputNumber := getUserInput()
	index, err := strconv.Atoi(inputNumber)
	if err != nil {
		return "", fmt.Errorf("введите число")
	}
	return getKeyByIndex(bookmarks, index)
}
