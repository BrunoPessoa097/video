package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Criar um label baseado no texto enviado
func Label(texto string) *widget.Label {
	return widget.NewLabelWithStyle(
		texto,
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)
}

func LabelTitulo(texto string) *widget.Label {
	labelTitulo := Label(texto)
	labelTitulo.SizeName = theme.SizeNameHeadingText

	return labelTitulo
}
