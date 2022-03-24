package dao

import (
	"gostrengthen/week02"
)

func Select() (err error, dataList []Agency) {
	err = week02.DB.QueryRow("SELECT * FROM `agency_info` WHERE  is_deleted=0 and  agency_name = '222' ").Scan(&dataList)
	return err, dataList
}
