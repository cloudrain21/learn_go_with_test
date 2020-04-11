package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// slow
func CheckWebsites1(wc WebsiteChecker, urls []string) map[string]bool {
	retResults := make(map[string]bool)

	for i := 0; i < len(urls); i++ {
		url := urls[i]
		retResults[url] = wc(url)
	}

	return retResults
}

// concurrency (use channel)
func CheckWebsites2(wc WebsiteChecker, urls []string) map[string]bool {
	retResults := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		ret := <-resultChannel
		retResults[ret.string] = ret.bool
	}

	return retResults
}
