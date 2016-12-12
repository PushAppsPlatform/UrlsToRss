# URLs To RSS

## Intro
This project was created from the need to create custom RSS feeds from some URLs. The goal was the enable an input of array of strings (each is a URL to an article), extract the meta data from each article and create a custom RSS feed from it.

## Installation

`go get github.com/pushapps/urlstorss`

## Usage
    // Your articles to be put in the RSS feed
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
	
	// Create your RSS feed
	feed := &feeds.Feed{
		Title:       "Testing Go RSS creator",
		Link:        &feeds.Link{Href: "http://noticieros.televisa.com/"},
		Description: "Testing Go RSS creator - Description",
		Created:     now,
	}

	// Add the items
	feed.Items = []*feeds.Item{}
	for i := 0; i < len(articles); i++ {
		feed.Items = append(feed.Items, helpers.UrlToRssItem(articles[i]))
	}

	// Atom format
	atom, err := feed.ToAtom()
	if err != nil {
		log.Fatal(err)
	}
	
	// RSS format
	rss, err := feed.ToRss()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rss)
	fmt.Println("\n\n\n\n")
	fmt.Println(atom)

## What's Next?

1. Wrap everything with an HTTP server
2. Download the RSS as a file directly
3. Add more support for RSS and ATOM custom formats

## Thanks

1. [gorilla/feeds](https://github.com/gorilla/feeds) - Used the basic classes and formats from this library

2. [dyatlov/go-opengraph](https://github.com/dyatlov/go-opengraph/) - Used for parsing the meta data from the articles
