package builder

import "net"

type Plugin struct {
	Name string
	Mode string // net/storage
}

type Subnet struct {
	Network net.IPNet
}

type Kubernetes struct {
	Name       string
	Version    string
	Mode       string
	Plugins    []Plugin
	Arch       string
	EntryPoint string
	Subnet     Subnet
}

type KubernetesBuilder interface {
	BuildPlugins()
	BuildSubnet()
	BuildOther()
}

type ABuilder struct{}

func (A *ABuilder) BuildPlugins() {
	panic("implement me")
}

func (A *ABuilder) BuildSubnet() {
	panic("implement me")
}

func (A *ABuilder) BuildOther() {
	panic("implement me")
}

type BBuilder struct{}

func (B *BBuilder) BuildPlugins() {
	panic("implement me")
}

func (B *BBuilder) BuildSubnet() {
	panic("implement me")
}

func (B *BBuilder) BuildOther() {
	panic("implement me")
}

type Director struct {
	// 控制创建流程
}

func (d *Director) MakeXK8s() {
	b := new(ABuilder)
	b.BuildPlugins()
	b.BuildSubnet()
}

func (d *Director) MakeYK8s() {
	b := new(BBuilder)
	b.BuildSubnet()
	b.BuildPlugins()
	b.BuildOther()
}

func t() {
	d := new(Director)
	d.MakeXK8s()
}
