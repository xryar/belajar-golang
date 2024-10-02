package main

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func GetStudents() []*Student {
	return students
}

func selectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}

	return nil
}

func init() {
	students = append(students, &Student{Id: "S001", Name: "Bourne", Grade: 2})
	students = append(students, &Student{Id: "S002", Name: "Bruce", Grade: 1})
	students = append(students, &Student{Id: "B021", Name: "Agent", Grade: 3})
	students = append(students, &Student{Id: "C011", Name: "Bourne", Grade: 2})
}
