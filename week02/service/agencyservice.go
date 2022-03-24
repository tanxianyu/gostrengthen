package service

import (
	"database/sql"
	"gostrengthen/week02/dao"

	"github.com/pkg/errors"
)

func Select() (err error, dataList []Agency) {
	err, dataList = dao.Select()
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.Wrap(err, "数据查询为空")
		}
	}
	return err, dataList
}
