from bs4 import BeautifulSoup
import requests

url = 'http://forecast.weather.gov/MapClick.php?lat=37.7772&lon=-122.4168'
res = requests.get(url)
page = BeautifulSoup(res.content, 'html.parser')
# seven_day = page.find(id='seven-day-forcast')
# forecast_items = seven_day.find_all(class_='tombstone-container')
# tonight = forecast_items[0]
# print(tonight.prettify())

# page = BeautifulSoup(res.content, 'html.parser')
seven_day = page.find(id="seven-day-forecast")
forecast_items = seven_day.find_all(class_="tombstone-container")
tonight = forecast_items[0]
print(tonight.prettify())
