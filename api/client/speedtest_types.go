package client

import (
	"log/slog"
)

type SpeedtestClient struct {
	V3Client *Client
	Logger   *slog.Logger
}
