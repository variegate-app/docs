package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Регистрируем обработчики для каждого роута
	for i := 1; i <= 10; i++ {
		path := fmt.Sprintf("/helloworld1", i)
		message := fmt.Sprintf("Hello, World %d", i)

		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, message)
		})
	}

	// Запускаем сервер на порту 8080
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
