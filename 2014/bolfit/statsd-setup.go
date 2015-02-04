import (
	"flag"
	"fmt"

	"github.com/cactus/go-statsd-client/statsd"
)

//START OMIT
var statsdHost = flag.String("statsd.host", "", "host address of the statsd collector. Default is empty which means statsd metrics are logged only.")
var statsdPort = flag.Int("statsd.port", 18125, "port address of the statsd collector.")
var statsdPrefix = flag.String("statsd.prefix", "dev.localhost.boqs", "environment prefix for statds metrics.")

var statsdClient StatsdAccess

func SetupStatsdClient() error {
	if len(*statsdHost) == 0 {
		return fmt.Errorf("no statsd.host given, no statsdclient available")
	}
	std, err := statsd.Dial(fmt.Sprintf("%s:%d", *statsdHost, *statsdPort), *statsdPrefix)
	if err != nil {
		return err
	}
	statsdClient = std
	return nil
}

//END OMIT