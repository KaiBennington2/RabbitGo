package users

import (
	outPorts "github.com/KaiBennington2/RabbitGo/internal/domain/ports"
	"github.com/KaiBennington2/RabbitGo/internal/infra/adpters/persistence/users"
	"github.com/KaiBennington2/RabbitGo/sdk/ports"
	"sync"
)

type UserRepos struct {
	mysql outPorts.IUserRepository
	mongo outPorts.IUserRepository
}

var onceREPO sync.Once
var repoInstance *UserRepos
var LoadREPOSITORIES = func(connMap map[string]ports.IConnectionDB) *UserRepos {
	onceREPO.Do(func() {
		instWithImpl := &UserRepos{
			mysql: users.NewUserMysqlRepo(ports.GetSqlExec(connMap, "db_mysql_conn")),
			mongo: users.NewUserMongoRepo(ports.GetNoSqlExec(connMap, "db_mongo_conn")),
		}
		repoInstance = instWithImpl
	})
	return repoInstance
}
