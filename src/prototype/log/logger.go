package log

import (
	"github.com/op/go-logging"
	"net/http"
)

var Own = logging.MustGetLogger("Own")

var LoggerHttpClient *http.Client
