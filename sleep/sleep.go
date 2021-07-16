package sleep

import (
	"fmt"
	"time"
)

func SleepsAndLog(seconds int32) {
	sleepFor := time.Duration(seconds) * time.Second
	dt := time.Now()
	fmt.Printf("Sleeping for '%.2f' seconds. Start Time: '%s'\n", sleepFor.Seconds(), dt.Format(time.RFC3339))

	time.Sleep(sleepFor)

	dt = time.Now()
	fmt.Printf("Finished sleeping at: '%s'\n", dt.Format(time.RFC3339))
}
