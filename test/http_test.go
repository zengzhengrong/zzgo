package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/zengzhengrong/zzgo/request"
)

func TestHtppQuery(t *testing.T) {
	url := "https://httpbin.org?"
	args := map[string]string{
		"a": "1",
		"b": "2",
	}
	result := request.HttpBuildQuery(args)
	fmt.Println(result)
	url = url + result
	fmt.Println(url)
	fmt.Println(strings.Index(url, "还"))
}

func TestRequest(t *testing.T) {
	h := map[string]string{
		"A": "a",
		"B": "b",
	}
	q := map[string]string{
		"a": "1",
		"b": "2",
	}
	body := map[string]string{
		"aa": "1",
		"ba": "2",
	}
	r, err := request.NewReuqest(
		http.MethodGet,
		"https://httpbin.org/get",
		request.WithHeader(h),
		request.WithBody(body),
		request.WithQuery(q),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}

func TestClient(t *testing.T) {
	h := map[string]string{
		"A": "a",
		"B": "b",
	}
	q := map[string]string{
		"a": "1",
		"b": "2",
	}
	body := map[string]string{
		"aa": "1",
		"ba": "2",
	}
	b, _ := json.Marshal(body)
	r, err := request.NewReuqest(
		http.MethodPost,
		"https://httpbin.org/post",
		request.WithHeader(h),
		request.WithBody(b),
		request.WithQuery(q),
	)
	if err != nil {
		panic(err)
	}
	client := request.NewClient(
		request.WithDebug(true),
		request.WithTimeOut(2*time.Second),
	)
	resp, err := client.HttpClient.Do(r.HttpReq)
	if err != nil {
		panic(err)
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
