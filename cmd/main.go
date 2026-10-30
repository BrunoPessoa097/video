package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/BrunoPessoa097/video/internal/ui/navegacao"
)

func main() {
	a := app.New()
	w := a.NewWindow("Video gestor")
	w.Resize(fyne.NewSize(800, 600))
	w.SetFixedSize(true)

	navegacao.Iniciar(w)

	w.ShowAndRun()
}
