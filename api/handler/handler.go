package handler

import (
    "net/http"
    "xingtu/service"
    "github.com/gin-gonic/gin"
)

type Handler struct {
    service *service.Service
}

func NewHandler(service *service.Service) *Handler {
    return &Handler{service: service}
}

type LoginRequest struct {
    OpenID    string `json:"openid"`
    Nickname  string `json:"nickname"`
    AvatarURL string `json:"avatar_url"`
}

// @Summary 用户登录
// @Description 用户通过微信登录
// @Tags 用户
// @Accept json
// @Produce json
// @Param body body LoginRequest true "登录请求"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /login [post]
func (h *Handler) Login(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user, err := h.service.Login(req.OpenID, req.Nickname, req.AvatarURL)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}

// @Summary 获取题目
// @Description 根据类目和类型获取题目
// @Tags 题目
// @Accept json
// @Produce json
// @Param category query string true "题目类目"
// @Param type query string true "题目类型"
// @Success 200 {array} model.Question
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /questions [get]
func (h *Handler) GetQuestions(c *gin.Context) {
    category := c.Query("category")
    questionType := c.Query("type")
    limit := 10 // 默认10道题
    questions, err := h.service.GetQuestions(category, questionType, limit)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, questions)
}

type SubmitTestRequest struct {
    Results []Result
}

type Result struct {
    QuestionId string
    Answer string
}

// @Summary 提交测试
// @Description 提交测试结果
// @Tags 测试
// @Accept json
// @Produce json
// @Param body body SubmitTestRequest true "测试结果"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /submit-test [post]
func (h *Handler) SubmitTest(c *gin.Context) {
    var req struct {
        UserID   uint   `json:"user_id"`
        Category string `json:"category"`
        Score    int    `json:"score"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.service.SubmitTest(req.UserID, req.Category, req.Score); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "success"})
}