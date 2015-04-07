package ui

import (
	"fmt"
	"sort"
	"unicode/utf8"

	"github.com/hashicorp/errwrap"
	"github.com/nsf/termbox-go"

	"github.com/dockpit/pit/client"
	"github.com/dockpit/pit/model"
	"github.com/dockpit/pit/server"
)

type Box struct {
	client *client.Docker
	model  *model.Model
	store  *Store
	svr    *server.Server

	currLn    int
	input     string
	selection int
	filtered  []*Entry

	//@todo depcrecated
	current *model.Isolation
	errors  []error
}

func NewBox(m *model.Model, svr *server.Server, client *client.Docker, store *Store) *Box {
	return &Box{
		client: client,
		model:  m,
		store:  store,
		svr:    svr,

		filtered: []*Entry{},
		currLn:   0,

		//@todo remove
		//@todo depcrecated
		errors: []error{},
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

	//@todo rerender on state change

	//handle redraw on user input events
	b.store.Sync()
	b.updateFiltered()
	b.Draw()
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

func (b *Box) ThrowError(err error) {
	b.errors = append(b.errors, err)
	b.Draw()
}

func (b *Box) Printf(format string, a ...interface{}) {
	b.PrintLn(fmt.Sprintf(format, a...))
}

func (b *Box) PrintLn(str string) {
	y := 0
	for _, c := range str {
		termbox.SetCell(y, b.currLn, c, termbox.ColorDefault, termbox.ColorDefault)
		y++
	}

	b.currLn++
}

func (b *Box) Draw() {

	//reset carriage and clear termbox
	b.currLn = 0
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		b.ThrowError(errwrap.Wrapf("Draw could not clear termbox: {{err}}", err))
		return
	}

	//print dockpit logo
	b.PrintLn(`    ___           _          _ _   `)
	b.PrintLn(`   /   \___   ___| | ___ __ (_) |_ `)
	b.PrintLn(`  / /\ / _ \ / __| |/ / '_ \| | __|`)
	b.PrintLn(` / /_// (_) | (__|   <| |_) | | |_ `)
	b.PrintLn(`/___,' \___/ \___|_|\_\ .__/|_|\__|`)
	b.Printf(`        Welcome!      |_| %s    `, b.svr.Version)
	b.PrintLn(``)

	//render status lines
	b.Printf(`Docker Host: %s (%s)`, b.store.State.DockerHostStatus, b.store.State.DockerHostAddress)
	b.Printf(`Web Interface: Online (%s)`, b.svr.URL()) //@todo replace by hostname

	//@todo render current isolation

	//reander header and or input
	b.PrintLn(``)
	if b.input == "" {
		b.PrintLn(`Isolations:`)
	} else {
		b.Printf(`Isolation '%s':`, b.input)
	}

	//render isolation or trigger message
	if len(b.store.State.Isolations) < 1 {
		b.PrintLn(``)
		b.PrintLn("\t" + ` Almost there! When you’re done you’ll be`)
		b.PrintLn("\t" + ` able handle your dependencies with`)
		b.PrintLn("\t" + ` ease, to configure the first:`)
		b.PrintLn(``)
		b.PrintLn("\t" + `  1. Open your favorite web browser`)
		b.Printf("\t"+`  2. Browse to %s`, b.svr.URL())
	} else {
		b.PrintLn(``)
		for i, m := range b.filtered {
			selchar := ` `
			if i == b.selection {
				selchar = `*`
			}

			b.Printf("\t"+`  [%s] %s`, selchar, m.Isolation.Name)
		}
	}

	termbox.Flush()
}

func (b *Box) updateFiltered() {
	b.filtered = []*Entry{}
	for _, iso := range b.store.State.Isolations {
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
