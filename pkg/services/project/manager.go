package project

func NewManager(store Store) *Manager {
	return &Manager{
		store: store,
	}
}

type Manager struct {
	store Store
}

func (m *Manager) UpdateTitle(id, title string) error {
	p, err := m.store.FindProject(id)
	if err != nil {
		return err
	}

	p.Title = title

	if err := m.store.UpdateProject(p); err != nil {
		return err
	}
	return nil
}
