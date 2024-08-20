package dto

import "github.com/bitwormhole/wpm-api/v1/dxo"

// Worktree 表示一颗工作树, 或者 submodule
type Worktree struct {
	ID dxo.WorktreeID `json:"id"`
	Base

	URN  dxo.URN `json:"urn"`
	Name string  `json:"name"` // 如果 name=="", 表示默认工作树

	State       dxo.FileState `json:"state"`        // 文件夹状态
	IsSubmodule bool          `json:"is_submodule"` // 是否为子模块
	IsDefault   bool          `json:"is_default"`   // 是否为默认工作树

	DotGitPath  string `json:"dot_git_path"` // .git 文件（或目录）路径
	WorkingDir  string `json:"working_dir"`  // 工作文件夹路径
	RegularPath string `json:"regular_path"` // == DotGitPath
	Path        string `json:"path"`         // == DotGitPath

	// Location dxo.LocationID    `json:"location"`
	// Class    dxo.LocationClass `json:"location_class"`

	OwnerRepository dxo.LocalRepositoryID `json:"owner_repository"`
}
