package userService

import (
	"evc/entity/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetByEmail(email string) (*user.User, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*user.User), args.Error(1)
}

func (m *MockRepository) GetByDate(limit, offset string, start, end time.Time) ([]user.User, error) {
	args := m.Called()
	result := args.Get(0)
	return result.([]user.User), args.Error(1)
}

func (m *MockRepository) GetAllUsers(limit, offset string) ([]user.User, error) {
	args := m.Called()
	result := args.Get(0)
	return result.([]user.User), args.Error(1)
}

func (m *MockRepository) Save(user2 *user.User) (*user.User, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*user.User), args.Error(1)
}

func TestSave(t *testing.T) {
	mockRepo := new(MockRepository)
	user := user.User{
		Id:        "2",
		FirstName: "Lawal",
		LastName:  "Bayo",
		Email:     "o@o.com",
		Date:      time.Now(),
	}
	mockRepo.On("Save").Return(&user, nil)

	testService := NewUserService(mockRepo)
	testRes, err := testService.Save(&user)
	assert.NotNil(t, testRes.Email)
	assert.Nil(t, err)
	assert.Equal(t, "Lawal", testRes.FirstName)
	assert.Equal(t, "o@o.com", testRes.Email)
}
