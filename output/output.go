package output

import (
	"io/ioutil"
	"log"
	"os"

	jww "github.com/spf13/jwalterweatherman"
)

var (
	DEBUG *log.Logger
	INFO  *log.Logger
	ERROR *log.Logger
	FATAL *log.Logger

	np *jww.Notepad
)

const DefaultVerbosity = jww.LevelInfo

func Init(verbosity int) {
	threshold := DefaultVerbosity
	if verbosity == 0 {
		threshold = jww.LevelFatal
	} else if verbosity == 1 {
		threshold = jww.LevelInfo
	} else if verbosity == 2 {
		threshold = jww.LevelDebug
	}

	np = jww.NewNotepad(threshold, threshold, os.Stdout, ioutil.Discard, "", 0)
	DEBUG = np.DEBUG
	INFO = np.INFO
	ERROR = np.ERROR
	FATAL = np.FATAL
}
