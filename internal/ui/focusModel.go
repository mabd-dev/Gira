package ui

type FocusState int

const (
	FocusProjects FocusState = iota
	FocusBoards
	FocusSprints
	FocusActiveSprint
)

func (m model) currentFocusState() FocusState {
	if len(m.focusStack) == 0 {
		return FocusProjects
	}
	return m.focusStack[len(m.focusStack)-1]
}

func (m *model) pushFocus(state FocusState) {
	m.focusStack = append(m.focusStack, state)
}

func (m *model) popFocus() {
	if len(m.focusStack) > 1 {
		m.focusStack = m.focusStack[:len(m.focusStack)-1]
	}
}
