package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsitesChecker(url string) bool{
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com" : true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds": false,
	}
	got := CheckWebsites(mockWebsitesChecker, websites)

	if !reflect.DeepEqual(want, got){
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

//using keyword benchmark to test system
func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}