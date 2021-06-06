package abstract_factory

type IKubernetes interface {
	GetApiServer() string
	GetName() string
}

type K8sFactory interface {
	Create(name, ip string) IKubernetes
}

type CloudK8s struct {
	UUID string
	IP   string
}

func (c *CloudK8s) GetApiServer() string {
	return c.IP
}

func (c *CloudK8s) GetName() string {
	return c.UUID
}

type LocalK8s struct {
	HostName string
	IP       string
}

func (l *LocalK8s) GetApiServer() string {
	return l.IP
}

func (l *LocalK8s) GetName() string {
	return l.HostName
}

type CloudK8sFactory struct{}

func (c CloudK8sFactory) Create(name, ip string) IKubernetes {
	return &CloudK8s{
		name,
		ip,
	}
}

type LocalK8sFactory struct{}

func (l LocalK8sFactory) Create(name, ip string) IKubernetes {
	return &LocalK8s{
		name + ip,
		ip,
	}
}

var (
	localFac LocalK8sFactory
	cloudFac CloudK8sFactory
)

func GetKubernetes(kind string, name, ip string) IKubernetes {
	switch kind {
	case "local":
		return localFac.Create(name, ip)
	case "cloud":
		return cloudFac.Create(name, ip)
	}
	panic("error kind")
}
