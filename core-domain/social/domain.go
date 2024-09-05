package social

import (
	"time"

	proto "github.com/corny90/social-app-pkg/grpc-proto"
	"github.com/gocql/gocql"
)

type Post struct {
	PostID        gocql.UUID     `json:"postId"`
	UserTargetID  gocql.UUID     `json:"userTargetId" validate:"required"`
	UserCreatorID gocql.UUID     `json:"userCreatorId"`
	CreatedAt     time.Time      `json:"createdAt"`
	Private       *bool          `json:"private" validate:"required"`
	PrivateType   string         `json:"privateType"` // "only_me", "only_to_target", "to_a_public"
	ContentText   string         `json:"contentText"`
	ContentStatus string         `json:"contentStatus"`
	LinkUrl       string         `json:"linkUrl"` // optional promoted post link
	Keywords      []string       `json:"keywords"`
	Media         []*proto.Media `json:"media"`
	Counters      PostCounter    `json:"counters"`
	IsLiked       bool           `json:"isLiked"`
	IsVirtual     bool           `json:"isVirtual"`
	Username      string         `json:"username"`
}

type PostCounter struct {
	PostID          gocql.UUID `json:"postId"`
	CounterLikes    int64      `json:"counterLikes"`
	CounterComments int64      `json:"counterComments"`
	CounterViews    int64      `json:"counterViews"`
}

type PostView struct {
	PostID    gocql.UUID `json:"postId"`
	UserID    gocql.UUID `json:"userId"`
	CreatedAt time.Time  `json:"createdAt"`
}

type PostViewJSON struct {
	PostID    string `json:"postId"`
	UserID    string `json:"userId"`
	CreatedAt string `json:"createdAt"`
}

type PostLike struct {
	PostID    gocql.UUID `json:"postId"`
	UserID    gocql.UUID `json:"userId"`
	CreatedAt time.Time  `json:"createdAt"`
}

type PostComment struct {
	CommentID gocql.UUID `json:"commentId"`
	PostID    gocql.UUID `json:"postId"`
	AuthorID  gocql.UUID `json:"authorId"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Content   string     `json:"content"`
}

type CommentReply struct {
	ParentCommentID gocql.UUID `json:"parent_commentId"`
	ReplyID         gocql.UUID `json:"replyId"`
	PostID          gocql.UUID `json:"postId"`
	AuthorID        gocql.UUID `json:"authorId"`
	CreatedAt       time.Time  `json:"createdAt"`
	Content         string     `json:"content"`
}

type CommentLike struct {
	CommentID gocql.UUID `json:"commentId"`
	UserID    gocql.UUID `json:"userId"`
	LikedAt   time.Time  `json:"likedAt"`
}

type MediaFile struct {
	PhotoBytes       []byte
	MimeType         string
	OriginalFilename string
}

type Follows struct {
	FollowingID gocql.UUID `json:"followingId"`
	FollowerID  gocql.UUID `json:"followerId"`
	CreatedAt   time.Time  `json:"createdAt"`
	Status      string     `json:"status"`
}

type FollowsCounter struct {
	UserID            gocql.UUID `json:"userId"`
	CounterFollowings int64      `json:"counterFollowings"`
	CounterFollowers  int64      `json:"counterFollowers"`
}
