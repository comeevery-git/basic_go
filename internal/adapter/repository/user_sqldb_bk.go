package repository

// // domain.repository 인터페이스의 구현체. 실제 데이터베이스 로직 처리

// import (
// 	"database/sql"

// 	"example.com/m/domain/model"
// 	"example.com/m/domain/repository"
// )

// // UserRepository 인터페이스를 구현하는 구조체
// type userRepository struct {
// 	// DB 연결
// 	db *sql.DB
// }

// // NewUserRepository는 새 UserRepository를 생성합니다.
// // 이 함수는 보통 main 함수에서 호출하여 UserRepository의 구현체를 생성합니다.
// func NewUserRepository(db *sql.DB) repository.UserRepository {
//     return &userRepository{
//         db: db,
//     }
// }


// func (r *userRepository) FindByID(id int) (*model.User, error) {
//     return nil, nil
// }

// func (r *userRepository) Save(user *model.User) error {
//     return nil
// }

// func (r *userRepository) Update(user *model.User) error {
//     return nil
// }

// func (r *userRepository) Delete(id int) error {
//     return nil
// }
// func (r *userRepository) GetAllUsers() ([]*model.User, error) {
// 	rows, err := r.db.Query("SELECT * FROM users")
//     if err != nil {
//         return nil, err
//     }
//     defer rows.Close()

//     var users []*model.User
//     for rows.Next() {
//         user := &model.User{}
//         err := rows.Scan(&user.ID, &user.UserName, &user.UserEmail)
//         if err != nil {
//             return nil, err
//         }
//         users = append(users, user)
//     }

//     if err = rows.Err(); err != nil {
//         return nil, err
//     }

//     return users, nil
// }