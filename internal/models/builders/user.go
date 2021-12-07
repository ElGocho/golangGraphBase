package builders

import (
	gql "sa_web_service/graph/model"
	"sa_web_service/internal/models"
	"sa_web_service/internal/models/consts"
	"fmt"
)

func UserFromGQL(qBuilder *gql.QueryBuilder) (builder *models.Builder){
	if qBuilder == nil {
		return
	}

	builder = &models.Builder{}

	getUserFilters(builder, qBuilder.Filter)
	GetUsersJoins(builder, qBuilder.Join)
	builder = getPagination(builder, qBuilder.Pagination)
	builder = getOrder(builder, qBuilder.Order)
	builder = getGroup(builder, qBuilder.Group)

	return
}

func getUserFilters(builder *models.Builder, filter map[string]interface{}) {
	if filter == nil {
		return
	}

	for k,v := range filter{
		getUserFilter(builder, k, v)
	}

}	

func GetUsersJoins(builder *models.Builder, param []string) *models.Builder{
	if len(param) == 0{
		return builder
	}

	for _,elem := range param{
		getUserJoin(builder, elem)		
	}
	
	return builder
}

func getUserJoin(builder *models.Builder, param string) *models.Builder{
	if param == ""{
		return builder
	}

	var join string

	switch(param){
		case "receipt":
			join = fmt.Sprintf(`
			Join %s as a on a.seller_id = %s.id
			Join %s as ra on ra.article_id = a.id
			Join %s on %s.id = ra.receipt_id
			`,
			consts.TableArticles,
			consts.TableUsers,
			consts.TableReceiptArticles,
			consts.TableReceipts,
			consts.TableReceipts,
		)

		default:
			return builder
	}

	builder.Joins = append(builder.Joins, join)
	
	return builder
}


func getUserFilter(builder *models.Builder, code string, param interface{} ) {
	where := models.Where{}

	switch(code){
		case "id": 
			where.Condition = "users.id = ?"
			where.Params = []interface{} { param };

			break
		case "name":
			where.Condition = "users.name like ?"
			where.Params = []interface{} { fmt.Sprintf("%c%v%c",'%',param,'%')}

			break
	}
	
	builder.Where = append(builder.Where, where)
}

