//Go + MYSQL+PDF
package main
import (
   	 "fmt"
    	"database/sql"
  	  _"github.com/go-sql-driver/mysql"
	"github.com/go-pdf/fpdf"
	"strconv"
 )
var Db *sql.DB

type Tag struct {
    	EmpNo  	string   	 `json:"eno"`
    	EmpName 	string   	 `json:"ename"`
	EmpDesig   	string    	`json:"edesig"`
    	EmpSal 		int    	 `json:"esal"`
	EmpAge   	int    	`json:"eage"`
}

func connectionDB(){
   	 var err error
   	 DbDriver := "mysql"
    	DbUser := "root"
    	DbPass := ""
    	DbName := "empdb"
   	 Db, err = sql.Open(DbDriver, DbUser+":"+DbPass+"@/"+DbName)
   	 if err != nil {
       		 panic(err.Error())
   	 }
}

func insertValue(){
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
       	 q = q + ","
	q = q + "'" +edesi + "'"
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

func generatePDF(){
    	fmt.Println("PDF Genetate = demo")
	pdf:=fpdf.New("P","mm","A4","")
	pdf.AddPage()
	pdf.SetFont("Arial","B",7)
	results, err := Db.Query("SELECT * FROM emptab")
    	if err != nil {
        		panic(err.Error()) 
    	}
   	 pdf.Cell(25,10, "Emp No") 
	pdf.Cell(25,10, "Emp Name") 
	pdf.Cell(25,10, "Emp Desig") 
	pdf.Cell(25,10, "Emp Salary")
	pdf.Cell(25,10, "Emp Age")
	pdf.Ln(8)
    	for results.Next() {
        		var tag Tag
        	err = results.Scan(&tag.EmpNo, &tag.EmpName,&tag.EmpDesig,&tag.EmpSal,&tag.EmpAge)
        		if err != nil {
            			panic(err.Error()) 
       		 }
        		pdf.Cell(25,10, tag.EmpNo) 
		pdf.Cell(25,10, tag.EmpName) 
		pdf.Cell(25,10, tag.EmpDesig) 
		pdf.Cell(25,10, strconv.Itoa(tag.EmpSal))
		pdf.Cell(25,10, strconv.Itoa(tag.EmpAge))
fmt.Println(tag.EmpNo,"\t",tag.EmpName,"\t",tag.EmpDesig,"\t",tag.EmpSal,"\t",tag.EmpAge)
		pdf.Ln(8)
        	}
	   	 err = pdf.OutputFileAndClose("ram.pdf")
	   	 if err!=nil{
		    fmt.Println(err.Error())
	    	}
    	}

    func selectValue(){
       	 results, err := Db.Query("SELECT * FROM emptab")
       	 if err != nil {
           	panic(err.Error()) 
        }
       	 for results.Next() {
           		 var tag Tag
            err = results.Scan(&tag.EmpNo, &tag.EmpName,&tag.EmpDesig,&tag.EmpSal,&tag.EmpAge)
            if err != nil {
                panic(err.Error()) 
            }
            fmt.Println(tag.EmpNo,"\t",tag.EmpName,"\t",tag.EmpDesig,"\t",tag.EmpSal,"\t",tag.EmpAge)
            }
    }

    func searchValue(){
        	var empnumber string 
        	var flag int = 0
       	fmt.Println("Enter Employee Number :")
        	fmt.Scan(&empnumber)
        	results, err := Db.Query("SELECT * FROM emptab  where empNo = " + "'" + empnumber + "'")
       	 if err != nil {
            		panic(err.Error()) 
        	}
        	for results.Next() { 
            	var tag Tag
            	err = results.Scan(&tag.EmpNo, &tag.EmpName,&tag.EmpDesig,&tag.EmpSal,&tag.EmpAge)
            if err != nil {
                panic(err.Error()) 
            }
            flag = 1
            fmt.Println(tag.EmpNo,"\t",tag.EmpName,"\t",tag.EmpDesig,"\t",tag.EmpSal,"\t",tag.EmpAge)
            }
            if flag == 0 {
                fmt.Println("Employee Number",empnumber , " Not Found")
            }
    }  
  func main() {
    	fmt.Println("Go MySQL ")
	connectionDB()
   	// insertValue()
   	// selectValue()
   	searchValue()  
   	// generatePDF()
   	 defer Db.Close()
}

