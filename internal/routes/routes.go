package routes

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Здесь может быть общий код для всех запросов...

	switch r.Method {
	case http.MethodGet:
		// Обработка GET запроса...

	case http.MethodPost:
		// Обработка POST запроса...

	case http.MethodOptions:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
