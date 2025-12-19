package config

import (
	"log/slog"

	"github.com/CHILLERAN/PersonalityAPI/internal/models"
)

type Application struct {
	TraitModel *models.TraitModel
	Logger *slog.Logger
}
