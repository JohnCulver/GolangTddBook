package concurrency

type WebsiteChecker func(string) bool

type result struct {
	url     string
	success bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChan := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChan <- result{url: u, success: wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChan
		results[r.url] = r.success
	}

	return results
}

// 2046234396 ns/op
