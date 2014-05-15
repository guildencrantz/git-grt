package main

import (
	"fmt"
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
	var client http.Client
	getQry := ""

	for k, v := range this.getVars {
		getQry += k + "=" + v + "&"
	}

	if len(getQry) > 0 {
		getQry = "?" + getQry[:len(getQry)-1]
	}

	resp, err := client.Get(this.protocol + "://" + this.domain + this.endpoint + getQry)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(this.method, this.protocol+"://"+this.domain+this.endpoint+getQry, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(this.digest)
	req.Header.Add("Authorization", this.digest)

	resp, err = client.Do(req)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	val := string(body)
	val = val[strings.Index(string(val), "\n"):]

	return val
}
