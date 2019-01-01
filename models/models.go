package models

import (
	"time"
	"fmt"
	"regexp"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql" 
	"github.com/zuoyong8/coins/config"
	"github.com/qor/validations"
	// "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/asaskevich/govalidator"
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

const (
	EMAIL_VALID_REGEX = "^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+"
	PHONE_VALID__REGEX = "^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}$"
)

var DB *gorm.DB

func InitDB() (*gorm.DB,error) {
	dbInfo,err := config.GetDbConectInfo()
	if err == nil{
		str := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",dbInfo.Username,dbInfo.Password,dbInfo.Dbname)
		db, err := gorm.Open("mysql", str)
		if err == nil {
			DB = db
			db.LogMode(true)
			DB.AutoMigrate(&Coins{})
			DB.AutoMigrate(&Users{})
			validations.RegisterCallbacks(DB)
			return db, err
		}
	}
	return nil,err
}

//Users CRUD
func validEmail(email string) (b bool) {
	reg := regexp.MustCompile(EMAIL_VALID_REGEX)
	return reg.MatchString(email)
}

func (user *Users)Validate(DB *gorm.DB){
	if !validEmail(user.Username){
		DB.AddError(validations.NewError(nil,"Username","用户名必须为邮箱"))
	}
}


func (user *Users) Insert() error {
	return DB.Create(user).Error
}

func (user *Users) Delete() error {
	return DB.Delete(user).Error
}

func GetUsersByUsername(username string)(*Users,error){
	var user Users
	err := DB.First(&user, "username = ?", username).Error
	return &user, err
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