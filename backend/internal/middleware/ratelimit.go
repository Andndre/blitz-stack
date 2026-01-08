package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// IPRateLimiter stores rate limiters for each IP address
type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

// NewIPRateLimiter creates a new IPRateLimiter
// r: requests per second
// b: burst size
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	i := &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}

	// Cleanup old IPs background routine (simple cleanup)
	go func() {
		for {
			time.Sleep(1 * time.Minute)
			i.mu.Lock()
			// Reset the map every minute to prevent memory leaks from old IPs
			i.ips = make(map[string]*rate.Limiter)
			i.mu.Unlock()
		}
	}()

	return i
}

// GetLimiter returns the rate limiter for the provided IP address
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter, exists := i.ips[ip]
	if !exists {
		limiter = rate.NewLimiter(i.r, i.b)
		i.ips[ip] = limiter
	}

	return limiter
}

// LimitMiddleware wraps the HTTP handler with rate limiting logic
func LimitMiddleware(next http.Handler) http.Handler {
	// Allow 5 requests per second, with a burst of 10
	limiter := NewIPRateLimiter(5, 10)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			// If we can't parse IP (e.g. mock test), fall back or allow
			ip = r.RemoteAddr
		}

		// If behind a proxy (Docker/Nginx), use X-Forwarded-For
		if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
			ip = xff
		}

		if !limiter.GetLimiter(ip).Allow() {
			http.Error(w, "429 Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
