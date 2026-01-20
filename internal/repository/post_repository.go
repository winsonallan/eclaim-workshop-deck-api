package repository

import (
	"eclaim-workshop-deck-api/internal/models"

	"gorm.io/gorm"
)

type SamplePostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *SamplePostRepository {
	return &SamplePostRepository{db: db}
}

// CREATE
func (r *SamplePostRepository) Create(post *models.SamplePost) error {
	return r.db.Create(post).Error
}

// READ - Get all posts
func (r *SamplePostRepository) FindAll() ([]models.SamplePost, error) {
	var posts []models.SamplePost
	err := r.db.Preload("User").Order("created_at desc").Find(&posts).Error
	return posts, err
}

// READ - Get posts by user
func (r *SamplePostRepository) FindByUserID(userID uint) ([]models.SamplePost, error) {
	var posts []models.SamplePost
	err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&posts).Error
	return posts, err
}

// READ - Get single post by ID
func (r *SamplePostRepository) FindByID(id uint) (*models.SamplePost, error) {
	var post models.SamplePost
	err := r.db.Preload("User").First(&post, id).Error
	return &post, err
}

// UPDATE
func (r *SamplePostRepository) Update(post *models.SamplePost) error {
	return r.db.Save(post).Error
}

// DELETE
func (r *SamplePostRepository) Delete(id uint) error {
	return r.db.Delete(&models.SamplePost{}, id).Error
}