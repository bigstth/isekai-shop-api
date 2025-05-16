package controller

import (
	"context"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/avast/retry-go"
	"github.com/bigstth/isekai-shop-api/config"
	"github.com/bigstth/isekai-shop-api/pkg/custom"
	_oauth2Exception "github.com/bigstth/isekai-shop-api/pkg/oauth2/exception"
	_oauth2Model "github.com/bigstth/isekai-shop-api/pkg/oauth2/model"
	_oauth2Service "github.com/bigstth/isekai-shop-api/pkg/oauth2/service"
	_playerModel "github.com/bigstth/isekai-shop-api/pkg/player/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type googleOAuth2Controller struct {
	oauth2Service _oauth2Service.OAuth2Service
	oauth2Conf    *config.OAuth2
	logger        echo.Logger
}

var (
	playerGoogleOAuth2Config *oauth2.Config
	adminGoogleOAuth2Config  *oauth2.Config
	once                     sync.Once

	oauth2AccessTokenCookieName  = "act"
	oauth2RefreshTokenCookieName = "rft"
	stateCookieName              = "state"

	letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func NewGoogleOAuth2Controller(
	oauth2Service _oauth2Service.OAuth2Service,
	oauth2Conf *config.OAuth2,
	logger echo.Logger,
) *googleOAuth2Controller {
	once.Do(func() {
		setGoogleOAuth2Config(oauth2Conf)
	})
	return &googleOAuth2Controller{
		oauth2Service: oauth2Service,
		oauth2Conf:    oauth2Conf,
		logger:        logger,
	}
}

func setGoogleOAuth2Config(oauth2Conf *config.OAuth2) {
	playerGoogleOAuth2Config = &oauth2.Config{
		ClientID:     oauth2Conf.ClientId,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.PlayerRedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Conf.Endpoints.AuthUrl,
			TokenURL:      oauth2Conf.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.Endpoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}

	adminGoogleOAuth2Config = &oauth2.Config{
		ClientID:     oauth2Conf.ClientId,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.AdminRedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Conf.Endpoints.AuthUrl,
			TokenURL:      oauth2Conf.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.Endpoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}
}

func (c *googleOAuth2Controller) PlayerLogin(ctx echo.Context) error {
	state := c.randomState()
	c.setCookie(stateCookieName, state, ctx)

	return ctx.Redirect(http.StatusFound, playerGoogleOAuth2Config.AuthCodeURL(state))
}

func (c *googleOAuth2Controller) AdminLogin(ctx echo.Context) error {
	state := c.randomState()
	c.setCookie(stateCookieName, state, ctx)

	return ctx.Redirect(http.StatusFound, adminGoogleOAuth2Config.AuthCodeURL(state))
}

func (c *googleOAuth2Controller) PlayerLoginCallback(ctx echo.Context) error {
	pctx := context.Background()

	if err := retry.Do(func() error {
		return c.callbackValidating(ctx)
	}, retry.Attempts(3), retry.Delay(3*time.Second)); err != nil {
		c.logger.Error("Failed to validate callback %s", err.Error())
		return custom.Error(ctx, http.StatusUnauthorized, err.Error())
	}

	token, err := playerGoogleOAuth2Config.Exchange(pctx, ctx.QueryParam("code"))
	if err != nil {
		c.logger.Error("Failed to exchange token %s", err.Error())
		return custom.Error(ctx, http.StatusUnauthorized, err.Error())
	}

	client := playerGoogleOAuth2Config.Client(pctx, token)
	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Error("Failed to get user info %s", err.Error())
		return custom.Error(ctx, http.StatusUnauthorized, err.Error())
	}

	playerCreatingRequest := &_playerModel.PlayerCreatingReq{
		ID:     userInfo.ID,
		Name:   userInfo.Name,
		Email:  userInfo.Email,
		Avatar: userInfo.Picture,
	}

	if err := c.oauth2Service.PlayerAccountCreating(playerCreatingRequest); err != nil {
		c.logger.Error("Failed to create player account %s", err.Error())
		return custom.Error(ctx, http.StatusUnauthorized, err.Error())
	}

	c.setSameSiteCookie(oauth2AccessTokenCookieName, token.AccessToken, ctx)
	c.setSameSiteCookie(oauth2RefreshTokenCookieName, token.RefreshToken, ctx)
	return ctx.JSON(http.StatusOK, &_oauth2Model.LoginResponse{Message: "Login Success"})
}

func (c *googleOAuth2Controller) AdminLoginCallback(ctx echo.Context) error {
	// ctx := context.Background()

	if err := retry.Do(func() error {
		return c.callbackValidating(ctx)
	}, retry.Attempts(3), retry.Delay(3*time.Second)); err != nil {
		c.logger.Error("Failed to validate callback %s", err.Error())
		return custom.Error(ctx, http.StatusUnauthorized, err.Error())
	}

	return ctx.JSON(http.StatusOK, &_oauth2Model.LoginResponse{Message: "Login Success"})
}

func (c *googleOAuth2Controller) Logout(ctx echo.Context) error {
	c.removeCookie(oauth2AccessTokenCookieName, ctx)
	c.removeCookie(oauth2RefreshTokenCookieName, ctx)
	c.removeCookie(stateCookieName, ctx)

	return ctx.NoContent(http.StatusNoContent)
}

func (c *googleOAuth2Controller) setCookie(name, value string, ctx echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Path = "/"
	cookie.HttpOnly = true
	// cookie.Secure = true
	// cookie.SameSite = http.SameSiteStrictMode
	ctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) removeCookie(name string, ctx echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = ""
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.MaxAge = -1
	// cookie.Secure = true
	// cookie.SameSite = http.SameSiteStrictMode
	ctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) setSameSiteCookie(name, value string, ctx echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Path = "/"
	cookie.HttpOnly = true
	// cookie.Secure = true
	cookie.SameSite = http.SameSiteStrictMode
	ctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) removeSameSiteCookie(name string, ctx echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = ""
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.MaxAge = -1
	// cookie.Secure = true
	cookie.SameSite = http.SameSiteStrictMode
	ctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) getUserInfo(client *http.Client) (*_oauth2Model.UserInfo, error) {
	resp, err := client.Get(c.oauth2Conf.UserInfoUrl)
	if err != nil {
		c.logger.Error("Failed to get user info %s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	userInfoInBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("Failed to read user info %s", err.Error())
		return nil, err
	}

	userInfo := new(_oauth2Model.UserInfo)
	if err := json.Unmarshal(userInfoInBytes, userInfo); err != nil {
		c.logger.Error("Failed to unmarshal user info %s", err.Error())
		return nil, err
	}

	return userInfo, nil
}

func (c *googleOAuth2Controller) callbackValidating(ctx echo.Context) error {
	state := ctx.QueryParam("state")
	stateFromCookie, err := ctx.Cookie(stateCookieName)

	if err != nil {
		c.logger.Error("state cookie not found")
		return &_oauth2Exception.UnauthorizedException{}
	}

	if state != stateFromCookie.Value {
		c.logger.Error("state not match")
		return &_oauth2Exception.UnauthorizedException{}
	}

	c.removeCookie(stateCookieName, ctx)
	return nil
}

func (c *googleOAuth2Controller) randomState() string {
	b := make([]byte, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
