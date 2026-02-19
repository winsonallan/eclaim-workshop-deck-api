package posts

import (
	"eclaim-workshop-deck-api/internal/common/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	service *Service
	log     *zap.Logger
}

func NewHandler(service *Service, log *zap.Logger) *Handler {
	return &Handler{service: service, log: log}
}

func (h *Handler) CreatePost(c *gin.Context) {
	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("user_id")

	post, err := h.service.CreatePost(userID.(uint), req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Post created successfully", gin.H{"post": ToPostResponse(post)})
}

func (h *Handler) GetAllPosts(c *gin.Context) {
	posts, err := h.service.GetAllPosts()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch posts")
		return
	}

	response.Success(c, http.StatusOK, "Posts retrieved successfully", gin.H{
		"posts": ToPostResponses(posts),
		"count": len(posts),
	})
}

func (h *Handler) GetMyPosts(c *gin.Context) {
	userID, _ := c.Get("user_id")

	posts, err := h.service.GetUserPosts(userID.(uint))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch posts")
		return
	}

	response.Success(c, http.StatusOK, "Posts retrieved successfully", gin.H{
		"posts": ToPostResponses(posts),
		"count": len(posts),
	})
}

func (h *Handler) GetPost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	post, err := h.service.GetPostByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "Post not found")
		return
	}

	response.Success(c, http.StatusOK, "Post retrieved successfully", gin.H{"post": ToPostResponse(post)})
}

func (h *Handler) UpdatePost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	var req UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("user_id")

	post, err := h.service.UpdatePost(uint(id), userID.(uint), req)
	if err != nil {
		if err.Error() == "post not found" {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		if err.Error() == "unauthorized: you can only update your own posts" {
			response.Error(c, http.StatusForbidden, err.Error())
			return
		}
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Post updated successfully", gin.H{"post": ToPostResponse(post)})
}

func (h *Handler) DeletePost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	userID, _ := c.Get("user_id")

	err = h.service.DeletePost(uint(id), userID.(uint))
	if err != nil {
		if err.Error() == "post not found" {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		if err.Error() == "unauthorized: you can only delete your own posts" {
			response.Error(c, http.StatusForbidden, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Post deleted successfully", nil)
}
