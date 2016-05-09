// Copyright 2016 @IrekRomaniuk. All rights reserved.
/*
	List files in given directory path and name pattern to verify that size is greater than minimum,
	then send notification by email (using 'relay')
	EXAMPLE using default name pattern, minimum size and email address
	$ ./backup-notify -path="/mnt/ftpbackup/ftpuser"
*/
package main

import (
	"fmt"
	//"os"
	"io/ioutil"
	"log"
	"flag"
	"strings"
	"gopkg.in/gomail.v2"
	"crypto/tls"
	//"html/template"
)

//backup-notify
func main() {
	//Run program with path name i.e. -path="C:/" , default is current directory
	backupdir := flag.String("path", ".", "Directory to look for files")
	//Run program with file name i.e. -name="mdsbk.tgz" , default is mdsbk.tgz
	namecontains := flag.String("name", "mdsbk.tgz", "File name or part of")
	//Run program with file size i.e. -size="200000" , default is 9000000
	backupsize := flag.Int64("size", 9000000, "Minimum file size")
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
