// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: comments.sql

package pgc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createComment = `-- name: CreateComment :one
INSERT INTO p_comment (post_id, user_id, ip, ip_loc, created_on)
VALUES ($1, $2, $3, $4, $5)
RETURNING id
`

type CreateCommentParams struct {
	PostID    int64
	UserID    int64
	Ip        string
	IpLoc     string
	CreatedOn int64
}

func (q *Queries) CreateComment(ctx context.Context, arg *CreateCommentParams) (int64, error) {
	row := q.db.QueryRow(ctx, createComment,
		arg.PostID,
		arg.UserID,
		arg.Ip,
		arg.IpLoc,
		arg.CreatedOn,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createCommentContent = `-- name: CreateCommentContent :one
INSERT INTO p_comment_content (comment_id, user_id, content, type, sort, created_on)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id
`

type CreateCommentContentParams struct {
	CommentID int64
	UserID    int64
	Content   string
	Type      int16
	Sort      int64
	CreatedOn int64
}

func (q *Queries) CreateCommentContent(ctx context.Context, arg *CreateCommentContentParams) (int64, error) {
	row := q.db.QueryRow(ctx, createCommentContent,
		arg.CommentID,
		arg.UserID,
		arg.Content,
		arg.Type,
		arg.Sort,
		arg.CreatedOn,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createCommentReply = `-- name: CreateCommentReply :one
INSERT INTO p_comment_reply (comment_id, user_id, content, at_user_id, ip, ip_loc, created_on)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id
`

type CreateCommentReplyParams struct {
	CommentID int64
	UserID    int64
	Content   string
	AtUserID  int64
	Ip        string
	IpLoc     string
	CreatedOn int64
}

func (q *Queries) CreateCommentReply(ctx context.Context, arg *CreateCommentReplyParams) (int64, error) {
	row := q.db.QueryRow(ctx, createCommentReply,
		arg.CommentID,
		arg.UserID,
		arg.Content,
		arg.AtUserID,
		arg.Ip,
		arg.IpLoc,
		arg.CreatedOn,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createThumbsUpDownComment = `-- name: CreateThumbsUpDownComment :one
INSERT INTO p_tweet_comment_thumbs (user_id, tweet_id, comment_id, 
    reply_id, is_thumbs_up, is_thumbs_down, comment_type, created_on)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id
`

type CreateThumbsUpDownCommentParams struct {
	UserID       int64
	TweetID      int64
	CommentID    int64
	ReplyID      pgtype.Int8
	IsThumbsUp   int16
	IsThumbsDown int16
	CommentType  int16
	CreatedOn    int64
}

func (q *Queries) CreateThumbsUpDownComment(ctx context.Context, arg *CreateThumbsUpDownCommentParams) (int64, error) {
	row := q.db.QueryRow(ctx, createThumbsUpDownComment,
		arg.UserID,
		arg.TweetID,
		arg.CommentID,
		arg.ReplyID,
		arg.IsThumbsUp,
		arg.IsThumbsDown,
		arg.CommentType,
		arg.CreatedOn,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const decrCommentReplyCount = `-- name: DecrCommentReplyCount :exec
UPDATE p_comment
SET reply_count=reply_count-1,
    modified_on=$1
WHERE id=$2 AND is_del=0
`

type DecrCommentReplyCountParams struct {
	ModifiedOn int64
	ID         int64
}

func (q *Queries) DecrCommentReplyCount(ctx context.Context, arg *DecrCommentReplyCountParams) error {
	_, err := q.db.Exec(ctx, decrCommentReplyCount, arg.ModifiedOn, arg.ID)
	return err
}

const deleteComment = `-- name: DeleteComment :exec

UPDATE p_comment SET deleted_on=$1, is_del=1 WHERE id=$2 AND is_del=0
`

type DeleteCommentParams struct {
	DeletedOn int64
	ID        int64
}

// ------------------------------------------------------------------------------
// comment_manage sql dml
// ------------------------------------------------------------------------------
func (q *Queries) DeleteComment(ctx context.Context, arg *DeleteCommentParams) error {
	_, err := q.db.Exec(ctx, deleteComment, arg.DeletedOn, arg.ID)
	return err
}

const deleteCommentReply = `-- name: DeleteCommentReply :exec
UPDATE p_comment_reply SET deleted_on=$1, is_del=1 WHERE id=$2 AND is_del=0
`

type DeleteCommentReplyParams struct {
	DeletedOn int64
	ID        int64
}

func (q *Queries) DeleteCommentReply(ctx context.Context, arg *DeleteCommentReplyParams) error {
	_, err := q.db.Exec(ctx, deleteCommentReply, arg.DeletedOn, arg.ID)
	return err
}

const deleteCommentThumbs = `-- name: DeleteCommentThumbs :exec
UPDATE p_tweet_comment_thumbs
SET deleted_on=$1, is_del=1
WHERE user_id=$2 AND tweet_id=$3 AND comment_id=$4 AND is_del=0
`

type DeleteCommentThumbsParams struct {
	DeletedOn int64
	UserID    int64
	TweetID   int64
	CommentID int64
}

func (q *Queries) DeleteCommentThumbs(ctx context.Context, arg *DeleteCommentThumbsParams) error {
	_, err := q.db.Exec(ctx, deleteCommentThumbs,
		arg.DeletedOn,
		arg.UserID,
		arg.TweetID,
		arg.CommentID,
	)
	return err
}

const deleteReplyThumbs = `-- name: DeleteReplyThumbs :exec
UPDATE p_tweet_comment_thumbs
SET deleted_on=$1, is_del=1
WHERE user_id=$2 AND comment_id=$3 AND reply_id=$4 AND is_del=0
`

type DeleteReplyThumbsParams struct {
	DeletedOn int64
	UserID    int64
	CommentID int64
	ReplyID   pgtype.Int8
}

func (q *Queries) DeleteReplyThumbs(ctx context.Context, arg *DeleteReplyThumbsParams) error {
	_, err := q.db.Exec(ctx, deleteReplyThumbs,
		arg.DeletedOn,
		arg.UserID,
		arg.CommentID,
		arg.ReplyID,
	)
	return err
}

const getCommentById = `-- name: GetCommentById :one
SELECT id, post_id, user_id, ip, ip_loc, created_on, modified_on, deleted_on, is_del, thumbs_up_count, thumbs_down_count, is_essence, reply_count FROM p_comment WHERE id=$1 AND is_del=0
`

func (q *Queries) GetCommentById(ctx context.Context, id int64) (*PComment, error) {
	row := q.db.QueryRow(ctx, getCommentById, id)
	var i PComment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.UserID,
		&i.Ip,
		&i.IpLoc,
		&i.CreatedOn,
		&i.ModifiedOn,
		&i.DeletedOn,
		&i.IsDel,
		&i.ThumbsUpCount,
		&i.ThumbsDownCount,
		&i.IsEssence,
		&i.ReplyCount,
	)
	return &i, err
}

const getCommentContentsByIds = `-- name: GetCommentContentsByIds :many
SELECT id, comment_id, user_id, content, type, sort, created_on, modified_on, deleted_on, is_del FROM p_comment_content WHERE comment_id = ANY($1::BIGINT[])
`

func (q *Queries) GetCommentContentsByIds(ctx context.Context, ids []int64) ([]*PCommentContent, error) {
	rows, err := q.db.Query(ctx, getCommentContentsByIds, ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*PCommentContent
	for rows.Next() {
		var i PCommentContent
		if err := rows.Scan(
			&i.ID,
			&i.CommentID,
			&i.UserID,
			&i.Content,
			&i.Type,
			&i.Sort,
			&i.CreatedOn,
			&i.ModifiedOn,
			&i.DeletedOn,
			&i.IsDel,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCommentCount = `-- name: GetCommentCount :one
SELECT count(*) FROM p_comment WHERE post_id=$1 AND is_del=0
`

func (q *Queries) GetCommentCount(ctx context.Context, postID int64) (int64, error) {
	row := q.db.QueryRow(ctx, getCommentCount, postID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getCommentRepliesByIds = `-- name: GetCommentRepliesByIds :many
SELECT id, comment_id, user_id, at_user_id, content, ip, ip_loc, created_on, modified_on, deleted_on, is_del, thumbs_up_count, thumbs_down_count 
FROM p_comment_reply 
WHERE comment_id = ANY($1::BIGINT[])
ORDER BY id ASC
`

func (q *Queries) GetCommentRepliesByIds(ctx context.Context, ids []int64) ([]*PCommentReply, error) {
	rows, err := q.db.Query(ctx, getCommentRepliesByIds, ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*PCommentReply
	for rows.Next() {
		var i PCommentReply
		if err := rows.Scan(
			&i.ID,
			&i.CommentID,
			&i.UserID,
			&i.AtUserID,
			&i.Content,
			&i.Ip,
			&i.IpLoc,
			&i.CreatedOn,
			&i.ModifiedOn,
			&i.DeletedOn,
			&i.IsDel,
			&i.ThumbsUpCount,
			&i.ThumbsDownCount,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCommentReplyById = `-- name: GetCommentReplyById :one
SELECT id, comment_id, user_id, at_user_id, content, ip, ip_loc, created_on, modified_on, deleted_on, is_del, thumbs_up_count, thumbs_down_count FROM p_comment_reply WHERE id=$1 AND is_del=0
`

func (q *Queries) GetCommentReplyById(ctx context.Context, id int64) (*PCommentReply, error) {
	row := q.db.QueryRow(ctx, getCommentReplyById, id)
	var i PCommentReply
	err := row.Scan(
		&i.ID,
		&i.CommentID,
		&i.UserID,
		&i.AtUserID,
		&i.Content,
		&i.Ip,
		&i.IpLoc,
		&i.CreatedOn,
		&i.ModifiedOn,
		&i.DeletedOn,
		&i.IsDel,
		&i.ThumbsUpCount,
		&i.ThumbsDownCount,
	)
	return &i, err
}

const getCommentReplyThumb = `-- name: GetCommentReplyThumb :one
SELECT id, user_id, tweet_id, comment_id, reply_id, comment_type, is_thumbs_up, is_thumbs_down, created_on, modified_on, deleted_on, is_del
FROM p_tweet_comment_thumbs
WHERE user_id=$1 AND tweet_id=$2 AND comment_id=$3 AND reply_id=$4 AND comment_type=1 AND is_del=0
`

type GetCommentReplyThumbParams struct {
	UserID    int64
	TweetID   int64
	CommentID int64
	ReplyID   pgtype.Int8
}

func (q *Queries) GetCommentReplyThumb(ctx context.Context, arg *GetCommentReplyThumbParams) (*PTweetCommentThumb, error) {
	row := q.db.QueryRow(ctx, getCommentReplyThumb,
		arg.UserID,
		arg.TweetID,
		arg.CommentID,
		arg.ReplyID,
	)
	var i PTweetCommentThumb
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TweetID,
		&i.CommentID,
		&i.ReplyID,
		&i.CommentType,
		&i.IsThumbsUp,
		&i.IsThumbsDown,
		&i.CreatedOn,
		&i.ModifiedOn,
		&i.DeletedOn,
		&i.IsDel,
	)
	return &i, err
}

const getCommentThumbs = `-- name: GetCommentThumbs :many
SELECT user_id, 
    tweet_id, 
    comment_id, 
    reply_id, 
    comment_type, 
    is_thumbs_up, 
    is_thumbs_down
FROM p_tweet_comment_thumbs
WHERE user_id=$1 AND tweet_id=$2
`

type GetCommentThumbsParams struct {
	UserID  int64
	TweetID int64
}

type GetCommentThumbsRow struct {
	UserID       int64
	TweetID      int64
	CommentID    int64
	ReplyID      pgtype.Int8
	CommentType  int16
	IsThumbsUp   int16
	IsThumbsDown int16
}

func (q *Queries) GetCommentThumbs(ctx context.Context, arg *GetCommentThumbsParams) ([]*GetCommentThumbsRow, error) {
	rows, err := q.db.Query(ctx, getCommentThumbs, arg.UserID, arg.TweetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetCommentThumbsRow
	for rows.Next() {
		var i GetCommentThumbsRow
		if err := rows.Scan(
			&i.UserID,
			&i.TweetID,
			&i.CommentID,
			&i.ReplyID,
			&i.CommentType,
			&i.IsThumbsUp,
			&i.IsThumbsDown,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDefaultComments = `-- name: GetDefaultComments :many
SELECT id, post_id, user_id, ip, ip_loc, created_on, modified_on, deleted_on, is_del, thumbs_up_count, thumbs_down_count, is_essence, reply_count 
FROM p_comment 
WHERE post_id=$1 AND is_del=0 
ORDER BY is_essence DESC, id ASC 
LIMIT $2 OFFSET $3
`

type GetDefaultCommentsParams struct {
	PostID int64
	Limit  int32
	Offset int32
}

func (q *Queries) GetDefaultComments(ctx context.Context, arg *GetDefaultCommentsParams) ([]*PComment, error) {
	rows, err := q.db.Query(ctx, getDefaultComments, arg.PostID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*PComment
	for rows.Next() {
		var i PComment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.Ip,
			&i.IpLoc,
			&i.CreatedOn,
			&i.ModifiedOn,
			&i.DeletedOn,
			&i.IsDel,
			&i.ThumbsUpCount,
			&i.ThumbsDownCount,
			&i.IsEssence,
			&i.ReplyCount,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getHotsComments = `-- name: GetHotsComments :many
SELECT c.id, c.post_id, c.user_id, c.ip, c.ip_loc, c.created_on, c.modified_on, c.deleted_on, c.is_del, c.thumbs_up_count, c.thumbs_down_count, c.is_essence, c.reply_count 
FROM p_comment c
LEFT JOIN p_comment_metric m
ON c.id=m.comment_id
WHERE c.post_id=$1 AND c.is_del=0 AND m.is_del=0
ORDER BY is_essence DESC, m.rank_score DESC, c.id DESC
LIMIT $2 OFFSET $3
`

type GetHotsCommentsParams struct {
	PostID int64
	Limit  int32
	Offset int32
}

func (q *Queries) GetHotsComments(ctx context.Context, arg *GetHotsCommentsParams) ([]*PComment, error) {
	rows, err := q.db.Query(ctx, getHotsComments, arg.PostID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*PComment
	for rows.Next() {
		var i PComment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.Ip,
			&i.IpLoc,
			&i.CreatedOn,
			&i.ModifiedOn,
			&i.DeletedOn,
			&i.IsDel,
			&i.ThumbsUpCount,
			&i.ThumbsDownCount,
			&i.IsEssence,
			&i.ReplyCount,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNewestComments = `-- name: GetNewestComments :many

SELECT id, post_id, user_id, ip, ip_loc, created_on, modified_on, deleted_on, is_del, thumbs_up_count, thumbs_down_count, is_essence, reply_count 
FROM p_comment 
WHERE post_id=$1 AND is_del=0 
ORDER BY is_essence DESC, id DESC 
LIMIT $2 OFFSET $3
`

type GetNewestCommentsParams struct {
	PostID int64
	Limit  int32
	Offset int32
}

// ------------------------------------------------------------------------------
// comment sql dml
// ------------------------------------------------------------------------------
func (q *Queries) GetNewestComments(ctx context.Context, arg *GetNewestCommentsParams) ([]*PComment, error) {
	rows, err := q.db.Query(ctx, getNewestComments, arg.PostID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*PComment
	for rows.Next() {
		var i PComment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.Ip,
			&i.IpLoc,
			&i.CreatedOn,
			&i.ModifiedOn,
			&i.DeletedOn,
			&i.IsDel,
			&i.ThumbsUpCount,
			&i.ThumbsDownCount,
			&i.IsEssence,
			&i.ReplyCount,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTweetCommentThumb = `-- name: GetTweetCommentThumb :one
SELECT id, user_id, tweet_id, comment_id, reply_id, comment_type, is_thumbs_up, is_thumbs_down, created_on, modified_on, deleted_on, is_del
FROM p_tweet_comment_thumbs
WHERE user_id=$1 AND tweet_id=$2 AND comment_id=$3 AND comment_type=0 AND is_del=0
`

type GetTweetCommentThumbParams struct {
	UserID    int64
	TweetID   int64
	CommentID int64
}

func (q *Queries) GetTweetCommentThumb(ctx context.Context, arg *GetTweetCommentThumbParams) (*PTweetCommentThumb, error) {
	row := q.db.QueryRow(ctx, getTweetCommentThumb, arg.UserID, arg.TweetID, arg.CommentID)
	var i PTweetCommentThumb
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TweetID,
		&i.CommentID,
		&i.ReplyID,
		&i.CommentType,
		&i.IsThumbsUp,
		&i.IsThumbsDown,
		&i.CreatedOn,
		&i.ModifiedOn,
		&i.DeletedOn,
		&i.IsDel,
	)
	return &i, err
}

const highlightComment = `-- name: HighlightComment :one
UPDATE p_comment
SET is_essence=1-is_essence,
    modified_on=$1
WHERE id=$2 AND user_id=$3 AND is_del=0
RETURNING is_essence
`

type HighlightCommentParams struct {
	ModifiedOn int64
	ID         int64
	UserID     int64
}

func (q *Queries) HighlightComment(ctx context.Context, arg *HighlightCommentParams) (int16, error) {
	row := q.db.QueryRow(ctx, highlightComment, arg.ModifiedOn, arg.ID, arg.UserID)
	var is_essence int16
	err := row.Scan(&is_essence)
	return is_essence, err
}

const incrCommentReplyCount = `-- name: IncrCommentReplyCount :exec
UPDATE p_comment
SET reply_count=reply_count+1,
    modified_on=$1
WHERE id=$2 AND is_del=0
`

type IncrCommentReplyCountParams struct {
	ModifiedOn int64
	ID         int64
}

func (q *Queries) IncrCommentReplyCount(ctx context.Context, arg *IncrCommentReplyCountParams) error {
	_, err := q.db.Exec(ctx, incrCommentReplyCount, arg.ModifiedOn, arg.ID)
	return err
}

const updateCommentThumbsCount = `-- name: UpdateCommentThumbsCount :exec
UPDATE p_comment
SET thumbs_up_count=$1, Thumbs_down_count=$2, modified_on=$3
WHERE id=$4 AND is_del=0
`

type UpdateCommentThumbsCountParams struct {
	ThumbsUpCount   int32
	ThumbsDownCount int32
	ModifiedOn      int64
	ID              int64
}

func (q *Queries) UpdateCommentThumbsCount(ctx context.Context, arg *UpdateCommentThumbsCountParams) error {
	_, err := q.db.Exec(ctx, updateCommentThumbsCount,
		arg.ThumbsUpCount,
		arg.ThumbsDownCount,
		arg.ModifiedOn,
		arg.ID,
	)
	return err
}

const updateReplyThumbsCount = `-- name: UpdateReplyThumbsCount :exec
UPDATE p_comment_reply
SET thumbs_up_count=$1, thumbs_down_count=$2, modified_on=$3
WHERE id=$4 AND is_del=0
`

type UpdateReplyThumbsCountParams struct {
	ThumbsUpCount   int32
	ThumbsDownCount int32
	ModifiedOn      int64
	ID              int64
}

func (q *Queries) UpdateReplyThumbsCount(ctx context.Context, arg *UpdateReplyThumbsCountParams) error {
	_, err := q.db.Exec(ctx, updateReplyThumbsCount,
		arg.ThumbsUpCount,
		arg.ThumbsDownCount,
		arg.ModifiedOn,
		arg.ID,
	)
	return err
}

const updateThumbsUpDownComment = `-- name: UpdateThumbsUpDownComment :exec
UPDATE p_tweet_comment_thumbs
SET is_thumbs_up=$1,
    is_thumbs_down=$2,
    modified_on=$3
WHERE id=$4 AND is_del=0
`

type UpdateThumbsUpDownCommentParams struct {
	IsThumbsUp   int16
	IsThumbsDown int16
	ModifiedOn   int64
	ID           int64
}

func (q *Queries) UpdateThumbsUpDownComment(ctx context.Context, arg *UpdateThumbsUpDownCommentParams) error {
	_, err := q.db.Exec(ctx, updateThumbsUpDownComment,
		arg.IsThumbsUp,
		arg.IsThumbsDown,
		arg.ModifiedOn,
		arg.ID,
	)
	return err
}
