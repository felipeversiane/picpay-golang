package converter

import (
	domain "github.com/felipeversiane/picpay-golang.git/internal"
	"github.com/felipeversiane/picpay-golang.git/internal/entity/response"
)

func ConvertUserDomainToResponse(
	userDomain domain.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:         userDomain.GetID(),
		Email:      userDomain.GetEmail(),
		FirstName:  userDomain.GetFirstName(),
		LastName:   userDomain.GetLastName(),
		IsMerchant: userDomain.GetIsMerchant(),
		Document:   userDomain.GetDocument(),
		Balance:    userDomain.GetBalance(),
		CreatedAt:  userDomain.GetCreatedAt(),
		UpdatedAt:  userDomain.GetUpdatedAt(),
	}
}
