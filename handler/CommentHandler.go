package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/ronak-tatvasoft/go-blog-api/model"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {

	type returnData struct {
		CommentId   int64  `json:"comment_id"`
		ErrorString string `json:"error"`
	}
	errorString := ""
	var commentId int64

	articleId, _ := strconv.Atoi(mux.Vars(r)["id"])

	isExist, err := model.IsArticleExist(articleId)
	if isExist == false || err != nil {
		errorString = "Article is not found."
		if err != nil {
			errorString = err.Error()
		}
		commentId = 0
		w.WriteHeader(http.StatusNotFound)
	} else {
		type CommentRequest struct {
			Nickname string `validate:"required"`
			Content  string `validate:"required"`
		}

		req := &CommentRequest{}
		err = json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			errorString = err.Error()
			commentId = 0
			w.WriteHeader(http.StatusBadRequest)
		} else {
			validate := validator.New()
			err := validate.Struct(req)
			if err != nil {
				commentId = 0
				errorString = err.Error()
				w.WriteHeader(http.StatusBadRequest)
			} else {
				commentId, err = model.NewComment(articleId, req.Nickname, req.Content)
				if err != nil {
					errorString = err.Error()
					w.WriteHeader(http.StatusInternalServerError)
				}
			}
		}
	}

	data := &returnData{
		CommentId:   commentId,
		ErrorString: errorString,
	}

	json.NewEncoder(w).Encode(data)
}

func CommentList(w http.ResponseWriter, r *http.Request) {

	errorString := ""
	articleId, _ := strconv.Atoi(mux.Vars(r)["id"])

	article, err := model.ArticleDetails(articleId)
	if err != nil {
		errorString = "Article is not found."
		w.WriteHeader(http.StatusNotFound)
	}

	comments, err := article.Comments()

	if err != nil {
		errorString = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	}

	type returnData struct {
		Comments    []model.Comment `json:"comments"`
		ErrorString string          `json:"error"`
	}

	data := &returnData{
		Comments:    comments,
		ErrorString: errorString,
	}

	json.NewEncoder(w).Encode(data)
}

func ReplyComment(w http.ResponseWriter, r *http.Request) {

	type returnData struct {
		CommentId   int64  `json:"comment_id"`
		ErrorString string `json:"error"`
	}
	errorString := ""
	var commentId int64

	articleId, _ := strconv.Atoi(mux.Vars(r)["id"])
	parentCommentId, _ := strconv.Atoi(mux.Vars(r)["commnetId"])

	isExist, err := model.IsArticleExist(articleId)
	if isExist == false || err != nil {
		errorString = "Article is not found."
		if err != nil {
			errorString = err.Error()
		}
		commentId = 0
		w.WriteHeader(http.StatusNotFound)
	} else {
		parentComment, err := model.CommentDetails(parentCommentId)
		if err != nil {
			errorString = "Parent comment is not found."
			w.WriteHeader(http.StatusNotFound)
		} else {
			type CommentRequest struct {
				Nickname string `validate:"required"`
				Content  string `validate:"required"`
			}

			req := &CommentRequest{}
			err = json.NewDecoder(r.Body).Decode(req)
			if err != nil {
				errorString = err.Error()
				commentId = 0
				w.WriteHeader(http.StatusBadRequest)
			} else {
				validate := validator.New()
				err := validate.Struct(req)
				if err != nil {
					commentId = 0
					errorString = err.Error()
					w.WriteHeader(http.StatusBadRequest)
				} else {
					commentId, err = parentComment.NewChildComment(articleId, req.Nickname, req.Content)
					if err != nil {
						errorString = err.Error()
						w.WriteHeader(http.StatusInternalServerError)
					}
				}
			}
		}

	}

	data := &returnData{
		CommentId:   commentId,
		ErrorString: errorString,
	}

	json.NewEncoder(w).Encode(data)
}
