package test

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "testing"

    "github.com/ronak-tatvasoft/go-blog-api/model"
)

func TestArtilceList(t *testing.T) {
    upMigrations()

    // insert one sample article
    articleId, _ := model.NewArticle("testName", "testTitle", "testContent")

    type ReturnData struct {
        Articles    []model.Article `json:"articles"`
        ErrorString string          `json:"error"`
    }

    req, _ := http.NewRequest("GET", "/articles", nil)
    response := executeRequest(req)

    var returnData ReturnData
    json.Unmarshal(response.Body.Bytes(), &returnData)

    if returnData.Articles[0].ID != articleId {
        t.Errorf("Expected Id not match with actual Id")
    }
    if returnData.Articles[0].Nickname != "testName" {
        t.Errorf("Expected nickname not match with actual nickname")
    }
    if returnData.Articles[0].Title != "testTitle" {
        t.Errorf("Expected title not match with actual title")
    }
    if returnData.Articles[0].Content != "testContent" {
        t.Errorf("Expected content not match with actual content")
    }
    checkResponseCode(t, http.StatusOK, response.Code)

    downMigrations()
}

func TestCreateArticle(t *testing.T) {
    upMigrations()

    type ReturnData struct {
        ArticleId   int64  `json:"article_id"`
        ErrorString string `json:"error"`
    }

    var jsonStr = []byte(`{"nickname":"testName","title":"testTitle","content":"testContent"}`)
    req, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(jsonStr))
    response := executeRequest(req)

    var returnData ReturnData
    json.Unmarshal(response.Body.Bytes(), &returnData)

    if returnData.ArticleId == 0 {
        t.Errorf("Article is not created")
    }
    if returnData.ErrorString != "" {
        t.Errorf("Article is not created")
    }
    checkResponseCode(t, http.StatusOK, response.Code)

    downMigrations()
}

func TestCreateArticleFailed(t *testing.T) {
    upMigrations()

    type ReturnData struct {
        ArticleId   int64  `json:"article_id"`
        ErrorString string `json:"error"`
    }

    var jsonStr = []byte(`{"nickname":"","title":"testTitle","content":"testContent"}`)
    req, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(jsonStr))
    response := executeRequest(req)

    var returnData ReturnData
    json.Unmarshal(response.Body.Bytes(), &returnData)

    if returnData.ArticleId != 0 {
        t.Errorf("Test case failed")
    }
    if returnData.ErrorString == "" {
        t.Errorf("Test case failed")
    }
    checkResponseCode(t, http.StatusBadRequest, response.Code)

    downMigrations()
}

func TestArticleDetails(t *testing.T) {
    upMigrations()

    type ReturnData struct {
        Article     model.Article `json:"article"`
        ErrorString string        `json:"error"`
    }

    // insert one sample article
    articleId, _ := model.NewArticle("testName", "testTitle", "testContent")

    req, _ := http.NewRequest("GET", fmt.Sprintf("/articles/%d", articleId), nil)
    response := executeRequest(req)

    var returnData ReturnData
    json.Unmarshal(response.Body.Bytes(), &returnData)

    if returnData.Article.ID != articleId {
        t.Errorf("Expected Id not match with actual Id")
    }
    if returnData.Article.Nickname != "testName" {
        t.Errorf("Expected nickname not match with actual nickname")
    }
    if returnData.Article.Title != "testTitle" {
        t.Errorf("Expected title not match with actual title")
    }
    if returnData.Article.Content != "testContent" {
        t.Errorf("Expected content not match with actual content")
    }

    checkResponseCode(t, http.StatusOK, response.Code)

    downMigrations()
}

func TestArticleDetailsFailed(t *testing.T) {
    upMigrations()

    type ReturnData struct {
        Article     model.Article `json:"article"`
        ErrorString string        `json:"error"`
    }

    req, _ := http.NewRequest("GET", fmt.Sprintf("/articles/%d", 1), nil)
    response := executeRequest(req)

    var returnData ReturnData
    json.Unmarshal(response.Body.Bytes(), &returnData)

    if returnData.Article.ID != 0 {
        t.Errorf("Test case failed")
    }

    if returnData.ErrorString == "" {
        t.Errorf("Test case failed")
    }

    checkResponseCode(t, http.StatusNotFound, response.Code)

    downMigrations()
}
