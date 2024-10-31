package middleware

import (
	"github.com/tomasen/realip"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

func RateLimit(enabled bool, burst, rps int) func(http.Handler) http.Handler {
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)

	f := func(next http.Handler) http.Handler {
		go func() {
			for {
				time.Sleep(time.Minute)

				mu.Lock()
				for ip, client := range clients {
					if time.Since(client.lastSeen) > 3*time.Minute {
						delete(clients, ip)
					}
				}
				mu.Unlock()
			}
		}()

		fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if enabled {
				ip := realip.FromRequest(r)

				mu.Lock()
				if _, found := clients[ip]; !found {
					clients[ip] = &client{
						limiter: rate.NewLimiter(rate.Limit(rps), burst),
					}
				}

				clients[ip].lastSeen = time.Now()
				limiter := clients[ip].limiter
				mu.Unlock()

				if !limiter.Allow() {
					w.Header().Set("Content-Type", "text/plain")
					w.WriteHeader(http.StatusTooManyRequests)
					return
				}
			}
			next.ServeHTTP(w, r)
		})
		return fn
	}
	return f
}
