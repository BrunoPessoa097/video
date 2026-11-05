package pessoa

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/BrunoPessoa097/video/internal/ui/components"
)

// lista das pessoas
func Listar() fyne.CanvasObject {
	// criar titulo
	titulo := components.Label("Lista de Pessoas")

	// mocks de cards pessoa
	pessoas := []components.Cards{
		{
			Icone: theme.AccountIcon(), Titulo: "Bruno", Subtitulo: "Dias livres",
		},
		{
			Icone: theme.AccountIcon(), Titulo: "Ligia", Subtitulo: "Dias livres",
		},
	}

	// criando modelo de container
	cardPessoa := container.NewVBox()

	// criando os cards
	for _, pessoa := range pessoas {
		linha := components.CardSimples(pessoa)
		cardPessoa.Add(linha)
	}

	// criando a listagem
	return container.NewVBox(
		container.NewCenter(titulo),
		container.NewPadded(cardPessoa),
	)
}
