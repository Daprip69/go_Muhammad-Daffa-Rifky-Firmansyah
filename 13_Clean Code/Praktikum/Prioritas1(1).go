//1. Kode berikut ini dituliskan tanpa mengikuti kebiasaan penulisan yang disarankan.
    //- Berapa banyak kekurangan dalam penulisan kode tersebut?
    //- Bagian mana saja terjadi kekurangan tersebut?
   // - Tuliskan alasan dari setiap kekurangan tersebut. Alasan bisa diberikan dalam bentuk komentar pada kode yang disediakan berikut.
    
package main

type user struct {
	id       int
	username int
	password int
}

type userservice struct {
	//penamaan tidak jelas sehingga tidak mudah dipahami, mengapa t?
	t []user
}

func (u userservice) getallusers() []user {
	//penamaan tidak jelas sehingga tidak mudah dipahami, mengapa u dan t?

	return u.t
}

func (u userservice) getuserbyid(id int) user {
	//penamaan tidak jelas sehingga tidak mudah dipahami, mengapa r, u, dan t?
	//r.id menunjukkan ketidakjelasan, apa itu r.id?
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}