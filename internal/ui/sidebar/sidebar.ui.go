package sidebar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/BrunoPessoa097/video/internal/ui/components"
)

// menu lateral
func SideBar(navegar func(string)) fyne.CanvasObject {
	// titulos lateral
	titulo := components.Label("Gestor Video")
	subtituloPessoa := components.Label("Pessoa")

	// criando menu lateral
	menu := container.NewVBox(
		titulo,
		widget.NewSeparator(),
		components.Botao(components.BotaoCampos{Texto: "Dashbord", Icone: theme.HomeIcon(), Acao: func() {
			navegar("dashbord")
		}}),
		subtituloPessoa,
		components.Botao(components.BotaoCampos{Texto: "Cadastro", Icone: theme.BrokenImageIcon(), Acao: func() {
			navegar("pessoa")
		}}),
		components.Botao(components.BotaoCampos{Texto: "Listagem", Icone: theme.BrokenImageIcon(), Acao: func() {
			navegar("listagem")
		}}),
	)

	// criando menu
	return container.NewPadded(menu)
}
