package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// strutura BotaoCampos: texto(string), icone(fyne.Resource) e acao(func)
type BotaoCampos struct {
	Texto string
	Icone fyne.Resource
	Acao  func()
}

// Botao principal
func Botao(b BotaoCampos) *widget.Button {
	return widget.NewButtonWithIcon(
		b.Texto,
		b.Icone,
		b.Acao,
	)
}
