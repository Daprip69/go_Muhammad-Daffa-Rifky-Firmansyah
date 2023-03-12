//Kode berikut ini dituliskan tanpa mengikuti kebiasaan penulisan yang disarankan. Ubahlah penulisan kode berikut menjadi lebih terbaca dan lebih rapi!
//setelah dirapikan

package main

type kendaraan struct {
    totalroda       int
    kecepatanperjam int
}

type mobil struct {
    kendaraan
}

func (m *mobil) berjalan() {
    m.tambahkecepatan(10)
}

func (m *mobil) tambahkecepatan(kecepatanbaru int) {
    m.kecepatanperjam += kecepatanbaru
}

func main() {
    mobilcepat := mobil{}
    mobilcepat.berjalan()
    mobilcepat.berjalan()
    mobilcepat.berjalan()

    mobillamban := mobil{}
    mobillamban.berjalan()
}
