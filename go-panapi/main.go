package main

import (
	"net/http"
	"crypto/tls"
	"io"
	"os"
	"log"
	"flag"
	"fmt"
)

var (
	URL = flag.String("api", "https://golang.org", "API URL") // or os.Getenv("USER") or os.Getenv("USERNAME")
)
/*
go-panapi -api="https://10.34.2.149/api/
 */
func init() { flag.Parse() }

func main() {
	flag.Usage = func() {
		fmt.Printf("Copyright 2016 @IrekRomaniuk. All rights reserved.\n")
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	fmt.Println(*URL)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(*URL)
	if err != nil {
		log.Fatal(err)
	} else {
		defer resp.Body.Close()
		_, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
}
