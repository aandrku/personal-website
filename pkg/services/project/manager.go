package project

import (
	"github.com/google/uuid"
)

func NewManager(store Store) *Manager {
	return &Manager{
		store: store,
	}
}

type Manager struct {
	store Store
}

func (m *Manager) FindProject(id string) (Project, error) {
	var p Project
	p, err := m.store.FindProject(id)
	if err != nil {
		return p, err
	}
	return p, nil

}

func (m *Manager) Projects() ([]Project, error) {
	p, err := m.store.Projects()
	if err != nil {
		return p, err
	}

	return p, err
}

func (m *Manager) CreateProject(project Project) error {
	project.ID = uuid.New()

	if err := m.store.CreateProject(project); err != nil {
		return err
	}
	return nil
}

func (m *Manager) UpdateProject(project Project) error {
	if err := m.store.UpdateProject(project); err != nil {
		return err
	}
	return nil
}

func (m *Manager) DeleteProject(id string) error {
	if err := m.store.DeleteProject(id); err != nil {
		return err
	}
	return nil
}
