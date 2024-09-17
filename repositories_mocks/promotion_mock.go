package mocks

import (
	"gotest/repositories"

	"github.com/stretchr/testify/mock"
)

type promotionRepositoryMock struct {
	mock.Mock
}

func NewPromotionRepositoryMock() *promotionRepositoryMock {
	return &promotionRepositoryMock{}
}

func (m *promotionRepositoryMock) GetPromotion() (repositories.Promotion, error) {
	args := m.Called()
	//catch type
	return args.Get(0).(repositories.Promotion), args.Error(1)
}
