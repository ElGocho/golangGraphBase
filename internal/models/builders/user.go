package builders

import (
	gql "sa_web_service/graph/model"
	"sa_web_service/internal/models"
	"fmt"
)

func UserFromGQL(qBuilder *gql.QueryBuilder) (builder *models.Builder){
	if qBuilder == nil {
		return
	}

	builder = &models.Builder{}

	getUserFilters(builder, qBuilder.Filter)
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

