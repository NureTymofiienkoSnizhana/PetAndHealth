package handlers

import (
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/data"
	"net/http"
)

const MasterDBContextKey = iota

func MongoDB(r *http.Request) data.MasterDB {
	return r.Context().Value(MasterDBContextKey).(data.MasterDB)
}
