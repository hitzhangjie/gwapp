package ctrl

import (
	"bytes"
	"html/template"
	"io/ioutil"
)

type Controller interface {
}

type BaseController struct {
	Template string
	Data     interface{}
}

func (c *BaseController) Render() ([]byte, error) {
	buf, err := ioutil.ReadFile(c.Template)
	if err != nil {
		return nil, err
	}

	tpl, err := template.New(c.Template).Parse(string(buf))
	if err != nil {
		return nil, err
	}

	dat := bytes.Buffer{}
	err = tpl.Execute(&dat, c.Data)
	if err != nil {
		return nil, err
	}
	return dat.Bytes(), nil
}
