package composite

// IOrganization 组织接口，实现统计人数的功能
type IOrganization interface {
	Count() int
}

type Employee struct {
	Name string
}

func (Employee) Count() int {
	return 1
}

type Department struct {
	Name string
	SubOrganizations []IOrganization
}

func (d Department) Count() int {
	c := 0
	for _, org := range d.SubOrganizations {
		c += org.Count()
	}
	return c
}

func (d *Department) AddSub(org IOrganization) {
	d.SubOrganizations = append(d.SubOrganizations, org)
}

func NewOrganization() IOrganization {
	root := &Department{Name: "root"}
	root.AddSub(&Employee{})
	root.AddSub(&Department{Name: "sub", SubOrganizations: []IOrganization{&Employee{}}})
	return root
}
