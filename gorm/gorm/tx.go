package gorm

import (
	"errors"
	"gorm.io/gorm"
)

// Transaction 普通事务
func Transaction() {
	t := teacherTemp
	t1 := teacherTemp
	// 返回错误就会回滚，返回nil就会提交
	DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&t).Error; err != nil {
			return err
		}
		if err := tx.Create(&t1).Error; err != nil {
			return err
		}
		return nil
	})
}

// 嵌套事务
func NestTransaction() {
	t1 := teacherTemp
	t2 := teacherTemp
	t3 := teacherTemp
	t4 := teacherTemp
	// 返回错误就会回滚，返回nil就会提交
	DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&t1).Error; err != nil {
			return err
		}
		// 回滚子事务,不影响大事务最终结果,也就是t1，t4没出错的情况下，整体事务不会回滚，只回滚当前子事务内容
		tx.Transaction(func(tx1 *gorm.DB) error {
			tx1.Create(t2)
			return errors.New("rollback t2")
		})
		tx.Transaction(func(tx2 *gorm.DB) error {
			if err := tx2.Create(&t3).Error; err != nil {
				return err
			}
			return nil
		})
		if err := tx.Create(&t4).Error; err != nil {
			return err
		}
		return nil
	})
}
