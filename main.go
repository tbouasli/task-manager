package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func main() {
	ui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	ui.Cursor = true
	ui.SetManagerFunc(layout)

	if err := ui.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		panic(err)
	}

	if err := ui.SetKeybinding("list", 'k', gocui.ModNone, cursorUp); err != nil {
		panic(err)
	}
	
	if err := ui.SetKeybinding("list", 'j', gocui.ModNone, cursorDown); err != nil {
		panic(err)
	}

	if err := ui.SetKeybinding("list", gocui.KeySpace, gocui.ModNone, toggleTask); err != nil {
		panic(err)
	}

	if err := ui.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err)
	}
}

func quit(ui *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}


func layout(ui *gocui.Gui) error {
	maxX, maxY := ui.Size()
	if v, err := ui.SetView("list", 1, 1, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Clear()
		v.Highlight = true
		v.SelFgColor = gocui.ColorGreen
		v.SelBgColor = gocui.ColorBlack

		renderTasks(ui, v)

		if _, err := ui.SetCurrentView("list"); err != nil {
			return err
		}

	}

	return nil
}

func renderTasks(ui *gocui.Gui, v *gocui.View) error {
	v.Clear()
	tasks := List()
	for _, task := range tasks {
		if task.Completed {
			fmt.Fprintf(v, "- [x] %s\n", task.Name)
		} else {
			fmt.Fprintf(v, "- [ ] %s\n", task.Name)
		}
	}

	return nil
}


func cursorUp(ui *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return nil 
			}
		}
	}

	return nil
}

func cursorDown(ui *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func toggleTask(ui *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	tasks := List()
	if cy >= len(tasks) {
		return nil
	}

	task := tasks[cy]
	if task.Completed {
		task.Uncomplete()
	} else {
		task.Complete()
	}

	tasks[cy] = task

	renderTasks(ui, v)

	return nil
}


