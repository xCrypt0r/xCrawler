const { promisify } = require('util');
const request = promisify(require('request'));
const cheerio = require('cheerio');
const URL = 'https://maple.gg/rank/dojang?page=';

function main() {
    getPages();
}

function getPages() {
    let users = [],
        maxPage = 5,
        promises = [];

    for (let page = 1; page <= maxPage; page++) {
        promises.push(getUsers(page, users));
    }

    Promise.all(promises)
        .then(() => {
            users
                .sort((x, y) => x.rank - y.rank)
                .forEach(user => {
                    console.log('[%s] %s %s', user.server, user.nickname, user.level);
                });
        });
}

async function getUsers(page, users) {
    let { body } = await request(URL + page),
        $ = cheerio.load(body);

    $('td.align-middle').not('.d-none').each((i, el) => {
        let it = $(el),
            nickname = it.find('.text-grape-fruit').text(),
            server = it.find('div.d-inline-block img').eq(1).attr('alt'),
            level = it.find('.font-size-14').eq(0).text();

        users.push({
            rank: (page - 1) * 20 + i + 1,
            nickname,
            server,
            level
        });
    });
}

main();
