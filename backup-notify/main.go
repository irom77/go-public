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
	//"text/template"
	"html/template"
	//"encoding/json"
	"bytes"
)

var (
	Version = "No Version Provided"
	BuildTime = ""
)
//backup-notify -v to check version, -h to get help
func main() {
	flag.Usage = func() {
		fmt.Printf("Copyright 2016 @IrekRomaniuk. All rights reserved.\n")
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
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
	//Run program with email i.e. -notify="me@somewhere" , default is to not email
	notify := flag.String("email", "", "Whom to notify")
	flag.Parse()
	if *version {
		fmt.Printf("App Version: %s\nBuild Time : %s\n", Version, BuildTime)
		os.Exit(0)
	}
	computername, _ := os.Hostname()
	fmt.Println(*namecontains, *backupdir, *backupsize, computername, *notify)
	files, err := ioutil.ReadDir(*backupdir)
	if err != nil {
		log.Fatal(err)
	}
	type File struct {
		FileName string
		FileSize int64
	}
	var Files []File
	for _, file := range files {
		if strings.Contains(file.Name(), *namecontains ) {
			Files =  append(Files,File{file.Name(),file.Size()})
			/*if int64(file.Size()) < *backupsize {
				fmt.Printf("File %s size is less than %d\n",
					file.Name(), *backupsize)
			}*/
		}
	}
	/*const tmpl = `
	File : {{.FileName | printf "%40s"}} Size: {{.FileSize | printf "%8d"}}`
	t := template.Must(template.New("file names and sizes").Parse(tmpl))
	for _, f := range Files {
		err := t.Execute(os.Stdout, f)
		if err != nil { panic(err) }
	}*/
	const tmplhtml = `
	<table>
	<tr style='text-align: left'>
  	<th>File</th>
  	<th>Size</th>
	</tr>
	{{range .}}
	<tr>
	<td>file {{.FileName}}</td>
	<td>size {{.FileSize}}</td>
	{{end}}
	</table>
	`
	buf := new(bytes.Buffer)
	t := template.Must(template.New("html table").Parse(tmplhtml))
	err = t.Execute(buf, Files)
	if err != nil { panic(err) }

	//output, _ := json.Marshal(Files)
	// Email me using relay
	if *notify != "" {
		//Template email
		m := gomail.NewMessage()
		m.SetHeader("From", "gopher@" + computername)
		m.SetHeader("To", *notify)
		m.SetHeader("Subject", os.Args[0] + " " + *backupdir)
		m.SetBody("text/html", buf.String())
		//m.SetBody("text/plain", string(output))
		fmt.Printf("\nSending email notification to %s:\n", *notify)
		d := gomail.Dialer{Host: "relay", Port: 25}
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	}
}
