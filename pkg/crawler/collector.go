package crawler

import (
	"context"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/gocolly/colly/v2"
)

// Collector searches for css, js, and images within a given link
// TODO improve for better performance
func Collector(ctx context.Context, urlLink string, projectPath string, cookieJar *cookiejar.Jar, proxyString string, userAgent string) (pagePath string, err error) {
	// create a new collector
	c := colly.NewCollector(colly.Async(true))
	setUpCollector(c, ctx, cookieJar, proxyString, userAgent)

	paths := strings.Split(urlLink, `com`)[1:]
	pagePath, _ = url.QueryUnescape(paths[len(paths)-1])
	fmt.Println("Page Path", "-->", pagePath)

	// search for all link tags that have a rel attribute that is equal to stylesheet - CSS
	//c.OnHTML("link[rel='stylesheet']", func(e *colly.HTMLElement) {
	//	// hyperlink reference
	//	link := e.Attr("href")
	//	// print css file was found
	//	fmt.Println("Css found", "-->", link)
	//	// extraction
	//	Extractor(e.Request.AbsoluteURL(link), projectPath)
	//})

	// search for all script tags with src attribute -- JS
	//c.OnHTML("script[src]", func(e *colly.HTMLElement) {
	//	// src attribute
	//	link := e.Attr("src")
	//	// Print link
	//	fmt.Println("Js found", "-->", link)
	//	// extraction
	//	Extractor(e.Request.AbsoluteURL(link), projectPath)
	//})

	// serach for all img tags with src attribute -- Images
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		// src attribute
		link := e.Attr("src")
		if strings.HasPrefix(link, "data:image") || strings.HasPrefix(link, "blob:") {
			return
		}
		// Print link
		fmt.Println("Img found", "-->", link)
		// extraction
		Extractor(e.Request.AbsoluteURL(link), projectPath)
	})

	//Before making a request
	c.OnRequest(func(r *colly.Request) {
		link := r.URL.String()
		//if urlLink == link {
		HTMLExtractor(link, projectPath, pagePath)
		//}
	})

	// Visit each urlLink and wait for stuff to load :)
	if err = c.Visit(urlLink); err != nil {
		return
	}
	c.Wait()
	return
}

type cancelableTransport struct {
	ctx       context.Context
	transport http.RoundTripper
}

func (t cancelableTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := t.ctx.Err(); err != nil {
		return nil, err
	}
	return t.transport.RoundTrip(req.WithContext(t.ctx))
}

func setUpCollector(c *colly.Collector, ctx context.Context, cookieJar *cookiejar.Jar, proxyString, userAgent string) {
	if cookieJar != nil {
		c.SetCookieJar(cookieJar)
	}
	if proxyString != "" {
		c.SetProxy(proxyString)
	} else {
		c.WithTransport(cancelableTransport{ctx: ctx, transport: http.DefaultTransport})
	}
	if userAgent != "" {
		c.UserAgent = userAgent
	}
}
