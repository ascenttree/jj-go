package update

import "encoding/xml"

type LastVersionResponse struct {
	XMLName xml.Name `xml:"lastversionnum"`
	Value   string   `xml:",chardata"`
}

func LatestVersionHandler(ctx *UpdateContext) {
	ctx.Server.Logger.Debug("Handling request for latest version")

	// The game refuses to start if the version isn't 0.0.0.9
	ctx.XML(200, &LastVersionResponse{
		Value: "0.0.0.9",
	})
}
