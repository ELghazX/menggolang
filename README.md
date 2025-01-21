# Belajar GO lang

## untuk menjalankan file gunakan ```go run "nama file"```di terminal

1. buat data project dengan perintah di commandline ```go mod <nama-project>```  [go.mod](go.mod)
2. stdout [Hello world](helloworld.go)
3. Tidak boleh ada banyak func main dalam satu project, atau tidak bisa dibuild ```go build``` [contoh](sample.go)
4. Cara [Deklarasi tipe data](typeDeclaration.go) dan membuat alias dari tipe data
5. [Tipe data number](number.go)
6. Membuat [variabel](variable.go)
7. [Konversi tipe data](datatypeconversion.go)
8. [Array](arya.go)
9. Tipe data [Slice](sliceAlterArya.go) lebih fleksibel dari array
10. Array asosiatif / dictionary python namanya [map](map.go)
11. percabangan [ifelse](ifelse.go)
12. percabangan switch [case](switch.go)
13. perulangan di golang menggunakan [for](for.go) untuk fungsi for loop, while loop, dan foreach
14. [break dan continue](breakAndContinue.go) untuk memilih/menghentikan iterasi dalam kondisi tertentu.
15. [function](function.go) prosedur, function return value, function return multiple value, function named return value
16. variadic [function](variadicFuncion.go) fungsi yang menerima banyak argumen dengan 1 parameter
17. di golang function bisa diperlakukan sebagai [value](functionValue.go) bisa dimasukkan ke variabel
18. function juga bisa dibuat di [parameter](functionAsParameter.go).
agar parameter tidak terlalu panjang, bisa dialiaskan seperti [ini](functionAsParameter2.go). lebih [lengkap](typeDeclaration.go)
19. agar lebih cepat dengan masukkan sebagai parameter kita bisa buat function [tanpa nama](anonymusFunction.go)
20. fungsi [rekursif](recursiveFunction.go)
21. [closure](closure.go), fitur di golang buat variabel berintraksi dengan variabel yang bukan scope nya
22. eksekusi fungsi setelah fungsi dengan [defer](defer.go) akan selalu dieksekusi bagaimanapun setelah func parent nya selesai dieksekusi error atau tidak, misal dihentikan dengna [panic()](panic.go), makanya bisa dimanfaatkan untuk menampung fungsi [recover](recover.go)
23. Representasi dari data disajikan di [struct](struct.go)
24. [Interface](interface.go) berisi sebuah kontrak untuk struct yang mau ngimplementasikannya
