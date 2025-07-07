package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/SomnathVN/students-api/internal/types"
	"github.com/SomnathVN/students-api/internal/config"
	
)

type FirestoreStorage struct {
	client *firestore.Client
}

func New(cfg *config.Config) (*FirestoreStorage, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, cfg.ProjectId)
	if err != nil {
		return nil, err
	}
	return &FirestoreStorage{client: client}, nil
}

func (fs *FirestoreStorage) CreateStudent(name string, email string, age int) (string, error) {
	ctx := context.Background()
	docRef, _, err := fs.client.Collection("students").Add(ctx, map[string]interface{}{
		"name":  name,
		"email": email,
		"age":   age,
	})
	if err != nil {
		return "", err
	}
	return docRef.ID, nil
}

func (fs *FirestoreStorage) GetStudentById(id string) (types.Student, error) {
	ctx := context.Background()
	doc, err := fs.client.Collection("students").Doc(id).Get(ctx)
	if err != nil {
		return types.Student{}, err
	}
	var student types.Student
	if err := doc.DataTo(&student); err != nil {
		return types.Student{}, err
	}
	student.Id = id
	return student, nil
}

func (fs *FirestoreStorage) GetStudents() ([]types.Student, error) {
	ctx := context.Background()
	iter := fs.client.Collection("students").Documents(ctx)
	var students []types.Student
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var student types.Student
		if err := doc.DataTo(&student); err != nil {
			return nil, err
		}
		student.Id = doc.Ref.ID
		students = append(students, student)
	}
	return students, nil
}

func (fs *FirestoreStorage) UpdateStudent(id string, name string, email string, age int) (types.Student, error) {
	ctx := context.Background()
	updates := map[string]interface{}{
		"name":  name,
		"email": email,
		"age":   age,
	}
	_, err := fs.client.Collection("students").Doc(id).Set(ctx, updates, firestore.MergeAll)
	if err != nil {
		return types.Student{}, err
	}
	return fs.GetStudentById(id)
}

func (fs *FirestoreStorage) DeleteStudent(id string) error {
	ctx := context.Background()
	_, err := fs.client.Collection("students").Doc(id).Delete(ctx)
	return err
} 