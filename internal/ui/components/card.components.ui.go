package components

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Cards struct {
	Icone          fyne.Resource
	Titulo         string
	Subtitulo      string
	Valor          int
	SubtituloValor string
}

// Criação de cards
func Card(c Cards) *widget.Card {
	headerTextTitulo := canvas.NewText(c.Titulo, color.White)
	headerTextTitulo.TextSize = 20
	headerTextTitulo.TextStyle.Bold = true
	headerTextTitulo.Alignment = fyne.TextAlignLeading

	headerTextSubtitulo := canvas.NewText(c.Subtitulo, color.Opaque)
	headerTextSubtitulo.TextSize = 14

	icone := widget.NewIcon(c.Icone)
	iconeContainer := container.NewGridWrap(
		fyne.NewSize(58, 58),
		icone,
	)

	headerTxt := container.NewVBox(
		headerTextTitulo,
		headerTextSubtitulo,
	)

	header := container.NewHBox(
		iconeContainer,
		headerTxt,
	)

	valor := canvas.NewText(strconv.Itoa(c.Valor), color.Opaque)
	valor.Alignment = fyne.TextAlignLeading
	valor.TextSize = 20

	valorSub := canvas.NewText(c.SubtituloValor, color.Opaque)
	valorSub.TextSize = 10

	corpo := container.NewVBox(
		header,
		container.NewPadded(valor),
		container.NewPadded(valorSub),
	)

	return widget.NewCard(
		"",
		"",
		corpo,
	)
}
