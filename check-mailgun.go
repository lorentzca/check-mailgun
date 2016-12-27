package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/mackerelio/checkers"
)

var opts struct {
	Apikey *string `short:"p" long:"apikey" required:"true" description:"Mailgun Api Key"`
	Domain *string `short:"d" long:"domain" required:"true" description:"Mailgun Domain"`
}

type domain struct {
	State string
}

type data struct {
	Domain domain
}

func mailgunEndPoint() string {
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}
	url := fmt.Sprintf("https://api.mailgun.net/v3/domains/%s", *opts.Domain)

	return url
}

func getMailgunState() string {
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

	var d data
	json.Unmarshal(body, &d)

	return d.Domain.State
}

func run() *checkers.Checker {
	st := getMailgunState()

	checkSt := checkers.OK
	if st != "active" {
		checkSt = checkers.CRITICAL
		msg := fmt.Sprintf("%s is dead\n", *opts.Domain)
		return checkers.NewChecker(checkSt, msg)
	}

	msg := fmt.Sprintf("%s is %s\n", *opts.Domain, st)
	return checkers.NewChecker(checkSt, msg)
}

func main() {
	ckr := run()
	ckr.Name = "State"
	ckr.Exit()
}
