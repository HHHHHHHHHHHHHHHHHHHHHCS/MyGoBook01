package Search

//默认的匹配器
type defaultMatcher struct {

}

//载入这个模块的时候自动注册一些东西
func init()  {
	var matcher defaultMatcher
	Register("default",matcher)
}

//实现了默认匹配器的行为
func (m defaultMatcher)Search(feed *Feed,searchTerm string)([]*Result,error) {
	return  nil,nil
}