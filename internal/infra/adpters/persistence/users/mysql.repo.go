package users

import (
	"context"
	"database/sql"
	"errors"
	"github.com/KaiBennington2/RabbitGo/internal/domain/entity"
	outPorts "github.com/KaiBennington2/RabbitGo/internal/domain/ports"
	"github.com/KaiBennington2/RabbitGo/sdk/ports"
	"strings"
)

var _ outPorts.IUserRepository = (*Mysql)(nil)

type Mysql struct {
	db ports.IExecutorSqlDB
}

func NewUserMysqlRepo(db ports.IExecutorSqlDB) *Mysql {
	return &Mysql{db: db}
}

func (rpo *Mysql) Save(ctx context.Context, e *entity.User) (*string, error) {
	_, err := rpo.db.ExecContext(ctx, SqlUserCreate, &e.Code, &e.Name)
	if err != nil {
		return nil, err
	}
	return &e.Code, nil
}

func (rpo *Mysql) Delete(ctx context.Context, code *string) (bool, error) {
	rs, err := rpo.db.ExecContext(ctx, SqlUserDeleteById, code)
	if err != nil {
		return false, err
	}

	rsp, err := rs.RowsAffected()
	if err != nil {
		return false, err
	}
	return rsp > 0, nil
}

func (rpo *Mysql) LogicalDelete(ctx context.Context, code *string) (bool, error) {
	rs, err := rpo.db.ExecContext(ctx, SqlUserDeletedLogicalById, code)
	if err != nil {
		return false, err
	}

	rsp, err := rs.RowsAffected()
	if err != nil {
		return false, err
	}
	return rsp > 0, nil
}

func (rpo *Mysql) ExistsById(ctx context.Context, code *string, knob ...bool) (bool, error) {
	var exists bool
	query := SqlUserExistsById
	if knob != nil {
		if knob[0] {
			query = strings.ReplaceAll(SqlUserExistsById, "date_deleted IS NULL AND", "")
		}
	}
	row := rpo.db.QueryRowContext(ctx, query, code)
	if err := row.Scan(&exists); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return exists, err
		}
		return exists, err
	}
	return exists, nil
}
