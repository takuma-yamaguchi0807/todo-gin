package usecase

import (
	"context"

	domainauth "github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/auth"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/user"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/common"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/dto"
)

// UserLoginUsecase はユーザログインのユースケースです。
type UserLoginUsecase struct{
	repo user.Repository
	tok  domainauth.TokenService
}

// NewUserLoginUsecase は UserLoginUsecase のコンストラクタです。
func NewUserLoginUsecase(repo user.Repository, tok domainauth.TokenService) *UserLoginUsecase {
	return &UserLoginUsecase{repo: repo, tok: tok}
}

// Execute は email/password で認証し、成功時にアクセストークンを返す。
func (uc *UserLoginUsecase) Execute(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	email, err := user.NewEmail(req.Email)
	if err != nil { return dto.LoginResponse{}, err }
	password, err := user.NewPassword(req.Password)
	if err != nil { return dto.LoginResponse{}, err }

	uid, ok, err := uc.repo.FindIdByEmailAndPassword(email, password)
	if err != nil { return dto.LoginResponse{}, err }
	if !ok {
		return dto.LoginResponse{}, common.New(common.Unauthorized, "credentials", "invalid email or password")
	}

	claims := domainauth.NewClaims(uid.String())
	token, err := uc.tok.Generate(claims)
	if err != nil { return dto.LoginResponse{}, err }
	return dto.LoginResponse{AccessToken: token}, nil
}