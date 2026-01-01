package rateLimiter

import "container/list"

type RequestLog struct {
	requests *list.List
}

func RateLimiter(timestamps []int64, ipAddresses []string, limit int, timeWindow int64) []int {
	requestLog := make(map[string]*RequestLog)
	result := make([]int, len(timestamps))

	for i, timestamp := range timestamps {
		ip := ipAddresses[i]
		if _, ok := requestLog[ip]; !ok {
			requestLog[ip] = &RequestLog{requests: list.New()}
		}
		log := requestLog[ip]

		// Remove outdated requests
		for log.requests.Len() > 0 {
			front := log.requests.Front()
			if front.Value.(int64) < timestamp-timeWindow {
				log.requests.Remove(front)
			} else {
				break
			}
		}

		// Check if we can accept the request
		if log.requests.Len() < limit {
			log.requests.PushBack(timestamp)
			result[i] = 1 // Accept request
		} else {
			result[i] = 0 // Reject request
		}

	}
	return result
}
