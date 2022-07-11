package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"student_data/Config"
	"student_data/Controllers"
	"student_data/Models"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestUpdateStudent(t *testing.T) {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	studentModel := &Models.Student{}
	Config.DB.AutoMigrate(studentModel)
	r := SetUpRouter()
	routePath := r.Group("/student-api")
	routePath.PUT("student/:id", Controllers.UpdateStudent)
	jsonStr := []byte(`{"name": "put_name", "lastname": "fire", "dob": "2021-03-01", "address": "21 jump street", "subject": "Chemistry", "Marks": 100}`)
	req, _ := http.NewRequest("PUT", "/student-api/student/"+fmt.Sprint(1), bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
