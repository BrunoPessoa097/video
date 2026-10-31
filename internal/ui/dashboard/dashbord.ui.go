package dashboard

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/BrunoPessoa097/video/internal/ui/components"
)

// dashbord principal
func Dashboard() fyne.CanvasObject {
	tituloLabel := components.LabelTitulo("Visão Geral")

	// cabeçalho
	header := container.NewCenter(
		tituloLabel,
	)

	// inicio dos cards
	pessoa := components.Card("Pessoas Cadastradas", "Equipe", "0 Pessoas")
	funcoes := components.Card("Funções Cadastradas", "Equipe", "0 Funções")
	dias := components.Card("Dias Cadastrado", "Planejados", "0 dias")
	cargos := components.Card("Cargos Cadastrado", "Existentes", "0 Cargos")

	// alinhamento dos cards
	cards := container.NewGridWithColumns(2, pessoa, funcoes, dias, cargos)

	return container.NewVBox(
		header,
		container.NewPadded(cards),
	)
}
