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

	"gorm.io/plugin/dbresolver"

	"github.com/TremblingV5/DouTok/applications/favorite/dal/model"
)

func newFavoriteCount(db *gorm.DB, opts ...gen.DOOption) favoriteCount {
	_favoriteCount := favoriteCount{}

	_favoriteCount.favoriteCountDo.UseDB(db, opts...)
	_favoriteCount.favoriteCountDo.UseModel(&model.FavoriteCount{})

	tableName := _favoriteCount.favoriteCountDo.TableName()
	_favoriteCount.ALL = field.NewAsterisk(tableName)
	_favoriteCount.VideoId = field.NewInt64(tableName, "video_id")
	_favoriteCount.Number = field.NewInt64(tableName, "number")
	_favoriteCount.CreatedAt = field.NewTime(tableName, "created_at")
	_favoriteCount.UpdatedAt = field.NewTime(tableName, "updated_at")

	_favoriteCount.fillFieldMap()

	return _favoriteCount
}

type favoriteCount struct {
	favoriteCountDo

	ALL       field.Asterisk
	VideoId   field.Int64
	Number    field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (f favoriteCount) Table(newTableName string) *favoriteCount {
	f.favoriteCountDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f favoriteCount) As(alias string) *favoriteCount {
	f.favoriteCountDo.DO = *(f.favoriteCountDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *favoriteCount) updateTableName(table string) *favoriteCount {
	f.ALL = field.NewAsterisk(table)
	f.VideoId = field.NewInt64(table, "video_id")
	f.Number = field.NewInt64(table, "number")
	f.CreatedAt = field.NewTime(table, "created_at")
	f.UpdatedAt = field.NewTime(table, "updated_at")

	f.fillFieldMap()

	return f
}

func (f *favoriteCount) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *favoriteCount) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 4)
	f.fieldMap["video_id"] = f.VideoId
	f.fieldMap["number"] = f.Number
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
}

func (f favoriteCount) clone(db *gorm.DB) favoriteCount {
	f.favoriteCountDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f favoriteCount) replaceDB(db *gorm.DB) favoriteCount {
	f.favoriteCountDo.ReplaceDB(db)
	return f
}

type favoriteCountDo struct{ gen.DO }

type IFavoriteCountDo interface {
	gen.SubQuery
	Debug() IFavoriteCountDo
	WithContext(ctx context.Context) IFavoriteCountDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFavoriteCountDo
	WriteDB() IFavoriteCountDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFavoriteCountDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFavoriteCountDo
	Not(conds ...gen.Condition) IFavoriteCountDo
	Or(conds ...gen.Condition) IFavoriteCountDo
	Select(conds ...field.Expr) IFavoriteCountDo
	Where(conds ...gen.Condition) IFavoriteCountDo
	Order(conds ...field.Expr) IFavoriteCountDo
	Distinct(cols ...field.Expr) IFavoriteCountDo
	Omit(cols ...field.Expr) IFavoriteCountDo
	Join(table schema.Tabler, on ...field.Expr) IFavoriteCountDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFavoriteCountDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFavoriteCountDo
	Group(cols ...field.Expr) IFavoriteCountDo
	Having(conds ...gen.Condition) IFavoriteCountDo
	Limit(limit int) IFavoriteCountDo
	Offset(offset int) IFavoriteCountDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFavoriteCountDo
	Unscoped() IFavoriteCountDo
	Create(values ...*model.FavoriteCount) error
	CreateInBatches(values []*model.FavoriteCount, batchSize int) error
	Save(values ...*model.FavoriteCount) error
	First() (*model.FavoriteCount, error)
	Take() (*model.FavoriteCount, error)
	Last() (*model.FavoriteCount, error)
	Find() ([]*model.FavoriteCount, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FavoriteCount, err error)
	FindInBatches(result *[]*model.FavoriteCount, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FavoriteCount) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFavoriteCountDo
	Assign(attrs ...field.AssignExpr) IFavoriteCountDo
	Joins(fields ...field.RelationField) IFavoriteCountDo
	Preload(fields ...field.RelationField) IFavoriteCountDo
	FirstOrInit() (*model.FavoriteCount, error)
	FirstOrCreate() (*model.FavoriteCount, error)
	FindByPage(offset int, limit int) (result []*model.FavoriteCount, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFavoriteCountDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f favoriteCountDo) Debug() IFavoriteCountDo {
	return f.withDO(f.DO.Debug())
}

func (f favoriteCountDo) WithContext(ctx context.Context) IFavoriteCountDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f favoriteCountDo) ReadDB() IFavoriteCountDo {
	return f.Clauses(dbresolver.Read)
}

func (f favoriteCountDo) WriteDB() IFavoriteCountDo {
	return f.Clauses(dbresolver.Write)
}

func (f favoriteCountDo) Session(config *gorm.Session) IFavoriteCountDo {
	return f.withDO(f.DO.Session(config))
}

func (f favoriteCountDo) Clauses(conds ...clause.Expression) IFavoriteCountDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f favoriteCountDo) Returning(value interface{}, columns ...string) IFavoriteCountDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f favoriteCountDo) Not(conds ...gen.Condition) IFavoriteCountDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f favoriteCountDo) Or(conds ...gen.Condition) IFavoriteCountDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f favoriteCountDo) Select(conds ...field.Expr) IFavoriteCountDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f favoriteCountDo) Where(conds ...gen.Condition) IFavoriteCountDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f favoriteCountDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IFavoriteCountDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f favoriteCountDo) Order(conds ...field.Expr) IFavoriteCountDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f favoriteCountDo) Distinct(cols ...field.Expr) IFavoriteCountDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f favoriteCountDo) Omit(cols ...field.Expr) IFavoriteCountDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f favoriteCountDo) Join(table schema.Tabler, on ...field.Expr) IFavoriteCountDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f favoriteCountDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFavoriteCountDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f favoriteCountDo) RightJoin(table schema.Tabler, on ...field.Expr) IFavoriteCountDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f favoriteCountDo) Group(cols ...field.Expr) IFavoriteCountDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f favoriteCountDo) Having(conds ...gen.Condition) IFavoriteCountDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f favoriteCountDo) Limit(limit int) IFavoriteCountDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f favoriteCountDo) Offset(offset int) IFavoriteCountDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f favoriteCountDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFavoriteCountDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f favoriteCountDo) Unscoped() IFavoriteCountDo {
	return f.withDO(f.DO.Unscoped())
}

func (f favoriteCountDo) Create(values ...*model.FavoriteCount) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f favoriteCountDo) CreateInBatches(values []*model.FavoriteCount, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f favoriteCountDo) Save(values ...*model.FavoriteCount) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f favoriteCountDo) First() (*model.FavoriteCount, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FavoriteCount), nil
	}
}

func (f favoriteCountDo) Take() (*model.FavoriteCount, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FavoriteCount), nil
	}
}

func (f favoriteCountDo) Last() (*model.FavoriteCount, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FavoriteCount), nil
	}
}

func (f favoriteCountDo) Find() ([]*model.FavoriteCount, error) {
	result, err := f.DO.Find()
	return result.([]*model.FavoriteCount), err
}

func (f favoriteCountDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FavoriteCount, err error) {
	buf := make([]*model.FavoriteCount, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f favoriteCountDo) FindInBatches(result *[]*model.FavoriteCount, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f favoriteCountDo) Attrs(attrs ...field.AssignExpr) IFavoriteCountDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f favoriteCountDo) Assign(attrs ...field.AssignExpr) IFavoriteCountDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f favoriteCountDo) Joins(fields ...field.RelationField) IFavoriteCountDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f favoriteCountDo) Preload(fields ...field.RelationField) IFavoriteCountDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f favoriteCountDo) FirstOrInit() (*model.FavoriteCount, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FavoriteCount), nil
	}
}

func (f favoriteCountDo) FirstOrCreate() (*model.FavoriteCount, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FavoriteCount), nil
	}
}

func (f favoriteCountDo) FindByPage(offset int, limit int) (result []*model.FavoriteCount, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f favoriteCountDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f favoriteCountDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f favoriteCountDo) Delete(models ...*model.FavoriteCount) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *favoriteCountDo) withDO(do gen.Dao) *favoriteCountDo {
	f.DO = *do.(*gen.DO)
	return f
}
