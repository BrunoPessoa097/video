package pessoa

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/BrunoPessoa097/video/internal/ui/components"
)

func Cadastro() fyne.CanvasObject {
	// titulo do painel
	titulo := components.Label("Cadastro de Pessoas")

	// criando os inputs
	nome := components.InputSimples("Digite seu nome")
	senha := components.InputPass("Digite sua senha")
	telefone := components.InputSimples("Digite seu telefone")

	// criando os labels
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

	// criando o formulárop
	formulario := components.Formulario(labels)

	// criando os botões
	botaoCadastar := components.Botao(components.BotaoCampos{Texto: "Cadastro", Icone: theme.ConfirmIcon(), Acao: func() {}})
	botaoReset := components.Botao(components.BotaoCampos{Texto: "Reset", Icone: theme.CancelIcon(), Acao: func() {}})

	// criando os formularios
	return container.NewVBox(
		container.NewCenter(titulo),
		formulario,
		container.NewPadded(botaoCadastar),
		container.NewPadded(botaoReset),
	)
}
