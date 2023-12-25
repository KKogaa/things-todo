package repositories

// import "github.com/KKogaa/things-todo-backend/infra/db"

type Priority struct {
	BaseModel
	Title string `json:"title"`
}

func GetAllPriorities() ([]Priority, error) {
	// var priorities []Priority
	// if err := db.Gorm.Find(&priorities).Error; err != nil {
	// 	return nil, err
	// }
	// return priorities, nil
	return nil, nil
}
