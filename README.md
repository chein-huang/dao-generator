<!--
 * @Author: huangcheng1 huangcheng1@sensetime.com
 * @Date: 2024-05-06 11:21:47
 * @LastEditors: huangcheng1 huangcheng1@sensetime.com
 * @LastEditTime: 2024-05-06 14:18:51
 * @FilePath: /dao-generator/README.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
# Table gen options
## StructName
Specify the name of the struct and indicate the generation of CRUD code. This is a required field.

Model: 
``` golang
// gen:"ApprovalInfo"
type ApprovalInfoWithAuth struct {
	...
}
```
Result:
``` golang
type ApprovalInfoDao interface {
	CreateApprovalInfo(ctx context.Context, approvalInfo *testfile.ApprovalInfo) error
	GetApprovalInfo(ctx context.Context, opts ...DaoOpt) (*testfile.ApprovalInfoWithAuth, error)
	SearchApprovalInfo(ctx context.Context, page, pageSize int, isAll bool, opts ...DaoOpt) ([]*testfile.ApprovalInfoWithAuth, int, error)
	UpdateApprovalInfo(ctx context.Context, opts ...DaoOpt) (*testfile.ApprovalInfoWithAuth, error)
}

```

## NameSnake
Specify the name in snake case, typically used for generating file names. By default, it is the snake case form of StructName.

Model:
``` golang
// gen:"ApprovalInfo,nameSnake:approvals"
type ApprovalInfoWithAuth struct {
	...
}
```
Result:
Will generate the following files
- approvals_crud_gorm.go
- dao_gorm.go
- errors_gorm.go
- transaction_gorm.go

## PackageAliceName
Used to specify the package alias where the model is located. By default, it is the package name of the file.

Model:
``` golang
// gen:"ApprovalInfo,packageAliceName:auth"
type ApprovalInfoWithAuth struct {
	...
}
```
Result:
``` golang
type ApprovalInfoDao interface {
	CreateApprovalInfo(ctx context.Context, approvalInfo *auth.ApprovalInfo) error
	GetApprovalInfo(ctx context.Context, opts ...DaoOpt) (*auth.ApprovalInfoWithAuth, error)
	SearchApprovalInfo(ctx context.Context, page, pageSize int, isAll bool, opts ...DaoOpt) ([]*auth.ApprovalInfoWithAuth, int, error)
	UpdateApprovalInfo(ctx context.Context, opts ...DaoOpt) (*auth.ApprovalInfoWithAuth, error)
}
```

## IsPreload
Used to indicate that the struct is in Preload form. We will parse the embedded fields, construct CRUD options for the corresponding fields, and then treat named fields as preload fields.

Model:
``` golang
// gen:"ApprovalInfo,flags:isPreload"
type ApprovalInfoWithAuth struct {
	*ApprovalInfo
	AuthInfo *AuthInfo `gorm:"foreignKey:AuthID"`
}
```

Result:
``` golang
func (*approvalInfoQueries) WithPreloadAuthInfo() DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		opt.tx = opt.tx.Preload("AuthInfo")
	})
}
```

# Field gen options
## Order
Used to indicate whether sorting for this field is required.

Model:
``` golang
// gen:"ApprovalInfo,flags:isPreload"
type ApprovalInfoWithAuth struct {
	*ApprovalInfo
	// gen:"orderBy:created_at DESC"
	AuthInfo *AuthInfo `gorm:"foreignKey:AuthID"`
}

type ApprovalInfo struct {
	ID string `gorm:"primaryKey;type:varchar(36)"`
	// gen:"flags:order"
	CreatedAt time.Time
    ...
}
```

Result:
``` golang
func (*approvalInfoQueries) WithOrderByCreatedAt(desc bool) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		orderStr := "created_at"
		if desc {
			orderStr += " DESC"
		}
		opt.tx = opt.tx.Order(orderStr)
	})
}
```

## Range
Used to indicate whether range retrieval for this field is supported.

Model:
``` golang
// gen:"ApprovalInfo,flags:isPreload"
type ApprovalInfoWithAuth struct {
	*ApprovalInfo
	// gen:"orderBy:created_at DESC"
	AuthInfo *AuthInfo `gorm:"foreignKey:AuthID"`
}

type ApprovalInfo struct {
	ID string `gorm:"primaryKey;type:varchar(36)"`
	// gen:"flags:order;range"
	CreatedAt time.Time
    ...
}
```

Result:
``` golang
func (*approvalInfoQueries) WithCreatedAtLt(v time.Time) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("created_at < ?", v)
        opt.tx = queryParams.where(opt.tx)
        opt.queryParams = append(opt.queryParams, queryParams)
    })
}

func (*approvalInfoQueries) WithCreatedAtLte(v time.Time) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("created_at <= ?", v)
        opt.tx = queryParams.where(opt.tx)
        opt.queryParams = append(opt.queryParams, queryParams)
    })
}

func (*approvalInfoQueries) WithCreatedAtGt(v time.Time) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("created_at > ?", v)
        opt.tx = queryParams.where(opt.tx)
        opt.queryParams = append(opt.queryParams, queryParams)
    })
}

func (*approvalInfoQueries) WithCreatedAtGte(v time.Time) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("created_at >= ?", v)
        opt.tx = queryParams.where(opt.tx)
        opt.queryParams = append(opt.queryParams, queryParams)
    })
}
```

## In
Used to indicate whether list retrieval for this field is supported.

Model:
``` golang
// gen:"ApprovalInfo,flags:isPreload"
type ApprovalInfoWithAuth struct {
	*ApprovalInfo
	// gen:"orderBy:created_at DESC"
	AuthInfo *AuthInfo `gorm:"foreignKey:AuthID"`
}

type ApprovalInfo struct {
	ID string `gorm:"primaryKey;type:varchar(36)"`
	// gen:"flags:in"
	State ApprovalState `gorm:"index"`
    ...
}
```

Result:
``` golang
func (*approvalInfoQueries) WithInState(vs []testfile.ApprovalState) DaoOpt {
	return daoOpt(func(opt *daoOptions) {
        queryParams := newQueryParams("state in (?)", vs)
		opt.tx = queryParams.where(opt.tx)
		opt.queryParams = append(opt.queryParams, queryParams)
	})
}
```

## OrderBy
Used to specify the sorting order of the preload results.

Model:
``` golang
// gen:"ApprovalInfo,flags:isPreload"
type ApprovalInfoWithAuth struct {
	*ApprovalInfo
	// gen:"orderBy:created_at DESC"
	AuthInfo *AuthInfo `gorm:"foreignKey:AuthID"`
}
```

Result:
``` golang
func (*approvalInfoQueries) WithPreloadAuthInfo() DaoOpt {
	return daoOpt(func(opt *daoOptions) {
		opt.tx = opt.tx.Preload("AuthInfo", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC")
		})
	})
}
```