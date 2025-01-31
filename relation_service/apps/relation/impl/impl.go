// @Author: Hexiaoming 2023/2/7
package impl

import (
	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	"github.com/Go-To-Byte/DouSheng/api_rooter/client/rpc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	userRpc "github.com/Go-To-Byte/DouSheng/user_center/client/rpc"

	"github.com/Go-To-Byte/DouSheng/relation_service/apps/relation"
	"github.com/Go-To-Byte/DouSheng/message_service/apps/message"
	messageRpc "github.com/Go-To-Byte/DouSheng/message_service/client/rpc"
)

var (
	impl = &relationServiceImpl{}
)

type relationServiceImpl struct {
	db *gorm.DB
	l  logger.Logger

	relation.UnimplementedServiceServer

	// 依赖Token的客户端
	tokenService token.ServiceClient

	// 依赖User 的客户端
	userServer user.ServiceClient

	// 依赖Message 的客户端
	messageService message.ServiceClient
}

func (s *relationServiceImpl) Init() error {

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	s.db = db
	s.l = zap.L().Named(relation.AppName)

	// 获取API网关的客户端[GRPC调用]
	apiRooter, err := rpc.NewApiRooterClientFromCfg()
	if err != nil {
		s.l.Errorf("relations: getRelationPo 出现错误：%s", err.Error())
		return err
	}
	s.tokenService = apiRooter.TokenService()

	// 获取用户中心的客户端[GRPC调用]
	userCenter, err := userRpc.NewUserCenterClientFromCfg()
	if err != nil {
		return err
	}
	s.userServer = userCenter.UserService()

	// 获取消息服务的客户端[GRPC调用]
	msgService, err := messageRpc.NewMessageServiceClientFromCfg()
	if err != nil {
		s.l.Errorf("relations: getRelationPo 出现错误：%s", err.Error())
		return err
	}
	s.messageService = msgService.MessageService()

	return nil
}

func (s *relationServiceImpl) Name() string {
	return relation.AppName
}

func (s *relationServiceImpl) Registry(server *grpc.Server) {
	relation.RegisterServiceServer(server, impl)
}

func init() {
	ioc.GrpcDI(impl)
}
