package util

import (
	"time"
)

func TimeNowJSTZone() time.Time {
	jstZone := time.FixedZone("Asia/Tokyo", 9*60*60)
	return time.Now().In(jstZone)
}
