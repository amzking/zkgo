package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(seachTerm string) {

	feeds, err := RetrieveFeeds()

	if err != nil {
		log.Fatal(err);
	}

	results := make(chan *Result)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// go程序终止时会关闭所有之前启动且还在运行的goroutine
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	go func() {
		waitGroup.Wait()
		close(results);
	}();

	Display(results);

}


func Register(feedType string, matcher Matcher) {

}