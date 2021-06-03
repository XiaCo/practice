package iterator

type Iterator interface {
	Next() interface{}
	HasNext() bool
}

type Articles []struct {
	ID      string
	Name    string
	Content string
}

type AscSearch struct {
	Articles
	nowLocation int
}

func (a *AscSearch) Next() interface{} {
	if a.HasNext() {
		a.nowLocation++
		return a.Articles[a.nowLocation]
	}
	return nil
}

func (a *AscSearch) HasNext() bool {
	if len(a.Articles)-1 > a.nowLocation {
		return true
	}
	return false
}

func NewAscSearch(art Articles) *AscSearch {
	return &AscSearch{
		art,
		-1,
	}
}

type DescSearch struct {
	Articles
	nowLocation int
}

func (d *DescSearch) Next() interface{} {
	if d.HasNext() {
		d.nowLocation--
		return d.Articles[d.nowLocation]
	}
	return nil
}

func (d *DescSearch) HasNext() bool {
	if d.nowLocation > 0 {
		return true
	}
	return false
}

func NewDescSearch(articles Articles) *DescSearch {
	return &DescSearch{
		articles,
		len(articles),
	}
}
