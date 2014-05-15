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
	body     string
}

func NewGrtCmd() grtCmd {
	var cmd grtCmd

	cmd.method = "GET"
	cmd.protocol = "http"
	cmd.domain = "gerrit.dev.returnpath.net"
	cmd.endpoint = "/a/changes/"
	cmd.getVars = map[string]string{
		"q": "status:open",
	}
	cmd.body = ""

	var client http.Client

	resp, err := client.Get(cmd.protocol + "://" + cmd.domain + cmd.endpoint)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(cmd.method, cmd.protocol+"://"+cmd.domain+cmd.endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	user := GetConfigValue("gerrit.user")
	pass := GetConfigValue("gerrit.pass")
	auth := GetAuthorization(user, pass, resp)
	cmd.digest = GetAuthString(auth, req.URL, req.Method, 1)

	return cmd
}

func (this grtCmd) Call() string {
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
