// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ir

import "cmd/compile/internal/types"

// A Package holds information about the package being compiled.
type Package struct {
	// Imports, listed in source order.
	// See golang.org/issue/31636.
	// 导入包，按源顺序列出
	Imports []*types.Pkg

	// Init functions, listed in source order.
	// 初始化函数，按源顺序列出
	Inits []*Func

	// Top-level declarations.
	// 顶级声明
	Decls []Node

	// Extern (package global) declarations.
	// 外部（包全局）声明
	Externs []Node

	// Assembly function declarations.
	// 汇编函数声明
	Asms []*Name

	// Cgo directives.
	// Cgo指令
	CgoPragmas [][]string

	// Variables with //go:embed lines.
	// 带有 //go:embed 行的变量
	Embeds []*Name

	// Exported (or re-exported) symbols.
	// 导出（或重新导出）的符号
	Exports []*Name
}
