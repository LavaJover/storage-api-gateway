package middleware

import(
	"net/http"
	"sync"
	"time"
)

var (
	rateLimit = make(map[string]int)
	mu sync.Mutex
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		mu.Lock()
		rateLimit[ip]++
		count := rateLimit[ip]
		mu.Unlock()

		if count > 10 {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		// Сбрасываем счётчик через 1 минуту
		time.AfterFunc(1*time.Minute, func() {
			mu.Lock()
			rateLimit[ip]--
			mu.Unlock()
		})

		next.ServeHTTP(w, r)
	})
}