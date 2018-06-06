package search

type defaultMatcher struct{}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

/**
 * deaultMatcher 这个结构体实现量Matcher 的接口
 * 若一个函数在声明的时候带有一个接收者，则意味着声明量一个方法，这个方法会和指定的接受者类型绑在一起。
 */
func (m defaultMatcher) Search(feed *Feed, searchTerm string)([]*Result, error) {
	return nil, nil
}
