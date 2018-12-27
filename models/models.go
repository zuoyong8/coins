package models

import (
	"time"

     "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Base struct{
	Id 				uint		`gorm:"primary_key"`	
}

type Coins struct{
	Base
	Currency 		string
	FromAddress 	string
	ToAddress		string
	Amount			float64
	TranferTime 	time.Time
	Confirmations 	int32
	BlockNumber		int64
	Txid			string		`gorm:"unique_index"`
}

type Users struct{
	Base
	Username		string		`gorm:"unique_index"`
	Pwdsalt		    string
	Password		string
	CreatAt			time.Time
}


var DB *gorm.DB

func InitDB() (*gorm.DB,error) {
	db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		DB = db
		db.LogMode(true)
		DB.AutoMigrate(&Coins{})
		DB.AutoMigrate(&Users{})
		return db, err
	}
	return db,err
}


//Users CRUD
func (users *Users) Insert() error{
	return DB.Create(users).Error
}

func (users *Users) Delete() error {
	return DB.Delete(users).Error
}

func GetUserByUsername(username string)(*Users,error){
	var users Users
	err := DB.First(&users, "username = ?", username).Error
	return &users, err
}

//Coins CRUD
func (coins *Coins) Insert() error {
	return DB.Create(coins).Error
}

func (coins *Coins) Update() error {
	return DB.Model(coins).Updates(map[string]interface{}{
		"currency":        coins.Currency,
	}).Error
}

func (coins *Coins) Delete() error {
	return DB.Delete(coins).Error
}

func GetCoinsByTxid(txid string)(*Coins,error){
	var coins Coins
	err := DB.First(&coins, "txid = ?", txid).Error
	return &coins, err
}

func GetCoinsByCurrency(currency string)([]*Coins,error){
	var coins []*Coins
	var err error
	err = DB.Where("currency=?",currency).Find(&coins).Error
	return coins,err
}

func Count() int {
	var count int
	DB.Model(&Coins{}).Count(&count)
	return count
}