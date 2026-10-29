package navegacao

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/BrunoPessoa097/video/internal/ui/dashboard"
)

// navegação de telas
func Iniciar(w fyne.Window) {
	conteudo := container.NewStack(
		dashboard.Dashboard(),
	)

	layout := container.NewBorder(nil, nil, nil, nil, conteudo)

	w.SetContent(layout)
}
