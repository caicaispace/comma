package l

import (
	"fmt"
	"time"

	"goaway/pkg/library/setting"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.App.LogPath, "")
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.App.LogPrefix,
		time.Now().Format(setting.App.TimeFormat),
		setting.App.LogExtension,
	)
}
