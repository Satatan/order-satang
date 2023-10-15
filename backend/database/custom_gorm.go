package database

import (
	"database/sql"

	"gorm.io/gorm"
)

type CustomGorm interface {
	Table(name string) CustomGorm
	Model(value interface{}) CustomGorm
	Order(value interface{}) CustomGorm
	Where(query interface{}, args ...interface{}) CustomGorm
	Joins(query string, args ...interface{}) CustomGorm
	Select(query interface{}, args ...interface{}) CustomGorm
	Scan(dest interface{}) CustomGorm
	First(out interface{}, where ...interface{}) CustomGorm
	Last(out interface{}, where ...interface{}) CustomGorm
	Preload(column string, conditions ...interface{}) CustomGorm
	Raw(sql string, values ...interface{}) CustomGorm
	Limit(limit int) CustomGorm
	Offset(offset int) CustomGorm
	Count(value *int64) CustomGorm
	Group(query string) CustomGorm
	Having(query interface{}, values ...interface{}) CustomGorm
	Omit(columns ...string) CustomGorm
	Or(query interface{}, args ...interface{}) CustomGorm

	Find(out interface{}, where ...interface{}) CustomGorm
	Take(out interface{}, where ...interface{}) CustomGorm
	Create(value interface{}) CustomGorm
	Updates(values interface{}) CustomGorm
	Update(column string, attrs interface{}) CustomGorm
	Save(value interface{}) CustomGorm
	Delete(value interface{}, where ...interface{}) CustomGorm

	Transaction(fc func(tx CustomGorm) error) (err error)
	Pluck(column string, value interface{}) CustomGorm

	Debug() CustomGorm
	Error() error
	RowsAffected() int64
	Rows() (*sql.Rows, error)
}

type grcGorm struct {
	db *gorm.DB
}

func Wrap(db *gorm.DB) CustomGorm {
	return &grcGorm{db: db}
}

func (d *grcGorm) Table(name string) CustomGorm {
	return Wrap(d.db.Table(name))
}

func (d *grcGorm) Model(value interface{}) CustomGorm {
	return Wrap(d.db.Model(value))
}

func (d *grcGorm) Order(value interface{}) CustomGorm {
	return Wrap(d.db.Order(value))
}

func (d *grcGorm) Where(query interface{}, args ...interface{}) CustomGorm {
	return Wrap(d.db.Where(query, args...))
}

func (d *grcGorm) Joins(query string, args ...interface{}) CustomGorm {
	return Wrap(d.db.Joins(query, args...))
}

func (d *grcGorm) Select(query interface{}, args ...interface{}) CustomGorm {
	return Wrap(d.db.Select(query, args...))
}

func (d *grcGorm) Scan(dest interface{}) CustomGorm {
	return Wrap(d.db.Scan(dest))
}

func (d *grcGorm) First(out interface{}, where ...interface{}) CustomGorm {
	return Wrap(d.db.First(out, where...))
}

func (d *grcGorm) Last(out interface{}, where ...interface{}) CustomGorm {
	return Wrap(d.db.Last(out, where...))
}

func (d *grcGorm) Preload(column string, conditions ...interface{}) CustomGorm {
	return Wrap(d.db.Preload(column, conditions...))
}

func (d *grcGorm) Raw(sql string, values ...interface{}) CustomGorm {
	return Wrap(d.db.Raw(sql, values...))
}

func (d *grcGorm) Limit(limit int) CustomGorm {
	return Wrap(d.db.Limit(limit))
}

func (d *grcGorm) Offset(offset int) CustomGorm {
	return Wrap(d.db.Offset(offset))
}

func (d *grcGorm) Count(value *int64) CustomGorm {
	return Wrap(d.db.Count(value))
}

func (d *grcGorm) Group(query string) CustomGorm {
	return Wrap(d.db.Group(query))
}

func (d *grcGorm) Having(query interface{}, values ...interface{}) CustomGorm {
	return Wrap(d.db.Having(query, values...))
}

func (d *grcGorm) Omit(columns ...string) CustomGorm {
	return Wrap(d.db.Omit(columns...))
}

func (d *grcGorm) Or(query interface{}, args ...interface{}) CustomGorm {
	return Wrap(d.db.Or(query, args...))
}

func (d *grcGorm) Find(out interface{}, where ...interface{}) CustomGorm {
	return Wrap(d.db.Find(out, where...))
}

func (d *grcGorm) Take(out interface{}, where ...interface{}) CustomGorm {
	return Wrap(d.db.Take(out, where...))
}

func (d *grcGorm) Create(value interface{}) CustomGorm {
	return Wrap(d.db.Create(value))
}

func (d *grcGorm) Updates(values interface{}) CustomGorm {
	return Wrap(d.db.Updates(values))
}

func (d *grcGorm) Update(column string, attrs interface{}) CustomGorm {
	return Wrap(d.db.Update(column, attrs))
}

func (d *grcGorm) Save(value interface{}) CustomGorm {
	return Wrap(d.db.Save(value))
}

func (d *grcGorm) Delete(value interface{}, where ...interface{}) CustomGorm {
	return Wrap(d.db.Delete(value, where...))
}

func (d *grcGorm) Transaction(fc func(tx CustomGorm) error) (err error) {
	return d.db.Transaction(func(tx *gorm.DB) error {
		return fc(Wrap(tx))
	})
}

func (d *grcGorm) Pluck(column string, value interface{}) CustomGorm {
	return Wrap(d.db.Pluck(column, value))
}

func (d *grcGorm) Debug() CustomGorm {
	return Wrap(d.db.Debug())
}

func (d *grcGorm) Error() error {
	return d.db.Error
}

func (d *grcGorm) RowsAffected() int64 {
	return d.db.RowsAffected
}

func (d *grcGorm) Rows() (*sql.Rows, error) {
	return d.db.Rows()
}
