    //  Token1    Palabra reservada
        //  Token2    Identificador
        //  Token3    Numero entero
        //  Token4    Numero con decimal
        //  Token5    Coma
        //  Token6    Punto y coma
        //  Token7    Operador aritmetico  +  -  /  *
        //  Token8    Operador logico    &&   ||  not
        //  Token9    Operador relacional   >  >=  <   <=    ==   !=
        //  Token10   Asignacion   =
package main

import (
	"io/ioutil"
	"fmt"
)

var imprimir = fmt.Println

func letra(x string) bool{

	if x >= "a" && x <= "z" || x >= "A" && x <= "Z" {
               return true
            }else{
             	return false
             	
             }

}

func  num(x string) bool {

		if x >= "1" && x <= "9" {
                 return true
             }else{
             	return false

             }
}

 
func main() {

PalRes:=[28]string{ "break" ,"default","func","interface"," select","case","defer","go","map",
"struct","chan","else","goto","package"," switch","const"," fallthrough","if ","range",
"type","continue","for","import","return","var","int","float","string"}

  text,err:= ioutil.ReadFile("hola.go")
 	var MaxText int = len(string(text))
 	var nLineas int =1
  var lexema string
  var encontrar bool
  var car string

  

	if err!=nil{
		imprimir("Hubo un err")
	}

	for  i := 0; i < MaxText; i++ {
             
            
               car = string(text)[i:i+1]
               if car==" "{
               lexema=""
               }else if car=="\n" {
                    lexema=""
               			nLineas++
                } else if letra(car) {
                      for {
                          if letra(car)||num(car)||car=="_"{
                          lexema+=car
                          car=string(text)[i:i+1]
                          }else{break}
                          if i<=MaxText{
                             break   
                          }         
                    } 
                          
                  encontrar= false//bandera para indentificar palabra clave o ide
                    for j:=0;j<len(PalRes);j++{
                          if lexema==PalRes[j]{
                            encontrar=true
                          }
                    }
                    if encontrar== true{
                       imprimir("Token1: Palabra Reservada: " + lexema)
                    }else if encontrar  == false{
                    imprimir("Token2: Identificador: " + lexema)
                }
                
                }else if num(car){
                      lexema+=car

                      imprimir("Token3: Numero entero: "+ lexema)

                }else if car ==","{
                	imprimir("Token5: Coma: "+ car)
                }else if car ==";"{
                	imprimir("Token6: Punto y Coma: "+ car)
                }else if car =="+"|| car=="-"|| car== "/" || car=="*"{
                	imprimir("Token7: Operador aritmetico: " + car )
                } else if car =="="{
                	imprimir("Token10: Asignacion: " +car)
                }else if car ==">" || car =="<" || car ==">=" || car=="<=" || car=="=="|| car =="!="{
                  imprimir("Token9: Operador Relacional:  " +car)
                }
                                
    }//FOR
	imprimir("No de Lineas ", nLineas)
}//FIN PROGRAMA
