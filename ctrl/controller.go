package ctrl

type Controller interface {
	Init()
}

type BaseController struct {
}

func (c *BaseController) Init() {

}
