package html

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// TODO: figure out what was done here at 4am
func arrange(projectDir string, pagePath string) error {
	VisitedPageSet[domain+pagePath] = struct{}{}

	indexfile := projectDir + "/" + pagePath + ".html"
	input, err := ioutil.ReadFile(indexfile)
	if err != nil {
		return err
	}
	lines := strings.Split(string(input), "\n")

	for index, line := range lines {
		b := []byte(line)
		r := bytes.NewReader(b)
		doc, err := goquery.NewDocumentFromReader(r)
		if err != nil {
			return err
		}
		// Replace JS links in HTML
		doc.Find("script[src]").Each(func(i int, s *goquery.Selection) {
			data, exists := s.Attr("src")
			if exists {
				file := filepath.Base(data)

				s.SetAttr("src", projectDir+"/js/"+file)
				if data, _ := s.Attr("src"); data != "" {
					lines[index] = fmt.Sprintf(`<script src="%s"></script>`, data)
				}
			}
		})

		// Replace CSS links in HTML
		doc.Find("link[rel='stylesheet']").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the hyperlink reference
			data, exists := s.Attr("href")
			if exists {
				file := filepath.Base(data)

				s.SetAttr("href", projectDir+"/css/"+file)
				if data, _ := s.Attr("href"); data != "" {
					lines[index] = fmt.Sprintf(`<link rel="stylesheet" type="text/css" href="%s">`, data)
				}
			}
		})

		// 替换上一页下一页链接
		doc.Find(".menu-item").Each(func(i int, s *goquery.Selection) {
			data, exists := s.Attr("href")
			if exists {
				original := lines[index]
				data, _ = url.QueryUnescape(data)
				toCrawlPage := domain + data
				if _, ok := VisitedPageSet[toCrawlPage]; !ok {
					ToVisitPageSet[toCrawlPage] = struct{}{}
				}
				newLink := projectDir + data + ".html"
				lines[index] = reLink.ReplaceAllString(original, `href=`+`"`+newLink+`"`)
			}
		})

		// Replace IMG links in HTML
		// TODO: is the regex necessary here?
		doc.Find("img[src]").Each(func(i int, s *goquery.Selection) {
			data, exists := s.Attr("src")
			if exists {
				original := lines[index]
				file := filepath.Base(data)
				s.SetAttr("src", projectDir+"/imgs/"+file)

				if data, _ := s.Attr("src"); data != "" {
					lines[index] = reSrc.ReplaceAllString(original, `src=`+data)
				}
			}
		})
	}

	delete(ToVisitPageSet, pagePath)

	output := strings.Join(lines, "\n")
	return ioutil.WriteFile(indexfile, []byte(output), 0777)
}

var (
	reSrc = regexp.MustCompile(`src\s*=\s*"(.+?)"`)

	reLink = regexp.MustCompile(`href=".*.md"`)

	domain = `https://learn.lianglianglee.com`

	ToVisitPageSet = map[string]struct{}{}
	VisitedPageSet = map[string]struct{}{}
)
