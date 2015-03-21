package ui

import (
	"fmt"
	"sort"
	"unicode/utf8"

	"github.com/hashicorp/errwrap"
	"github.com/nsf/termbox-go"

	"github.com/dockpit/ppit/client"
	"github.com/dockpit/ppit/model"
	"github.com/dockpit/ppit/server"
)

type Box struct {
	client *client.Docker
	model  *model.Model
	svr    *server.Server

	input     string
	errors    []error
	selection int
	current   *model.Isolation
	isos      []*model.Isolation
	filtered  []*Entry
}

func NewBox(m *model.Model, svr *server.Server, client *client.Docker) *Box {
	return &Box{
		client: client,
		model:  m,
		svr:    svr,

		isos:     []*model.Isolation{},
		filtered: []*Entry{},
		errors:   []error{},
	}
}

func (b *Box) Init() error {
	err := termbox.Init()
	if err != nil {
		return errwrap.Wrapf("Failed to initialize termbox: {{err}}", err)
	}

	termbox.SetInputMode(termbox.InputEsc)
	return nil
}

func (b *Box) Run() error {
	defer termbox.Close()

	//handle redraw on model events
	go func() {
		for range b.model.Events {
			b.updateData()
		}
	}()

	//handle redraw on user input events
	b.updateData()
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc, termbox.KeyEnd, termbox.KeyCtrlC:
				return nil
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				b.removeInput()
			case termbox.KeySpace:
				b.addInput(' ')
			case termbox.KeyArrowDown:
				b.selectionDown()
			case termbox.KeyArrowUp:
				b.selectionUp()
			case termbox.KeyEnter:
				b.Select()
			default:
				if ev.Ch != 0 {
					b.addInput(ev.Ch)
				}
			}

			b.Draw()
		case termbox.EventError:
			return errwrap.Wrapf("Termbox EventError: {{err}}", ev.Err)
		}
	}

	return fmt.Errorf("main event loop ended unexpectedly")
}

func (b *Box) ThrowError(err error) {
	b.errors = append(b.errors, err)
	b.Draw()
}

func (b *Box) PrintLn(y int, msg string) {
	x := 0
	for _, c := range msg {
		termbox.SetCell(x, y, c, termbox.ColorDefault, termbox.ColorDefault)
		x++
	}
}

func (b *Box) Select() {
	if b.selection >= len(b.filtered) {
		return
	}

	b.current = b.filtered[b.selection].Isolation
	err := b.client.Switch(b.current)
	if err != nil {
		b.ThrowError(errwrap.Wrapf("Failed to switch isolation: {{err}}", err))
	}
}

func (b *Box) Draw() {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		b.ThrowError(errwrap.Wrapf("Draw could not clear termbox: {{err}}", err))
		return
	}

	b.PrintLn(0, fmt.Sprintf("Dockpit Web UI Running at %s (press Esc to exit)", b.svr.URL()))
	b.PrintLn(1, fmt.Sprintf("Docker Host: %s (%s)", b.client.Host, b.client.Info.OperatingSystem))

	if b.current != nil {
		b.PrintLn(2, fmt.Sprintf("Current Isolation: %s", b.current.Name))
	} else {
		b.PrintLn(2, fmt.Sprintf("Current Isolation: none"))
	}

	if len(b.errors) > 0 {
		b.PrintLn(3, fmt.Sprintf("Errors: %s", b.errors))
	}

	if len(b.isos) == 0 {
		b.PrintLn(4, "[ This project currently has no isolations, use the web UI to create your first. ]")
	} else {
		b.PrintLn(4, fmt.Sprintf("Available Isolations: %s", b.input))
		for i, m := range b.filtered {
			if i == b.selection {
				b.PrintLn(i+5, fmt.Sprintf("* %s (%d)", m.Isolation.Name, m.Distance))
			} else {
				b.PrintLn(i+5, fmt.Sprintf("- %s (%d)", m.Isolation.Name, m.Distance))
			}

		}
	}

	termbox.Flush()
}

func (b *Box) updateData() {
	var err error
	b.isos, err = b.model.GetAllIsolations()
	if err != nil {
		b.ThrowError(errwrap.Wrapf("Failed to update isolations: {{err}}", err))
		return
	}

	b.updateFiltered()
	b.Draw()
}

func (b *Box) updateFiltered() {
	b.filtered = []*Entry{}
	for _, iso := range b.isos {
		d := Levenshtein(b.input, iso.Name, DefaultCost)
		b.filtered = append(b.filtered, &Entry{d, iso})
	}

	if b.input != "" {
		sort.Sort(ByDistance(b.filtered))
	}

	b.selection = 0
}

func (b *Box) addInput(r rune) {
	var buf [utf8.UTFMax]byte
	n := utf8.EncodeRune(buf[:], r)
	b.input = b.input + string(buf[:n])
	b.updateFiltered()
}

func (b *Box) removeInput() {
	if len(b.input) == 0 {
		return
	}

	b.input = b.input[:len(b.input)-1]
	b.updateFiltered()
}

func (b *Box) selectionUp() {
	if b.selection == 0 {
		return
	}

	b.selection--
}

func (b *Box) selectionDown() {
	if b.selection >= len(b.filtered)-1 {
		return
	}
	b.selection++
}
