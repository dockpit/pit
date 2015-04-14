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

func green(str string) string {
	return fmt.Sprintf(string(0x01)+"%s"+string(0x00), str)
}

func red(str string) string {
	return fmt.Sprintf(string(0x02)+"%s"+string(0x00), str)
}

func yellow(str string) string {
	return fmt.Sprintf(string(0x03)+"%s"+string(0x00), str)
}

type Box struct {
	client *client.Docker
	model  *model.Model
	store  *Store
	svr    *server.Server

	currLn    int
	input     string
	selection int
	filtered  []*Entry

	currentColor termbox.Attribute
}

func NewBox(m *model.Model, svr *server.Server, client *client.Docker, store *Store) *Box {
	return &Box{
		client: client,
		model:  m,
		store:  store,
		svr:    svr,

		filtered: []*Entry{},
		currLn:   0,

		currentColor: termbox.ColorDefault,
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

	//redraw on syncs
	go func() {
		for range b.store.Syncs {
			b.updateFiltered()
			b.Draw()
		}
	}()

	//handle redraw on user input events
	b.store.Sync()
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

	b.store.SwitchTo(b.filtered[b.selection].Isolation)
}

func (b *Box) Printf(format string, a ...interface{}) {
	b.PrintLn(fmt.Sprintf(format, a...))
}

func (b *Box) PrintLn(str string) {
	w, _ := termbox.Size()

	y := 0
	for _, c := range str {
		switch c {
		case rune(0x01):
			b.currentColor = termbox.ColorGreen
			continue
		case rune(0x02):
			b.currentColor = termbox.ColorRed
			continue
		case rune(0x03):
			b.currentColor = termbox.ColorYellow
			continue
		case rune(0x00):
			b.currentColor = termbox.ColorDefault
			continue
		}

		//jump to next line if message doesnt fit
		if y == w {
			y = 0
			b.currLn++
		}

		termbox.SetCell(y, b.currLn, c, b.currentColor, termbox.ColorDefault)
		y++
	}

	b.currLn++
}

func (b *Box) Draw() {

	//reset carriage and clear termbox
	b.currLn = 0
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		b.PrintLn(errwrap.Wrapf("Draw could not clear termbox: {{err}}", err).Error())
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
	b.Printf(`Web Interface: %s (%s)`, green("OK"), b.svr.URL()) //@todo replace by hostname
	b.Printf(`Current Isolation: %s (%s)`, b.store.State.CurrentIsolationStatus, b.store.State.CurrentIsolationName)

	//@todo render current isolation

	//reander header and or input
	b.PrintLn(``)
	if b.input == "" {
		b.PrintLn(`Isolations:`)
	} else {
		b.Printf(`Isolation '%s':`, b.input)
	}

	//render isolation or trigger message
	if len(b.store.State.Deps) < 1 {
		b.PrintLn(``)
		b.PrintLn("\t" + ` Almost there! When you’re done you’ll be`)
		b.PrintLn("\t" + ` able handle your dependencies with`)
		b.PrintLn("\t" + ` ease, to configure the first:`)
		b.PrintLn(``)
		b.PrintLn("\t" + `  1. Open your favorite web browser`)
		b.Printf("\t"+`  2. Browse to %s`, b.svr.URL())
		b.PrintLn("\t" + `  3.Follow the instructions on screen`)
	} else {
		b.PrintLn(``)
		for i, m := range b.filtered {
			selchar := ` `
			if i == b.selection {
				selchar = `*`
			}

			if m.Distance == -1 {
				b.PrintLn(``)
				b.Printf("\t"+`  [%s] %s`, selchar, "Stop all isolations")
			} else {
				b.Printf("\t"+`  [%s] %s`, selchar, m.Isolation.Name)
			}

		}

		b.PrintLn(``)
		b.PrintLn(``)
		b.PrintLn(`Use Esc or Ctrl-C to exit`)
	}

	//print errors
	b.PrintLn(``)
	if len(b.store.State.Errors) > 0 {
		b.PrintLn(red("Errors:"))
		for _, err := range b.store.State.Errors {
			b.PrintLn("  - " + err.Error())
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

	b.filtered = append(b.filtered, &Entry{
		Distance:  -1,
		Isolation: nil,
	})

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
