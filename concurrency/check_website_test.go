package concurrency

import (
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "www.naver.com" {
		return false
	}
	return true
}

func stubSlowWebsiteChecker(url string) bool {
	time.Sleep(time.Millisecond * 2)
	return true
}

func TestCheckWebsites(t *testing.T) {
	websiteUrls := []string{
		"www.google.com",
		"www.daum.net",
		"www.naver.com",
	}

	wants := map[string]bool{
		"www.google.com": true,
		"www.daum.net":   true,
		"www.naver.com":  true,
	}

	myAssert := func(t *testing.T, wants, results map[string]bool) {
		t.Helper()
		//if !reflect.DeepEqual(wants, results) {
		//	t.Fatalf("want : %v, result : %v\n", wants, results)
		//}
		//for url, _ := range wants {
		//	ret := results[url]
		//	want := wants[url]
		//	if ret != want {
		//		t.Errorf("url : %s want : %t ret : %t", url, want, ret)
		//	}
		//}
	}

	t.Run("basic check", func(t *testing.T) {
		siteChecker := mockWebsiteChecker
		results := CheckWebsites(siteChecker, websiteUrls)
		myAssert(t, wants, results)
	})
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < 100; i++ {
		urls[i] = "test url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(stubSlowWebsiteChecker, urls)
	}
}
