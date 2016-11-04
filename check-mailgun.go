package main

import (
	_ "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	Apikey *string `short:"p" long:"apikey" required:"true" description:"Mailgun Api Key"`
	Domain *string `short:"d" long:"domain" required:"true" description:"Mailgun Domain"`
}

type data struct {
	Domain              domain
	ReceivingDnsRecords []receivingDnsRecords `json:"receiving_dns_records"`
	SendingDnsRecords   []sendingDnsRecords   `json:"sending_dns_records"`
}

type domain struct {
	CreatedAt        string `json:"created_at"`
	Name             string
	RequireTls       string `json:"require_tls"`
	SkipVerification string `json:"skip_verification"`
	SmtpLogin        string `json:"smtp_login"`
	SmtpPassword     string `json:"smtp_password"`
	SpamAction       string `json:"spam_action"`
	State            string
	Type             string
	Wildcard         string
}

type receivingDnsRecords struct {
	Priority   string
	RecordType string `json:"record_type"`
	Valid      string
	Value      string
}

type sendingDnsRecords struct {
	Name       string
	RecordType string `json:"record_type"`
	Valid      string
	Value      string
}

func url() string {
	flags.Parse(&opts)
	url := fmt.Sprintf("https://api.mailgun.net/v3/domains/%s", *opts.Domain)

	return url
}

func httpBody() string {
	client := &http.Client{}
	url := url()

	req, err := http.NewRequest("GET", url, nil)
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

	return string(body)
}

func main() {
	body := httpBody()
	fmt.Println(body)

	//jsonStr := json.NewDecoder(body)
	//fmt.Println(body)
	//d := data
	//jsonStr.Decode(&d)
	//fmt.Println("%+v\n", d.Domain)
}
