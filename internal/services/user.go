package services

import (
	"gorm.io/gorm"

	gql "sa_web_service/graph/model"
	"sa_web_service/internal/models"
	"sa_web_service/internal/models/consts"
	"sa_web_service/internal/models/builders"
)

func RegisterUser(tx *gorm.DB, user *models.User) (*models.User, error){
	db := models.NewSession(tx)

	err := configUser(db, user, consts.StateActUser)
	
	if err != nil {
		return nil, err
	}

	err = user.Create(db).Error

	return user, err
}

func Users(tx *gorm.DB, qBuilder *gql.QueryBuilder, pBuilder *models.Builder) (resp models.Users){
	builder := &models.Builder{}
	priority := models.Priority1
	db := models.NewSession(tx)	

	if qBuilder != nil{
		b := builders.UserFromGQL(qBuilder)

		builder.Merge(b, &priority)
	}

	builder.Merge(pBuilder, &priority)

	resp.Find(db, builder)

	return
}

func configUser(db *gorm.DB, user *models.User, state consts.State) error {
	builder := models.SelectITable(consts.TableUsers, "id")

	err := user.Table.Get(db, builder).Error

	if err != nil {
		return err
	}

	user.TableID = user.Table.ID

	builder = models.SelectIState(state, user.TableID, "id")

	err = user.State.Get(db, builder).Error

	if err != nil {
		return err
	}

	user.StateID = user.State.ID

	return nil
}
