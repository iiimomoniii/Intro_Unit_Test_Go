package tests

import (
	"gotest/repositories"
	mocks "gotest/repositories_mocks"
	"gotest/services"
	"gotest/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

// how to test
// go test ./tests -run="TestPromotionCalculateDiscount" -v
func TestPromotionCalculateDiscount(t *testing.T) {

	//prepair for test
	type testCase struct {
		caseName         string
		purchPurchaseMin int
		discountPercent  int
		amount           int
		expected         int
	}

	//test case amount >= promotion.PurchaseMin
	getPromotionCase := []testCase{
		{caseName: "positive applied 100", purchPurchaseMin: 100, discountPercent: 20, amount: 100, expected: 80},
		{caseName: "positive applied 200", purchPurchaseMin: 100, discountPercent: 20, amount: 200, expected: 160},
		{caseName: "positive applied 300", purchPurchaseMin: 100, discountPercent: 20, amount: 300, expected: 240},
		{caseName: "positive not applied 50", purchPurchaseMin: 100, discountPercent: 20, amount: 50, expected: 50},
	}

	for _, c := range getPromotionCase {
		t.Run(c.caseName, func(t *testing.T) {

			//loop test function GetPromotion
			//Arrage
			//เตรียมของ
			promotionRepository := mocks.NewPromotionRepositoryMock()
			promotionRepository.On("GetPromotion").Return(repositories.Promotion{
				ID:              1,
				PurchaseMin:     c.purchPurchaseMin,
				DiscountPercent: c.discountPercent,
			}, nil)

			promotionService := services.NewPromotionService(promotionRepository)

			//Art
			//เริ่มทำงาน
			discount, _ := promotionService.CalculateDiscount(c.amount)
			expected := c.expected

			//Asset
			assert.Equal(t, expected, discount)
		})

		//loop test function  amount <= 0
		t.Run("amountLesthan0", func(t *testing.T) {
			//Arrage
			//เตรียมของ
			promotionRepository := mocks.NewPromotionRepositoryMock()
			promotionRepository.On("GetPromotion").Return(repositories.Promotion{
				ID:              1,
				PurchaseMin:     100,
				DiscountPercent: 20,
			}, nil)

			promotionService := services.NewPromotionService(promotionRepository)

			//Art
			//เริ่มทำงาน
			_, err := promotionService.CalculateDiscount(0)

			//Asset
			assert.ErrorIs(t, err, utils.ErrZeroAmount)
			//ignore function GetPromotion when amount <= 0
			//when call function s.promotionRepository.GetPromotion() before
			//call function check amount <= 0
			promotionRepository.AssertNotCalled(t, "GetPromotion")

		})

	}

}
