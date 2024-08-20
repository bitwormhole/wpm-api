package dto

import "github.com/bitwormhole/wpm-api/v1/dxo"

// Project ...
type Project struct {
	ID dxo.ProjectID `json:"id"`
	Base

	Name        string              `json:"name"`
	Description string              `json:"description"`
	Type        dxo.ContentTypeName `json:"type"`

	ConfigFileName string `json:"config_file_name"`
	PathInWorktree string `json:"path_in_worktree"`
	ProjectDir     string `json:"project_dir"`
	RegularPath    string `json:"regular_path"`
	Path           string `json:"path"`

	// Location dxo.LocationID    `json:"location"`
	// Class    dxo.LocationClass `json:"location_class"`

	OwnerRepository dxo.LocalRepositoryID `json:"owner_repository"`
	Group           dxo.ProjectGroupName  `json:"group"`
	State           dxo.FileState         `json:"state"`
	IsDir           bool                  `json:"is_dir"`
	IsFile          bool                  `json:"is_file"`
	Tags            dxo.StringList        `json:"tags"`
}
