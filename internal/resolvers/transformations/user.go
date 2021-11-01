package transformations

import (
	"sa_web_service/internal/models"
	gql "sa_web_service/graph/model"
)


func UserInputToUserModel(userInput *gql.UserInput) (*models.User, error) {
	user := &models.User{
		Name:	userInput.Name,
		Email:	userInput.Email,
		Identification:	userInput.Identification,
		Image:	userInput.Image,
		Password:	userInput.Password,
		RoleID:	uint(userInput.RoleID),
	}

	return user, nil
}

func UserModelToUserGQL(user *models.User) (*gql.User, error){
	gqlUser := &gql.User{
		ID:	user.ID.String(),
		Name:	user.Name,
		Email:	user.Email,
		Identification:	user.Identification,
		Image:	user.Image,
		Password:	user.Password,
		RoleID:	int(user.RoleID),
	}

	return gqlUser, nil
}



