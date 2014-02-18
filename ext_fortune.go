/**
 * Description
 *   Display the fortune either from an online source or using 'fortune' command
 *
 * Dependencies:
 *   Needs following packages/commands to be installed:
 *      - fortune
 *      - html2text
 *
 * Configuration:
 *   None
 *
 * Commands:
 *   fortune
 *   fortune me
 *   show my fortune
 *   local fortune
 *   online fortune
 *   short fortune
 *
 * Author:
 *   @vishaltelangre
 */

package main

import (
	"bytes"
	"code.google.com/p/go.net/html"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func init() {
	var s = spec{
		func(inp string) bool {
			lowerInp := strings.ToLower(inp)
			return strings.Contains(lowerInp, "fortune")
		},
		func(inp string) {
			lowerInp := strings.ToLower(inp)
			tellFortune(lowerInp)
		},
	}
	specList = append(specList, s)
}

func tellFortune(inp string) {
	if strings.Contains(inp, "online") || strings.Contains(inp, "live") {
		response, err := http.Get("http://www.shlomifish.org/humour/fortunes/show.cgi?mode=random")
		if err != nil {
			fmt.Println("Online fortune is stuck. Error:", err)
			return
		}
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Something went wrong, try again. Error:", err)
			return
		}
		doc, err := html.Parse(strings.NewReader(string(contents)))
		if err != nil {
			fmt.Println("Something went wrong, try again. Error:", err)
			return
		}
		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "blockquote" || n.Data == "pre" {
				var buf bytes.Buffer
				if err := html.Render(&buf, n); err != nil {
					// should never happen
					panic(err)
				}
				reader := strings.NewReader(buf.String())
				fmt.Println(Text(reader))
				return
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
		f(doc)
	} else if strings.Contains(inp, "short") {
		out, err := exec.Command("fortune", "-s").Output()
		if err != nil {
			fmt.Println("Can't tell your fortune right now!", err)
			return
		}
		fmt.Printf("%s", out)
	} else {
		out, err := exec.Command("fortune").Output()
		if err != nil {
			fmt.Println("Can't tell your fortune right now!", err)
			return
		}
		fmt.Printf("%s", out)
	}
}

func extract(node *html.Node, buff *bytes.Buffer) {
	if node.Type == html.TextNode {
		data := strings.Trim(node.Data, "\r\n ")
		if data != "" {
			buff.WriteString("\n")
			buff.WriteString(data)
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extract(c, buff)
	}
}

func Text(reader io.Reader) string {
	var buffer bytes.Buffer
	doc, err := html.Parse(reader)
	if err != nil {
		log.Fatal(err)
	}
	extract(doc, &buffer)
	return buffer.String()
}
