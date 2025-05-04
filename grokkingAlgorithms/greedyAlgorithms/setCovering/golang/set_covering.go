package main

import "fmt"

// Задача о покрытии множества
func main() {
	// Список штатов
	states_needed := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}

	var stations = make(map[string][]string)
	stations["kone"] = []string{"id", "nv", "ut"}
	stations["ktwo"] = []string{"wa", "id", "mt"}
	stations["kthree"] = []string{"or", "nv", "ca"}
	stations["kfour"] = []string{"nv", "ut"}
	stations["kfive"] = []string{"ca", "az"}

	// Ключи - названия станций, значения - сокращенные обозначения штатов, входящих в зону охвата
	station_key := []string{"kone", "ktwo", "kthree", "kfour", "kfive"}

	var final_stations []string // Итоговой набор станций, покрывающий все штаты

	for len(states_needed) > 0 {
		var best_station string     // Станция, которая покрывает больше всего штатов, не входящих в зону покрытия
		var states_covered []string // Список штатов, обслуживаемые этой станцией, и которые еще не входят в зону покрытия

		for _, station := range station_key {
			states := stations[station]                    // Список штатов, обслуживаемых этой станцией
			var covered = equaldata(states_needed, states) // Список штатов, обслуживаемые этой станцией, и которые еще не входят в зону покрытия
			// Проверяем, если количество штатов, обслуживаемых этой станцией,
			// и которые еще не входят в зону покрытия больше, чем количество штатов,
			// обслуживаемых предыдущей станцией, то записываем ее в качестве лучшей
			if len(covered) > len(states_covered) {
				best_station = station
				states_covered = covered
			}
		}
		// Удаляем из списка штатов, обслуживаемых этой станцией, и которые еще не входят в зону покрытия
		states_needed = removedata(states_needed, states_covered)
		// Добавляем станцию в итоговый набор станций
		final_stations = append(final_stations, best_station)

	}

	fmt.Println("final_stations:", final_stations)
}

// equaldata функция для нахождения пересечения двух массивов строк.
func equaldata(states_needed []string, states []string) []string {
	var covered []string

	for _, state_needed := range states_needed {
		for _, state := range states {
			// Если штаты совпадают, то добавляем его в список обслуживаемых станцией
			if state_needed == state {
				covered = append(covered, state_needed)
			}
		}
	}
	return covered
}

// removedata функция для удаления элементов из массива строк.
func removedata(states_needed []string, states_covered []string) []string {
	for _, state_covered := range states_covered {
		states_needed = remove(states_needed, state_covered)
	}
	return states_needed
}

// remove функция для удаления элемента из массива строк.
func remove(states_needed []string, state_covered string) []string {
	for i, state_needed := range states_needed {
		// Если штаты совпадают, то удаляем его из списка обслуживаемых станцией
		if state_covered == state_needed {
			return append(states_needed[:i], states_needed[i+1:]...)
		}
	}
	return states_needed
}
