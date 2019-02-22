package model

// orm docs: http://doc.gorm.io/database.html

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "time"
)

type Property struct {
	gorm.Model
	Title    string  `gorm:"unique" json:"title"`
	Url      string  `gorm:"unique" json:"url"`
	Location string  `json:"location"`
	Active   bool    `json:"active"`
	Asset    []Asset `gorm:"ForeignKey:PropertyID" json:"assets"`
}

func (p *Property) SetActive() {
	p.Active = true
}

func (p *Property) SetInactive() {
	p.Active = false
}

type Asset struct {
	gorm.Model
	Url        string `gorm:"unique" json:"url"`
	Priority   string `gorm:"type:ENUM('0', '1', '2', '3');default:'0'" json:"priority"`
	Active     bool   `json:"active"`
	PropertyID uint   `json:"property_id"`
}

// migration
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&Property{}, &Asset{})
	db.Model(&Asset{}).AddForeignKey("property_id", "properties(id)", "CASCADE", "RESTRICT")
	return db
}
