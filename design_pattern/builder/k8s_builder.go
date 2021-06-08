package builder

// k8s中 sigs.k8s.io/controller-runtime包 Builder类
// 它是这种写法来构建一个对象

type Product struct {
	ID       string
	Score    int
	ColorRGB [3]uint8
}

type Builder struct {
	Name     string
	Score    int
	ColorRGB [3]uint8
}

func (b *Builder) WithName(name string) *Builder {
	b.Name = name
	return b
}

func (b *Builder) WithScore(score int) *Builder {
	b.Score = score
	return b
}

func (b *Builder) WithColorRGB(rgb [3]uint8) *Builder {
	b.ColorRGB = rgb
	return b
}

func (b *Builder) Build() Product {
	return Product{
		b.Name,
		b.Score,
		b.ColorRGB,
	}
}

// 用户代码如下
// 当服务端升级了Product，修改了NewP，这段代码可以兼容
func GetProduct1() Product {
	var name string
	var score int
	var rgb = [3]uint8{}
	return (&Builder{}).
		WithScore(score).
		WithColorRGB(rgb).
		WithName(name).
		Build()
}

// 如果服务端代码用New来构建
func NewP(name string, score int, rgb [3]uint8) Product {
	return Product{
		name, score, rgb,
	}
}

// 用户代码如下
// 当服务端升级了Product，修改了NewP，这段代码就得修改了
func GetProduct2() Product {
	var name string
	var score int
	var rgb = [3]uint8{}
	return NewP(name, score, rgb)
}
