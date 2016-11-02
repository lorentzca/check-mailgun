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
