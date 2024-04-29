package dao

import (
	"context"
	"database/sql"
	"time"

	"encoding/json"
	"fmt"
	"strings"

	"gorm.io/gorm"

	"github.com/pkg/errors"
)

const (
	// slowSqlThreshold 慢sql阈值
	slowSQLThreshold = time.Millisecond * 200
	// logSQLMaxLength 日志中打印的sql最大长度
	logSQLMaxLength = 1000
)

type GenDao interface {
    ApprovalInfoDao
	WithTransaction(ctx context.Context, opts ...*sql.TxOptions) context.Context
	// 重复调用实际只会执行一次
	FinishTransaction(ctx context.Context) error
	Transaction(ctx context.Context, f func(ctx context.Context) error, opts ...*sql.TxOptions) (err error)
}

type dao struct {
	db *gorm.DB
}

type DaoOpt interface {
	Apply(opt *daoOptions)
}

type daoOpt func(opt *daoOptions)

func (f daoOpt) Apply(opt *daoOptions) {
	f(opt)
}

type daoOptions struct {
	tx          *gorm.DB
	queryParams queryParamsList
	updateMap   map[string]interface{}
}

func (opts *daoOptions) setUpdateMap(key string, value interface{}) {
	opts.updateMap[key] = value
}

func (opts *daoOptions) wrapError(err error) error {
	builder := strings.Builder{}
	if len(opts.queryParams) > 0 {
		queryParams, mErr := json.Marshal(opts.queryParams)
		if mErr != nil {
			return errors.Wrap(err, mErr.Error())
		}

		builder.WriteString(fmt.Sprintf("queryParams: %s. ", queryParams))
	}

	if len(opts.updateMap) > 0 {
		updateMap, mErr := json.Marshal(opts.updateMap)
		if mErr != nil {
			return errors.Wrap(err, mErr.Error())
		}

		builder.WriteString(fmt.Sprintf("updateMap: %s. ", updateMap))
	}

	msg := builder.String()
	if len(msg) > 0 {
		return errors.Wrap(err, msg)
	} else {
		return errors.WithStack(err)
	}
}

type queryParams struct {
	SQL  string
	Args []interface{}
}

func newQueryParams(sql string, args ...interface{}) *queryParams {
	return &queryParams{
		SQL:  sql,
		Args: args,
	}
}

func (p *queryParams) where(tx *gorm.DB) *gorm.DB {
	return tx.Where(p.SQL, p.Args...)
}

type queryParamsList []*queryParams

func (d *dao) fromDaoOptions(ctx context.Context, opts ...DaoOpt) *daoOptions {
	opt := &daoOptions{
		tx:          d.tryGetTx(ctx),
		queryParams: queryParamsList{},
		updateMap:   make(map[string]interface{}),
	}
	for _, o := range opts {
		o.Apply(opt)
	}
	return opt
}
