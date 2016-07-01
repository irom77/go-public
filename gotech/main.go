package main

import (
	"log"
	"os/exec"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"gopkg.in/gomail.v2"
	"bytes"
	"flag"
)
var (
	Version = "No Version Provided"
	BuildTime = ""
)
//gotech.exe -username="irek" -password="********"
func main() {
	flag.Usage = func() {
		fmt.Printf("Copyright 2016 @IrekRomaniuk. All rights reserved.\n")
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	user := flag.String("username", ".", "Username")
	pass := flag.String("password", ".", "Password")
	notify := flag.String("email", "", "Who to send result to")
	addr := flag.String("address", "8.8.8.9", "Destination to reach")
	flag.Parse()
	fmt.Println(*user, *pass, *notify, *addr)
	//api := "https://pan1/esp/restapi.esp?type=op&cmd=%3Crequest%3E%3Ctech-support%3E%3Cdump%3E%3C/dump%3E%3C/tech-support%3E%3C/request%3E&key="
	//key := "LUFRPT1jRVRDTmo1VVpCZ2wwa3hCU1Roc1pWUVh0VTA9QU5jREpOWVFCaFBXbW5xZ214UU9zQT09"
	// send two pings and send the ouput to STDOUT
	count := 0

	for {
		output, err := exec.Command("ping", "-n", "2", *addr).Output()
		if err != nil {
			//log.Fatal(err)
			fmt.Println("Host not reachable\n")
			if *notify != "" {
				send_email_alert("Host not reachable", *notify)
			}
			//tech_support(api,key)
			tech_cisco(*user,*pass)
			os.Exit(1)
		} else {
			count += 1
			log.Printf("Ping count %d ------>  %s\n", count, output)
		}
	}

}
func send_email_alert(subj, addr string) {
	computername, _ := os.Hostname()
	m := gomail.NewMessage()
	m.SetHeader("From", "gopher@" + computername)
	m.SetHeader("To", addr)
	m.SetHeader("Subject", subj)
	m.SetBody("text/html", "")
	fmt.Printf("\nSending email notification to %s:\n", addr)
	d := gomail.Dialer{Host: "relay", Port: 25}
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func tech_support(api, key string) {
	conf := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}
	client := &http.Client{Transport: conf}
	resp, err := client.Get(api + key)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(1)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//fmt.Println(os.Stdout, string(body))
		fmt.Printf("%s\n", string(body))
	}
}

func tech_cisco(userid, password string) {
	conf := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}
	url := "http://10.34.1.11/ins"
	var payload = []byte(`{"ins_api":
		{
    	"version": "1.0",
    	"type": "cli_show",
    	"chunk": "0",
    	"sid": "1",
    	"input": "show version",
    	"output_format": "json"
  		}
  	}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.SetBasicAuth(userid,password)
	//resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Transport: conf}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}