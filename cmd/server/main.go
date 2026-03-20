package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fvaiiii/archiver/internal/archive"
	"github.com/fvaiiii/archiver/internal/lz77"
	"github.com/gin-gonic/gin"
)

const (
	maxFileSize   = 50 << 20
	defaultWindow = 32768
)

func main() {
	r := gin.Default()

	// r.Use(corsMiddleware())

	api := r.Group("/api")
	{
		api.POST("/compress", compressHandler)
		api.POST("/decompress", decompressHandler)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	log.Println("Gin server is running on port :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func compressHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "file not received" + err.Error()})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "file opening error"})
		return
	}

	defer src.Close()

	data, err := io.ReadAll(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "file opening error"})
		return
	}

	tokens := lz77.Compress(data, defaultWindow)

	tmpPath := filepath.Join(os.TempDir(), file.Filename+".arc")
	if err := archive.WriteArchive(tmpPath, tokens); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "compression error: " + err.Error()})
		return
	}
	defer os.Remove(tmpPath)

	c.FileAttachment(tmpPath, file.Filename+".arc")
}

func decompressHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "file not received" + err.Error()})
		return
	}

	tmpFile, err := os.CreateTemp("", "uploaded-*.arc")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create temporary file: " + err.Error()})
		return
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create temporary file: " + err.Error()})
		return
	}
	defer src.Close()

	if _, err := io.Copy(tmpFile, src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "file copy error: " + err.Error()})
		return
	}

	tokens, err := archive.ReadArchive(tmpFile.Name())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid archive" + err.Error()})
		return
	}

	data := lz77.Decompress(tokens)

	name := file.Filename
	if filepath.Ext(name) == ".arc" {
		name = name[:len(name)-4]
	}

	c.DataFromReader(http.StatusOK, int64(len(data)), "application/octet-stream", bytes.NewReader(data), map[string]string{
		"Content-Disposition": fmt.Sprintf(`attachment; filename="%s"`, name),
	})

}

// func corsMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(http.StatusOK)
// 			return
// 		}

// 		c.Next()
// 	}
// }
