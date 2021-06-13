package handle

import (
	"fmt"
	"net/http"
	"short-url/methods"
	"short-url/short"
	"time"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong")

}
func Acortar(w http.ResponseWriter, r *http.Request) {
	//estrae las variables del request
	url := r.FormValue("url")
	var urlShort string
	var nuevaurl short.DataUrl
	shortGen:=false
	_, err := methods.UrlFormat(url, "larga")
	if err != nil {
		fmt.Fprintf(w, " %v ", err)
	} else {
		//aca tenemos que generar la funcion  para el aleatorio
		for shortGen == false {
			urlShort = methods.GetRandonString()

			//busco si la url short esta asignada
			if !short.ValUrlExits(urlShort, "short") {
				shortGen = true
			}
		}
		//vamos a validar si ya a la url solicirtante si esta y solo debe ser actulizado la fecha mas la url
		dataurl := short.Get(url,"large")
		urlShort="http://ml.c/"+urlShort
		if len(dataurl.Acortada) > 0 {
			nuevaurl = short.DataUrl{url, urlShort, dataurl.Creada, time.Now().String()}
		} else {
			nuevaurl = short.DataUrl{url, urlShort, time.Now().String(), ""}
		}
		nuevaurl.Save()

		// al generar el aleatorio debemos de consultar en la base a ver si ya no esta asigando y sino lo mandamos a genear de vuelta

		fmt.Fprintf(w, "%v", urlShort)

	}

}
func Revertir(w http.ResponseWriter, r *http.Request) {
	//estrae las variables del request
	url := r.FormValue("url")
	_, err := methods.UrlFormat(url, "corta")
	if err != nil {
		fmt.Fprintf(w, " %v ", err)
	} else {

		dataurl := short.Get(url, "corta")
		fmt.Fprintf(w, "%v", dataurl.Original)
	}
}
