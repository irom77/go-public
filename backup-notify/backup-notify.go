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
//List files with given name pattern and directory path to verify that size is greater than minimum
//Run i.e. $ ./backup-notify -path="/mnt/ftpbackup/ftpuser"
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
