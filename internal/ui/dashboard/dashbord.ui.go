package dashboard

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
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
	pessoa := components.Card(components.Cards{Icone: theme.AccountIcon(), Titulo: "Pessoa", Subtitulo: "Equipe de Produção", Valor: 0, SubtituloValor: "Total de Pessoas Registradas"})
	funcoes := components.Card(components.Cards{Icone: theme.MediaVideoIcon(), Titulo: "Função", Subtitulo: "Função Ativas", Valor: 0, SubtituloValor: "Total de Funções Registradas"})
	dias := components.Card(components.Cards{Icone: theme.CalendarIcon(), Titulo: "Datas", Subtitulo: "Datas Escolhidas", Valor: 0, SubtituloValor: "Total de Datas Escolhidas"})
	cargos := components.Card(components.Cards{Icone: theme.ComputerIcon(), Titulo: "Escolhidos", Subtitulo: "Escalas das pessoas", Valor: 0, SubtituloValor: "Total de Escalas Registradas"})

	// alinhamento dos cards
	cards := container.NewGridWithColumns(2, pessoa, funcoes, dias, cargos)

	return container.NewVBox(
		header,
		container.NewPadded(cards),
	)
}
