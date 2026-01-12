package rateLimiter

import (
	"container/list"
	"net"
	"net/http"
	"sync"
	"time"
)

type RequestLog struct {
	requests *list.List
}

type Limiter struct {
	limit      int
	timeWindow time.Duration
	clients    map[string]*RequestLog
	mu         sync.Mutex
}

func NewLimiter(limit int, timeWindow time.Duration) *Limiter {
	return &Limiter{
		limit:      limit,
		timeWindow: timeWindow,
		clients:    make(map[string]*RequestLog),
	}
}

func (l *Limiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "invalid ip", http.StatusInternalServerError)
			return
		}

		now := time.Now().UnixNano()
		window := int64(l.timeWindow)

		l.mu.Lock()
		defer l.mu.Unlock()

		if _, ok := l.clients[ip]; !ok {
			l.clients[ip] = &RequestLog{requests: list.New()}
		}

		log := l.clients[ip]

		// remove old requests
		for log.requests.Len() > 0 {
			front := log.requests.Front()
			if front.Value.(int64) < now-window {
				log.requests.Remove(front)
			} else {
				break
			}
		}

		if log.requests.Len() >= l.limit {
			http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		log.requests.PushBack(now)
		next.ServeHTTP(w, r)
	})
}
