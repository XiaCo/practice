package abstract_factory

import (
	"context"
)

type Zone interface {
	SubCompanyFactory
	K8sFactory
}

// 外包公司
type OutTree struct{}

func (o OutTree) CreateCompany(ctx context.Context) SubCompany {
	panic("implement me")
}

func (o OutTree) Create(name, ip string) IKubernetes {
	panic("implement me")
}

// 集团内部
type InTree struct{}

func (i InTree) CreateCompany(ctx context.Context) SubCompany {
	panic("implement me")
}

func (i InTree) Create(name, ip string) IKubernetes {
	panic("implement me")
}
