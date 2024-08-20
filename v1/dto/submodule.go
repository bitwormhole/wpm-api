package dto

// Submodule 表示一颗 submodule
type Submodule struct {
	Worktree

	Active  bool   `json:"is_active"`
	RawPath string `json:"raw_path"`
	URL     string `json:"url"`
}
