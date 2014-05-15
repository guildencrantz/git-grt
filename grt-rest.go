package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type grtCmd struct {
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

	return cmd
}

func (this grtCmd) Call() string {
	var client http.Client
	getQry := ""

	for k, v := range this.getVars {
		getQry += k + "=" + v + "&"
	}

	getQry = getQry[:len(getQry)-1]

	resp, err := client.Get(this.protocol + "://" + this.domain + this.endpoint + getQry)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", "http://gerrit.dev.returnpath.net/a/changes/?q=status:open", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	user := GetConfigValue("gerrit.user")
	pass := GetConfigValue("gerrit.pass")

	auth := GetAuthorization(user, pass, resp)
	digest := GetAuthString(auth, req.URL, req.Method, 1)
	fmt.Println(digest)
	req.Header.Add("Authorization", digest)

	resp, err = client.Do(req)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
