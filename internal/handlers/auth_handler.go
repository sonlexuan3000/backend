package handlers

import (
    "net/http"
    
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    
    "backend/internal/models"
    "backend/internal/repository"
    "backend/internal/utils"
)

type LoginRequest struct {
    Username string `json:"username" binding:"required,min=3,max=50"`
}

type AuthResponse struct {
    Token string       `json:"token"`
    User  models.User  `json:"user"`
}


func Login(c *gin.Context) {
    var req LoginRequest
    
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    user, err := repository.FindUserByUsername(req.Username)
    
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            newUser := models.User{
                Username: req.Username,
            }
            
            if err := repository.CreateUser(&newUser); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
                return
            }
            
            user = &newUser
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
            return
        }
    }
    
    token, err := utils.GenerateToken(user.ID, user.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }
    
    c.JSON(http.StatusOK, AuthResponse{
        Token: token,
        User:  *user,
    })
}

func GetMe(c *gin.Context) {
    userID, _ := c.Get("user_id")
    
    user, err := repository.GetUserByID(userID.(uint))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    
    c.JSON(http.StatusOK, user)
}