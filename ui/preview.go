package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Preview struct {
	View    *tview.Pages
	Records *tview.Table
	Schemas *tview.Table
}

func NewPreview() *Preview {
	records := tview.NewTable()

	records.
		Select(0, 0).
		SetFixed(1, 0).
		SetSelectable(true, true).
		SetTitle("Records (Ctrl-R)").
		SetTitleAlign(tview.AlignLeft)

	schemas := tview.NewTable()

	schemas.
		Select(0, 0).
		SetFixed(1, 0).
		SetSelectable(true, true).
		SetTitle("Schemas (Ctrl-S)").
		SetTitleAlign(tview.AlignLeft)

	pages := tview.NewPages()

	pages.
		SetTitle("Preview (Records: Ctrl-R | Schemas: Ctrl-S)").
		SetBorder(true).
		SetTitleAlign(tview.AlignLeft)

	pages.AddPage("records", records, true, true)
	pages.AddPage("schemas", schemas, true, false)

	p := &Preview{
		View:    pages,
		Records: records,
		Schemas: schemas,
	}

	return p
}

func (p *Preview) SwitchPage(page string) {
	p.View.SwitchToPage(page)
}

func (p *Preview) SetRecords(records [][]*string) {
	p.Records.Clear().ScrollToBeginning()

	for i, row := range records {
		for j, col := range row {
			var cellValue string
			cellColor := tcell.ColorWhite
			notSelectable := false

			if col != nil {
				cellValue = *col
			}

			// カラム名の色はレコードと異なるものを指定
			if i == 0 {
				cellColor = tcell.ColorNavy
			}

			p.Records.SetCell(
				i, j,
				&tview.TableCell{
					Text:          cellValue,
					Color:         cellColor,
					NotSelectable: notSelectable,
				},
			)
		}
	}
}

func (p *Preview) SetSchemas(schemas [][]*string) {
	p.Schemas.Clear().ScrollToBeginning()

	for i, row := range schemas {
		for j, col := range row {
			var cellValue string
			cellColor := tcell.ColorWhite
			notSelectable := false

			if col != nil {
				cellValue = *col
			}

			// カラム名の色はレコードと異なるものを指定
			if i == 0 {
				cellColor = tcell.ColorNavy
			}

			p.Schemas.SetCell(
				i, j,
				&tview.TableCell{
					Text:          cellValue,
					Color:         cellColor,
					NotSelectable: notSelectable,
				},
			)
		}
	}
}
