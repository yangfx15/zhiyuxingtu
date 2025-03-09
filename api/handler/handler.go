package handler

import (
    "errors"
    "net/http"
    "encoding/json"
    "xingtu/service"
    "xingtu/model"
    "github.com/gin-gonic/gin"
)

type Handler struct {
    service *service.Service
}

func NewHandler(service *service.Service) *Handler {
    return &Handler{service: service}
}

type LoginRequest struct {
    Code string `json:"code"`
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

    // 调用微信接口获取 openid
    openID, err := getWeChatOpenID(req.Code)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // 存储用户信息
    user := &model.User{
        OpenID:    openID,
        Nickname:  req.Nickname,
        AvatarURL: req.AvatarURL,
    }
    if err := h.service.CreateUser(user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // 返回用户信息
    c.JSON(http.StatusOK, user)
}

// 调用微信接口获取 openid
func getWeChatOpenID(code string) (string, error) {
    // 替换为你的微信 AppID 和 AppSecret
    appID := "wxfe068e6872b90c7a"
    appSecret := "5229e1f5cc6ac21ac061bd6049a44bf5"
    url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + appID + "&secret=" + appSecret + "&js_code=" + code + "&grant_type=authorization_code"

    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    var result struct {
        OpenID     string `json:"openid"`
        SessionKey string `json:"session_key"`
        ErrCode    int    `json:"errcode"`
        ErrMsg     string `json:"errmsg"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return "", err
    }
    if result.ErrCode != 0 {
        return "", errors.New(result.ErrMsg)
    }
    return result.OpenID, nil
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
    Results []Result `json:"results"`
}

type Result struct {
    QuestionId string `json:"questionId"`
    Answer string `json:"answer"`
}

type TestResultResp struct {
    Data TestResultData `json:"data"`
}

type TestResultData struct {
    Score int `json:"score"`
}

// @Summary 提交测试
// @Description 提交测试结果
// @Tags 测试
// @Accept json
// @Produce json
// @Param body body SubmitTestRequest true "测试结果"
// @Success 200 {object} TestResultResp
// @Failure 400 {object} TestResultResp
// @Failure 500 {object} TestResultResp
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

    data := TestResultData {Score: 80}

    resp := TestResultResp{Data: data}

    c.JSON(http.StatusOK, resp)
}