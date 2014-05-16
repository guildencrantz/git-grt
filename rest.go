package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type grtCmd struct {
	digest   string
	method   string
	protocol string
	domain   string
	endpoint string
	getVars  map[string]string
	rawForm  string
	body     string
}

func NewGrtCmd(method, endpoint string) *grtCmd {
	cmd := &grtCmd{
		method:   method,
		protocol: "http",
		domain:   "gerrit.dev.returnpath.net",
		endpoint: endpoint,
	}

	cmd.SetDigest()

	return cmd
}

func (this *grtCmd) GetDigest() string {
	var client http.Client

	resp, err := client.Get(this.protocol + "://" + this.domain + this.endpoint)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(this.method, this.protocol+"://"+this.domain+this.endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	user := GetConfigValue("gerrit.user")
	pass := GetConfigValue("gerrit.pass")
	auth := GetAuthorization(user, pass, resp)

	return GetAuthString(auth, req.URL, req.Method, 1)
}

func (this *grtCmd) SetDigest() {
	this.digest = this.GetDigest()
}

func (this *grtCmd) Call() string {
	var form string
	switch {
	case this.rawForm != "":
		form = "?" + this.rawForm
	case len(this.getVars) > 0:
		form = "?"
		for k, v := range this.getVars {
			form += k + "=" + v + "&"
		}
		form = form[:len(form)-1] // Trim trailing ampersand
	}

	var client http.Client
	resp, err := client.Get(this.protocol + "://" + this.domain + this.endpoint + form)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(this.method, this.protocol+"://"+this.domain+this.endpoint+form, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", this.digest)

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: Need to check the return status.
	// If, for example, you get Unauthorized the slice below will be out of bounds.

	val := string(body)
	val = val[strings.Index(string(val), "\n"):]

	return val
}
