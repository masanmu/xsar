package xsar

import (
	"fmt"
	"github.com/xsar/module"
	"reflect"
)

type Module module.ModuleConfig

func (m *Module) ListModule() error {
	fmt.Println("enalbe modules:")
	t := reflect.TypeOf(m)

	for i := 0; i <= 100; i++ {
		field := t.Elem().Field(i)
		if len(field.Index) > 0 {
			fmt.Printf("|----%s\n", field.Name)
		}
	}
	return nil
}
