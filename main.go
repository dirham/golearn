package main

import (

	"fmt"

	"github.com/dirham/golearn/paket/arithmetic"

)
// implementasi penggunaan libarari atau paket dengan golang.
// untuk menggunakan paket yang terdapat dalam folder cukup arahkan ke path folder.
// semua nama "package" dalam setiap file yang berada dalam folder yang sama harus sama.
func main() {

	sumRes := arithmetic.Sum(5, 6)

	subRes := arithmetic.Subtract(10, 5)
	fmt.Printf("hasil penambahan adalah %d dan pengurangan %d\n", sumRes,subRes)

}
