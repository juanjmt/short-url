package short

type DataUrl struct {
	Original	string
	Acortada	string
	Creada		string
	Consultada	string
}
type allDataUrl []DataUrl

var urls = allDataUrl{}

type Repo interface {
	Save(ularge string)(string, error)
	Get(ushort string)(string, error)

}

func (d DataUrl) Save() allDataUrl {
	//antes de guardar primero evalu si esta y si lo encuentra lo elimina
	for i, url:= range urls {
		if url.Original==d.Original{
			urls = append(urls[:i], urls[i + 1:]...)
		}
	}
	urls=append(urls,d)
	return urls
}

func Get(url string, tipo string) DataUrl{
	var res DataUrl
	for _, u := range urls {
		if tipo =="large" {
			if u.Original == url {
				return u
			}
		}else {
			if u.Acortada == url {
				return u
			}
		}
	}
	return res

}
func ValUrlExits(urla string, tipo string) bool{
	for _, u := range urls {
		if tipo=="short"{
			if u.Acortada==urla{
				return true
			}
		}else{
			if u.Original==urla{
				return true
			}
		}
	}
	return false
}



