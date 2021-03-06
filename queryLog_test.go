package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/fredallotey2000/ciq/internal"
	"github.com/fredallotey2000/ciq/utils"
)

func TestNumberOfUsers(t *testing.T) {
	queryLogs, err := internal.NewQueryLogs()
	results := queryLogs.NumberOfUsers()
	if results != 6 || err != nil {
		t.Errorf("Output %v not equal to expected %v", results, 6)
	}

}

func benchmarkNumberOfUsers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queryLogs, _ := internal.NewQueryLogs()
		results := queryLogs.NumberOfUsers()
		fmt.Print(results)
	}

}
func BenchmarkNumberOfUsers(b *testing.B) {
	benchmarkNumberOfUsers(b)
}

func TestNumberOfUploadsLargerThan(t *testing.T) {
	queryLogs, err := internal.NewQueryLogs()
	results := queryLogs.NumberOfUploadsLargerThan(85, "upload")
	if results != 7 || err != nil {
		t.Errorf("Output %v not equal to expected %v", results, 7)
	}

}

func benchmarkNumberOfUploadsLargerThan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queryLogs, _ := internal.NewQueryLogs()
		results := queryLogs.NumberOfUploadsLargerThan(85, "upload")
		fmt.Print(results)
	}
}

func BenchmarkNumberOfUploadsLargerThan(b *testing.B) {
	benchmarkNumberOfUploadsLargerThan(b)
}

func TestNumberOfDownloadsLargerThan(t *testing.T) {
	queryLogs, err := internal.NewQueryLogs()
	results := queryLogs.NumberOfUploadsLargerThan(50, "download")
	if results != 137 || err != nil {
		t.Errorf("Output %v not equal to expected %v", results, 137)
	}
}

func benchmarkNumberOfDownloadsLargerThan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queryLogs, _ := internal.NewQueryLogs()
		results := queryLogs.NumberOfUploadsLargerThan(50, "download")
		fmt.Print(results)
	}
}

func BenchmarkNumberOfDownloadsLargerThan(b *testing.B) {
	benchmarkNumberOfDownloadsLargerThan(b)
}

func TestNumberOfUploadsPerUserPerDate(t *testing.T) {
	queryLogs, err2 := internal.NewQueryLogs()
	parsedDate, err := time.Parse(utils.InputdateFormat, "Apr-19-2020")
	if err != nil {
		t.Error()
	}
	results := queryLogs.NumberOfUploadsPerUserPerDate("rosannaM", parsedDate, "download")
	if results != 10 || err2 != nil {
		t.Errorf("Output %v not equal to expected %v", results, 10)
	}
}
func benchmarkNumberOfUploadsPerUserPerDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queryLogs, _ := internal.NewQueryLogs()
		parsedDate, _ := time.Parse(utils.InputdateFormat, "Apr-19-2020")
		results := queryLogs.NumberOfUploadsPerUserPerDate("rosannaM", parsedDate, "download")
		fmt.Print(results)
	}
}
func BenchmarkNumberOfUploadsPerUserPerDate(b *testing.B) {
	benchmarkNumberOfUploadsPerUserPerDate(b)
}

func TextNumberOfDownsPerUserPerDate(t *testing.T) {
	queryLogs, err2 := internal.NewQueryLogs()
	parsedDate, err := time.Parse(utils.InputdateFormat, "Apr-19-2020")
	if err != nil {
		t.Error()
	}
	results := queryLogs.NumberOfUploadsPerUserPerDate("rosannaM", parsedDate, "upload")
	if results != 0 || err2 != nil {
		t.Errorf("Output %v not equal to expected %v", results, 0)
	}
}
func benchmarkNumberOfDownsPerUserPerDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queryLogs, _ := internal.NewQueryLogs()
		parsedDate, _ := time.Parse(utils.InputdateFormat, "Apr-19-2020")
		results := queryLogs.NumberOfUploadsPerUserPerDate("rosannaM", parsedDate, "upload")
		fmt.Print(results)
	}
}

func BenchmarkNumberOfDownsPerUserPerDate(b *testing.B) {
	benchmarkNumberOfDownsPerUserPerDate(b)
}
func ExampleGetDataBankLocation() {
	queryLogs, _ := internal.NewQueryLogs()
	results := queryLogs.NumberOfUploadsLargerThan(50, "download")
	fmt.Println(results)
	// Output: 137
}
