package main

import "fmt"

var graph = make(map[string][]string) // string - имя друга, []string - друзья этого друга

func main() {
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	search("you")
}

// search поиск в ширину. O(n) + O(m) = O(n + m)
func search(name string) bool {
	var search_queue []string                           // Создание новой очереди
	search_queue = append(search_queue, graph[name]...) // Добавление элементов в очередь поиска
	var searched []string                               // С помощью этого массива вы можете отслеживать, каких людей вы искали раньше
	var person string                                   // С помощью этой переменной вы можете извлекать людей из очереди поиска

	for len(search_queue) != 0 { // Пока очередь не пуста
		// Удаление человека из очереди поиска и его сохранение в переменную person
		person = search_queue[0]
		search_queue = search_queue[1:]
		// Если человек еще не был проверен ранее, проверьте, является ли он продавцом манго
		if person_not_in_searched(person, searched) {
			if person_is_seller(person) {
				fmt.Println(person + " is mango seller!")
				return true
			}

			// Если человек не является продавцом манго, добавьте его друзей в очередь поиска
			search_queue = append(search_queue, graph[person]...)
			// Добавьте человека в массив, чтобы отслеживать, кто уже был проверен
			searched = append(searched, person)

		}
	}

	return false
}

// person_is_seller проверяет на продавца манго.
func person_is_seller(name string) bool {
	// Если последняя буква имени человека - 'm', то он продавец манго
	return name[len(name)-1] == 'm'
}

// person_not_in_searched проверяет, не включался ли человек в поиск ранее.
func person_not_in_searched(person string, searched []string) bool {
	for _, n := range searched {
		// Если человек уже был проверен ранее, верните false
		if n == person {
			return false
		}
	}
	return true
}
