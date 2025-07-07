package types


type Student struct{
	// Id    int64  `json:"id" firestore:"id,omitempty"`
	Id    string  `json:"id" firestore:"id,omitempty"`
	Name  string `json:"name" validate:"required" firestore:"name"`
	Email string `json:"email" validate:"required" firestore:"email"`
	Age   int    `json:"age" validate:"required" firestore:"age"`
}