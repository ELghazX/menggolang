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
25. Puncak tipe data itu [any](aniani.go) atau interface{}
26. [nil](nil.go) itu null, value yang hanya bisa di tipe data interface ,function, map, slice, pointer, dan channel
27. konversi dari ani ke tipe data statis dengan [Assertion](typeAssertions.go)
28. Secara default pengisian variabel dsb menggunakan metode pass by value, untuk menggunakan pass by reference kita gunakan [pointer](pointer.go) untuk memodifikasi memori yang sama
29. operator [asterisk](pointer.go). * intinya mengubah semua yang pernah menunjuk memori yang sama ke  penunjuk ke memori baru (pindah memori)
buat pointer bisa pakai method [new](new.go)
30. pointer di [function](pointerFunction.go) karena default nya pass by value kita pakai pointer untuk mengubah nilai memori aslinya. untuk [method](pointerMethod.go) sangat direkomendasikan pakai pointer
31. [package](helper/helper.go) digunakan untuk mengorganisir file program, dalam membuatkan function harus diawali dengan [huruf besar](helper/helper.go). untuk menggunakan function di package lain bisa pakai keyword [import](import.go)
32. [Access Modifier](helper/helper.go)di golang ga pakai keyword public private protected dsb. untuk menentukan access modifier nya ialah dari nama funct/variabel yang digunakan.
jika diawali dengan huruf besar maka func/var itu bisa diakses dari [package lain](import.go)
33. ada namanya [init function](database/mysql.go) yang akan pertama dieksekusi ketika package nya [import](init.go). karena di golang var yang dibuat atau dalam hal ini import ini harus digunakan, jika kita tidak ingin mengakses function apa apa dari yang diimport maka golang akan protes. misal kita ada function [init](internal/internal.go). dan kita akses di package lain. agar tidak hilang / golang protes maka kita tambahkan blank identifier (_)  di awal [nama package yang import](init.go)
