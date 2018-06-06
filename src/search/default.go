package search

type defaultMatcher struct{}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

/**
 * deaultMatcher 这个结构体实现量Matcher 的接口
 */
func (m defaultMatcher) Search(feed *Feed, searchTerm string)([]*Result, error) {
	return nil, nil
}
