package basic

import (
	"sa_web_service/internal/models"
	"sa_web_service/internal/models/consts"
)

func GetWTables() []*models.WTable{
	return []*models.WTable{
		&models.WTable{
			Name: consts.TableUsers.Name(),
			Code: string(consts.TableUsers),
			Status: true,
			States: GetWStatesUser(),
		},
		&models.WTable{
			Name: consts.TableArticles.Name(),
			Code: string(consts.TableArticles),
			Status: true,
		},
		&models.WTable{
			Name: consts.TablePrices.Name(),
			Code: string(consts.TablePrices),
			Status: true,
		},
		&models.WTable{
			Name: consts.TableReceipts.Name(),
			Code: string(consts.TableReceipts),
			Status: true,
		},
		&models.WTable{
			Name: consts.TableReceiptArticles.Name(),
			Code: string(consts.TableReceiptArticles),
			Status: true,
		},
	}
}


