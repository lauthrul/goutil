package router

import (
	"github.com/buaazp/fasthttprouter"
	"liushubox/controller"
	"liushubox/controller/md5tool"
	"liushubox/controller/timetool"
	"liushubox/controller/urltool"
)

func Init() *fasthttprouter.Router {
	router := fasthttprouter.New()
	get(router, "/debug/ver", controller.Ver)
	get(router, "/md5/md5", md5tool.Md5)
	get(router, "/md5/md5file", md5tool.Md5File)
	get(router, "/time/add", timetool.Add)
	get(router, "/time/diff", timetool.Diff)
	get(router, "/time/convert", timetool.Covert)
	get(router, "/url/encode", urltool.Encode)
	return router
}
