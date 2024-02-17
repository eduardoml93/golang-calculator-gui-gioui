package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type calcState struct {
	num1Edit, num2Edit         widget.Editor
	opAdd, opSub, opMul, opDiv widget.Clickable
	result                     string
}

func main() {
	go func() {
		w := app.NewWindow()
		if err := run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops
	state := &calcState{}
	state.num1Edit.SingleLine, state.num2Edit.SingleLine = true, true

	for {
		e := w.NextEvent()

		switch e := e.(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Spacing: layout.SpaceAround}.Layout(gtx,
						layout.Rigid(material.Editor(th, &state.num1Edit, "Número  1").Layout),
						layout.Rigid(material.Editor(th, &state.num2Edit, "Número  2").Layout),
					)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceAround}.Layout(gtx,
						layout.Rigid(state.buttonLayout(th, &state.opAdd, "+", 0)),
						layout.Rigid(state.buttonLayout(th, &state.opSub, "-", 1)),
						layout.Rigid(state.buttonLayout(th, &state.opMul, "*", 2)),
						layout.Rigid(state.buttonLayout(th, &state.opDiv, "/", 3)),
					)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Body1(th, state.result).Layout(gtx)
				}),
			)

			e.Frame(gtx.Ops)
		case key.Event:
			if e.Name == key.NameEscape {
				return nil
			}
		}
	}
}

func (s *calcState) buttonLayout(th *material.Theme, clickable *widget.Clickable, label string, opIndex int) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		btn := material.Button(th, clickable, label)
		if clickable.Clicked(gtx) {
			s.calculate(opIndex)
		}
		// Ensure that the function always returns a layout.Dimensions
		return btn.Layout(gtx)
	}
}

func (s *calcState) calculate(opIndex int) {
	num1, _ := strconv.ParseFloat(s.num1Edit.Text(), 64) // Ignore err1
	num2, _ := strconv.ParseFloat(s.num2Edit.Text(), 64) // Ignore err2
	switch opIndex {
	case 0:
		s.result = fmt.Sprintf("Resultado: %.2f", num1+num2)
	case 1:
		s.result = fmt.Sprintf("Resultado: %.2f", num1-num2)
	case 2:
		s.result = fmt.Sprintf("Resultado: %.2f", num1*num2)
	case 3:
		if num2 == 0 {
			s.result = "Erro: Divisão por zero."
			return
		}
		s.result = fmt.Sprintf("Resultado: %.2f", num1/num2)
	default:
		s.result = "Erro: Operação inválida."
	}
}
