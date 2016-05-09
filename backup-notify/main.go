// Copyright 2016 @IrekRomaniuk. All rights reserved.
/*
	List files in given directory path and name pattern to verify that size is greater than minimum,
	then send notification by email (using 'relay')
	EXAMPLE using default name pattern, minimum size and email address
	$ ./backup-notify -h
	$ ./backup-notify -path="/mnt/ftpbackup/ftpuser" -name=txt
*/
package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"flag"
	"strings"
	"gopkg.in/gomail.v2"
	"crypto/tls"
	//"html/template"
)

var (
	Version = "No Version Provided"
	BuildTime = ""
)
//backup-notify -v to check version
func main() {
	/*flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("Example\n")
		flag.PrintDefaults()
	}*/
	//go run -ldflags "-X main.Version=1.0.1" github.com\irom77\go-public\backup-notify
	/*go build -ldflags "-X main.BuildTime=`date -u +.%Y%m%d.%H%M%S` -X main.Version=1.0.1"
	github.com\irom77\go-public\backup-notify*/
	//Run program with path name i.e. -path="C:/" , default is current directory
	backupdir := flag.String("path", ".", "Directory to look for files")
	//Run program with file name i.e. -name="mdsbk.tgz" , default is mdsbk.tgz
	namecontains := flag.String("name", "mdsbk.tgz", "File name or part of it")
	//Run program with file size i.e. -size="200000" , default is 9000000
	backupsize := flag.Int64("size", 9000000, "Minimum file size")
	version := flag.Bool("v", false, "Prints current version")
	flag.Parse()
	if *version {
		fmt.Printf("App Version: %s\nBuild Time : %s\n", Version, BuildTime)
		os.Exit(0)
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "gopher@mycomputer")
	m.SetHeader("To", "iromaniuk@commonwealth.com")
	m.SetHeader("Subject", "backup-notify")
	m.SetBody("text/html", "Hello <b>me</b>!")
	m.SetBody("text/plain", "Hello!")
	flag.Parse();
	fmt.Println(*namecontains, *backupdir, *backupsize)
	files, err := ioutil.ReadDir(*backupdir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), *namecontains ) {
			fmt.Println(file.Name(), file.Size())
			if int64(file.Size()) < *backupsize {
				fmt.Printf("File %s size is less than %d",
					file.Name(), *backupsize)
			}
		}

	}

	// Send the email to me using relay
	d := gomail.Dialer{Host: "relay", Port: 25}
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
