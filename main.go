package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type user struct {
	rank     int
	nickname string
	server   string
	level    string
}

var URL string = "https://maple.gg/rank/dojang?page="
var users = make([]user, 0)
var c = make(chan []user)

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

	sort.SliceStable(users, func(x, y int) bool {
		return users[x].rank < users[y].rank
	})

	for _, user := range users {
		fmt.Printf("[%s] %s %s\n", user.server, user.nickname, user.level)
	}

	return 0
}

func getUsers(p int, c chan []user) {
	_users := make([]user, 0)

	res, err := http.Get(URL + strconv.Itoa(p))

	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	checkErr(err)

	doc.Find("td.align-middle").Not(".d-none").Each(func(i int, s *goquery.Selection) {
		nickname := s.Find(".text-grape-fruit").Text()
		server, _ := s.Find("div.d-inline-block img").Eq(1).Attr("alt")
		level := s.Find(".font-size-14").Eq(0).Text()

		_users = append(_users, user{
			rank:     (p-1)*20 + i + 1,
			nickname: nickname,
			server:   server,
			level:    level,
		})
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
