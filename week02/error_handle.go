package week02

import (
	"database/sql"
	perrors "github.com/pkg/errors"
)

/*
第二周作业：
    我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

答：我认为需要wrap这个error并抛给上层
*/

// 引擎层查询可能返回sql.ErrNoRows错误
func query() (*sql.Rows, error) {
	return nil, sql.ErrNoRows
}

// dao层调用query并不知道返回什么错误，所以需要先wrap再网上抛
func dao() error {
	_, err := query()
	if err != nil {
		return perrors.Wrap(err, "query db error")
	}

	return nil
}

// 业务层拿到错误，判断是不是sql.ErrNoRows错误
func service() error {
	err := dao()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		} else {
			return err
		}
	}

	return nil
}

