package components

import "fyne.io/fyne/v2/widget"

func InputSimples(placeholder string) *widget.Entry {
	input := widget.NewEntry()
	input.SetPlaceHolder(placeholder)
	return input
}

func InputPass(placeholder string) *widget.Entry {
	input := widget.NewPasswordEntry()
	input.SetPlaceHolder(placeholder)
	return input
}
