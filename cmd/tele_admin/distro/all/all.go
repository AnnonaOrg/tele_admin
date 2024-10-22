package all

import (
	_ "github.com/umfaka/umfaka_core/internal/log"

	_ "github.com/umfaka/umfaka_core/internal/features/about_features"
	_ "github.com/umfaka/umfaka_core/internal/features/ping_features"
	_ "github.com/umfaka/umfaka_core/internal/features/start_features"

	_ "github.com/umfaka/umfaka_core/internal/features/callback"
	_ "github.com/umfaka/umfaka_core/internal/features/text"

	_ "github.com/umfaka/umfaka_core/internal/features/ban_features"
)
