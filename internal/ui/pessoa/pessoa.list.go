package pessoa

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/BrunoPessoa097/video/internal/ui/components"
)

func Listar() fyne.CanvasObject {
	titulo := components.Label("Lista de Pessoas")

	pessoas := []components.Cards{
		{
			Icone: theme.AccountIcon(), Titulo: "Bruno", Subtitulo: "Dias livres",
		},
		{
			Icone: theme.AccountIcon(), Titulo: "Ligia", Subtitulo: "Dias livres",
		},
	}

	cardPessoa := container.NewVBox()

	for _, pessoa := range pessoas {
		linha := components.CardSimples(pessoa)
		cardPessoa.Add(linha)
	}

	return container.NewVBox(
		container.NewCenter(titulo),
		container.NewPadded(cardPessoa),
	)
}
