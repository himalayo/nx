package util

import (
	gd "github.com/himalayo/nx/gamedata"

	"github.com/himalayo/nx/cmd/nx/spinner"
)

func LoadGameData(mgr gd.Manager, message string, types ...gd.Type) error {
	return spinner.DoErr(message, func() error {
		return mgr.Load(types...)
	})
}

func LoadTexts(mgr gd.Manager) error {
	return LoadGameData(mgr, "Loading external texts...", gd.GameDataTexts)
}

func LoadFurni(mgr gd.Manager) error {
	return LoadGameData(mgr, "Loading furni data...", gd.GameDataFurni)
}

func LoadFigure(mgr gd.Manager) error {
	return LoadGameData(mgr, "Loading figure data...", gd.GameDataFigure)
}
