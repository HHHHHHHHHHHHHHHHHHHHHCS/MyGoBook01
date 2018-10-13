package OtherLib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	GResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:'visibleUrl'`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}

	GResponse struct {
		ResponseData struct {
			Results []GResult `json:"results"`
		} `json:"responseData"`
	}
)

func MyJsonMain() {
	Test05()
}

func Test01() {
	url := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer resp.Body.Close()

	var gr GResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Println(gr)
}

func Test02() {
	var jsonString = `{
	"responseData": {
		"results": [{
				"GsearchResultClass": "GwebSearch",
				"unescapedUrl": "https://www.reddit.com/r/golang",
				"url": "https://www.reddit.com/r/golang",
				"visibleUrl": "www.reddit.com",
				"cacheUrl": "http://www.google.com/search?q=cache:W...",
				"title": "r/<b>Golang</b> - Reddit",
				"titleNoFormatting": "r/Golang - Reddit",
				"content": "First Open Source <b>Golang%u..."
			},
			{
				"GsearchResultClass": "GwebSearch",
				"unescapedUrl": "http://tour.golang.org/",
				"url": "http://tour.golang.org/",
				"visibleUrl": "tour.golang.org",
				"cacheUrl": "http://www.google.com/search?q=cache:O...",
				"title": "A Tour of Go",
				"titleNoFormatting": "A Tour of Go",
				"content": "Welcome to a tour of the Go programming ..."
			}
		]
	}
}`

	var gr GResponse
	json.Unmarshal([]byte(jsonString), &gr)
	fmt.Println(gr)
}

func Test03() {
	type T struct {
		A int
		B string
		C byte
		D struct {
			D_A int
			D_B string
		}
		E bool
	}

	t := &T{
		A: 1,
		B: "Test b",
		C: 255,
		D: struct {
			D_A int
			D_B string
		}{
			D_A: 123123,
			D_B: "asdasd",
		},
		E: true,
	}

	data, _ := json.Marshal(t)
	fmt.Println(string(data))
}

func Test04() {
	type T struct {
		Aa string `json:"aa"`
	}

	var t T

	str := `{"aa":"a123a"}`
	json.Unmarshal([]byte(str), &t)
	fmt.Println(t)
}

func Test05() {
	var JSON = `
{
  "name": "Gopher",
  "title": "programmer",
  "contact": {
    "home": "415.333.3333",
    "cell": "415.555.5555"
  } 
}
`

	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	for k, v := range c {
		fmt.Println(k, v)
	}

	fmt.Println("==================================")

	fmt.Println("Name:", c["name"])
	fmt.Println("Title:", c["title"])
	fmt.Println("Contact:")
	fmt.Println("H:", c["contact"].(map[string]interface{})["home"])
	fmt.Println("H:", c["contact"].(map[string]interface{})["cell"])
}
