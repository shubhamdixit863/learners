package services

type EmployeeServiceV1 struct {
	Employees []string
}

func NewEmployeeServiceV1() Service {
	return &EmployeeServiceV1{
		Employees: make([]string, 4),
	}
}

func (s *EmployeeServiceV1) GetData() string {
	return s.Employees[0]
}

func (s *EmployeeServiceV1) SearchData() string {
	return s.Employees[0]
}

func (s *EmployeeServiceV1) AddData(data string) error {
	return nil
}
