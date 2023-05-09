package main

import "fmt"

type student struct {
	id    int
	name  string
	marks []int
}

func main() {

	employee := student{
		id:   1,
		name: "tushar",
		marks: []int{
			56,
			67,
			78,
		},
	}

	fmt.Println(employee)

	emp := student{}
	emp.id = 23
	emp.name = "tushar"
	emp.marks = []int{23, 23}
	fmt.Println(emp)

	
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

}
