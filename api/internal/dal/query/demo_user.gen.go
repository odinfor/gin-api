// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gin-api/internal/dal/model"
)

func newDemoUser(db *gorm.DB) demoUser {
	_demoUser := demoUser{}

	_demoUser.demoUserDo.UseDB(db)
	_demoUser.demoUserDo.UseModel(&model.DemoUser{})

	tableName := _demoUser.demoUserDo.TableName()
	_demoUser.ALL = field.NewField(tableName, "*")
	_demoUser.ID = field.NewInt32(tableName, "id")
	_demoUser.CreatedAt = field.NewTime(tableName, "created_at")
	_demoUser.UpdatedAt = field.NewTime(tableName, "updated_at")
	_demoUser.DeletedAt = field.NewField(tableName, "deleted_at")
	_demoUser.Name = field.NewString(tableName, "name")
	_demoUser.Address = field.NewString(tableName, "address")
	_demoUser.Age = field.NewInt32(tableName, "age")

	_demoUser.fillFieldMap()

	return _demoUser
}

type demoUser struct {
	demoUserDo demoUserDo

	ALL       field.Field
	ID        field.Int32
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Name      field.String
	Address   field.String
	Age       field.Int32

	fieldMap map[string]field.Expr
}

func (d demoUser) Table(newTableName string) *demoUser {
	d.demoUserDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d demoUser) As(alias string) *demoUser {
	d.demoUserDo.DO = *(d.demoUserDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *demoUser) updateTableName(table string) *demoUser {
	d.ALL = field.NewField(table, "*")
	d.ID = field.NewInt32(table, "id")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.DeletedAt = field.NewField(table, "deleted_at")
	d.Name = field.NewString(table, "name")
	d.Address = field.NewString(table, "address")
	d.Age = field.NewInt32(table, "age")

	d.fillFieldMap()

	return d
}

func (d *demoUser) WithContext(ctx context.Context) *demoUserDo { return d.demoUserDo.WithContext(ctx) }

func (d demoUser) TableName() string { return d.demoUserDo.TableName() }

func (d demoUser) Alias() string { return d.demoUserDo.Alias() }

func (d *demoUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *demoUser) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 7)
	d.fieldMap["id"] = d.ID
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["deleted_at"] = d.DeletedAt
	d.fieldMap["name"] = d.Name
	d.fieldMap["address"] = d.Address
	d.fieldMap["age"] = d.Age
}

func (d demoUser) clone(db *gorm.DB) demoUser {
	d.demoUserDo.ReplaceDB(db)
	return d
}

type demoUserDo struct{ gen.DO }

func (d demoUserDo) Debug() *demoUserDo {
	return d.withDO(d.DO.Debug())
}

func (d demoUserDo) WithContext(ctx context.Context) *demoUserDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d demoUserDo) Clauses(conds ...clause.Expression) *demoUserDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d demoUserDo) Returning(value interface{}, columns ...string) *demoUserDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d demoUserDo) Not(conds ...gen.Condition) *demoUserDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d demoUserDo) Or(conds ...gen.Condition) *demoUserDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d demoUserDo) Select(conds ...field.Expr) *demoUserDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d demoUserDo) Where(conds ...gen.Condition) *demoUserDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d demoUserDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *demoUserDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d demoUserDo) Order(conds ...field.Expr) *demoUserDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d demoUserDo) Distinct(cols ...field.Expr) *demoUserDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d demoUserDo) Omit(cols ...field.Expr) *demoUserDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d demoUserDo) Join(table schema.Tabler, on ...field.Expr) *demoUserDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d demoUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) *demoUserDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d demoUserDo) RightJoin(table schema.Tabler, on ...field.Expr) *demoUserDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d demoUserDo) Group(cols ...field.Expr) *demoUserDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d demoUserDo) Having(conds ...gen.Condition) *demoUserDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d demoUserDo) Limit(limit int) *demoUserDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d demoUserDo) Offset(offset int) *demoUserDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d demoUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *demoUserDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d demoUserDo) Unscoped() *demoUserDo {
	return d.withDO(d.DO.Unscoped())
}

func (d demoUserDo) Create(values ...*model.DemoUser) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d demoUserDo) CreateInBatches(values []*model.DemoUser, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d demoUserDo) Save(values ...*model.DemoUser) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d demoUserDo) First() (*model.DemoUser, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DemoUser), nil
	}
}

func (d demoUserDo) Take() (*model.DemoUser, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DemoUser), nil
	}
}

func (d demoUserDo) Last() (*model.DemoUser, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DemoUser), nil
	}
}

func (d demoUserDo) Find() ([]*model.DemoUser, error) {
	result, err := d.DO.Find()
	return result.([]*model.DemoUser), err
}

func (d demoUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DemoUser, err error) {
	buf := make([]*model.DemoUser, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d demoUserDo) FindInBatches(result *[]*model.DemoUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d demoUserDo) Attrs(attrs ...field.AssignExpr) *demoUserDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d demoUserDo) Assign(attrs ...field.AssignExpr) *demoUserDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d demoUserDo) Joins(fields ...field.RelationField) *demoUserDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d demoUserDo) Preload(fields ...field.RelationField) *demoUserDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d demoUserDo) FirstOrInit() (*model.DemoUser, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DemoUser), nil
	}
}

func (d demoUserDo) FirstOrCreate() (*model.DemoUser, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DemoUser), nil
	}
}

func (d demoUserDo) FindByPage(offset int, limit int) (result []*model.DemoUser, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d demoUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d *demoUserDo) withDO(do gen.Dao) *demoUserDo {
	d.DO = *do.(*gen.DO)
	return d
}