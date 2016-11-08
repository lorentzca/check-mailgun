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

func mailgunEndPoint() string {
	flags.Parse(&opts)
	url := fmt.Sprintf("https://api.mailgun.net/v3/domains/%s", *opts.Domain)

	return url
}

func mailgunState() string {
	client := &http.Client{}

	req, err := http.NewRequest("GET", mailgunEndPoint(), nil)
	if err != nil {
		fmt.Println(err)
	}

	req.SetBasicAuth("api", *opts.Apikey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	type domain struct {
		State string
	}

	type data struct {
		Domain domain
	}

	var d data
	json.Unmarshal(body, &d)

	return d.Domain.State
}

func main() {
	fmt.Println(mailgunState())
}
