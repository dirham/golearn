package main
import(
	"fmt"
)
func main(){
	var a = "hello"
	fmt.Println("nilai dari 'a' adalah : ", a)
	fmt.Print("dereferencing atau mengambil nilai pointer dari variabel \n dapat dilakukan denga cara \n menambahkan simbol & pada awal variable contoh : \n '&a' : ", &a)
}
