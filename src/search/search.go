package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {

	feeds, err := RetrieveFeeds()

	if err != nil {
		log.Fatal(err);
	}

	results := make(chan *Result)

	var waitGroup sync.WaitGroup

	// WaitGroup 跟踪goroutine的工作是否完成，是一个信号计数量
	waitGroup.Add(len(feeds))

	// _为占位符，索引值，当我们所调用的函数返回多个值时，不需要其中某个值，可用下划线将其忽略
	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// go启动了一个匿名函数
		// 指针变量可以方便的在函数之间共享数据
		// searchTerm results 闭包方式访问
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			// 计数递减
			waitGroup.Done()
		}(matcher, feed)
	}

	go func() {
		// 阻塞直到计数为0
		waitGroup.Wait()
		close(results);
	}();

	Display(results);

}

// 在函数传值，所有的变量都是值传递，指针为内存地址，传递内存地址，指向同一份数据
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already regiestered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}