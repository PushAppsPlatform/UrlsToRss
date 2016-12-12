package main

import (
	"time"
	"github.com/pushapps/urlstorss/feeds"
	"log"
	"fmt"
	"github.com/pushapps/urlstorss/helpers"
)

func main()  {

	articles := []string{"http://noticieros.televisa.com/economia/2016-12-10/operativo-rampa-revisara-salud-pilotos-aeropuertos-pais/",
		"http://noticieros.televisa.com/economia/2016-12-09/malasia-petronas-operara-aguas-profundas-mexico/",
		"http://noticieros.televisa.com/economia/2016-12-08/mezcla-mexicana-avanza-49-centavos-barril-cierra-43-81/",
		"http://noticieros.televisa.com/economia/2016-12-09/dolar-cierra-semana-ligera-alza/",
		"http://noticieros.televisa.com/economia/2016-12-08/crece-economia-estados-industriales-caen-petroleros-inegi/",
		"http://noticieros.televisa.com/economia/2016-12-08/precios-consumidor-suben-0-78-ciento-noviembre-inegi/",
		"http://noticieros.televisa.com/economia/2016-12-09/juncker-lanza-advertencia-paises-bloque-quieran-ir-libre/",
		"http://noticieros.televisa.com/economia/2016-11-10/juncker-pide-trump-aclare-posturas-comercio-clima-otan/",
		"http://noticieros.televisa.com/economia/2016-12-09/china-fija-principales-tareas-economicas-2017/",}

	now := time.Now()
	feed := &feeds.Feed{
		Title:       "Testing Go RSS creator",
		Link:        &feeds.Link{Href: "http://noticieros.televisa.com/"},
		Description: "Testing Go RSS creator - Description",
		Created:     now,
	}

	feed.Items = []*feeds.Item{}
	for i := 0; i < len(articles); i++ {
		feed.Items = append(feed.Items, helpers.UrlToRssItem(articles[i]))
	}

	atom, err := feed.ToAtom()
	if err != nil {
		log.Fatal(err)
	}

	rss, err := feed.ToRss()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rss)
	fmt.Println("\n\n\n\n")
	fmt.Println(atom)

}
