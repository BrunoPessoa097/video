package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func Botao(texto string, icone fyne.Resource, acao func()) *widget.Button {
	return widget.NewButtonWithIcon(
		texto,
		icone,
		acao,
	)
}
