package posts

import "time"

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostResponse struct {
	PostNo      uint      `json:"post_no"`
	PostTitle   string    `json:"post_title"`
	PostContent string    `json:"post_content"`
	UserNo      uint      `json:"user_no"`
	IsLocked    bool      `json:"is_locked"`
	Author      AuthorDTO `json:"author"`
	CreatedAt   time.Time `json:"created_date"`
	UpdatedAt   time.Time `json:"last_modified_date"`
}

type AuthorDTO struct {
	UserNo   uint   `json:"user_no"`
	UserName string `json:"user_name"`
}

func ToPostResponse(post *Post) PostResponse {
	return PostResponse{
		PostNo:      post.PostNo,
		PostTitle:   post.PostTitle,
		PostContent: post.PostContent,
		UserNo:      post.UserNo,
		IsLocked:    post.IsLocked,
		Author: AuthorDTO{
			UserNo:   post.User.UserNo,
			UserName: post.User.Name,
		},
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

func ToPostResponses(posts []Post) []PostResponse {
    responses := make([]PostResponse, len(posts))
    for i, post := range posts {
        responses[i] = ToPostResponse(&post)
    }
    return responses
}