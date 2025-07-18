package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func extractXpathPredicate(xpathExpression string) int {
	start := strings.Index(xpathExpression, "[")
	end := strings.Index(xpathExpression, "]")

	if start != -1 && end != -1 && start < end {
		predicate, err := strconv.Atoi(xpathExpression[start+1 : end])
		if err != nil {
			return -1
		}
		return predicate
	}

	return -1
}

func extractXpathNodeName(xpathExpression string) string {
	end := strings.Index(xpathExpression, "[")

	if end == -1 {
		return xpathExpression
	}

	return xpathExpression[:end]
}

func getElementByXpath(node *html.Node, xpath string) *html.Node {
	if xpath[0] == '/' {
		xpath = xpath[1:]
	}

	xpathElements := strings.SplitSeq(xpath, "/")

	resultElement := node

	for xpathElement := range xpathElements {
		var didFindElement bool = false

		xpathNodeName := extractXpathNodeName(xpathElement)
		count := 1
		xpathPredicate := extractXpathPredicate(xpathElement)

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			if child.Type != html.ElementNode {
				continue
			}

			if child.Data == xpathNodeName {

				if xpathPredicate == -1 || count == xpathPredicate {
					node = child
					didFindElement = true
					resultElement = node
					break
				}

				count++
			}
		}

		if !didFindElement {
			fmt.Println("could not find: " + xpathElement)
			return nil
		}
	}

	return resultElement
}

func main() {
	resp, err := http.Get("https://scrapeme.live/shop/")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		log.Fatalln(err)
	}

	element := getElementByXpath(doc, "/html/body/div[1]/div[2]/div/div[2]/main/ul/li[3]/a[1]/h2")
	fmt.Println(element.FirstChild.Data)
}
