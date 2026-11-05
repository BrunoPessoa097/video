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
	// tela inicial
	conteudo := container.NewStack(
		dashboard.Dashboard(),
	)

	// navegação de telas
	navegar := func(tela string) {
		switch tela {
		case "dashbord":
			conteudo.Objects = []fyne.CanvasObject{
				dashboard.Dashboard(),
			}
		case "pessoa":
			conteudo.Objects = []fyne.CanvasObject{
				pessoa.Cadastro(),
			}
		case "listagem":
			conteudo.Objects = []fyne.CanvasObject{
				pessoa.Listar(),
			}
		}

		conteudo.Refresh()
	}

	// titulo da tela
	tituloLabel := components.LabelTitulo("DashBoard")

	// cabeçalho
	header := container.NewCenter(
		tituloLabel,
	)

	// menu lateral
	menu := sidebar.SideBar(navegar)

	// corpo
	corpo := container.NewHSplit(menu, conteudo)
	corpo.SetOffset(0.22)

	// tela principal
	layoutPrincipal := container.NewBorder(
		header, nil, nil, nil, corpo,
	)

	// telaa crição
	w.SetContent(layoutPrincipal)
}
