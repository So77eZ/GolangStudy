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
// - 4. Выход.
// При выборе пункта 1, выводится список закладок, если их нет, то сообщение "Закладок нет".
// При выборе пункта 2, 2 поля ввода: название закладки и ссылка. После ввода, закладка добавляется в мапу.
// При выборе пункта 3, выводится список закладок с их индексами, пользователь вводит индекс закладки,
// которую хочет удалить. После удаления выводится сообщение "Закладка удалена".
// При выборе пункта 4, программа завершает работу.

func main() {
	bookmarks := make(map[string]string)
	bookmarks["google"] = "https://www.google.com"
	bookmarks["youtube"] = "https://www.youtube.com"
	bookmarks["github"] = "https://github.com"
	bookmarks["stackoverflow"] = "https://stackoverflow.com"
	bookmarks["reddit"] = "https://www.reddit.com"
	bookmarks["wikipedia"] = "https://www.wikipedia.org"
	bookmarks["lolkek"] = "https://lolkek.ru"
	bookmarks["yandex"] = "https://yandex.ru"

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

func menuNavigate(choice int, bookmarks map[string]string) {

	switch choice {
	case 1:
		showBookmarks(bookmarks)
	case 2:
		fmt.Println("\nВведите название закладки: ")

		bookmarkKey := strings.ToLower(getUserInput())
		if checkBookmarkExists(bookmarks, bookmarkKey) {
			fmt.Printf("\nЗакладка '%s' уже существует\n", bookmarkKey)
			return
		}
		fmt.Println("Введите ссылку закладки: ")
		bookmarkValue := getUserInput()
		bookmarks[bookmarkKey] = bookmarkValue

		fmt.Printf("\nЗакладка '%s' добавлена\n", bookmarkKey)
	case 3:
		if len(bookmarks) == 0 {
			fmt.Println("Удалять нечего, закладок нет")
		} else {
			showBookmarks(bookmarks)
			fmt.Println("\nВведите номер закладки, которую хотите удалить: ")
			inputNumber := getUserInput()
			index, err := strconv.Atoi(inputNumber)
			if err != nil {
				fmt.Println("\nОшибка: введите число")
				return
			}

			bookmarkKey, err := getKeyByIndex(bookmarks, index)
			if err != nil {
				fmt.Printf("\nОшибка: %s\n", err)
				return
			} else {
				delete(bookmarks, bookmarkKey)
				fmt.Printf("\nЗакладка '%s' удалена\n", bookmarkKey)
			}
		}
	case 4:
		if len(bookmarks) == 0 {
			fmt.Println("Редактировать нечего, закладок нет")
		} else {
			showBookmarks(bookmarks)
			fmt.Println("\nВведите номер закладки, которую хотите редактировать: ")
			inputNumber := getUserInput()
			index, err := strconv.Atoi(inputNumber)
			if err != nil {
				fmt.Println("\nОшибка: введите число")
				return
			}

			oldKey, err := getKeyByIndex(bookmarks, index)
			if err != nil {
				fmt.Printf("\nОшибка: %s\n", err)
				return
			}
			oldValue := bookmarks[oldKey]
			fmt.Printf("\nВы выбрали закладку '%s' - '%s'\n", oldKey, oldValue)

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

			newKey := oldKey
			newValue := oldValue

			if editChoiceInt == 1 || editChoiceInt == 3 {
				fmt.Println("\nВведите новое название закладки: ")
				newKeyInput := getUserInput()
				if newKeyInput != "" {
					newKey = strings.ToLower(newKeyInput)
					if newKey != oldKey && checkBookmarkExists(bookmarks, newKey) {
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

			if newKey != oldKey {
				delete(bookmarks, oldKey)
				bookmarks[newKey] = newValue
			} else {
				bookmarks[oldKey] = newValue
			}

			fmt.Printf("\nЗакладка '%s' обновлена на '%s' - '%s'\n", oldKey, newKey, newValue)
		}
	}
}

func showBookmarks(bookmarks map[string]string) {

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

func checkBookmarkExists(bookmarks map[string]string, bookmarkKey string) bool {
	_, exists := bookmarks[bookmarkKey]
	return exists
}

func getSortedKeys(bookmarks map[string]string) []string {
	keys := make([]string, 0, len(bookmarks))
	for key := range bookmarks {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func getKeyByIndex(bookmarks map[string]string, index int) (string, error) {
	if index < 1 || index > len(bookmarks) {
		return "", fmt.Errorf("индекс должен быть от 1 до %d", len(bookmarks))
	}
	keys := getSortedKeys(bookmarks)
	return keys[index-1], nil
}
