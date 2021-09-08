package utils

import (
	"log"
	"net/http"
)

func HandleErr(err error, str string) {
	if err != nil {
		log.Fatalln(str, ": ", err)
	}
}

func HandleError(w http.ResponseWriter, err error, str string) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(str, ": ", err)
	}
}

// func ToBytes(i interface{}) []byte {
// 	var aBuffer bytes.Buffer
// 	encoder := gob.NewEncoder(&aBuffer)
// 	HandleErr(encoder.Encode(i))
// 	return aBuffer.Bytes()
// }

// func FromBytes(i interface{}, data []byte) {
// 	encoder := gob.NewDecoder(bytes.NewReader(data))
// 	HandleErr(encoder.Decode(i))
// }


// func Hash(i interface{}) string {
// 	s := fmt.Sprintf("%v",i)
// 	hash := sha256.Sum256([]byte(s))
// 	return fmt.Sprintf("%x",hash)
// }