package domiannameindex

import (
	"bufio"
	"os"
	"testing"
)

func BenchmarkInsert(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	f, err := os.Open("domains.txt")
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	domains := make([]string, 0, 100)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		domains = append(domains, scanner.Text())
	}

	for i := 0; i < b.N; i++ {
		t := New()
		for _, domain := range domains {
			t.Insert(domain)
		}
	}
}

func BenchmarkFind(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	t := New()
	t.Insert("example.com")
	t.Insert("*.example.com")
	t.Insert("*.name.example.com")
	t.Insert("google.com")
	for i := 0; i < b.N; i++ {
		t.Find("google.com")
	}
}
