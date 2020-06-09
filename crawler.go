package main

import (
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

// Product from mercadolivre
type Product struct {
	Name  string `json:"name"`
	Link  string `json:"link"`
	Price string `json:"price"`
	Store string `json:"store"`
	State string `json:"state"`
}

func searchProduct(search *Search) []Product {
	var result = make(chan []Product)

	go extractData(search, result)

	return <-result
}

func extractData(search *Search, result chan []Product) {
	var products = make([]Product, 0)
	var c = colly.NewCollector()

	c.OnHTML("#searchResults", func(e *colly.HTMLElement) {
		e.ForEachWithBreak(".results-item", func(_ int, elem *colly.HTMLElement) bool {
			if len(products) >= search.Limit {
				return false
			}

			product := Product{
				Name:  getProductName(elem),
				Link:  getProductLink(elem),
				Price: getProductPrice(elem),
				Store: getProductStore(elem),
				State: getProductState(elem),
			}

			products = append(products, product)

			return true
		})

		result <- products
	})

	c.OnError(func(response *colly.Response, err error) {
		log.Fatalln("[main.searchProduct]", err)
	})

	c.Visit(os.Getenv("URL_MERCADOLIVRE_SEARCH_ITEM") + search.Text)
}

func getProductName(elem *colly.HTMLElement) string {
	return elem.ChildText(".main-title")
}

func getProductLink(elem *colly.HTMLElement) string {
	href := elem.ChildAttr(".item__info-link", "href")

	if strings.TrimSpace(href) == "" {
		href = elem.ChildAttr(".item__info-title", "href")
	}

	return href
}

func getProductPrice(elem *colly.HTMLElement) string {
	price := elem.ChildText(".item__info-link > span")
	price = strings.Replace(price, "R$", "", -1)
	return getProductPriceWithDecimals(price, elem)
}

func getProductPriceWithDecimals(price string, elem *colly.HTMLElement) string {
	if strings.TrimSpace(price) == "" {
		price = elem.ChildText(".price__fraction")
		decimals := elem.ChildText(".price__decimals")

		if strings.TrimSpace(decimals) != "" {
			price = price + "." + decimals
		}
	}

	return price
}

func getProductStore(elem *colly.HTMLElement) string {
	return elem.ChildText(".item__brand-title-tos")
}

func getProductState(elem *colly.HTMLElement) string {
	state := elem.ChildText(".item__status > .item__condition")

	if strings.TrimSpace(state) == "" {
		state = elem.ChildText(".item__status > .item__title').text().trim() || html.find('a > div > div.item__title")
	}

	return state
}
