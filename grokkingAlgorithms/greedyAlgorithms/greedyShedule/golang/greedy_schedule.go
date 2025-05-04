package main

import (
	"fmt"
	"sort"
)

// Lesson представляет структуру урока
type Lesson struct {
	Subject string
	Start   float64
	End     float64
}

// Задача составления расписания уроков
func main() {
	lessons := []Lesson{
		{"рисование", 9, 10},
		{"английский", 9.5, 10.5},
		{"математика", 10, 11},
		{"информатика", 10.5, 11.5},
		{"музыка", 11, 12},
	}

	schedule := GreedySchedule(lessons)
	fmt.Println("Оптимальное расписание уроков:")
	for _, lesson := range schedule {
		fmt.Printf("%s: %.1f-%.1f\n", lesson.Subject, lesson.Start, lesson.End)
	}
}

// GreedySchedule реализует жадный алгоритм для составления расписания
func GreedySchedule(lessons []Lesson) []Lesson {
	// Сортируем уроки по времени окончания
	sort.Slice(lessons, func(i, j int) bool {
		return lessons[i].End < lessons[j].End
	})
	// [{рисование, 9, 10} {английский, 9.5, 10.5} {математика, 10, 11} {информатика, 10.5, 11.5} {музыка, 11, 12}]

	var result []Lesson
	lastLessonEnd := 0.0 // Время окончания предыдущего урока

	for _, lesson := range lessons {
		// Добавляем урок, если его время начала больше времени окончания предыдущего урока
		if lesson.Start >= lastLessonEnd {
			result = append(result, lesson)
			lastLessonEnd = lesson.End
		}
	}

	return result
}
