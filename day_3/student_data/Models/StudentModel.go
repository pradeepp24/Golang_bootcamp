package Models

type Student struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	DOB      string `json:"dob"`
	Address  string `json:"address"`
	Subject  string `json:"subject"`
	Marks    uint   `json:"marks"`
}

func (b *Student) TableName() string {
	return "student"
}
