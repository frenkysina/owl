package dao

import (
	"github.com/ibanyu/owl/service/task"
)

type BackupImpl struct {
}

var BackupDAO BackupImpl

func (BackupImpl) AddBackup(backup *task.DbInjectionBackup) (int64, error) {
	err := GetDB().Create(backup).Error
	return backup.ID, err
}

func (BackupImpl) UpdateBackup(backup *task.DbInjectionBackup) error {
	return GetDB().Model(backup).Where("id = ?", backup.ID).Updates(backup).Error
}

func (BackupImpl) GetBackupInfoById(id int64) (*task.DbInjectionBackup, error) {
	var backup task.DbInjectionBackup
	return &backup, GetDB().Where("id = ?", id).First(&backup).Error
}
