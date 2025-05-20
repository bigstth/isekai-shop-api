package controller

import (
	"context"
	"net/http"

	"github.com/bigstth/isekai-shop-api/pkg/custom"
	_oauth2Exception "github.com/bigstth/isekai-shop-api/pkg/oauth2/exception"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

func (c *googleOAuth2Controller) PlayerAuthorizing(ctx echo.Context, next echo.HandlerFunc) error {
	pctx := context.Background()
	tokenSource, err := c.getTokenSource(ctx)
	if err != nil {
		return custom.Error(ctx, http.StatusUnauthorized, err.Error())
	}

	if !tokenSource.Valid() {
		tokenSource, err = c.playerTokenRefreshing(ctx, tokenSource)
		if err != nil {
			return custom.Error(ctx, http.StatusUnauthorized, err.Error())
		}
	}

	client := playerGoogleOAuth2Config.Client(pctx, tokenSource)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		return custom.Error(ctx, http.StatusUnauthorized, err.Error())
	}

	if !c.oauth2Service.IsPlayerExist(userInfo.ID) {
		return custom.Error(ctx, http.StatusForbidden, "No permission")
	}

	ctx.Set("playerID", userInfo.ID)

	return next(ctx)
}

func (c *googleOAuth2Controller) AdminAuthorizing(ctx echo.Context, next echo.HandlerFunc) error {
	pctx := context.Background()
	tokenSource, err := c.getTokenSource(ctx)
	if err != nil {
		return custom.Error(ctx, http.StatusUnauthorized, err.Error())
	}

	if !tokenSource.Valid() {
		tokenSource, err = c.adminTokenRefreshing(ctx, tokenSource)
		if err != nil {
			return custom.Error(ctx, http.StatusUnauthorized, err.Error())
		}
	}

	client := adminGoogleOAuth2Config.Client(pctx, tokenSource)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		return custom.Error(ctx, http.StatusUnauthorized, err.Error())
	}

	if !c.oauth2Service.IsAdminExist(userInfo.ID) {
		return custom.Error(ctx, http.StatusForbidden, "No permission")
	}

	ctx.Set("adminID", userInfo.ID)

	return next(ctx)
}

func (c *googleOAuth2Controller) playerTokenRefreshing(ctx echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	pctx := context.Background()

	updatedToken, err := playerGoogleOAuth2Config.TokenSource(pctx, token).Token()
	if err != nil {
		return nil, &_oauth2Exception.UnauthorizedException{}
	}

	c.setSameSiteCookie(oauth2AccessTokenCookieName, updatedToken.AccessToken, ctx)
	c.setSameSiteCookie(oauth2RefreshTokenCookieName, updatedToken.RefreshToken, ctx)
	return updatedToken, nil
}

func (c *googleOAuth2Controller) adminTokenRefreshing(ctx echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	pctx := context.Background()

	updatedToken, err := adminGoogleOAuth2Config.TokenSource(pctx, token).Token()
	if err != nil {
		return nil, &_oauth2Exception.UnauthorizedException{}
	}

	c.setSameSiteCookie(oauth2AccessTokenCookieName, updatedToken.AccessToken, ctx)
	c.setSameSiteCookie(oauth2RefreshTokenCookieName, updatedToken.RefreshToken, ctx)
	return updatedToken, nil
}

func (c *googleOAuth2Controller) getTokenSource(ctx echo.Context) (*oauth2.Token, error) {
	accessToken, err := ctx.Cookie(oauth2AccessTokenCookieName)

	if err != nil {
		return nil, &_oauth2Exception.UnauthorizedException{}
	}

	refreshToken, err := ctx.Cookie(oauth2RefreshTokenCookieName)

	if err != nil {
		return nil, &_oauth2Exception.UnauthorizedException{}
	}

	return &oauth2.Token{
		AccessToken:  accessToken.Value,
		RefreshToken: refreshToken.Value,
	}, nil
}
