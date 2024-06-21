package logic

import (
	"context"
	"micro/api/user/internal/svc"
	"micro/api/user/internal/types"
	"micro/rpc/user/userRPC"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteUserRequest) (resp *types.DeleteUserResponse, err error) {
	if _, err = l.svcCtx.UserService.DeleteUser(l.ctx, &userRPC.DeleteUserRequest{Id: req.ID}); err != nil {
		return &types.DeleteUserResponse{
			Code: 500,
			Msg:  err.Error(),
		}, nil
	} else {
		return &types.DeleteUserResponse{
			Code: 200,
			Msg:  "删除成功",
		}, nil
	}
}
