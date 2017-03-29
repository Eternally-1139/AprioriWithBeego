package models

import (
	"github.com/astaxie/beego/orm"
)

type BankData struct {
	Id            	string		`orm:"pk"`
	Age 		int   		`orm:"null"`
	Sex 		string   	`orm:"null"`
	Region 		string   	`orm:"null"`
	Income 		float64   	`orm:"null"`
	Married 	string   	`orm:"null"`
	Children 	int      	`orm:"null"`
	Car	 	string   	`orm:"null"`
	SaveAct	 	string   	`orm:"null"`
}


func (this *BankData) Insert() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *BankData) Read() error {
	return orm.NewOrm().QueryTable("bank_data").Filter("id", this.Id).RelatedSel().One(this)
}

func (this *BankData) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

