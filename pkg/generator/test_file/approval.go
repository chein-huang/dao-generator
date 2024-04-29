/*
 * @Author: huangcheng1 huangcheng1@sensetime.com
 * @Date: 2024-04-28 10:49:04
 * @LastEditors: huangcheng1 huangcheng1@sensetime.com
 * @LastEditTime: 2024-04-29 16:50:24
 * @FilePath: /dao-generator/resource/testmodel/approval.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package testfile

import (
	"time"

	"gorm.io/gorm"
)

// gen:"ApprovalInfo,tableName:approval_info,flags:isPreload"
// gen:"package:github.com/chein-huang/dao-generator/pkg/generator/test_file,packageAliceName:testfile"
type ApprovalInfoWithAuth struct {
	*ApprovalInfo
	// gen:"orderBy:created_at DESC"
	AuthInfo *AuthInfo `gorm:"foreignKey:AuthID"`
}
type ApprovalInfo struct {
	ID        string `gorm:"primaryKey;type:varchar(36)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	CreatorID   string `gorm:"type:varchar(36);index"`
	CreatorName string `gorm:"type:varchar(256);index"`
	TenantID    string `gorm:"type:varchar(36);index"`

	State ApprovalState `gorm:"index"`

	// 数据id
	RepoID         string    `gorm:"type:varchar(36);index"`
	SourceRepoName string    `gorm:"type:varchar(256);index"`
	RepoInfo       *RepoInfo `gorm:"type:MEDIUMTEXT"`
	TaskID         string    `gorm:"type:varchar(36);index"`
	// task name
	TaskName        string   `gorm:"type:varchar(256);index"`
	TaskType        TaskType `gorm:"type:tinyint(1);index"`
	SourceDataspace string   `gorm:"type:varchar(256);index"`
	AuthID          string   `gorm:"type:varchar(36)"`
}

// 授权信息
type AuthInfo struct {
	ID        string `gorm:"primaryKey;type:varchar(36)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	CreatorID   string `gorm:"index"`
	CreatorName string `gorm:"type:varchar(256)"`
	TenantID    string `gorm:"type:varchar(36);index"`
}
