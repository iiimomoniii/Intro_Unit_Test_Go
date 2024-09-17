package services

import (
	"gotest/repositories"
	"gotest/utils"
)

type PromotionService interface {
	CalculateDiscount(amount int) (int, error)
}

type promotionService struct {
	promotionRepository repositories.PromotionRepository
}

func NewPromotionService(promotionRepository repositories.PromotionRepository) PromotionService {
	return promotionService{promotionRepository: promotionRepository}
}

func (s promotionService) CalculateDiscount(amount int) (int, error) {

	if amount <= 0 {
		return 0, utils.ErrZeroAmount
	}
	promotion, err := s.promotionRepository.GetPromotion()
	if err != nil {
		return 0, utils.ErrRepository
	}

	if amount >= promotion.PurchaseMin {
		return amount - (promotion.DiscountPercent * amount / 100), nil
	}
	return amount, nil
}
