package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Card(titulo, texto, valor string) *widget.Card {

	tituloCard := container.NewHBox(
		widget.NewIcon(theme.AccountIcon()),
		widget.NewLabelWithStyle(
			texto,
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
	)

	return widget.NewCard(
		titulo,
		"",
		container.NewVBox(
			tituloCard,
			widget.NewLabel(valor),
		),
	)
}
