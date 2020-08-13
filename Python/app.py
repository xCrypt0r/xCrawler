#!/usr/bin/python3

import requests
import re
from collections import namedtuple
from operator import attrgetter
from threading import Thread

URL = 'https://maple.gg/rank/dojang?page='
User = namedtuple('User', 'rank nickname server level')
Users = []

def main():
	get_pages()

def get_pages():
	threads = []
	max_page = 5

	for page in range(1, max_page + 1):
		th = Thread(target=get_users, args=(page,))
		
		th.start()
		threads.append(th)

	for th in threads:
		th.join()


	for user in sorted(Users, key=attrgetter('rank')):
		print [', user.server, '] ', user.nickname, ' ', user.level

def get_users(page):
	rgx = '/world/ico_world_.+"(.+)">.+>(.+)</a></span>[^@]+?font-size-14">(.+)<'
	r = requests.get(URL + str(page))
	i = 0
	
	for (servers, nicknames, levels) in re.findall(rgx, r.text):
		user = User(
			rank=(page - 1) * 20 + (i + 1),
			nickname=nicknames,
			server=servers,
			level=levels
		)

		i += 1
		Users.append(user)
		
if __name__ == "__main__":
	main()
