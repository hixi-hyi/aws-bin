//usr/bin/env go run "$0" "$@" ; exit
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/smartystreets/go-aws-auth"
)

func main() {
	var url, body, method string
	flag.StringVar(&url, "u", "", "")
	flag.StringVar(&method, "m", "", "")
	flag.StringVar(&body, "b", "", "")
	flag.Parse()
	if url == "" && method == "" {
		flag.PrintDefaults()
		return
	}
	req, _ := http.NewRequest(method, url, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	do(req)
}

func do(req *http.Request) {
	awsauth.Sign(req, awsauth.Credentials{
		AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		SecurityToken:   os.Getenv("AWS_SESSION_TOKEN"),
	})
	{
		dreq, _ := httputil.DumpRequestOut(req, true)
		fmt.Fprintf(os.Stderr, "%s\n\n", dreq)
	}
	client := new(http.Client)
	resp, _ := client.Do(req)
	{
		dresp, _ := httputil.DumpResponse(resp, false)
		fmt.Fprintf(os.Stderr, "%s\n", dresp)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", body)
}
