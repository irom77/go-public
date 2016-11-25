package archive

import (
	"fmt"
	"io"
	"github.com/andybalholm/cascadia"
	"os"
	"strings"
	"golang.org/x/net/html"
	"log"
)

func parseSimple(tag string, b io.Reader) {
	z := html.NewTokenizer(b)
	content := []string{}
	for z.Token().Data != tag {
		tt := z.Next()
		if tt == html.StartTagToken {
			t := z.Token()
			if t.Data == "value" {
				inner := z.Next()
				if inner == html.TextToken {
					text := (string)(z.Text())
					t := tag + "," + strings.TrimSpace(text)
					content = append(content, t)
				}
			}
		}
	}
	// Print to check the slice's content
	fmt.Println(content,len(content))
}

func parseTags (tag string, b io.Reader) {
	doc, err := html.Parse(b)
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tag {
			fmt.Println(n.Data)
			fmt.Println(n.FirstChild.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

func parseText(b io.Reader) {  //after http://golang-examples.tumblr.com/post/47426518779/parse-html
	d := html.NewTokenizer(b)
	for {
		// token type
		tokenType := d.Next()
		if tokenType == html.ErrorToken {
			return
		}
		token := d.Token()
		switch tokenType {
		case html.StartTagToken: // <tag>
		// type Token struct {
		//     Type     TokenType
		//     DataAtom atom.Atom
		//     Data     string
		//     Attr     []Attribute
		// }
		//
		// type Attribute struct {
		//     Namespace, Key, Val string
		// }
		case html.TextToken: // text between start and end tag
			fmt.Println(token.Data)
		case html.EndTagToken: // </tag>
		case html.SelfClosingTagToken: // <tag/>

		}
	}
}

func parseCascadia(tag string, b io.Reader) {
	doc, err := html.Parse(b)
	if err != nil {
		log.Fatal(err)
	}

	body := cascadia.MustCompile(tag).MatchFirst(doc)
	html.Render(os.Stdout, body)
	//fmt.Println(body)
}
