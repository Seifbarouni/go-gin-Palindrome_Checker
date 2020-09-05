package main

import (
	"fmt"
	"net/http"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
)

var str string

func palindrome(str string)bool{
	newStr:=trim(str)
	fmt.Println(newStr)
	length:=len(newStr)

	 bound:=(length/2)+1
	var i int
	for i =0; i < bound; i++ {
		if newStr[i]!=newStr[length-1-i]{
			return false
		}
	}
	return true
}
func trim(s string)string{
	res:=strings.Trim(s,"!@#$%^&*()*?><.,;''/| ")
	for _,char:=range res{
		if !unicode.IsLetter(char) && !unicode.IsDigit(char){
			strings.ReplaceAll(res,string(char),"")
		}
	}
	
	return strings.ToLower(strings.ReplaceAll(res," ",""))
}


func main() {
	r:=gin.Default()
	r.LoadHTMLGlob("template/*.html")
	r.Static("/views","./views")

	r.GET("/",func(c *gin.Context){
		c.Redirect(http.StatusFound,"/Pal")
	})
	r.GET("/Pal",func(c *gin.Context ){
		c.HTML(200,"index.html",str)
		str = ""
	})
	r.POST("/Pal",func(c *gin.Context){
		strToCheck:=c.PostForm("txt")
		if strToCheck ==""{
			c.Redirect(http.StatusFound,"/Pal")
			return
		}
		valid:=palindrome(strToCheck)

		if valid == true{
			str=fmt.Sprintf("'%s' is a Palindrome",strToCheck)
		}else{
			str=fmt.Sprintf("'%s' is not a Palindrome",strToCheck)
		}
		c.Redirect(http.StatusFound,"/Pal")

	})

	r.Run()
}