package repository

import (
    "backend/internal/database"
    "backend/internal/models"
)

func GetPostsByTopicID(topicID uint) ([]models.Post, error) {
    var posts []models.Post
    result := database.DB.
        Preload("Creator").
        Where("topic_id = ?", topicID).
        Order("created_at DESC").
        Find(&posts)
    return posts, result.Error
}

func GetPostByID(id uint) (*models.Post, error) {
    var post models.Post
    result := database.DB.
        Preload("Creator").
        Preload("Topic").
        First(&post, id)
    return &post, result.Error
}

func CreatePost(post *models.Post) error {
    return database.DB.Create(post).Error
}

func UpdatePost(post *models.Post) error {
    return database.DB.Save(post).Error
}

func DeletePost(id uint) error {
    return database.DB.Delete(&models.Post{}, id).Error
}

func CountPostsByTopicID(topicID uint) (int64, error) {
    var count int64
    result := database.DB.Model(&models.Post{}).Where("topic_id = ?", topicID).Count(&count)
    return count, result.Error
}