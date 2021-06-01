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
	GetResult() *Kubernetes
}

type ABuilder struct {
	k8s *Kubernetes
}

func (A *ABuilder) BuildPlugins() {
	panic("implement me")
}

func (A *ABuilder) BuildSubnet() {
	panic("implement me")
}

func (A *ABuilder) BuildOther() {
	panic("implement me")
}

func (A *ABuilder) GetResult() *Kubernetes {
	return A.k8s
}

type BBuilder struct {
	k8s *Kubernetes
}

func (B *BBuilder) BuildPlugins() {
	panic("implement me")
}

func (B *BBuilder) BuildSubnet() {
	panic("implement me")
}

func (B *BBuilder) BuildOther() {
	panic("implement me")
}

func (B *BBuilder) GetResult() *Kubernetes {
	return B.k8s
}

type Director struct {
	// 控制创建流程
}

func (d *Director) MakeXK8s() *Kubernetes {
	b := new(ABuilder)
	b.BuildPlugins()
	b.BuildSubnet()
	return b.GetResult()
}

func (d *Director) MakeYK8s() *Kubernetes {
	b := new(BBuilder)
	b.BuildSubnet()
	b.BuildPlugins()
	b.BuildOther()
	return b.GetResult()
}
