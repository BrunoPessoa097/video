package navegacao

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/BrunoPessoa097/video/internal/ui/components"
	"github.com/BrunoPessoa097/video/internal/ui/dashboard"
	"github.com/BrunoPessoa097/video/internal/ui/pessoa"
	"github.com/BrunoPessoa097/video/internal/ui/sidebar"
)

// navegação de telas
func Iniciar(w fyne.Window) {
	conteudo := container.NewStack(
		dashboard.Dashboard(),
	)

	navegar := func(tela string) {
		switch tela {
		case "Dashbord":
			conteudo.Objects = []fyne.CanvasObject{
				dashboard.Dashboard(),
			}
		case "Pessoa":
			conteudo.Objects = []fyne.CanvasObject{
				pessoa.Cadastro(),
			}
		}

		conteudo.Refresh()
	}

	tituloLabel := components.Label("DashBoard")

	header := container.NewCenter(
		tituloLabel,
	)

	menu := sidebar.SideBar(navegar)

	corpo := container.NewHSplit(menu, conteudo)
	corpo.SetOffset(0.22)

	layoutPrincipal := container.NewBorder(
		header, nil, nil, nil, corpo,
	)

	w.SetContent(layoutPrincipal)
}
