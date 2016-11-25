package maintable

import (
	"net/http"
	"crypto/tls"
	"log"
	"golang.org/x/net/html"
	"strings"
	"fmt"
)

func main() {
	const CMD = "&cmd=<show><running><resource-monitor><second></second></resource-monitor></running></show>"
	const API = "&key="
	const IP = "https://10.34.2.21/esp/restapi.esp?type=op"

	URL := IP + CMD + API

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	/* else {
		defer resp.Body.Close()
		_, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}*/
	b :=resp.Body
	defer b.Close() // close Body when the function returns
	//parseTags("dp2", b)
	//parseText(b)
	//parseCascadia("value", b)
	//parseGoQuery(b)
	z := html.NewTokenizer(b)
	content := []string{}
	for z.Token().Data != "dp2" {
		tt := z.Next()
		if tt == html.StartTagToken {
			t := z.Token()
			if t.Data == "value" {
				inner := z.Next()
				if inner == html.TextToken {
					text := (string)(z.Text())
					t := strings.TrimSpace(text)
					content = append(content, t)
				}
			}
		}
	}
	// Print to check the slice's content
	fmt.Println(content)

}