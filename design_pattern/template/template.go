package template

type Crawler interface {
	Download() ([]byte, error)
	Unmarshaler([]byte) interface{}
	Save(interface{}) error
}

type CrawlerTemplate struct {
	Crawler
}

func (t *CrawlerTemplate) Run() error {
	html, err := t.Download()
	if err != nil {
		return err
	}
	instance := t.Unmarshaler(html)
	return t.Save(instance)
}

type SinaNewsCrawler struct{}

func (s SinaNewsCrawler) Download() ([]byte, error) {
	// request news.sina
	panic("implement me")
}

func (s SinaNewsCrawler) Unmarshaler(bytes []byte) interface{} {
	// 反序列化
	panic("implement me")
}

func (s SinaNewsCrawler) Save(i interface{}) error {
	// 保存到mongo news库 sina表
	panic("implement me")
}

type TencentNewsCrawler struct{}

func (t TencentNewsCrawler) Download() ([]byte, error) {
	// request news.tencent
	return nil, nil
}

func (t TencentNewsCrawler) Unmarshaler(bytes []byte) interface{} {
	// 反序列化
	panic("implement me")
}

func (t TencentNewsCrawler) Save(i interface{}) error {
	// 保存到mongo news库 tencent表
	panic("implement me")
}
