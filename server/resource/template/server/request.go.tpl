package request

import (
	"server/model/{{.Package}}"
	"server/model/common/request"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    request.PageInfo
}
