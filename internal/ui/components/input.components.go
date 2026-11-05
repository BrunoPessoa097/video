package components

import "fyne.io/fyne/v2/widget"

// input simples
func InputSimples(placeholder string) *widget.Entry {
	input := widget.NewEntry()
	input.SetPlaceHolder(placeholder)
	return input
}

// input para senha
func InputPass(placeholder string) *widget.Entry {
	input := widget.NewPasswordEntry()
	input.SetPlaceHolder(placeholder)
	return input
}
