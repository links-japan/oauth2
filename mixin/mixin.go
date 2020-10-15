package mixin

import (
	"golang.org/x/oauth2"
)

// Endpoint is Mixin's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://mixin.one/oauth/authorize",
	TokenURL: "https://api.mixin.one/oauth/token",
	AuthStyle: oauth2.AuthStyleMixin,
}

var BackUpEndpoint = oauth2.Endpoint{
	AuthURL:  "https://mixin-www.zeromesh.net/oauth/authorize",
	TokenURL: "https://mixin-api.zeromesh.net/oauth/token",
	AuthStyle: oauth2.AuthStyleMixin,
}
