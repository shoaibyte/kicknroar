// Package web provides the embedded frontend build (web/dist) for serving the SPA.
// Build order: run "yarn build" in web/ so that web/dist exists before "go build".
package web

import "embed"

// DistFS is the embedded filesystem of the built frontend (web/dist).
// It is populated when building the frontend with "yarn build" before go build.
//
//go:embed all:dist
var DistFS embed.FS
