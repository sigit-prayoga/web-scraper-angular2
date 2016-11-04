package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rs/cors"
	goji "goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
)

func main() {
	//initiate new multiplexer
	mux := goji.NewMux()
	//register all handler for each end point
	mux.HandleFuncC(pat.Get("/go/api/scraping"), getHandler())
	//to allow cross origin
	handler := cors.Default().Handler(mux)
	//finally, listen and serve in designated host and port
	http.ListenAndServe(":3001", handler)
}

func getHandler() goji.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, req *http.Request) {
		//get url from query params
		keyword := req.URL.Query().Get("keyword")
		//create a container for all info
		categories := scraping(keyword)
		//ready to return the list
		json, error := json.Marshal(categories)
		if error != nil {
			//What's wrong, stop it!
			log.Fatal(error)
		}
		//prepare for the response
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		//send OK status code through header, which is 200 in case you wonder
		w.WriteHeader(http.StatusOK)
		//write the list into json as the resonse body
		w.Write(json)
	}
}

func scraping(keyword string) []string {
	fmt.Printf("Scraping... : %s \n", keyword)
	//change whitespace into '+' sign, we might have a better way to do it later
	keyword = strings.Replace(keyword, " ", "+", -1)
	//get the HTML DOM
	url := "https://www.bukalapak.com/products?utf8=%E2%9C%93&source=navbar&from=omnisearch&search_source=omnisearch_organic&search%5Bkeywords%5D=" + keyword
	doc, err := goquery.NewDocument(url)
	if err != nil {
		//continue
		log.Fatal(err)
	}

	var categories []string
	max := 2
	//happy scraping
	doc.Find("ul.js-tree.tree").EachWithBreak(func(idx int, el *goquery.Selection) bool {
		el.ChildrenFiltered("li").EachWithBreak(func(index int, s *goquery.Selection) bool {
			category := s.ChildrenFiltered("a").First().Text()
			categories = append(categories, category)
			if len(categories) >= max {
				//break loop
				return false
			}
			//keep going
			return true
		})

		if len(categories) >= max {
			//break loop
			return false
		}

		//keep going
		return true
	})

	return categories
}
