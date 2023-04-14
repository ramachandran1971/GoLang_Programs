package main
import (
"gopkg.in/gomail.v2"
"github.com/go-pdf/fpdf"
"fmt"
)
func main(){
	m := gomail.NewMessage()
	m.SetHeader("From", "trichysramachandran@gmail.com")
	m.SetHeader("To",   "trichysramachandran@gmail.com","ramachandran@udc.ac.in")
	m.SetAddressHeader("Cc", "udccshod@gmail.com", "Ravi")
	m.SetHeader("Subject", "Welcome to Gollanguge !")
	m.SetBody("text/html", "Hello <b>Ramachandran</b> and <i><h1>Computer Science</h1></i>")
	m.Attach(genratePDF())
	d := gomail.NewDialer("smtp.gmail.com", 587, "trichysramachandran@gmail.com", "drfkonxopmkutjse")

// Send the email to Bob, Cora and Dan.
if err := d.DialAndSend(m); err != nil {
	panic(err)
}
}
func genratePDF() string{
var data = [][]string{
		{"SNo", "RolnNo", "Name",  "M1", "M2", "M3", "M4", "M5", "M6"},
		{"1", "1011", "Siva",  "87", "57", "76", "55", "77", "78"},
		{"2", "1012", "Ravi",  "47", "73", "66", "47", "65", "72"},
		{"3", "1013", "Hari",  "77", "71", "57", "77", "77", "71"},
		{"4", "1014", "Guna",  "77", "73", "57", "77", "78", "77"},
		{"5", "1015", "Latha", "77", "37", "77", "77", "77", "77"},
		{"6", "1016", "Selva", "77", "57", "77", "67", "77", "77"},
		{"7", "1017", "Ramachandran", "77", "44", "87", "77", "77", "77"},
		{"8", "1018", "Kumar", "77", "79", "77", "77", "77", "77"},
} 
	 
	pdf:=fpdf.New("P","mm","A4","")
	pdf.AddPage()
	pdf.SetFont("Arial","B",9)
	 
	for _,row := range data{
		for _,colcell:= range row {
			   
			 pdf.Cell(25,10, colcell)  
		}
		pdf.Ln(8)
	}
	filename:="ram1.pdf"	
	err:= pdf.OutputFileAndClose(filename)
	if err!=nil{
		fmt.Println(err.Error())
	}
	return filename
}
