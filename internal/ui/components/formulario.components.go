package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// CampoFormulario: Label(string), Input(*widget.Entry)
type CampoFormulario struct {
	Label string
	Input *widget.Entry
}

// formulario
func Formulario(campos []CampoFormulario) fyne.CanvasObject {
	// criar o container de formulário
	formulario := container.NewVBox()

	for _, campo := range campos {
		// criando label dos inputs
		label := widget.NewLabel(campo.Label)

		// criando do do input
		linha := container.New(
			layout.NewFormLayout(),
			label,
			campo.Input,
		)

		// formulario criar
		formulario.Add(linha)
	}

	return formulario
}
