package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/ronak-tatvasoft/go-blog-api/model"
)

func ArtilceList(w http.ResponseWriter, r *http.Request) {

	v := r.URL.Query()
	page, _ := strconv.Atoi(v.Get("page"))

	offset := 0
	limit := 20

	if page != 0 {
		offset = (limit * (page - 1))
	}

	articles, err := model.ArticlesList(offset, limit)

	errorString := ""
	if err != nil {
		errorString = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	}

	type returnData struct {
		Articles    []model.Article `json:"articles"`
		ErrorString string          `json:"error"`
	}

	data := &returnData{
		Articles:    articles,
		ErrorString: errorString,
	}

	json.NewEncoder(w).Encode(data)
}

func ArticleDetails(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	article, err := model.ArticleDetails(id)

	errorString := ""
	if err != nil {
		errorString = err.Error()
		w.WriteHeader(http.StatusNotFound)
	}

	type returnData struct {
		Article     model.Article `json:"article"`
		ErrorString string        `json:"error"`
	}

	data := &returnData{
		Article:     article,
		ErrorString: errorString,
	}

	json.NewEncoder(w).Encode(data)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {

	type returnData struct {
		ArticleId   int64  `json:"article_id"`
		ErrorString string `json:"error"`
	}
	errorString := ""
	var articleId int64

	type ArticleRequest struct {
		Nickname string `validate:"required"`
		Title    string `validate:"required"`
		Content  string `validate:"required"`
	}

	req := &ArticleRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		errorString = err.Error()
		articleId = 0
		w.WriteHeader(http.StatusBadRequest)
	} else {
		validate := validator.New()
		err := validate.Struct(req)
		if err != nil {
			articleId = 0
			errorString = err.Error()
			w.WriteHeader(http.StatusBadRequest)
		} else {
			articleId, err = model.NewArticle(req.Nickname, req.Title, req.Content)
			if err != nil {
				errorString = err.Error()
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}

	data := &returnData{
		ArticleId:   articleId,
		ErrorString: errorString,
	}

	json.NewEncoder(w).Encode(data)
}
