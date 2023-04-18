package main

import (
	"fmt"

	"github.com/gocolly/colly"

	"math/rand"
)

type movie struct {
	title  string
	rating string
	imgURL string
}

// generate_movies returns a list of movie structs
// where the movies are scraped from the IMDB website using colly
func generate_movies() []movie {

	c := colly.NewCollector()

	var movies []movie

	c.OnHTML("tbody.lister-list tr", func(h *colly.HTMLElement) {

		new_movie := movie{
			title:  h.ChildText("a[href]"),
			rating: h.ChildText("td.ratingColumn strong"),
			imgURL: h.ChildAttr("img", "src"),
		}

		movies = append(movies, new_movie)

	})

	c.Visit("https://www.imdb.com/chart/moviemeter/?ref_=nv_mv_mpm")

	return movies

}

func main() {
	var movies []movie = generate_movies()

	var rand_num int = rand.Intn(len(movies))
	var random_movie movie = movies[rand_num]

	fmt.Println(random_movie)

}
