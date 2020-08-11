#include <iostream>
#include <sstream>
#include <regex>
#include <vector>
#include <thread>
#include <algorithm>
#include <curl/curl.h>

using namespace std;

const char* URL = "https://maple.gg/rank/dojang?page=";

struct User
{
    int rank;
    string nickname;
    string server;
    string level;
};

int getPages();
int getUsers(int page, vector<User>& users);
bool sortByRank(const User& x, const User& y);
size_t WriteCallback(void *contents, size_t size, size_t nmemb, void *userp);

int main()
{
    getPages();

    return 0;
}

int getPages()
{
    vector<User> users;
    vector<thread> threads;
    int maxPage = 5;

    for (int page = 1; page <= maxPage; ++page)
    {
        threads.emplace_back(thread(getUsers, page, ref(users)));
    }

    for (auto& th : threads)
    {
        th.join();
    }

    sort(users.begin(), users.end(), sortByRank);

    for (const auto& user : users)
    {
        cout << "[" << user.server << "] " << user.nickname << " " << user.level << endl;
    }

    return 0;
}

int getUsers(int page, vector<User>& users)
{
    CURL *curl = curl_easy_init();
    CURLcode res;
    string readBuffer;
    stringstream url;

    if (!curl)
    {
        cout << "Curl initialization failure" << endl;

        return 128;
    }

    url << ::URL << page;
    curl_easy_setopt(curl, CURLOPT_URL, url.str().c_str());
    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, WriteCallback);
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, &readBuffer);

    res = curl_easy_perform(curl);

    if (res != CURLE_OK)
    {
        cerr << "Error during curl request: " << curl_easy_strerror(res) << endl;
    }

    regex rgx("/world/ico_world_.+\"(.+)\">.+>(.+)</a></span>[^]+?font-size-14\">(.+)<", regex::icase);
    smatch match;
    int i = 0;

    while (regex_search(readBuffer, match, rgx))
    {
        User user = { (page - 1) * 20 + (i++ + 1), match.str(2), match.str(1), match.str(3) };

        users.push_back(user);

        readBuffer = match.suffix();
    }

    curl_easy_cleanup(curl);
    curl_global_cleanup();

    return 0;
}

bool sortByRank(const User& x, const User& y)
{
    return x.rank < y.rank;
}

size_t WriteCallback(void *contents, size_t size, size_t nmemb, void *userp)
{
    ((string*)userp)->append((char*)contents, size * nmemb);

    return size * nmemb;
}
