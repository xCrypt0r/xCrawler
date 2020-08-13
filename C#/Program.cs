using System;
using System.Net;
using System.Text;
using System.Text.RegularExpressions;
using System.Collections;
using System.Collections.Generic;
using System.Threading;
using System.Linq;

namespace MaplestoryScrapper {
    struct User {
        public int Rank;
        public string Nickname;
        public string Server;
        public string Level;
    }

    class Program {
        public static readonly string URL = @"https://maple.gg/rank/dojang?page=";

        public static void Main(string[] args) {
            GetPages();
        }

        public static void GetPages() {
            List<User> users = new List<User>();
            List<Thread> threads = new List<Thread>();
            int maxPage = 5;

            for (int page = 1; page <= maxPage; page++) {
                Thread thread = new Thread(delegate() {
                    GetUsers(page, ref users);
                });

                thread.Start();
                threads.Add(thread);
            }

            foreach (Thread th in threads) {
                th.Join();
            }

            users = (from user in users
                    orderby user.Rank
                    select user).ToList();

            foreach (User user in users) {
                Console.WriteLine("[{0}] {1} {2}", user.Nickname, user.Server, user.Level);
            }
        }

        public static void GetUsers(int page, ref List<User> users) {
            string html = GetHtml(Program.URL + page);
            Regex rgx = new Regex("/world/ico_world_.+\"(.+)\">.+>(.+)</a></span>[^@]+?font-size-14\">(.+)<", RegexOptions.IgnoreCase);
            MatchCollection match = rgx.Matches(html);
            int i = 0;

            foreach (Match m in match) {
                User user = new User() {
                    Rank = (page - 1) * 20 + (i++ + 1),
                    Nickname = m.Groups[1].Value,
                    Server = m.Groups[2].Value,
                    Level = m.Groups[3].Value
                };

                users.Add(user);
            }
        }

        public static string GetHtml(string url) {
            WebClient client = new WebClient();

            return client.DownloadString(url);
        }
    }
}
