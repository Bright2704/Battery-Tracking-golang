package entity

type RelatedProcess struct {
	

	ProcessNameIn ProcessNameIn `json: "processNameIn" bson: "processNameIn"`
	ProcessNameOut ProcessNameOut `json: "processNameOut" bson: "processNameOut"`
	ProcessTimeIn ProcessTimeIn `json: "processTimeIn" bson: "processTimeIn"`
	ProcessTimeOut ProcessTimeOut `json: "processTimeOut" bson: "processTimeOut"`
	EmployeeNameIn EmployeeNameIn `json: "employeeNameIn" bson: "employeeNameIn"`
	EmployeeNameOut EmployeeNameOut `json: "employeeNameOut" bson: "employeeNameOut"`
}

