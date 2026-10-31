package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type CampoFormulario struct {
	Label string
	Input *widget.Entry
}

func Formulario(campos []CampoFormulario) fyne.CanvasObject {
	formulario := container.NewVBox()

	for _, campo := range campos {
		label := widget.NewLabel(campo.Label)

		linha := container.New(
			layout.NewFormLayout(),
			label,
			campo.Input,
		)

		formulario.Add(linha)
	}

	return formulario
}
