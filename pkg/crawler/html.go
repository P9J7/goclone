package crawler

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// HTMLExtractor ...
func HTMLExtractor(link string, projectPath string, pagePath string) {
	fmt.Println("Extracting --> ", link)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// get the html body
	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}

	// Close the body once everything else is compled
	defer resp.Body.Close()

	fileName := projectPath + pagePath + ".html"
	// 获取目录路径
	dir := filepath.Dir(fileName)
	// 检查目录是否存在
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 如果目录不存在，则创建它
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	// get the project name and path we use the path to
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	htmlData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	f.Write(htmlData)

}
