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

	// _为占位符，索引值，当我们所调用的函数返回多个值时，不需要其中某个值，可用下划线将其忽略
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
		// WaitGroup 跟踪goroutine的工作是否完成，是一个信号计数量
		waitGroup.Wait()
		close(results);
	}();

	Display(results);

}


func Register(feedType string, matcher Matcher) {

}