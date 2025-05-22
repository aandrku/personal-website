package project

import "github.com/aandrku/portfolio-v2/pkg/model"

func NewManager() *Manager {
	return &Manager{}
}

type Manager struct {
	projects []model.Project
}

func (m *Manager) Projects() []model.Project {
	return m.projects
}

func (m *Manager) AddProject(project model.Project) {
	m.projects = append(m.projects, project)
}
