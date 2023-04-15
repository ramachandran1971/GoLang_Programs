package main
import (
	"fmt"
	"time"
	"log"
	"strconv"
	"database/sql"
    _"github.com/go-sql-driver/mysql"
	"github.com/jung-kurt/gofpdf"
	"github.com/xuri/excelize/v2" 
	"gopkg.in/gomail.v2"
 
)
var Db *sql.DB
var data[][] string

var (
    EmpNo       string      
    EmpName     string      
	EmpDesig    string      
    EmpSal      int        
	EmpAge      int        
)

func connectionDB(){
	var err error
    DbDriver := "mysql"
    DbUser := "root"
    DbPass := ""
    DbName := "empdb"
    Db, err = sql.Open(DbDriver, DbUser+":"+DbPass+"@/"+DbName)
    if err != nil {
        fmt.Println(err.Error())
    }
}
func insertData(){
	var eno,ename,edesi string
    var esal,eage int
    var q string =""
    fmt.Println("Enter Employee Number : ")
    fmt.Scan(&eno)
    fmt.Println("Enter Employee Name : ")
    fmt.Scan(&ename)
    fmt.Println("Enter Employee designation : ")
    fmt.Scan(&edesi)
    fmt.Println("Enter Employee Salary : ")
    fmt.Scan(&esal)
    fmt.Println("Enter Employee Age : ")
    fmt.Scan(&eage)
    q = "INSERT INTO emptab VALUES(" 
    q = q + "'"+ eno + "'"
    q = q + ","
    q = q + "'" + ename +"' "
    q = q +","
    q = q + "'" +edesi + "'"
    q = q + ","
    q = q + strconv.Itoa(esal)
    q = q + ","
    q = q + strconv.Itoa(eage)
    q = q + ")"
    insert, err := Db.Query(q)
    if err != nil {
        panic(err.Error())
    }
    defer insert.Close() 
	
}
func displayData(){
	data= append(data,[]string{"    Emp NO","    Emp Name","    Emp Designation","Emp Salary    ","Emp Age    "})	
    	results, err := Db.Query("SELECT * FROM emptab")
    	if err != nil {
        		panic(err.Error()) 
    	}
		for results.Next() {
        	 	err = results.Scan(&EmpNo, &EmpName,&EmpDesig,&EmpSal,&EmpAge)
        		if err != nil {
            			panic(err.Error()) 
       		 }
			data= append(data,[]string{EmpNo,EmpName,EmpDesig,strconv.Itoa(EmpSal),strconv.Itoa(EmpAge)})	
		}
		
		for _,line := range data{
			for _,str := range line{
				fmt.Printf("%10s" ,str)
			}
			fmt.Println("")
		}
		 
}
func generateExcel() string{
	f := excelize.NewFile() 
	sheetName := "Employee Report" 
	f.SetSheetName("Sheet1", sheetName) 
	for i, row := range data { 
		startCell, err := excelize.JoinCellName("A", i+1) 
		if err != nil { 
		 fmt.Println(err.Error()) 
		 return ""
		} 
	     
		if err := f.SetSheetRow(sheetName, startCell, &row); err != nil { 
		 fmt.Println(err.Error()) 
		 return ""
		} 
	   } 
	   var excelFileName = "EmpData.xlsx"
	   if err := f.SaveAs(excelFileName); err != nil { 
		fmt.Println(err.Error()) 
	   } 
	   return excelFileName
}
func generatePDF() string{
	pdf:= gofpdf.New("L","mm","A4","")
	pdf.AddPage()
	pdf.SetFont("Arial","B",20)
	pdf.Cell(10,40,"PDF Report")
	pdf.Ln(32)
	pdf.SetFont("Arial","",12)
	pdf.SetFillColor(255,255,255)
	align := []string{"L","L","L","R","R"}
    for _,line := range data{
		for i,str := range line{
			pdf.CellFormat(40,10,str,"1",0,align[i],false,0,"")
		}
		pdf.Ln(-1)
	}
	var pdfFileName = "EmpData.pdf"
	err := pdf.OutputFileAndClose(pdfFileName)
	if err != nil {
		log.Fatalf("Error Saving pdf File : %s",err)
		}
		return pdfFileName
}
func sendMail(){
	m := gomail.NewMessage()
	m.SetHeader("From", "trichysramachandran@gmail.com")
	m.SetHeader("To",   "trichysramachandran@gmail.com","ramachandran@udc.ac.in")
	m.SetAddressHeader("Cc", "udccshod@gmail.com", "Ravi")
	m.SetHeader("Subject", "Welcome to GoLanguge !")
	m.SetBody("text/html", "Hello <b>Ramachandran</b> and <i><h1>Computer Science</h1></i>")
	m.Attach(generatePDF())
	m.Attach(generateExcel())
	 
	d := gomail.NewDialer("smtp.gmail.com", 587, "trichysramachandran@gmail.com", "wssbqfydkwcdyauz")
    if err := d.DialAndSend(m); err != nil {
	panic(err)
}
}
func main(){
	fmt.Println("Mysql + PDF + Excel + SendMail",time.Now())
	connectionDB()
	//insertData()
	displayData()
	//generateExcel()
	//generatePDF()
	defer sendMail()

}