package controllers

import (
	"first/initializers"
	"first/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePostHandler handles the creation of a new post
func CreatePostHandler(c *gin.Context) {
    var post models.Post
	initializers.ConnectToDB()
	defer initializers.CloseConnection() // Ensure to close the connection after the function finishes
    if err := c.BindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	
    // Insert the new post into the database
    _, err := initializers.DB.Exec("INSERT INTO posts (title, body) VALUES (?, ?)", post.Title, post.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"+err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully"})
}
// GetAllPostsHandler retrieves all posts from the database
func GetAllPostsHandler(c *gin.Context) {
    var posts []models.Post
	initializers.ConnectToDB()
	defer initializers.CloseConnection() // Ensure to close the connection after the function finishes
    if initializers.DB == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not established"})
        return
    }

    rows, err := initializers.DB.Query("SELECT * FROM posts")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "1 Failed to retrieve posts"+err.Error()})
        return
    }
    //defer rows.Close()
	
    for rows.Next() {
        var post models.Post
        err := rows.Scan(&post.ID, &post.Title, &post.Body)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Issue with retrieving values from the current row of the result set"+err.Error()})
            return
        }
        posts = append(posts, post)
    }

    if err := rows.Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "3 Failed to retrieve posts"+err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func GetPostByIdHandler(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.ConnectToDB()
	defer initializers.CloseConnection() // Ensure to close the connection after the function finishes
	if initializers.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not established"})
		return
	}
	err := initializers.DB.QueryRow("SELECT * FROM posts WHERE id = ?", id).Scan(&post.ID, &post.Title, &post.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve post"+err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}

func DeletePostByIdHandler(c *gin.Context) {
	id := c.Param("id")
	initializers.ConnectToDB()
	defer initializers.CloseConnection() // Ensure to close the connection after the function finishes
	if initializers.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not established"})
		return
	}
	_, err := initializers.DB.Exec("DELETE FROM posts WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"+err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}