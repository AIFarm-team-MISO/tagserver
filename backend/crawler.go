package backend

type Crawler struct{}

func NewCrawler() *Crawler {
	return &Crawler{}
}

// ✅ 반드시 첫 글자가 대문자여야 바인딩 가능
func (c *Crawler) FetchTags(keyword string) []string {
	return []string{"#예시", "#태그", "#인기상품"}
}
