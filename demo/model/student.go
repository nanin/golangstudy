package student

type Student struct {
	Name string `json:"nickname"` //标签，用反射机制实现
	Age  int
}

type StudentSlice []Student

//实现Interface接口里的Len方法
func (s StudentSlice) Len() int {
	return len(s)
}

//实现Interface接口里的Less方法
func (s StudentSlice) Less(i, j int) bool {
	//返回true 则升序，false为降序
	// return s[i].Name < s[j].Name //按姓名排序
	return s[i].Age < s[j].Age //按年龄排序
}

//实现Interface接口里的Swap方飞
func (s StudentSlice) Swap(i, j int) {
	// tmp := s[i]
	// s[i] = s[j]
	// s[j] = tmp
	s[i], s[j] = s[j], s[i]
}
