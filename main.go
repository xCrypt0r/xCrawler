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
var users = make([]string, 0)
var c = make(chan []string)

func main() {
	getPages()
}

func getPages() int {
	maxPage := 5

	for p := 1; p <= maxPage; p++ {
		go getUsers(p, c)
	}

	for i := 0; i < maxPage; i++ {
		users = append(users, <-c...)
	}

	fmt.Println(strings.Join(users, "\n"))

	return 0
}

func getUsers(p int, c chan []string) {
	_users := make([]string, 0)

	res, err := http.Get(URL + strconv.Itoa(p))

	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	checkErr(err)

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
