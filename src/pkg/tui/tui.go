package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pablotrianda/til/src/pkg/db"
	"github.com/pablotrianda/til/src/pkg/pager"
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)
)

type item struct {
	title, desc, note string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.desc }

type listKeyMap struct {
	insertItem key.Binding
	deleteItem key.Binding
	editItem   key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		insertItem: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "add item"),
		),
		editItem: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "edit item"),
		),
		deleteItem: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "delete item"),
		),
	}
}

type model struct {
	list     list.Model
	selected string
	keys     *listKeyMap
}

func newModel(items []list.Item) model {

	listKeys := newListKeyMap()

	tilList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	tilList.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.insertItem,
			listKeys.editItem,
			listKeys.deleteItem,
		}
	}

	mod := model{
		list: tilList,
		keys: newListKeyMap(),
	}
	mod.list.Title = "TIL - TODAY I LEARN"

	return mod
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "a":
			return m, tea.Quit

		case "e":
			return m, tea.Quit

		case "d":
			return deleteItem(m, nil)

		case "enter", " ":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.selected = i.note
				pager.Pager(m.selected)

			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func deleteItem(actualModel model, cmd tea.Cmd) (tea.Model, tea.Cmd) {
	actualItem := actualModel.list.SelectedItem().(item)

	db.DeleteTil(actualItem.title)
	updatedList := castTilToItems(db.ListAll())
	actualModel.list.SetItems(updatedList)

	return actualModel, nil
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func castTilToItems(tils []db.Til) []list.Item {
	var items []list.Item

	for _, t := range tils {
		items = append(
			items,
			item{title: t.Title, desc: t.CreatedAt.Format("2006-01-02"), note: t.Note},
		)
	}

	return items
}

func List(title string, searchResult []db.Til) {
	items := castTilToItems(searchResult)

	p := tea.NewProgram(newModel(items), tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
