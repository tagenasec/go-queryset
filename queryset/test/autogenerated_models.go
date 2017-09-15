package test

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jirfag/go-queryset/queryset/tmp"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set BlogQuerySet

// BlogQuerySet is an queryset type for Blog
type BlogQuerySet struct {
	db *gorm.DB
}

// NewBlogQuerySet constructs new BlogQuerySet
func NewBlogQuerySet(db *gorm.DB) BlogQuerySet {
	return BlogQuerySet{
		db: db.Model(&Blog{}),
	}
}

func (qs BlogQuerySet) w(db *gorm.DB) BlogQuerySet {
	return NewBlogQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) All(ret *[]Blog) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *Blog) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) CreatedAtEq(createdAt time.Time) BlogQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) CreatedAtGt(createdAt time.Time) BlogQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) CreatedAtGte(createdAt time.Time) BlogQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) CreatedAtLt(createdAt time.Time) BlogQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) CreatedAtLte(createdAt time.Time) BlogQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) CreatedAtNe(createdAt time.Time) BlogQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Blog) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) Delete() error {
	return qs.db.Delete(Blog{}).Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) DeletedAtEq(deletedAt time.Time) BlogQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) DeletedAtGt(deletedAt time.Time) BlogQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) DeletedAtGte(deletedAt time.Time) BlogQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) DeletedAtIsNotNull() BlogQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) DeletedAtIsNull() BlogQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) DeletedAtLt(deletedAt time.Time) BlogQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) DeletedAtLte(deletedAt time.Time) BlogQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) DeletedAtNe(deletedAt time.Time) BlogQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) GetUpdater() BlogUpdater {
	return NewBlogUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) IDEq(ID uint) BlogQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) IDGt(ID uint) BlogQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) IDGte(ID uint) BlogQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) IDIn(ID uint, IDRest ...uint) BlogQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id IN (?)", iArgs))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) IDLt(ID uint) BlogQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) IDLte(ID uint) BlogQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) IDNe(ID uint) BlogQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) IDNotIn(ID uint, IDRest ...uint) BlogQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", iArgs))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) Limit(limit int) BlogQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// NameEq is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) NameEq(name string) BlogQuerySet {
	return qs.w(qs.db.Where("name = ?", name))
}

// NameIn is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) NameIn(name string, nameRest ...string) BlogQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name IN (?)", iArgs))
}

// NameNe is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) NameNe(name string) BlogQuerySet {
	return qs.w(qs.db.Where("name != ?", name))
}

// NameNotIn is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) NameNotIn(name string, nameRest ...string) BlogQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name NOT IN (?)", iArgs))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs BlogQuerySet) One(ret *Blog) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) OrderAscByCreatedAt() BlogQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) OrderAscByDeletedAt() BlogQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) OrderAscByID() BlogQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) OrderAscByUpdatedAt() BlogQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) OrderDescByCreatedAt() BlogQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) OrderDescByDeletedAt() BlogQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) OrderDescByID() BlogQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) OrderDescByUpdatedAt() BlogQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u BlogUpdater) SetCreatedAt(createdAt time.Time) BlogUpdater {
	u.fields[string(BlogDBSchema.CreatedAt)] = createdAt
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u BlogUpdater) SetID(ID uint) BlogUpdater {
	u.fields[string(BlogDBSchema.ID)] = ID
	return u
}

// SetName is an autogenerated method
// nolint: dupl
func (u BlogUpdater) SetName(name string) BlogUpdater {
	u.fields[string(BlogDBSchema.Name)] = name
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u BlogUpdater) SetUpdatedAt(updatedAt time.Time) BlogUpdater {
	u.fields[string(BlogDBSchema.UpdatedAt)] = updatedAt
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u BlogUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) UpdatedAtEq(updatedAt time.Time) BlogQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) UpdatedAtGt(updatedAt time.Time) BlogQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) UpdatedAtGte(updatedAt time.Time) BlogQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) UpdatedAtLt(updatedAt time.Time) BlogQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) UpdatedAtLte(updatedAt time.Time) BlogQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs BlogQuerySet) UpdatedAtNe(updatedAt time.Time) BlogQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// ===== END of query set BlogQuerySet

// ===== BEGIN of Blog modifiers

type blogDBSchemaField string

// BlogDBSchema stores db field names of Blog
var BlogDBSchema = struct {
	ID        blogDBSchemaField
	CreatedAt blogDBSchemaField
	UpdatedAt blogDBSchemaField
	DeletedAt blogDBSchemaField
	Name      blogDBSchemaField
}{

	ID:        blogDBSchemaField("id"),
	CreatedAt: blogDBSchemaField("created_at"),
	UpdatedAt: blogDBSchemaField("updated_at"),
	DeletedAt: blogDBSchemaField("deleted_at"),
	Name:      blogDBSchemaField("name"),
}

// Update updates Blog fields by primary key
func (o *Blog) Update(db *gorm.DB, fields ...blogDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":         o.ID,
		"created_at": o.CreatedAt,
		"updated_at": o.UpdatedAt,
		"deleted_at": o.DeletedAt,
		"name":       o.Name,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := string(f)
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update Blog %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// BlogUpdater is an Blog updates manager
type BlogUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewBlogUpdater creates new Blog updater
func NewBlogUpdater(db *gorm.DB) BlogUpdater {
	return BlogUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Blog{}),
	}
}

// ===== END of Blog modifiers

// ===== BEGIN of query set PostQuerySet

// PostQuerySet is an queryset type for Post
type PostQuerySet struct {
	db *gorm.DB
}

// NewPostQuerySet constructs new PostQuerySet
func NewPostQuerySet(db *gorm.DB) PostQuerySet {
	return PostQuerySet{
		db: db.Model(&Post{}),
	}
}

func (qs PostQuerySet) w(db *gorm.DB) PostQuerySet {
	return NewPostQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) All(ret *[]Post) error {
	return qs.db.Find(ret).Error
}

// BlogIsNotNull is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) BlogIsNotNull() PostQuerySet {
	return qs.w(qs.db.Where("blog IS NOT NULL"))
}

// BlogIsNull is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) BlogIsNull() PostQuerySet {
	return qs.w(qs.db.Where("blog IS NULL"))
}

// Count is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *Post) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) CreatedAtEq(createdAt time.Time) PostQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) CreatedAtGt(createdAt time.Time) PostQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) CreatedAtGte(createdAt time.Time) PostQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) CreatedAtLt(createdAt time.Time) PostQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) CreatedAtLte(createdAt time.Time) PostQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) CreatedAtNe(createdAt time.Time) PostQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) Delete() error {
	return qs.db.Delete(Post{}).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Post) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) DeletedAtEq(deletedAt time.Time) PostQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) DeletedAtGt(deletedAt time.Time) PostQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) DeletedAtGte(deletedAt time.Time) PostQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) DeletedAtIsNotNull() PostQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) DeletedAtIsNull() PostQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) DeletedAtLt(deletedAt time.Time) PostQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) DeletedAtLte(deletedAt time.Time) PostQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) DeletedAtNe(deletedAt time.Time) PostQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) GetUpdater() PostUpdater {
	return NewPostUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) IDEq(ID uint) PostQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) IDGt(ID uint) PostQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) IDGte(ID uint) PostQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) IDIn(ID uint, IDRest ...uint) PostQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id IN (?)", iArgs))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) IDLt(ID uint) PostQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) IDLte(ID uint) PostQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) IDNe(ID uint) PostQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) IDNotIn(ID uint, IDRest ...uint) PostQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", iArgs))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) Limit(limit int) PostQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs PostQuerySet) One(ret *Post) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) OrderAscByCreatedAt() PostQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) OrderAscByDeletedAt() PostQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) OrderAscByID() PostQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) OrderAscByUpdatedAt() PostQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) OrderDescByCreatedAt() PostQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) OrderDescByDeletedAt() PostQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) OrderDescByID() PostQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) OrderDescByUpdatedAt() PostQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// PreloadBlog is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) PreloadBlog() PostQuerySet {
	return qs.w(qs.db.Preload("Blog"))
}

// PreloadUser is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) PreloadUser() PostQuerySet {
	return qs.w(qs.db.Preload("User"))
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u PostUpdater) SetCreatedAt(createdAt time.Time) PostUpdater {
	u.fields[string(PostDBSchema.CreatedAt)] = createdAt
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u PostUpdater) SetID(ID uint) PostUpdater {
	u.fields[string(PostDBSchema.ID)] = ID
	return u
}

// SetStr is an autogenerated method
// nolint: dupl
func (u PostUpdater) SetStr(str tmp.StringDef) PostUpdater {
	u.fields[string(PostDBSchema.Str)] = str
	return u
}

// SetTitle is an autogenerated method
// nolint: dupl
func (u PostUpdater) SetTitle(title string) PostUpdater {
	u.fields[string(PostDBSchema.Title)] = title
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u PostUpdater) SetUpdatedAt(updatedAt time.Time) PostUpdater {
	u.fields[string(PostDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SetUser is an autogenerated method
// nolint: dupl
func (u PostUpdater) SetUser(user User) PostUpdater {
	u.fields[string(PostDBSchema.User)] = user
	return u
}

// StrEq is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) StrEq(str tmp.StringDef) PostQuerySet {
	return qs.w(qs.db.Where("str = ?", str))
}

// StrIn is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) StrIn(str tmp.StringDef, strRest ...tmp.StringDef) PostQuerySet {
	iArgs := []interface{}{str}
	for _, arg := range strRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("str IN (?)", iArgs))
}

// StrNe is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) StrNe(str tmp.StringDef) PostQuerySet {
	return qs.w(qs.db.Where("str != ?", str))
}

// StrNotIn is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) StrNotIn(str tmp.StringDef, strRest ...tmp.StringDef) PostQuerySet {
	iArgs := []interface{}{str}
	for _, arg := range strRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("str NOT IN (?)", iArgs))
}

// TitleEq is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) TitleEq(title string) PostQuerySet {
	return qs.w(qs.db.Where("title = ?", title))
}

// TitleIn is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) TitleIn(title string, titleRest ...string) PostQuerySet {
	iArgs := []interface{}{title}
	for _, arg := range titleRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("title IN (?)", iArgs))
}

// TitleNe is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) TitleNe(title string) PostQuerySet {
	return qs.w(qs.db.Where("title != ?", title))
}

// TitleNotIn is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) TitleNotIn(title string, titleRest ...string) PostQuerySet {
	iArgs := []interface{}{title}
	for _, arg := range titleRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("title NOT IN (?)", iArgs))
}

// Update is an autogenerated method
// nolint: dupl
func (u PostUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) UpdatedAtEq(updatedAt time.Time) PostQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) UpdatedAtGt(updatedAt time.Time) PostQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) UpdatedAtGte(updatedAt time.Time) PostQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) UpdatedAtLt(updatedAt time.Time) PostQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) UpdatedAtLte(updatedAt time.Time) PostQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs PostQuerySet) UpdatedAtNe(updatedAt time.Time) PostQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// ===== END of query set PostQuerySet

// ===== BEGIN of Post modifiers

type postDBSchemaField string

// PostDBSchema stores db field names of Post
var PostDBSchema = struct {
	ID        postDBSchemaField
	CreatedAt postDBSchemaField
	UpdatedAt postDBSchemaField
	DeletedAt postDBSchemaField
	Blog      postDBSchemaField
	User      postDBSchemaField
	Title     postDBSchemaField
	Str       postDBSchemaField
}{

	ID:        postDBSchemaField("id"),
	CreatedAt: postDBSchemaField("created_at"),
	UpdatedAt: postDBSchemaField("updated_at"),
	DeletedAt: postDBSchemaField("deleted_at"),
	Blog:      postDBSchemaField("blog"),
	User:      postDBSchemaField("user"),
	Title:     postDBSchemaField("title"),
	Str:       postDBSchemaField("str"),
}

// Update updates Post fields by primary key
func (o *Post) Update(db *gorm.DB, fields ...postDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":         o.ID,
		"created_at": o.CreatedAt,
		"updated_at": o.UpdatedAt,
		"deleted_at": o.DeletedAt,
		"blog":       o.Blog,
		"user":       o.User,
		"title":      o.Title,
		"str":        o.Str,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := string(f)
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update Post %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// PostUpdater is an Post updates manager
type PostUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewPostUpdater creates new Post updater
func NewPostUpdater(db *gorm.DB) PostUpdater {
	return PostUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Post{}),
	}
}

// ===== END of Post modifiers

// ===== BEGIN of query set UserQuerySet

// UserQuerySet is an queryset type for User
type UserQuerySet struct {
	db *gorm.DB
}

// NewUserQuerySet constructs new UserQuerySet
func NewUserQuerySet(db *gorm.DB) UserQuerySet {
	return UserQuerySet{
		db: db.Model(&User{}),
	}
}

func (qs UserQuerySet) w(db *gorm.DB) UserQuerySet {
	return NewUserQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) All(ret *[]User) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *User) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtEq(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtGt(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtGte(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtLt(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtLte(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtNe(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *User) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) Delete() error {
	return qs.db.Delete(User{}).Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtEq(deletedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtGt(deletedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtGte(deletedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtIsNotNull() UserQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtIsNull() UserQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtLt(deletedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtLte(deletedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtNe(deletedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// EmailEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) EmailEq(email string) UserQuerySet {
	return qs.w(qs.db.Where("email = ?", email))
}

// EmailIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) EmailIn(email string, emailRest ...string) UserQuerySet {
	iArgs := []interface{}{email}
	for _, arg := range emailRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("email IN (?)", iArgs))
}

// EmailNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) EmailNe(email string) UserQuerySet {
	return qs.w(qs.db.Where("email != ?", email))
}

// EmailNotIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) EmailNotIn(email string, emailRest ...string) UserQuerySet {
	iArgs := []interface{}{email}
	for _, arg := range emailRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("email NOT IN (?)", iArgs))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) GetUpdater() UserUpdater {
	return NewUserUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDEq(ID uint) UserQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDGt(ID uint) UserQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDGte(ID uint) UserQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDIn(ID uint, IDRest ...uint) UserQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id IN (?)", iArgs))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDLt(ID uint) UserQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDLte(ID uint) UserQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDNe(ID uint) UserQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDNotIn(ID uint, IDRest ...uint) UserQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", iArgs))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) Limit(limit int) UserQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// NameEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) NameEq(name string) UserQuerySet {
	return qs.w(qs.db.Where("name = ?", name))
}

// NameIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) NameIn(name string, nameRest ...string) UserQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name IN (?)", iArgs))
}

// NameNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) NameNe(name string) UserQuerySet {
	return qs.w(qs.db.Where("name != ?", name))
}

// NameNotIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) NameNotIn(name string, nameRest ...string) UserQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name NOT IN (?)", iArgs))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs UserQuerySet) One(ret *User) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderAscByCreatedAt() UserQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderAscByDeletedAt() UserQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderAscByID() UserQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderAscByUpdatedAt() UserQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderDescByCreatedAt() UserQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderDescByDeletedAt() UserQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderDescByID() UserQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderDescByUpdatedAt() UserQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetCreatedAt(createdAt time.Time) UserUpdater {
	u.fields[string(UserDBSchema.CreatedAt)] = createdAt
	return u
}

// SetEmail is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetEmail(email string) UserUpdater {
	u.fields[string(UserDBSchema.Email)] = email
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetID(ID uint) UserUpdater {
	u.fields[string(UserDBSchema.ID)] = ID
	return u
}

// SetName is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetName(name string) UserUpdater {
	u.fields[string(UserDBSchema.Name)] = name
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetUpdatedAt(updatedAt time.Time) UserUpdater {
	u.fields[string(UserDBSchema.UpdatedAt)] = updatedAt
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u UserUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtEq(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtGt(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtGte(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtLt(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtLte(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtNe(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// ===== END of query set UserQuerySet

// ===== BEGIN of User modifiers

type userDBSchemaField string

// UserDBSchema stores db field names of User
var UserDBSchema = struct {
	ID        userDBSchemaField
	CreatedAt userDBSchemaField
	UpdatedAt userDBSchemaField
	DeletedAt userDBSchemaField
	Name      userDBSchemaField
	Email     userDBSchemaField
}{

	ID:        userDBSchemaField("id"),
	CreatedAt: userDBSchemaField("created_at"),
	UpdatedAt: userDBSchemaField("updated_at"),
	DeletedAt: userDBSchemaField("deleted_at"),
	Name:      userDBSchemaField("name"),
	Email:     userDBSchemaField("email"),
}

// Update updates User fields by primary key
func (o *User) Update(db *gorm.DB, fields ...userDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":         o.ID,
		"created_at": o.CreatedAt,
		"updated_at": o.UpdatedAt,
		"deleted_at": o.DeletedAt,
		"name":       o.Name,
		"email":      o.Email,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := string(f)
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update User %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// UserUpdater is an User updates manager
type UserUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewUserUpdater creates new User updater
func NewUserUpdater(db *gorm.DB) UserUpdater {
	return UserUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&User{}),
	}
}

// ===== END of User modifiers

// ===== END of all query sets
