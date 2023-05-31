package entity

type RelatedProcess struct {
	ID uint `json: "id" bson: "id"`

	ProcessNameIn ProcessNameIn `json: "processNameIn" bson: "processNameIn"`
	ProcessNameOut ProcessNameOut `json: "processNameOut" bson: "processNameOut"`
	ProcessTimeIn ProcessTimeIn `json: "processTimeIn" bson: "processTimeIn"`
	ProcessTimeOut ProcessTimeOut `json: "processTimeOut" bson: "processTimeOut"`
	EmployeeNameIn EmployeeNameIn `json: "employeeNameIn" bson: "employeeNameIn"`
	EmployeeNameOut EmployeeNameOut `json: "employeeNameOut" bson: "employeeNameOut"`
}