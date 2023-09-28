package test

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "testing"

    "github.com/ronak-tatvasoft/go-blog-api/model"
)

func TestCommentList(t *testing.T) {
    upMigrations()

    // insert one sample article and sample comment
    articleId, _ := model.NewArticle("testName", "testTitle", "testContent")
    commentId, _ := model.NewComment(int(articleId), "testName", "testContent")

    type ReturnData struct {
        Comments    []model.Comment `json:"comments"`
        ErrorString string          `json:"error"`
    }

    req, _ := http.NewRequest("GET", fmt.Sprintf("/articles/%d/comments", articleId), nil)
    response := executeRequest(req)

    var returnData ReturnData
    json.Unmarshal(response.Body.Bytes(), &returnData)

    if returnData.Comments[0].ID != commentId {
        t.Errorf("Expected Id not match with actual Id")
    }
    if returnData.Comments[0].Nickname != "testName" {
        t.Errorf("Expected nickname not match with actual nickname")
    }
    if returnData.Comments[0].Content != "testContent" {
        t.Errorf("Expected content not match with actual content")
    }
    checkResponseCode(t, http.StatusOK, response.Code)

    downMigrations()
}

func TestCreateComment(t *testing.T) {
    upMigrations()

    // insert one sample article
    articleId, _ := model.NewArticle("testName", "testTitle", "testContent")

    type ReturnData struct {
        CommentId   int64  `json:"comment_id"`
        ErrorString string `json:"error"`
    }

    var jsonStr = []byte(`{"nickname":"testName","content":"testContent"}`)
    req, _ := http.NewRequest("POST", fmt.Sprintf("/articles/%d/comments", articleId), bytes.NewBuffer(jsonStr))
    response := executeRequest(req)

    var returnData ReturnData
    json.Unmarshal(response.Body.Bytes(), &returnData)

    if returnData.CommentId == 0 {
        t.Errorf("Comment is not created")
    }
    if returnData.ErrorString != "" {
        t.Errorf("Comment is not created")
    }
    checkResponseCode(t, http.StatusOK, response.Code)

    downMigrations()
}

func TestCreateCommentFailed(t *testing.T) {
    upMigrations()

    // insert one sample article
    articleId, _ := model.NewArticle("testName", "testTitle", "testContent")

    type ReturnData struct {
        CommentId   int64  `json:"comment_id"`
        ErrorString string `json:"error"`
    }

    var jsonStr = []byte(`{"nickname":"","content":"testContent"}`)
    req, _ := http.NewRequest("POST", fmt.Sprintf("/articles/%d/comments", articleId), bytes.NewBuffer(jsonStr))
    response := executeRequest(req)

    var returnData ReturnData
    json.Unmarshal(response.Body.Bytes(), &returnData)

    if returnData.CommentId != 0 {
        t.Errorf("Test case failed")
    }
    if returnData.ErrorString == "" {
        t.Errorf("Test case failed")
    }
    checkResponseCode(t, http.StatusBadRequest, response.Code)

    downMigrations()
}

func TestReplyComment(t *testing.T) {
    upMigrations()

    // insert one sample article and parent comment
    articleId, _ := model.NewArticle("testName", "testTitle", "testContent")
    parentCommentId, _ := model.NewComment(int(articleId), "testName", "testContent")

    type ReturnData struct {
        CommentId   int64  `json:"comment_id"`
        ErrorString string `json:"error"`
    }

    var jsonStr = []byte(`{"nickname":"testName2","content":"testContent2"}`)
    req, _ := http.NewRequest("POST", fmt.Sprintf("/articles/%d/comments/%d/reply", articleId, parentCommentId), bytes.NewBuffer(jsonStr))
    response := executeRequest(req)

    var returnData ReturnData
    json.Unmarshal(response.Body.Bytes(), &returnData)

    if returnData.CommentId == 0 {
        t.Errorf("Comment is not created")
    }
    if returnData.ErrorString != "" {
        t.Errorf("Comment is not created")
    }
    checkResponseCode(t, http.StatusOK, response.Code)

    downMigrations()
}

func TestReplyCommentFailed(t *testing.T) {
    upMigrations()

    // insert one sample article and parent comment
    articleId, _ := model.NewArticle("testName", "testTitle", "testContent")
    parentCommentId, _ := model.NewComment(int(articleId), "testName", "testContent")

    type ReturnData struct {
        CommentId   int64  `json:"comment_id"`
        ErrorString string `json:"error"`
    }

    var jsonStr = []byte(`{"nickname":"","content":"testContent2"}`)
    req, _ := http.NewRequest("POST", fmt.Sprintf("/articles/%d/comments/%d/reply", articleId, parentCommentId), bytes.NewBuffer(jsonStr))
    response := executeRequest(req)

    var returnData ReturnData
    json.Unmarshal(response.Body.Bytes(), &returnData)

    if returnData.CommentId != 0 {
        t.Errorf("Test case is failed")
    }
    if returnData.ErrorString == "" {
        t.Errorf("Test case is failed")
    }
    checkResponseCode(t, http.StatusBadRequest, response.Code)

    downMigrations()
}
