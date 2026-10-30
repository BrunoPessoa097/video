package sidebar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/BrunoPessoa097/video/internal/ui/components"
)

func SideBar(navegar func(string)) fyne.CanvasObject {
	titulo := components.Label("Gestor Video")

	menu := container.NewVBox(
		titulo,
		widget.NewSeparator(),
		components.Botao("Dashbord", theme.HomeIcon(), func() {
			navegar("Dashbord")
		}),
		components.Botao("Pessoa", theme.BrokenImageIcon(), func() {
			navegar("Pessoa")
		}),
	)

	return container.NewPadded(menu)
}
