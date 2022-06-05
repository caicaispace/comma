package pool

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func doPoolV2Test(pool *PoolV2) error {
	testURL, err := url.Parse("http://127.0.0.1:8080/")
	if err != nil {
		return err
	}
	httpClient := pool.Get()
	resp, err := httpClient.Do(&http.Request{
		URL:    testURL,
		Method: "GET",
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	defer pool.Recycle(httpClient)
	_, err = ioutil.ReadAll(resp.Body)
	return nil
}

func TestNewV2Demo(t *testing.T) {
	pool := NewV2(25, 200)

	urlString := "https://www.baidu.com/"

	reqURL, err := url.Parse(urlString)
	if err != nil {
		fmt.Printf("[ERROR] Unable to parse: %s", urlString)
		os.Exit(1)
	}

	httpClient := pool.Get()

	resp, err := httpClient.Do(&http.Request{
		URL:    reqURL,
		Method: "GET",
	})
	if err != nil {
		fmt.Printf("[ERROR] Unable to perform request: %s", err)
		os.Exit(2)
	}
	defer pool.Recycle(httpClient)
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[ERROR] Unable to read response body: %s", err)
		os.Exit(3)
	}
}

func BenchmarkPoolV2_DoPool_3_20(b *testing.B) {
	b.Logf("[INFO] Starting BenchmarkPoolV2_DoPool_3_20")
	pool := NewV2(3, 20)
	for n := 0; n < b.N; n++ {
		if err := doPoolV2Test(pool); err != nil {
			b.Error("pool: ", err)
		}
	}
	b.Logf("[INFO] Completed BenchmarkPoolV2_DoPool_3_20")
}
