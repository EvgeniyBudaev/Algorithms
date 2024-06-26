/*
С помощью жадного алгоритма написать код для составления списка уроков, максимум из которых можн провести в одном классе.
*/

const lessons = [
    { subject: 'рисование', start: 9, end: 10 },
    { subject: 'английский', start: 9.5, end: 10.5 },
    { subject: 'математика', start: 10, end: 11 },
    { subject: 'информатика', start: 10.5, end: 11.5 },
    { subject: 'музыка', start: 11, end: 12 },
];

function greedySchedule(lessons) {
    lessons.sort((a, b) => a.end - b.end); // сортируем уроки по завершению
    const result = [];
    let lastLessonEnd = 0; // счетчик последнего урока

    for (const lesson of lessons) {
        if (lesson.start >= lastLessonEnd) {
            result.push(lesson);
            lastLessonEnd = lesson.end; // добавляем уроки, которые закончились раньше последнего урока
        }
    }

    return result;
}

console.log(greedySchedule(lessons));
/*
[
    { subject: 'рисование', start: 9, end: 10 },
    { subject: 'математика', start: 10, end: 11 },
    { subject: 'музыка', start: 11, end: 12 },
]
*/
