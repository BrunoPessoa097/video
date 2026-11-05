package components

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// estrutura de Cards: Icone(fyne.Resource), Titulo(string), Subtitulo(string), Valor(int), SubtituloValor(string)
type Cards struct {
	Icone          fyne.Resource
	Titulo         string
	Subtitulo      string
	Valor          int
	SubtituloValor string
}

// Criação de cards
func Card(c Cards) *widget.Card {
	// titulo do card
	headerTextTitulo := canvas.NewText(c.Titulo, color.White)
	headerTextTitulo.TextSize = 20
	headerTextTitulo.TextStyle.Bold = true
	headerTextTitulo.Alignment = fyne.TextAlignLeading

	//subtitulo
	headerTextSubtitulo := canvas.NewText(c.Subtitulo, color.Opaque)
	headerTextSubtitulo.TextSize = 14

	//icone
	icone := widget.NewIcon(c.Icone)
	iconeContainer := container.NewGridWrap(
		fyne.NewSize(58, 58),
		icone,
	)

	// container do header titulo
	headerTxt := container.NewVBox(
		headerTextTitulo,
		headerTextSubtitulo,
	)

	// container do header
	header := container.NewHBox(
		iconeContainer,
		headerTxt,
	)

	// valor
	valor := canvas.NewText(strconv.Itoa(c.Valor), color.Opaque)
	valor.Alignment = fyne.TextAlignLeading
	valor.TextSize = 20

	// subtitulo do valor
	valorSub := canvas.NewText(c.SubtituloValor, color.Opaque)
	valorSub.TextSize = 10

	// corpo do container
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

func CardSimples(c Cards) fyne.CanvasObject {
	// titulo
	headerTextTitulo := canvas.NewText(c.Titulo, color.White)
	headerTextTitulo.TextSize = 20
	headerTextTitulo.TextStyle.Bold = true
	headerTextTitulo.Alignment = fyne.TextAlignLeading

	// subtitulo
	headerTextSubtitulo := canvas.NewText(c.Subtitulo, color.Opaque)
	headerTextSubtitulo.TextSize = 14

	// header de titulo
	headerTxt := container.NewVBox(
		headerTextTitulo,
		headerTextSubtitulo,
	)

	// icone
	icone := widget.NewIcon(c.Icone)
	iconeContainer := container.NewGridWrap(
		fyne.NewSize(58, 58),
		icone,
	)

	// botões
	botaoCad := Botao(BotaoCampos{Texto: "Atualizar", Icone: theme.ConfirmIcon(), Acao: func() {}})
	botaoRes := Botao(BotaoCampos{Texto: "Excluir", Icone: theme.CancelIcon(), Acao: func() {}})

	// container dos botões
	corpoBotoes := container.NewHBox(
		botaoCad,
		botaoRes,
	)

	// grid de botoes
	grid := container.NewGridWrap(
		fyne.NewSize(199, 32),
		corpoBotoes,
	)

	// corpo do card
	header := container.NewBorder(
		nil,            // Top
		nil,            // Bottom
		iconeContainer, // Left
		grid,           // Right
		headerTxt,      // Center
	)

	return widget.NewCard(
		"",
		"",
		header,
	)
}
