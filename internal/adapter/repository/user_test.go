package repository

import (
	"testing"

	"example.com/m/domain/model"
	"example.com/m/domain/repository"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_FindByID(t *testing.T) {
    mockDB := new(gormmocks.GormMock)

    mockDB.On("First", mock.Anything, mock.Anything).Return(func(user *model.User, id int) *gorm.DB {
		if id == 1 {
            *user = model.User{ID: 1, UserName: "Lydia"}
            return &gorm.DB{}
        }
        return &gorm.DB{Error: errors.New("User Not found.")}
    })

    r := &userRepository{db: mockDB}

    // Case 1 - 존재하는 ID
    user, err := r.FindByID(1)
    assert.NoError(t, err)
    assert.Equal(t, &model.User{ID: 1, UserName: "Lydia"}, user)

	// Case 2 - 존재하지 않는 ID
    user, err = r.FindByID(2)
    assert.Error(t, err)
    assert.Nil(t, user)

    mockDB.AssertCalled(t, "First", mock.Anything, 1)
    mockDB.AssertCalled(t, "First", mock.Anything, 2)
}