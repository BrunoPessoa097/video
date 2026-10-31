package pessoa

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/BrunoPessoa097/video/internal/ui/components"
)

type CamposInput struct {
	nome     *widget.Entry
	senha    *widget.Entry
	telefone *widget.Entry
}

func Cadastro() fyne.CanvasObject {
	titulo := components.Label("Cadastro de Pessoas")

	nome := components.InputSimples("Digite seu nome")
	senha := components.InputPass("Digite sua senha")
	telefone := components.InputSimples("Digite seu telefone")

	labels := []components.CampoFormulario{
		{
			Label: "Nome", Input: nome,
		},
		{
			Label: "Senha", Input: senha,
		}, {
			Label: "Telefo", Input: telefone,
		},
	}

	formulario := components.Formulario(labels)

	botaoCadastar := components.Botao("Cadastro", theme.ConfirmIcon(), func() {})
	botaoReset := components.Botao("Reset", theme.CancelIcon(), func() {})

	return container.NewVBox(
		container.NewCenter(titulo),
		formulario,
		container.NewPadded(botaoCadastar),
		container.NewPadded(botaoReset),
	)
}
