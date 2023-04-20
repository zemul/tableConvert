package main

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/atotto/clipboard"
	"os"
	"strings"
)

func main() {
	fmt.Println("Paste your HTML Table data  here ")
	var md strings.Builder
	var html strings.Builder
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		html.WriteString(input)
		if strings.Contains(input, "</table>") {
			break
		}
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html.String()))
	if err != nil {
		panic(err)
	}

	doc.Find("thead tr th").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		md.WriteString("| ")
		md.WriteString(text)
		md.WriteString(" ")
	})
	md.WriteString("|\n")

	doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		s.Find("td").Each(func(j int, s2 *goquery.Selection) {
			text := s2.Text()
			md.WriteString("| ")
			md.WriteString(text)
			md.WriteString(" ")
		})
		md.WriteString("|\n")
	})

	fmt.Println()
	fmt.Println(md.String())
	err = clipboard.WriteAll(md.String())
	if err != nil {
		return
	}
	fmt.Println("🍺 write clipboard success")
}
