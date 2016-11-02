package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	Apikey *string `short:"p" long:"apikey" required:"true" description:"Mailgun Api Key"`
	Domain *string `short:"d" long:"domain" required:"true" description:"Mailgun Domain"`
}

func main() {
	flags.Parse(&opts)

	client := &http.Client{}
	url := fmt.Sprintf("https://api.mailgun.net/v3/domains/%s", *opts.Domain)
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("api", *opts.Apikey)
	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	fmt.Println(string(body))
}
