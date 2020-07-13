package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var URL string = "https://maple.gg/rank/dojang?page="

func main() {
	getPages()
}

func getPages() int {
	users := make([]string, 0)
	c := make(chan []string)

	for p := 1; p <= 5; p++ {
		res, err := http.Get(URL + strconv.Itoa(p))

		checkErr(err)
		checkCode(res)

		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)

		checkErr(err)

		go getUsers(doc, c)

		users = append(users, <-c...)
	}

	fmt.Println(strings.Join(users, "\n"))

	return 0
}

func getUsers(doc *goquery.Document, c chan []string) {
	_users := make([]string, 0)

	doc.Find("span .text-grape-fruit").Each(func(i int, s *goquery.Selection) {
		_users = append(_users, s.Text())
	})

	c <- _users
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
