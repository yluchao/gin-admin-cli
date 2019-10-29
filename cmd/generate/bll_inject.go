package generate

import (
	"context"
	"fmt"
	"strings"
)

func getBllInjectFileName(dir string) string {
	fullname := fmt.Sprintf("%s/internal/app/bll/impl/impl.go", dir)
	return fullname
}

// 插入bll注入文件
func insertBllInject(ctx context.Context, pkgName, dir, name, comment string) error {
	fullname := getBllInjectFileName(dir)

	injectContent := fmt.Sprintf("container.Provide(internal.New%s, dig.As(new(bll.I%s)))", name, name)
	injectStart := 0
	insertFn := func(line string) (data string, flag int, ok bool) {
		if injectStart == 0 && strings.Contains(line, "container *dig.Container") {
			injectStart = 1
			return
		}

		if injectStart == 1 && strings.Contains(line, "return") {
			injectStart = -1
			data = injectContent
			flag = -1
			ok = true
			return
		}

		return "", 0, false
	}

	err := insertContent(fullname, insertFn)
	if err != nil {
		return err
	}

	fmt.Printf("文件[%s]写入成功\n", fullname)

	return execGoFmt(fullname)
}
