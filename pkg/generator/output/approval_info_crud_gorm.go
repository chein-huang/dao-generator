
package dao

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
    "github.com/chein-huang/dao-generator/pkg/generator/test_file"
)

type ApprovalInfoDao interface {
	CreateApprovalInfo(ctx context.Context, approvalInfo *testfile.ApprovalInfo) error
	// CreateApprovalSample(ctx context.Context)
	GetApprovalInfo(ctx context.Context, opts ...DaoOpt) (*testfile.ApprovalInfoWithAuth, error)
	SearchApprovalInfo(ctx context.Context, page, pageSize int, isAll bool, opts ...DaoOpt) ([]*testfile.ApprovalInfoWithAuth, int, error)
	UpdateApprovalInfo(ctx context.Context, opts ...DaoOpt) (*testfile.ApprovalInfoWithAuth, error)
}

func (d *dao) CreateApprovalInfo(ctx context.Context, v *testfile.ApprovalInfo) error {
	tx := d.tryGetTx(ctx)

	err := tx.Create(&v).Error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (d *dao) GetApprovalInfo(ctx context.Context, opts ...DaoOpt) (*testfile.ApprovalInfoWithAuth, error) {
	opt := d.fromDaoOptions(ctx, opts...)

	var v testfile.ApprovalInfoWithAuth
	err := opt.tx.Take(&v).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = ErrApprovalInfoNotFound
		}
		return nil, opt.wrapError(err)
	}

	return &v, nil
}

func (d *dao) SearchApprovalInfo(ctx context.Context, page, pageSize int, isAll bool, opts ...DaoOpt) ([]*testfile.ApprovalInfoWithAuth, int, error) {
	opt := d.fromDaoOptions(ctx, opts...)

	count := int64(0)
	err := opt.tx.Model(&testfile.ApprovalInfo{}).Count(&count).Error
	if err != nil {
		return nil, 0, opt.wrapError(err)
	}

	if !isAll {
		opt.tx = opt.tx.Offset((page - 1) * pageSize).Limit(pageSize)
	}

	v := []*testfile.ApprovalInfoWithAuth{}
	err = opt.tx.Find(&v).Error
	if err != nil {
		return nil, 0, opt.wrapError(err)
	}

	return v, int(count), nil
}

func (d *dao) UpdateApprovalInfo(ctx context.Context, opts ...DaoOpt) (*testfile.ApprovalInfoWithAuth, error) {
	opt := d.fromDaoOptions(ctx, opts...)

	v, err := d.GetApprovalInfo(ctx, opts...)
	if err != nil {
		return nil, err
	}

	if len(opt.updateMap) == 0 {
		return v, nil
	}
	err = opt.tx.Model(v.ApprovalInfo).Updates(opt.updateMap).Error
	if err != nil {
		return nil, opt.wrapError(err)
	}

	return v, nil
}
