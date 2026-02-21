package workspaceassets

import "embed"

const Root = "workspace"

// FS contains bundled onboarding workspace templates.
//
//go:embed workspace/**
var FS embed.FS
