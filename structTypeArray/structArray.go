package structTypeArray

import "fmt"

type dArray struct {
	name   string
	rollno int
}

func StructArray() {
	studentarray := []dArray{
		{
		name:   "Rishu",
		rollno: 87,
	},
	{
		name:   "Rishu",
		rollno: 87,
	},
	}
	studentarray = append(studentarray, dArray{
		name:   "Rishu",
		rollno: 87,
	})
	fmt.Println(studentarray)
}
