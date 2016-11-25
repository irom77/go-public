package table

import (
	"strings"
	"fmt"
	"golang.org/x/net/html"
)
//https://codedump.io/share/nKGnx8Fkvn6w/1/how-to-convert-html-table-to-array-with-golang
var body = strings.NewReader(`
        <html>
        <body>
        <table>
        <tr>
        <td>Row 1, Content 1<td>
        <td>Row 1, Content 2<td>
        <td>Row 1, Content 3<td>
        <td>Row 1, Content 4<td>
        </tr>
        <tr>
        <td>Row 2, Content 1<td>
        <td>Row 2, Content 2<td>
        <td>Row 2, Content 3<td>
        <td>Row 2, Content 4<td>
        </tr>
        </table>
        </body>
        </html>`)

func main() {
	z := html.NewTokenizer(body)
	content := []string{}
	// While have not hit the </html> tag
	for z.Token().Data != "html" {
		tt := z.Next()
		if tt == html.StartTagToken {
			t := z.Token()
			if t.Data == "td" {
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

