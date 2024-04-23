package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"golang.org/x/net/html"
)


var (
	_, b, _, _  = runtime.Caller(0)
	ProjectRoot = filepath.Dir(b)
	FILENAME = filepath.Join(ProjectRoot, "entries.html")
)

func main() {

	// Get the <body> node
	doc := GetRoot()
	elemBody:= GetNodeWithData(doc, "body")
	_ = elemBody
}

func GetRoot() *html.Node {

	// Open the HTML input file
	fp, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	// Parse the HTML into a tree of HTML nodes
	doc, err := html.Parse(fp)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func GetNodeWithData(node *html.Node, data string) *html.Node {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Data == data {
			return child
		}
	}
	return nil
}
