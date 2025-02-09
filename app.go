package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"wails-app/backfront/util"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type UserConfig struct {
	WailUserName  string `env:"WAILUSERNAME,required"`
	WailsPassword string `env:"WAILSPASSWORD,required"`
}

var (
	user UserConfig
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	WailUserName := os.Getenv("WAILUSERNAME")
	if len(WailUserName) == 0 {
		WailUserName = "admin"
	}

	WailsPassword := os.Getenv("WAILSPASSWORD")
	if len(WailsPassword) == 0 {
		WailsPassword = "123456"
	}
	user.WailUserName = WailUserName
	user.WailsPassword = WailsPassword
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Login(req LoginRequest) LoginResponse {
	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		return LoginResponse{Code: util.FAIL_CODE, Message: "参数校验失败"}
	}
	// 简单验证逻辑（生产环境需加密处理）
	if req.Username == user.WailUserName && req.Password == user.WailsPassword {
		return LoginResponse{
			Code:    util.SUCESS_CODE,
			Message: "登录成功",
		}
	}
	return LoginResponse{
		Code:    util.FAIL_CODE,
		Message: "用户名或密码错误",
	}
}

// Login 方法定义
type LoginRequest struct {
	Username string `validate:"required,min=5" json:"username"`
	Password string `validate:"required,min=6" json:"password"`
}

type LoginResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
