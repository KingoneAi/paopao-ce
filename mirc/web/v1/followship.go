package v1

import (
	. "github.com/alimy/mir/v4"
	. "github.com/alimy/mir/v4/engine"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

func init() {
	Entry[Followship]()
}

// Followship 关注者模式 服务
type Followship struct {
	Chain `mir:"-"`
	Group `mir:"v1"`

	// FollowUser 关注用户
	FollowUser func(Post, Context, web.FollowUserReq) `mir:"/user/follow"`

	// UnfollowUser  取消关注用户
	UnfollowUser func(Post, Context, web.UnfollowUserReq) `mir:"/user/unfollow"`

	// ListFollows 获取用户的关注列表
	ListFollows func(Get, Context, web.ListFollowsReq) web.ListFollowsResp `mir:"/user/follows"`

	// ListFollowings 获取用户的追随者列表
	ListFollowings func(Get, Context, web.ListFollowingsReq) web.ListFollowingsResp `mir:"/user/followings"`
}
