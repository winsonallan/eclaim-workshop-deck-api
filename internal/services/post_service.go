package services

import (
	"eclaim-workshop-deck-api/internal/models"
	"eclaim-workshop-deck-api/internal/repository"
	"errors"
)

type PostService struct {
	postRepo *repository.SamplePostRepository
}

func NewPostService(postRepo *repository.SamplePostRepository) *PostService {
	return &PostService{postRepo: postRepo}
}

// CREATE
func (s *PostService) CreatePost(userID uint, title, content string) (*models.SamplePost, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}
	if content == "" {
		return nil, errors.New("content is required")
	}

	post := &models.SamplePost{
		Title:   title,
		Content: content,
		UserID:  userID,
	}

	if err := s.postRepo.Create(post); err != nil {
		return nil, err
	}

	return post, nil
}

// READ - Get all posts
func (s *PostService) GetAllPosts() ([]models.SamplePost, error) {
	return s.postRepo.FindAll()
}

// READ - Get user's posts
func (s *PostService) GetUserPosts(userID uint) ([]models.SamplePost, error) {
	return s.postRepo.FindByUserID(userID)
}

// READ - Get single post
func (s *PostService) GetPostByID(id uint) (*models.SamplePost, error) {
	return s.postRepo.FindByID(id)
}

// UPDATE
func (s *PostService) UpdatePost(postID, userID uint, title, content string) (*models.SamplePost, error) {
	// Get existing post
	post, err := s.postRepo.FindByID(postID)
	if err != nil {
		return nil, errors.New("post not found")
	}

	// Check if user owns the post
	if post.UserID != userID {
		return nil, errors.New("unauthorized: you can only update your own posts")
	}

	// Update fields
	if title != "" {
		post.Title = title
	}
	if content != "" {
		post.Content = content
	}

	if err := s.postRepo.Update(post); err != nil {
		return nil, err
	}

	return post, nil
}

// DELETE
func (s *PostService) DeletePost(postID, userID uint) error {
	// Get existing post
	post, err := s.postRepo.FindByID(postID)
	if err != nil {
		return errors.New("post not found")
	}

	// Check if user owns the post
	if post.UserID != userID {
		return errors.New("unauthorized: you can only delete your own posts")
	}

	return s.postRepo.Delete(postID)
}