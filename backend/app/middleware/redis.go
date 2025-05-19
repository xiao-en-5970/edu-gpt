package middleware

func GetPrefix(prefix string,key string)(string){
	return prefix+":"+key;
}
