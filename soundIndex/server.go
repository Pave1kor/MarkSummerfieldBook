package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

// Обработчик главной страницы
func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/window.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
		return
	}

	if r.Method == "POST" {
		r.ParseForm()
		input := r.FormValue("user_input")

		// Вызываем функцию обработки данных
		result := processData(input)

		sessionData, err := json.Marshal(result)
		if err != nil {
			log.Printf("Ошибка сериализации: %v", err)
			http.Error(w, "Ошибка сериализации данных", http.StatusInternalServerError)
			return
		}

		// Сохраняем данные в сессии
		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Ошибка создания сессии", http.StatusInternalServerError)
			return
		}

		session.Values["result"] = sessionData
		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   60,   // 1 час
			HttpOnly: true, // защита от XSS
		}

		if err := session.Save(r, w); err != nil {
			http.Error(w, "Ошибка сохранения сессии", http.StatusInternalServerError)
			log.Printf("Ошибка сохранения сессии: %v", err)
			return
		}

		tmpl.Execute(w, result)
		return
	}

	tmpl.Execute(w, nil)
}

// Обработчик страницы test
func testHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	result, ok := session.Values["result"].([]byte)
	if !ok {
		http.Error(w, "Нет данных", http.StatusBadRequest)
		return
	}

	var data []FormData
	err := json.Unmarshal(result, &data)
	if err != nil {
		log.Printf("Ошибка десериализации: %v", err)
		http.Error(w, "Ошибка десериализации данных", http.StatusInternalServerError)
		return
	}

	// Преобразуем данные в таблицу
	table := convertToTable(data)

	// Загружаем шаблон и передаем данные
	temp, err := template.ParseFiles("templates/test.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
		return
	}

	temp.Execute(w, table)
}

func main() {
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	http.HandleFunc("/", handler)
	http.HandleFunc("/test", testHandler)
	log.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
