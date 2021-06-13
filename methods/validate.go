package methods

import (
	"fmt"
	"math/rand"
	"net/url"
)

var urlValLarga =[]string {"mercadolibre.com"}
var urlValCorta =[]string {"ml.c"}


func UrlFormat(urlstr string, format string) (path string,  err error){
	//evaluamos que tenga formato a url
	var respArrval bool
	valor, errFormat := url.Parse(urlstr)

	if errFormat!=nil{
		err=fmt.Errorf("URL Format is invalid")
		return "", err
	}
	//evaluamos que este en el array de url validas
	if format=="larga"{
		respArrval=valValorInArray(urlValLarga,valor.Host)
	}else{
		respArrval=valValorInArray(urlValCorta,valor.Host)
	}

	if !respArrval{
		err=fmt.Errorf("The hostname is invalid")
		return  "", err
	}
	return valor.Path, nil

}
/**
	busco si la url short esta asignada
 */
func valValorInArray (source []string, value string) bool {
	for _, item := range source {
		if item == value {
			return true
		}
	}
	return false
}


func GetRandonString() string{
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)

}