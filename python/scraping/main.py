from bs4 import BeautifulSoup
import requests

res = requests.get(
    "http://dataquestio.github.io/web-scraping-pages/ids_and_classes.html")

# print(f'status code: {res.status_code}')
# print(res.content)

# Parsing the page
page = BeautifulSoup(res.content, 'html.parser')

# find element by class and id
p = page.find_all('p', class_='outer-text')
# print(len(p))
print(p[0].prettify())

first = page.find(id='first')
print(first.text)

# find element using css selector


########################
# print(html)
# print(html.prettify())

# elements = list(html.children)
# print(elements)
# print(elements[0])

# # find()
# h1 = html.find('h1')
# # print(h1)
# # print(h1.text)

# # find_all()
# h3 = html.find_all('h3')
# # print(h2)
# # print(h2[0])
# for news in h3:
#     print(news)
# children = list(page.children)
# print(len(children))
# print(children[2])

# cls0 = [type(item) for item in list(page.children)]
# print(cls0)
