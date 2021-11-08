package models

import (
	"gorm.io/gorm"
	"fmt"
)

type Builder struct {
	Table	string	`json:"table"`
	Select	interface{} `json:"select"`
	Where	[]Where	`json:"where"`
	Pagination *Pagination `json:"pagination"`
	Order *string	`json:"order"`
	Group *string	`json:"group"`
}

type Where struct {
	Code	string
	Condition interface{}
	Params []interface{}
}

type Pagination struct {
	Page *int
	Skip *int
	Take *int
}

func NewSession(tx *gorm.DB) *gorm.DB{
	return tx.Session(&gorm.Session{NewDB: true})
}

func BuilderORMQuery(db *gorm.DB, builder *Builder) *gorm.DB{
	if(builder == nil){
		return  db
	}

	db = BuilderORMTable(db,builder)

	db = BuilderORMSelect(db, builder)

	db = BuilderORMOrder(db, builder)

	db = BuilderORMGroup(db, builder)

	db = BuilderORMWhere(db, builder)

	db = BuilderORMPagination(db, builder)

	return db
}

func BuilderORMTable(db *gorm.DB, builder *Builder) *gorm.DB{
	if(builder == nil || builder.Table == ""){
		return db
	}

	db = db.Table(builder.Table)

	return db
}

func BuilderORMSelect(db *gorm.DB, builder *Builder) *gorm.DB{
	if(builder == nil || builder.Select == nil){
		return db
	}

	db = db.Select(builder.Select)
	
	return db
}

func BuilderORMOrder(db *gorm.DB, builder *Builder) *gorm.DB{
	if builder == nil || builder.Order == nil{
		return db
	}

	db = db.Order(*builder.Order)

	return db
}

func BuilderORMGroup(db *gorm.DB, builder *Builder) *gorm.DB{
	if builder == nil || builder.Group == nil{
		return db
	}

	db = db.Group(*builder.Group)

	return db
}

func BuilderORMWhere(db *gorm.DB, builder *Builder) *gorm.DB{
	if builder == nil{
		return db
	}

	for _,w := range builder.Where {
		db = db.Where(w.Condition, w.Params...)	
	} 

	return db
}

func BuilderORMPagination(db *gorm.DB, builder *Builder) *gorm.DB{
	if builder == nil || builder.Pagination == nil{
		return db
	}

	var skip, take int

	if builder.Pagination.Page != nil && builder.Pagination.Skip != nil {
		skip = *builder.Pagination.Page * *builder.Pagination.Skip
	}else if builder.Pagination.Skip != nil {
		skip = *builder.Pagination.Skip
	}

	if builder.Pagination.Take != nil {
		take = *builder.Pagination.Take
	}

	db = db.Offset(skip).Limit(take)

	return db
}

func (b *Builder) Merge(param *Builder, priority *Priority) {
	if param == nil {
		return
	}

	priority.Valid()

	b.mergeWhere(param.Where, *priority)	

	b.mergePagination(param.Pagination, *priority)

	b.mergeOrder(param.Order, *priority)

	b.mergeGroup(param.Group, *priority)
}

func (b *Builder) mergeWhere(param []Where, priority Priority){
	for _,v := range param {
		OK, index, _ := b.WhereContains(v, Priority3); 
		if OK{
			switch(priority){
				case Priority1:
					break
				case Priority2:
					b.Where, _ = RemoveWhereFromSlice(b.Where, index)

					b.Where = append(b.Where, v)
				default:
					b.Where = append(b.Where, v)
			}
		}else{
			b.Where = append(b.Where, v)
		}
	}
} 

func (b *Builder)mergePagination(param *Pagination, priority Priority){
	if param == nil {
		return
	}

	switch(priority){
		case Priority1: 
			if b.Pagination == nil {
				b.Pagination = param
				break
			}

			if b.Pagination.Page == nil {
				b.Pagination.Page = param.Page
			}

			if b.Pagination.Skip == nil {
				b.Pagination.Skip = param.Skip
			}

			if b.Pagination.Take == nil {
				b.Pagination.Take = param.Take
			}
		case Priority2: 
			if b.Pagination == nil {
				b.Pagination = param
				break
			}

			if param.Page != nil {
				b.Pagination.Page = param.Page
			}

			if param.Skip != nil {
				b.Pagination.Skip = param.Skip
			}

			if param.Take != nil {
				b.Pagination.Take = param.Take
			}
		default :
			break
	}
}

func (b *Builder)mergeOrder(param *string, priority Priority){
	if param == nil{
		return
	}else if b.Order == nil {
		b.Order = param
	}

	switch(priority){
		case Priority1: break
		case Priority2: 
			b.Order = param
		default:
			order := fmt.Sprintf("%s,%s",*b.Order, *param)
			b.Order = &order
	}
}

func (b *Builder)mergeGroup(param *string, priority Priority){
	if param == nil{
		return
	}else if b.Group == nil {
		b.Group = param
	}

	switch(priority){
		case Priority1: break
		case Priority2: 
			b.Group = param
		default:
			group := fmt.Sprintf("%s,%s",*b.Group, *param)
			b.Group = &group
	}
}

func (b *Builder)WhereContains(param Where, priority Priority) (bool, int, []Where){
	if b == nil {
		return false, 0, nil
	}

	wheres := []Where{}

	for index, v := range b.Where {

		if v.Code == param.Code && v.Code != ""{
			if priority == Priority1 || priority == Priority3{
				return true, index, nil
			}

			wheres = append(wheres, v)
		}

		if v.Condition == param.Condition {
			if priority == Priority2 || priority == Priority3{
				return true, index, nil
			} 

			wheres = append(wheres, v)
		}

	}

	return false, 0, wheres
}

func RemoveWhereFromSlice(slice []Where, index int) ([]Where, error) {
	final := len(slice)-1

	if index > final || index < 0{
		return slice, fmt.Errorf("Indice fuera de rango %d", index)
	}

	slice[index] = slice[final]

	return slice[:final], nil
}

func RemoveWhereFromSliceOrdenado(slice []Where, index int) ([]Where, error) {
	if index > len(slice) || index < 0{
		return slice, fmt.Errorf("Indice fuera de rango %d", index)
	}

	return append(slice[:index], slice[index+1:]...), nil
}

func SelectOnlyID() *Builder{
	return &Builder{ Select: "id" }
}
