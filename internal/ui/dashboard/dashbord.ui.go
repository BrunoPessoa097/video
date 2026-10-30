package dashboard

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/BrunoPessoa097/video/internal/ui/components"
)

// dashbord principal
func Dashboard() fyne.CanvasObject {
	tituloLabel := components.Label("Visão Geral")
	subtitulo := widget.NewLabel("Painel de gerenciamento")

	// cabeçalho
	header := container.NewHBox(
		tituloLabel,
		layout.NewSpacer(),
		subtitulo,
	)

	// inicio dos cards
	pessoa := components.Card("Pessoas Cadastradas", "Equipe", "0 Pessoas")
	funcoes := components.Card("Funções Cadastradas", "Equipe", "0 Funções")
	dias := components.Card("Dias Cadastrado", "Planejados", "0 dias")
	cargos := components.Card("Cargos Cadastrado", "Existentes", "0 Cargos")

	// alinhamento dos cards
	cards1 := container.NewGridWithColumns(2, pessoa, funcoes)
	cards2 := container.NewGridWithColumns(2, dias, cargos)

	// container cards
	cards := container.NewVBox(cards1, cards2)

	return container.NewVBox(
		header,
		cards,
	)
}
