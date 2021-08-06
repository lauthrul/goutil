package router

import (
	"github.com/buaazp/fasthttprouter"
	"liushubox/controller"
	"liushubox/controller/md5tool"
	"liushubox/controller/timetool"
)

func Init() *fasthttprouter.Router {
	router := fasthttprouter.New()
	get(router, "/debug/ver", controller.Ver)
	get(router, "/md5/checksum", md5tool.CheckSum)
	get(router, "/time/add", timetool.Add)
	get(router, "/time/diff", timetool.Diff)
	return router
}
