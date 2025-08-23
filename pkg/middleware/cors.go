package middleware

import (
	"log"
	"net/http"
)

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := w.Header()

		origin := r.Header.Get("Origin")

		log.Printf("CORS request: Origin=%s, Method=%s",
			origin,
			r.Method)

		if origin == "" {
			next.ServeHTTP(w, r)
			return
		}

		// TODO: Настроить для прода
		// // Разрешенные origin'ы
		// allowedOrigins := []string{
		//     "http://localhost:3000",
		//     "https://example.com",
		//     "https://*.example.com", // Поддомены
		// }

		// // Получаем origin из запроса
		// origin := r.Header.Get("Origin")

		// // Проверяем, разрешен ли origin
		// var allowedOrigin string
		// for _, o := range allowedOrigins {
		//     if o == origin ||
		//        (strings.Contains(o, "*") && strings.HasSuffix(origin, strings.TrimPrefix(o, "*"))) {
		//         allowedOrigin = origin
		//         break
		//     }
		// }

		// // Если origin не найден в разрешенных, используем первый или пустую строку
		// if allowedOrigin == "" && len(allowedOrigins) > 0 {
		//     allowedOrigin = allowedOrigins[0]
		// }

		// // Устанавливаем CORS headers
		// w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)

		// Настройка на этапе разработки
		header.Set("Access-Control-Allow-Origin", "*")

		header.Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			header.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Real-IP")
			header.Set("Access-Control-Max-Age", "86400") // 24 hours
		}

		// Для preflight запросов (OPTIONS) сразу возвращаем ответ
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
