package project

type Store interface {
	Projects() []Project
	FindProject(id string) (Project, error)
	CreateProject(project Project) error
	UpdateProject(project Project) error
	DeleteProject(project Project) error
}
