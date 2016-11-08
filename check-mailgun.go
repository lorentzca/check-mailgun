package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	Apikey *string `short:"p" long:"apikey" required:"true" description:"Mailgun Api Key"`
	Domain *string `short:"d" long:"domain" required:"true" description:"Mailgun Domain"`
}

func url() string {
	flags.Parse(&opts)
	url := fmt.Sprintf("https://api.mailgun.net/v3/domains/%s", *opts.Domain)

	return url
}

func httpBody() []byte {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url(), nil)
	if err != nil {
		fmt.Println(err)
	}
	req.SetBasicAuth("api", *opts.Apikey)

	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err)
	}

	body, err3 := ioutil.ReadAll(res.Body)
	if err3 != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	return body
}

func main() {
	type domain struct {
		State string
	}

	type data struct {
		Domain domain
	}

	body := httpBody()

	var d data
	json.Unmarshal(body, &d)
	fmt.Println(d.Domain.State)
}
