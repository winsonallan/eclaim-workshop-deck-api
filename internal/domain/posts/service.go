package posts

import "errors"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreatePost(userID uint, req CreatePostRequest) (*Post, error) {
	if req.Title == "" {
		return nil, errors.New("title is required")
	}
	if req.Content == "" {
		return nil, errors.New("content is required")
	}

	post := &Post{
		PostTitle:   req.Title,
		PostContent: req.Content,
		UserNo:  userID,
	}

	if err := s.repo.Create(post); err != nil {
		return nil, err
	}

	return s.repo.FindByID(post.PostNo)
}

func (s *Service) GetAllPosts() ([]Post, error) {
	return s.repo.FindAll()
}

func (s *Service) GetUserPosts(userID uint) ([]Post, error) {
	return s.repo.FindByUserID(userID)
}

func (s *Service) GetPostByID(id uint) (*Post, error) {
	return s.repo.FindByID(id)
}

func (s *Service) UpdatePost(postID, userID uint, req UpdatePostRequest) (*Post, error) {
	post, err := s.repo.FindByID(postID)
	if err != nil {
		return nil, errors.New("post not found")
	}

	if post.UserNo != userID {
		return nil, errors.New("unauthorized: you can only update your own posts")
	}

	if req.Title != "" {
		post.PostTitle = req.Title
	}
	if req.Content != "" {
		post.PostContent = req.Content
	}

	if err := s.repo.Update(post); err != nil {
		return nil, err
	}

	return post, nil
}

func (s *Service) DeletePost(postID, userID uint) error {
	post, err := s.repo.FindByID(postID)
	if err != nil {
		return errors.New("post not found")
	}

	if post.UserNo != userID {
		return errors.New("unauthorized: you can only delete your own posts")
	}

	return s.repo.Delete(postID)
}