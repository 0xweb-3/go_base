package b

import (
	"fmt"
	"github.com/0xweb-3/go_base/02_internal/b/internal/c"
)

type D struct {
}

func d() {
	c := c.C{}
	fmt.Println(c)
}
