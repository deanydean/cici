package watch

import (
	"fmt"
	"os"
	"path"

	"github.com/deanydean/clockwork/core/triggers"
	"github.com/deanydean/clockwork/core/utils"
	"github.com/deanydean/clockwork/core/watchfiles"
)

// Start the watch service
func Start() {
	var cwd, err = os.Getwd()
	if err != nil {
		fmt.Println("Can't to get working directory, cannot read Watchfile")
		return
	}

	utils.SetGlobalLogLevel(utils.LogDebug)

	var watchFileName = path.Join(cwd, "Watchfile")
	var watcher = watchfiles.GetWatcherFor(&watchFileName)

	if watcher == nil {
		fmt.Println("If I could, I'd watch stuff from", watchFileName)
		return
	}

	watcher.Watch(triggers.NewTextReporterTrigger("Something happened!"))
}
