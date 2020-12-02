package main

import (
	"database/sql"
	"fmt"
)

type MyError struct {
	Code uint32
	Msg string
	Err error
}

func (e *MyError)Error()  string{
	return fmt.Sprintf("error:errCode = %d info = %v",e.Code,e.Msg)
}

func (e *MyError)NewError(code uint32,msg string,err error)  *MyError{
	return &MyError{Code:code,Msg: msg,Err: err}
}



//dao
type MyDb struct {
	db *sql.DB
}

type Info struct {

}
func (d *MyDb)  find(condition string) (*Info,*MyError){
	var info *Info
	err := d.db.QueryRow(sql).Scan(info)
	if err != nil {
		if err == sql.ErrNoRows {
			return _,MyError.NewError(uint32(404),"info not found",err)
		}
		return _,MyError.NewError(uint32(500),"sys error",err)
	}
	return info,nil

}


//service
type MyService struct {
	dao *MyDb
}
type MyData struct {

}
func (ms *MyService)  GetInfo() (*MyData,error){
	info,err := ms.dao.find("")
	if err != nil {
		if err.Code != uint32(404){
			return errors.Wrap(err,err.Msg)
		}else{
			//do not info
			return
		}
	}
	return info,nil
}



//app

func Get()  {
	var s *MyService
	info,err :=s.GetInfo()
	if err != nil {
		fmt.Printf("%v",err)
		return
	}
}




func main() {
	
}
