package repository

import (
    "database/sql"
    "example.com/m/model"
)

// java 에서처럼 implements 키워드를 사용하여 명시적으로 구현할 필요 없이
// 특정 구조체가 인터페이스의 모든 메서드를 가지고있다면 그 구조체는 자동으로 그 인터페이스를 구현한다.
type UserRepository interface {
    FindByID(id int) (*model.User, error) // 대문자이니 public 접근 가능
    Save(user *model.User) error // 에러를 반환하면 이 메서드를 호출하는 코드에서 처리필요, 에러가 발생하지 않으면 nil 반환
    Update(user *model.User) error
    Delete(id int) error
	GetAllUsers() ([]*model.User, error)
}

// UserRepository 인터페이스를 구현하는 구조체
type userRepository struct {
	// DB 연결
	db *sql.DB
}

// NewUserRepository는 새 UserRepository를 생성합니다.
// 이 함수는 보통 main 함수에서 호출하여 UserRepository의 구현체를 생성합니다.
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}


func (r *userRepository) FindByID(id int) (*model.User, error) {
    return nil, nil
}

func (r *userRepository) Save(user *model.User) error {
    return nil
}

func (r *userRepository) Update(user *model.User) error {
    return nil
}

func (r *userRepository) Delete(id int) error {
    return nil
}
func (r *userRepository) GetAllUsers() ([]*model.User, error) {
	rows, err := r.db.Query("SELECT * FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []*model.User
    for rows.Next() {
        user := &model.User{}
        err := rows.Scan(&user.ID, &user.Name, &user.Email)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return users, nil
}