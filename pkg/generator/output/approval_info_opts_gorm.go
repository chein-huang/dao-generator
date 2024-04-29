
package dao

import (
	"time"

	"gorm.io/gorm"

	"gitlab.bj.sensetime.com/elementary/graviton/graviton-data-compliance-service/pkg/model"
)

type Queries struct{}

// ApprovalInfoQs 这种方式避免其他开发者用错option
var ApprovalInfoQs Queries
func (*Queries) WithID(v string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		queryParams := newQueryParams("id = ?", v)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithInID(vs []string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("id in (?)", vs)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithCreatedAt(v time.Time) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		queryParams := newQueryParams("created_at = ?", v)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithInCreatedAt(vs []time.Time) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("created_at in (?)", vs)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithUpdatedAt(v time.Time) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		queryParams := newQueryParams("updated_at = ?", v)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithInUpdatedAt(vs []time.Time) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("updated_at in (?)", vs)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithDeletedAt(v gorm.DeletedAt) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		queryParams := newQueryParams("deleted_at = ?", v)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithInDeletedAt(vs []gorm.DeletedAt) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("deleted_at in (?)", vs)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithCreatorID(v string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		queryParams := newQueryParams("creator_id = ?", v)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithInCreatorID(vs []string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("creator_id in (?)", vs)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithCreatorName(v string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		queryParams := newQueryParams("creator_name = ?", v)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithInCreatorName(vs []string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("creator_name in (?)", vs)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithTenantID(v string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		queryParams := newQueryParams("tenant_id = ?", v)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithInTenantID(vs []string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("tenant_id in (?)", vs)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithState(v ApprovalState) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		queryParams := newQueryParams("state = ?", v)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithInState(vs []ApprovalState) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("state in (?)", vs)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithRepoID(v string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		queryParams := newQueryParams("repo_id = ?", v)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithInRepoID(vs []string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("repo_id in (?)", vs)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithSourceRepoName(v string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		queryParams := newQueryParams("source_repo_name = ?", v)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithInSourceRepoName(vs []string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("source_repo_name in (?)", vs)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithRepoInfo(v RepoInfo) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		queryParams := newQueryParams("repo_info = ?", v)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithInRepoInfo(vs []RepoInfo) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("repo_info in (?)", vs)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithTaskID(v string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		queryParams := newQueryParams("task_id = ?", v)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithInTaskID(vs []string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("task_id in (?)", vs)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithTaskName(v string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		queryParams := newQueryParams("task_name = ?", v)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithInTaskName(vs []string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("task_name in (?)", vs)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithTaskType(v TaskType) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		queryParams := newQueryParams("task_type = ?", v)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithInTaskType(vs []TaskType) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("task_type in (?)", vs)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithSourceDataspace(v string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		queryParams := newQueryParams("source_dataspace = ?", v)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithInSourceDataspace(vs []string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("source_dataspace in (?)", vs)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithAuthID(v string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		queryParams := newQueryParams("auth_id = ?", v)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithInAuthID(vs []string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("auth_id in (?)", vs)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

func (*Queries) WithPreloadAuthInfo() DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		opt.tx = opt.tx.Preload("AuthInfo", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC")
		})
	})
}
func (*Queries) WithPreloadID() DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        opt.tx = opt.tx.Preload("ID")
	})
}
func (*Queries) WithPreloadCreatedAt() DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        opt.tx = opt.tx.Preload("CreatedAt")
	})
}
func (*Queries) WithPreloadUpdatedAt() DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        opt.tx = opt.tx.Preload("UpdatedAt")
	})
}
func (*Queries) WithPreloadDeletedAt() DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        opt.tx = opt.tx.Preload("DeletedAt")
	})
}
func (*Queries) WithPreloadCreatorID() DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        opt.tx = opt.tx.Preload("CreatorID")
	})
}
func (*Queries) WithPreloadCreatorName() DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        opt.tx = opt.tx.Preload("CreatorName")
	})
}
func (*Queries) WithPreloadTenantID() DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        opt.tx = opt.tx.Preload("TenantID")
	})
}
func (*Queries) WithPreloadState() DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        opt.tx = opt.tx.Preload("State")
	})
}
func (*Queries) WithPreloadRepoID() DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        opt.tx = opt.tx.Preload("RepoID")
	})
}
func (*Queries) WithPreloadSourceRepoName() DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        opt.tx = opt.tx.Preload("SourceRepoName")
	})
}
func (*Queries) WithPreloadRepoInfo() DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        opt.tx = opt.tx.Preload("RepoInfo")
	})
}
func (*Queries) WithPreloadTaskID() DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        opt.tx = opt.tx.Preload("TaskID")
	})
}
func (*Queries) WithPreloadTaskName() DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        opt.tx = opt.tx.Preload("TaskName")
	})
}
func (*Queries) WithPreloadTaskType() DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        opt.tx = opt.tx.Preload("TaskType")
	})
}
func (*Queries) WithPreloadSourceDataspace() DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        opt.tx = opt.tx.Preload("SourceDataspace")
	})
}
func (*Queries) WithPreloadAuthID() DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        opt.tx = opt.tx.Preload("AuthID")
	})
}
// WARNING: 不建议用这个方式，需要确保期望查询的字段的值不为零值
func (*Queries) WithModel(v *testfile.ApprovalInfo) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		opt.tx = opt.tx.Where(v)
		// 这里只是记录
		queryParams := newQueryParams("model", v)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}

type Updates struct{}

// ApprovalInfoUpdates 这种方式避免其他开发者用错option
var ApprovalInfoUpdates Updates
func (*Updates) WithUpdateID(v string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		opt.setUpdateMap("id", v)
	})
}
func (*Updates) WithUpdateCreatedAt(v time.Time) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		opt.setUpdateMap("created_at", v)
	})
}
func (*Updates) WithUpdateUpdatedAt(v time.Time) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		opt.setUpdateMap("updated_at", v)
	})
}
func (*Updates) WithUpdateDeletedAt(v gorm.DeletedAt) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		opt.setUpdateMap("deleted_at", v)
	})
}
func (*Updates) WithUpdateCreatorID(v string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		opt.setUpdateMap("creator_id", v)
	})
}
func (*Updates) WithUpdateCreatorName(v string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		opt.setUpdateMap("creator_name", v)
	})
}
func (*Updates) WithUpdateTenantID(v string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		opt.setUpdateMap("tenant_id", v)
	})
}
func (*Updates) WithUpdateState(v ApprovalState) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		opt.setUpdateMap("state", v)
	})
}
func (*Updates) WithUpdateRepoID(v string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		opt.setUpdateMap("repo_id", v)
	})
}
func (*Updates) WithUpdateSourceRepoName(v string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		opt.setUpdateMap("source_repo_name", v)
	})
}
func (*Updates) WithUpdateRepoInfo(v RepoInfo) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		opt.setUpdateMap("repo_info", v)
	})
}
func (*Updates) WithUpdateTaskID(v string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		opt.setUpdateMap("task_id", v)
	})
}
func (*Updates) WithUpdateTaskName(v string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		opt.setUpdateMap("task_name", v)
	})
}
func (*Updates) WithUpdateTaskType(v TaskType) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		opt.setUpdateMap("task_type", v)
	})
}
func (*Updates) WithUpdateSourceDataspace(v string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		opt.setUpdateMap("source_dataspace", v)
	})
}
func (*Updates) WithUpdateAuthID(v string) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		opt.setUpdateMap("auth_id", v)
	})
}
