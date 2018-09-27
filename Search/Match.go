package Search

import (
	"fmt"
	"log"
)

type Result struct {
	Field   string
	Content string
}

type Matcher interface {
	Search(feed *Feed, SearchTerm string) ([]*Result, error)
}

//为每个数据源单独启动goroutine来执行这个函数  并发的执行搜索
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	//对特定的匹配器执行搜索
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range searchResults {
		results <- result
	}
}

func Display(results chan *Result)  {
	for result:=range results{
		fmt.Printf("%s:\n%s\n\n",result.Field,result.Content)
	}
}