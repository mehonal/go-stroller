package main

import (
	"fmt"
	"os"
	"github.com/gocolly/colly"
)

func scrapeWebResults(keyword string) (titles []string, descriptions []string, err error){
    var url string = "https://www.google.com/search?q=" + keyword + "&num=" + fmt.Sprintf("%d", 100) 
    c := colly.NewCollector(
        colly.AllowedDomains("google.com"),
        colly.AllowedDomains("www.google.com"),
    )
    c.OnHTML(".egMi0 h3", func(e *colly.HTMLElement) { // all titles; other useful ones may be: kCrYT and egMi0
        fmt.Print("Title: " + e.Text + "\n")
        if (len(titles) < 100) {
            titles = append(titles, e.Text)
        }
    })
    c.OnHTML(".BNeawe.s3v9rd.AP7Wnd", func(e *colly.HTMLElement) { // all descriptions
        fmt.Print("Desc: " + e.Text + "\n") // e.DOM.Html() if HTML is needed
        if (len(descriptions) < 100) {
            descriptions = append(descriptions, e.Text)
        }
    })
    err = c.Visit(url)
    if err != nil {
        fmt.Print(err)
    }
    return titles, descriptions, err
}

func main() {
    fmt.Print("Running Stroller...")
    titles, descriptions, err := scrapeWebResults("golang")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Titles: %v\n", titles)
    fmt.Printf("Descriptions: %v\n", descriptions)
    fmt.Printf("Environment variable: %s\n", os.Getenv("TEST"))
}
