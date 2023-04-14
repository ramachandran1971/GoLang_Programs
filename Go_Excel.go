package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	fmt.Println("Excel Demo [Text File To Excel]")
	f := excelize.NewFile()
	sheetName := "Sheet1"
	f.SetSheetName("Sheet1", sheetName)
	data := [][]interface{}{
		{"Student Exam Score"},
		{"Type: Cycle Test-I", nil, nil, "Month : March-2023", nil, nil, "Class: I B.Sc (Computer Science)"},
		{"S_NO", "Roll_No", "Name", "Sub1", "Sub2", "Sub3", "Average"},
		{1, 1011, "Siva", 87, 57, 76},
		{2, 1012, "Ravi", 47, 73, 66},
		{3, 1013, "Hari", 77, 71, 57},
		{4, 1014, "Guna", 77, 73, 57},
		{5, 1015, "Latha", 77, 37, 77},
		{6, 1016, "Selva", 77, 57, 77},
		{7, 1017, "Nalan", 77, 44, 87},
		{8, 1018, "Kumar", 77, 79, 77},
	}
	for i, row := range data {
		startCell, err := excelize.JoinCellName("A", i+1)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if err := f.SetSheetRow(sheetName, startCell, &row); err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	formulaType, ref1 := excelize.STCellFormulaTypeShared, "G4:G11"
	if err := f.SetCellFormula(sheetName, "G4", "=ROUND(SUM(D4:F4)/3,0)",
		excelize.FormulaOpts{Ref: &ref1, Type: &formulaType}); err != nil {
		fmt.Println(err.Error())
		return
	}
	if err := f.AddChart(sheetName, "L1", &excelize.Chart{
		Type: "col3DClustered",
		Series: []excelize.ChartSeries{
			{
				Name:       "Sheet1!$C$4", 
				Categories: "Sheet1!$C$4:$C$11",
				Values:     "Sheet1!$C$4:$C$11",	  		
			}},
		Title: excelize.ChartTitle{
			Name:"Result Analysis",
		},
	}); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SaveAs("Four.xlsx"); err != nil {
		fmt.Println(err.Error())
	}

}
