// This file is for create new users and find users by username

package repository

import (
    "backend/internal/database"
    "backend/internal/models"
)


func FindUserByUsername(username string) (*models.User, error) {
    var user models.User
    result := database.DB.Where("username = ?", username).First(&user)
    return &user, result.Error
}


func CreateUser(user *models.User) error {
    return database.DB.Create(user).Error
}


func GetUserByID(id uint) (*models.User, error) {
    var user models.User
    result := database.DB.First(&user, id)
    return &user, result.Error
}