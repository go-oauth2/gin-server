package ginserver

import (
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/server"
)

// SetTokenType token type
func SetTokenType(tokenType string) {
	gServer.Config.TokenType = tokenType
}

// SetAllowGetAccessRequest to allow GET requests for the token
func SetAllowGetAccessRequest(allow bool) {
	gServer.Config.AllowGetAccessRequest = allow
}

// SetAllowedResponseType allow the authorization types
func SetAllowedResponseType(types ...oauth2.ResponseType) {
	gServer.Config.AllowedResponseTypes = types
}

// SetAllowedGrantType allow the grant types
func SetAllowedGrantType(types ...oauth2.GrantType) {
	gServer.Config.AllowedGrantTypes = types
}

// SetClientInfoHandler get client info from request
func SetClientInfoHandler(handler server.ClientInfoHandler) {
	gServer.ClientInfoHandler = handler
}

// SetClientAuthorizedHandler check the client allows to use this authorization grant type
func SetClientAuthorizedHandler(handler server.ClientAuthorizedHandler) {
	gServer.ClientAuthorizedHandler = handler
}

// SetClientScopeHandler check the client allows to use scope
func SetClientScopeHandler(handler server.ClientScopeHandler) {
	gServer.ClientScopeHandler = handler
}

// SetUserAuthorizationHandler get user id from request authorization
func SetUserAuthorizationHandler(handler server.UserAuthorizationHandler) {
	gServer.UserAuthorizationHandler = handler
}

// SetPasswordAuthorizationHandler get user id from username and password
func SetPasswordAuthorizationHandler(handler server.PasswordAuthorizationHandler) {
	gServer.PasswordAuthorizationHandler = handler
}

// SetRefreshingScopeHandler check the scope of the refreshing token
func SetRefreshingScopeHandler(handler server.RefreshingScopeHandler) {
	gServer.RefreshingScopeHandler = handler
}

// SetResponseErrorHandler response error handling
func SetResponseErrorHandler(handler server.ResponseErrorHandler) {
	gServer.ResponseErrorHandler = handler
}

// SetInternalErrorHandler internal error handling
func SetInternalErrorHandler(handler server.InternalErrorHandler) {
	gServer.InternalErrorHandler = handler
}

// SetExtensionFieldsHandler in response to the access token with the extension of the field
func SetExtensionFieldsHandler(handler server.ExtensionFieldsHandler) {
	gServer.ExtensionFieldsHandler = handler
}

// SetAccessTokenExpHandler set expiration date for the access token
func SetAccessTokenExpHandler(handler server.AccessTokenExpHandler) {
	gServer.AccessTokenExpHandler = handler
}

// SetAuthorizeScopeHandler set scope for the access token
func SetAuthorizeScopeHandler(handler server.AuthorizeScopeHandler) {
	gServer.AuthorizeScopeHandler = handler
}
