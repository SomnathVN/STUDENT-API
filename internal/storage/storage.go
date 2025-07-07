package storage

import "github.com/SomnathVN/students-api/internal/types"

type Storage interface {
	//CreateStudent(name string, email string, age int) (int64, error)
    //GetStudentById(id int64) (types.Student, error)
	CreateStudent(name string, email string, age int) (string, error)
	GetStudentById(id string) (types.Student, error)
	GetStudents() ([]types.Student, error)
	//UpdateStudent(id int64,name string, email string, age int) (types.Student, error)
    //DeleteStudent(id int64) (error)
	UpdateStudent(id string, name string, email string, age int) (types.Student, error)
	DeleteStudent(id string) error
}