package main

import (
   "fmt"
   "time"
   "github.com/gocolly/colly"
)
type Course struct {
   Name     string
   Rating    string
   University string
}
func main() {
   c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)

   courses := make([]Course, 0)
   c.OnHTML("li.cds-9.css-0.cds-11.cds-grid-item.cds-56.cds-64.cds-76 div.css-ilhc4l", func(e *colly.HTMLElement){
	item := Course{}
	e.ForEach("h2.cds-119.css-h1jogs.cds-121", func(i int, h *colly.HTMLElement) {
		
		item.Name = h.Text
		item.Rating = e.ChildText("div.css-pn23ng")
		item.University = e.ChildText("div.cds-9.css-1xdhyk6.cds-11.cds-grid-item")
		// fmt.Println(item)
		c.SetRequestTimeout(120 * time.Second)
		courses = append(courses, item)
	})
   })

    c.OnRequest(func(r *colly.Request) {
       fmt.Println("Visiting", r.URL)
   })

   c.OnResponse(func(r *colly.Response) {
       fmt.Println("Got a response from", r.Request.URL)
   })

   c.OnError(func(r *colly.Response, e error) {
       fmt.Println("Got this error:", e)
   })

   c.Visit("https://www.coursera.org/courses?query=python")

   for _,val:=range courses{
	fmt.Println("--------------------------------")
	fmt.Println("Course name -",val.Name,"\nProvided by -",val.University,"\nRatings",val.Rating)
	
   }
}