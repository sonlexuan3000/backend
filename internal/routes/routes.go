package routes

import (
    "github.com/gin-gonic/gin"
    
    "backend/internal/handlers"
    "backend/internal/middleware"
)

func SetupRoutes(r *gin.Engine) {
    // Public routes
    auth := r.Group("/auth")
    {
        auth.POST("/login", handlers.Login)
    }
    
    // API routes
    api := r.Group("/api")
    {
        // Topics routes - DÙNG CÙNG TÊN PARAMETER ":id"
        api.GET("/topics", handlers.GetAllTopics)
        api.GET("/topics/:id/posts", handlers.GetPostsByTopicID)  // ← ĐẶT TRƯỚC
        api.GET("/topics/:id", handlers.GetTopicByID)
        api.POST("/topics", middleware.AuthRequired(), handlers.CreateTopic)
        api.PUT("/topics/:id", middleware.AuthRequired(), handlers.UpdateTopic)
        api.DELETE("/topics/:id", middleware.AuthRequired(), handlers.DeleteTopic)
        
        // Posts routes - DÙNG CÙNG TÊN PARAMETER ":id"
        api.GET("/posts/:id/comments", handlers.GetCommentsByPostID)  // ← ĐẶT TRƯỚC
        api.GET("/posts/:id", handlers.GetPostByID)
        api.POST("/posts", middleware.AuthRequired(), handlers.CreatePost)
        api.PUT("/posts/:id", middleware.AuthRequired(), handlers.UpdatePost)
        api.DELETE("/posts/:id", middleware.AuthRequired(), handlers.DeletePost)
        
        // Comments routes
        api.GET("/comments/:id", handlers.GetCommentByID)
        api.POST("/comments", middleware.AuthRequired(), handlers.CreateComment)
        api.PUT("/comments/:id", middleware.AuthRequired(), handlers.UpdateComment)
        api.DELETE("/comments/:id", middleware.AuthRequired(), handlers.DeleteComment)
        
        // User routes
        api.GET("/me", middleware.AuthRequired(), handlers.GetMe)
    }
}