Вам предстоит сделать поисковик, в котором можно искать репозитории в GitHub по
их названию. Допишите код в файл index.js, чтобы при нажатии на кнопку «Найти»
отправлялся запрос на https://api.nomoreparties.co/github-search?q= и текст
поиска, который пользователь ввёл в форму. Например, запрос
https://api.nomoreparties.co/github-search?q=reactвозвращает объект вида:
{
    total_count: 1748551
    incomplete_results: false
    items: [
        {
            id: 10270250,
            name: "react",
            full_name: "facebook/react",
            html_url: "https://github.com/facebook/react",
            description: "A declarative, efficient, and flexible JavaScript library for building user interfaces.",
            stargazers_count: 160193,
            language: "JavaScript",
            ...
        },
        ...
    ]
}

Запрос должен отправляться по событию 'submit' на форме. Используйте для решения
задачи fetch().then().catch()конструкцию. Для успешного решения задачи 
необходимо выполнить следующие действия:
Перед отправкой запроса нужно вызывать функцию onSubmitStart(). Эта 
вспомогательная функция нужна для «оживления» интерфейса.
Функцию renderCount(total_count) следует вызывать только при наличии результатов
 поиска. Вместе с ней нужно добавить в resultsContainer найденные репозитории, 
 вызвав для каждого из них функцию template(item). Обратите внимание, функция 
 renderCount принимает аргумент из ответа сервера: total_count.
Функцию renderEmptyResults() нужно вызывать только при отсутствии результатов по 
запрошенному названию, то есть когда total_count равен нулю.
На случай ошибки, произошедшей на стороне сервера или же внутри клиентского 
кода, — вызывайте функцию renderError(). Вызов этой функции добавит 
пользовательскому интерфейсу информативности.
