package web

import (
	"fmt"
	"net/http"
	"time"

	"webook/internal/domain"
	"webook/internal/service"

	regexp "github.com/dlclark/regexp2"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

const (
	emailRegexPattern    = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,72}$`
)

type UserHandler struct {
	emailRegexExp    *regexp.Regexp
	passwordRegexExp *regexp.Regexp

	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		emailRegexExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRegexExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		svc:              svc,
	}
}

func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.POST("/signup", h.SignUp)
	ug.POST("/login", h.Login)
	ug.POST("/edit", h.Edit)
	ug.GET("/profile", h.Profile)
}

func (h *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}

	req := SignUpReq{}
	if err := ctx.Bind(&req); err != nil {
		return
	}

	isEmail, err := h.emailRegexExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系統錯誤")
		return
	}
	if !isEmail {
		ctx.String(http.StatusOK, "郵箱不正確")
		return
	}

	if req.Password != req.ConfirmPassword {
		ctx.String(http.StatusOK, "兩次輸入的密碼不同")
		return
	}

	isPassword, err := h.passwordRegexExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系統錯誤")
		return
	}
	if !isPassword {
		ctx.String(http.StatusOK, "密碼必須包含數字, 特殊字符, 並且長度不能小於 8 位")
	}

	err = h.svc.Signup(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})

	switch err {
	case nil:
		ctx.String(http.StatusOK, "Success")

	case service.ErrDuplicateEmail:
		ctx.String(http.StatusOK, "email重複")
	default:
		fmt.Errorf("Failed to create user: %s", err.Error())
		ctx.String(http.StatusOK, "Failed to create user")
	}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req Req
	if err := ctx.Bind(&req); err != nil {
		fmt.Errorf("Failed to parse body, err: %s", err.Error())
		return
	}

	user, err := h.svc.Login(ctx, req.Email, req.Password)
	switch err {
	case nil:
		uc := UserClaims{
			Uid:       user.ID,
			UserAgent: ctx.GetHeader("User-Agent"),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 1)),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, uc)
		tokenStr, err := token.SignedString(JWTKey)
		if err != nil {
			ctx.String(http.StatusOK, "系統錯誤")
			return
		}
		ctx.Header("x-jwt-token", tokenStr)
		ctx.String(http.StatusOK, "登錄成功")
	case service.ErrInvalidUserOrPassword:
		ctx.String(http.StatusOK, "用戶名或密碼錯誤")
	default:
		ctx.String(http.StatusOK, "系統錯誤")
	}

}

func (h *UserHandler) Edit(ctx *gin.Context) {
}

func (h *UserHandler) Profile(ctx *gin.Context) {
	us := ctx.MustGet("user").(UserClaims)
	ctx.String(http.StatusOK, "Profile, uid: %d", us.Uid)
}

var JWTKey = []byte("vYGCs=Y=vfSTmpM$D?+.zKUwr#bnYV6pysE2H?BD9f-37N$bFv%$V8ErrqT+AVTZ")

type UserClaims struct {
	jwt.RegisteredClaims
	Uid       int64
	UserAgent string
}
