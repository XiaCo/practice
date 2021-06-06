package abstract_factory

import "context"

type SubCompany interface {
	GetEmployeeCount() int
}

type SubCompanyFactory interface {
	CreateCompany(ctx context.Context) SubCompany
}

type WestCompany struct {
	EmployeeCount int
	Province      string
}

func (w *WestCompany) GetEmployeeCount() int {
	return w.EmployeeCount
}

type ForeignCompany struct {
	EmployeeCount int
	CountryName   string
}

func (f *ForeignCompany) GetEmployeeCount() int {
	return f.EmployeeCount
}
