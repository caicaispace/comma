package pool

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
)

// Set up a quick HTTP server for local testing
func init() {
	http.HandleFunc("/",
		func(respOut http.ResponseWriter, reqIn *http.Request) {
			defer reqIn.Body.Close()
		},
	)

	go func() {
		log.Fatalf("%s\n", http.ListenAndServe(":8080", nil))
	}()
}

// doPoolTest is the workhorse testing function
func doPoolTest(client Client) error {
	testURL, err := url.Parse("restful://127.0.0.1:8080/")
	if err != nil {
		return err
	}

	var doFunc func(req *http.Request) (*http.Response, error)

	switch cp := client.(type) {
	case *Pool:
		doFunc = cp.DoPool
	case *http.Client:
		doFunc = cp.Do
	}

	resp, err := doFunc(&http.Request{URL: testURL})
	if err != nil {
		return err
	}

	_, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	return nil
}

func TestNewDemo(t *testing.T) {
	standardLibClient := &http.Client{}

	pool := New(standardLibClient, 25, 200)

	urlString := "https://www.baidu.com/"

	reqURL, err := url.Parse(urlString)
	if err != nil {
		fmt.Printf("[ERROR] Unable to parse: %s", urlString)
		os.Exit(1)
	}

	resp, err := pool.Do(&http.Request{
		URL:    reqURL,
		Method: "GET",
	})
	if err != nil {
		fmt.Printf("[ERROR] Unable to perform request: %s", err)
		os.Exit(2)
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[ERROR] Unable to read response body: %s", err)
		os.Exit(3)
	}
}

// TestNew tests creating a PClient
func TestNew(t *testing.T) {
	t.Logf("[INFO] Starting TestNew")

	standardLibClient := &http.Client{}

	_ = New(standardLibClient, 25, 200) // Max 25 connections, 200 requests-per-second
	_ = New(standardLibClient, 0, 0)    // Why do this? Just use http.Client
	_ = New(standardLibClient, -1, -1)  // What

	t.Logf("[INFO] Completed TestNew")
}

// TestPClient_Do tests performing a drop-in http.Client with pooling
func TestPool_Do(t *testing.T) {
	t.Logf("[INFO] Starting TestPool_Do")
	pool := New(&http.Client{}, 0, 0)

	if err := doPoolTest(pool); err != nil {
		t.Error("pool: ", err)
	}
	t.Logf("[INFO] Completed TestPool_Do")
}

// TestPool_DoPool tests performing a request with the pooling logic
func TestPool_DoPool(t *testing.T) {
	t.Logf("[INFO] Starting TestPool_DoPool")
	pool := New(&http.Client{}, 0, 0)

	if err := doPoolTest(pool); err != nil {
		t.Error("pool: ", err)
	}

	pool2 := New(&http.Client{}, 25, 200)

	if err := doPoolTest(pool2); err != nil {
		t.Error("pool: ", err)
	}

	pool3 := New(&http.Client{}, -1, -1)

	if err := doPoolTest(pool3); err != nil {
		t.Error("pool: ", err)
	}
	t.Logf("[INFO] Completed TestPool_DoPool")
}

// BenchmarkPool_Do benchmarks the pooling logic
func BenchmarkPool_Do(b *testing.B) {
	b.Logf("[INFO] Starting BenchmarkPool_Do")
	pool := New(&http.Client{}, 0, 0)

	for n := 0; n < b.N; n++ {
		if err := doPoolTest(pool); err != nil {
			b.Error("pool: ", err)
		}
	}
	b.Logf("[INFO] Completed BenchmarkPool_Do")
}

// BenchmarkPool_DoPool benchmarks the pooling logic
func BenchmarkPool_DoPool(b *testing.B) {
	b.Logf("[INFO] Starting BenchmarkPool_DoPool")
	pool := New(&http.Client{}, 0, 0)

	for n := 0; n < b.N; n++ {
		if err := doPoolTest(pool); err != nil {
			b.Error("pool: ", err)
		}
	}
	b.Logf("[INFO] Completed BenchmarkPool_DoPool")
}

// BenchmarkBaseline benchmarks a request with the standard library http.Client
func BenchmarkBaseline(b *testing.B) {
	b.Logf("[INFO] Starting BenchmarkBaseline")
	stdClient := &http.Client{}

	for n := 0; n < b.N; n++ {
		if err := doPoolTest(stdClient); err != nil {
			b.Error("pool: ", err)
		}
	}
	b.Logf("[INFO] Completed BenchmarkBaseline")
}

// BenchmarkPool_DoPool_10_0 benchmarks the pooling logic
func BenchmarkPool_DoPool_10_0(b *testing.B) {
	b.Logf("[INFO] Starting BenchmarkPool_DoPool_10_0")
	pool := New(&http.Client{}, 10, 0)

	for n := 0; n < b.N; n++ {
		if err := doPoolTest(pool); err != nil {
			b.Error("pool: ", err)
		}
	}
	b.Logf("[INFO] Completed BenchmarkPool_DoPool_10_0")
}

// BenchmarkPool_DoPool_0_10 benchmarks the pooling logic
func BenchmarkPool_DoPool_0_10(b *testing.B) {
	b.Logf("[INFO] Starting BenchmarkPool_DoPool_0_10")
	pool := New(&http.Client{}, 0, 10)

	for n := 0; n < b.N; n++ {
		if err := doPoolTest(pool); err != nil {
			b.Error("pool: ", err)
		}
	}
	b.Logf("[INFO] Completed BenchmarkPool_DoPool_0_10")
}

// BenchmarkPool_DoPool_10_10 benchmarks the pooling logic
func BenchmarkPool_DoPool_10_10(b *testing.B) {
	b.Logf("[INFO] Starting BenchmarkPool_DoPool_10_10")
	pool := New(&http.Client{}, 10, 10)

	for n := 0; n < b.N; n++ {
		if err := doPoolTest(pool); err != nil {
			b.Error("pool: ", err)
		}
	}
	b.Logf("[INFO] Completed BenchmarkPool_DoPool_10_10")
}

// BenchmarkPool_DoPool_10_100 benchmarks the pooling logic
func BenchmarkPool_DoPool_10_100(b *testing.B) {
	b.Logf("[INFO] Starting BenchmarkPool_DoPool_10_100")
	pool := New(&http.Client{}, 10, 100)

	for n := 0; n < b.N; n++ {
		if err := doPoolTest(pool); err != nil {
			b.Error("pool: ", err)
		}
	}
	b.Logf("[INFO] Completed BenchmarkPool_DoPool_10_100")
}

// BenchmarkPool_DoPool_10_200 benchmarks the pooling logic
func BenchmarkPool_DoPool_10_200(b *testing.B) {
	b.Logf("[INFO] Starting BenchmarkPool_DoPool_10_200")
	pool := New(&http.Client{}, 10, 200)

	for n := 0; n < b.N; n++ {
		if err := doPoolTest(pool); err != nil {
			b.Error("pool: ", err)
		}
	}
	b.Logf("[INFO] Completed BenchmarkPool_DoPool_10_200")
}

// BenchmarkPool_DoPool_20_100 benchmarks the pooling logic
func BenchmarkPool_DoPool_20_100(b *testing.B) {
	b.Logf("[INFO] Starting BenchmarkPool_DoPool_20_100")
	pool := New(&http.Client{}, 20, 100)

	for n := 0; n < b.N; n++ {
		if err := doPoolTest(pool); err != nil {
			b.Error("pool: ", err)
		}
	}
	b.Logf("[INFO] Completed BenchmarkPool_DoPool_20_100")
}

// BenchmarkPool_DoPool_20_200 benchmarks the pooling logic
func BenchmarkPool_DoPool_20_200(b *testing.B) {
	b.Logf("[INFO] Starting BenchmarkPool_DoPool_20_200")
	pool := New(&http.Client{}, 20, 200)

	for n := 0; n < b.N; n++ {
		if err := doPoolTest(pool); err != nil {
			b.Error("pool: ", err)
		}
	}
	b.Logf("[INFO] Completed BenchmarkPool_DoPool_20_200")
}

// BenchmarkPool_DoPool_30_100 benchmarks the pooling logic
func BenchmarkPool_DoPool_30_100(b *testing.B) {
	b.Logf("[INFO] Starting BenchmarkPool_DoPool_30_100")
	pool := New(&http.Client{}, 30, 100)

	for n := 0; n < b.N; n++ {
		if err := doPoolTest(pool); err != nil {
			b.Error("pool: ", err)
		}
	}
	b.Logf("[INFO] Completed BenchmarkPool_DoPool_30_100")
}

// BenchmarkPool_DoPool_30_200 benchmarks the pooling logic
func BenchmarkPool_DoPool_30_200(b *testing.B) {
	b.Logf("[INFO] Starting BenchmarkPool_DoPool_30_200")
	pool := New(&http.Client{}, 30, 200)
	for n := 0; n < b.N; n++ {
		if err := doPoolTest(pool); err != nil {
			b.Error("pool: ", err)
		}
	}
	b.Logf("[INFO] Completed BenchmarkPool_DoPool_30_200")
}
