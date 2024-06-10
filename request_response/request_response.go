package requestresponse

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

type EmployeeList struct {
	Employee []Employee `json:"employee"`
	Count    int        `json:"count"`
}

type GenericResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
