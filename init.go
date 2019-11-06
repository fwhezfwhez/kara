package kara

import "sync"

func init() {
	KaraPool = sync.Map{}
}
