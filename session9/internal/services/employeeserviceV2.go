package services

type EmployeeServiceV2 struct {
	Employees []string
	PayOuts   []int
}

func (e EmployeeServiceV2) GetData() string {
	return e.Employees[0]
}

func (e EmployeeServiceV2) SearchData() string {
	//TODO implement me
	panic("implement me")
}

func (e EmployeeServiceV2) AddData(data string) error {
	//TODO implement me
	panic("implement me")
}

func NwEmployeeServiceV2() Service {
	return &EmployeeServiceV2{
		Employees: make([]string, 4),
		PayOuts:   nil,
	}
}
