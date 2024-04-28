/*
 * @Author: huangcheng1 huangcheng1@sensetime.com
 * @Date: 2024-04-28 12:16:50
 * @LastEditors: huangcheng1 huangcheng1@sensetime.com
 * @LastEditTime: 2024-04-28 12:18:23
 * @FilePath: /dao-generator/pkg/generator/test_file/simple.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package testfile

import (
	"database/sql"
	"database/sql/driver"
)

type ApprovalState int

const (
	ApprovalStateUndefined ApprovalState = iota
	// [zh] 审批中
	// [en] pending
	ApprovalStatePending
	// [zh] 通过
	// [en] approved
	ApprovalStateApproved
	// [zh] 驳回
	// [en] rejected
	ApprovalStateRejected
	// [zh] 取消
	// [en] canceled
	ApprovalStateCanceled
)

// 上传任务类型。
type TaskType int

const (
	// 本地上传。
	TaskTypeLocalUpload TaskType = iota
	// 对象存储上传。
	TaskTypeAossUpload
	// web端上传。
	TaskTypeWebUpload
	// senseData上传。
	TaskTypeSensedataUpload
	// 数据脱敏。
	TaskTypeMosaicUpload
	// 导出到本地。
	TaskTypeLocalDownload
	// 导出到对象存储。
	TaskTypeAossDownload
	// 编辑文件。
	TaskTypeUpdateFile
	// 新建文件。
	TaskTypeAddFile
	// 删除文件。
	TaskTypeDeleteFile
	// 导出到Ceph。
	TaskTypeCephDownload
	// 导出到Luster。
	TaskTypeLustreDownload
)

type RepoInfo struct {
	// 数据名称
	ID string `json:"id,omitempty"`
	// 数据名称
	Name string `json:"name,omitempty"`
	// 数据大小
	Size int64 `json:"size,omitempty"`
	// 数据文件数量
	FileNum int64 `json:"file_num,omitempty"`
}

var _ driver.Valuer = (*RepoInfo)(nil)
var _ sql.Scanner = (*RepoInfo)(nil)

func (i *RepoInfo) Value() (driver.Value, error) {
	return value(i)
}

func (i *RepoInfo) Scan(src any) error {
	return scan(i, src)
}
