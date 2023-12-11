package globals

import (
	"time"
)

var (
	last_upload time.Time
)

func UploadCompleted() {
	last_upload = time.Now()
}
func LastUploadCompleted() time.Time {
	return last_upload
}












































































