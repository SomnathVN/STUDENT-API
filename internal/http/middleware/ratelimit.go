package middleware

import (
	"net/http"
	"sync"
	"time"
)

type visitor struct {
	lastSeen time.Time
	requests int
}

var visitors = make(map[string]*visitor)
var mu sync.Mutex

const rateLimit = 60
const window = time.Minute

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		mu.Lock()
		v, exists := visitors[ip]
		if !exists || time.Since(v.lastSeen) > window {
			v = &visitor{lastSeen: time.Now(), requests: 1}
			visitors[ip] = v
		} else {
			v.requests++
			v.lastSeen = time.Now()
		}
		count := v.requests
		mu.Unlock()
		if count > rateLimit {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
} 