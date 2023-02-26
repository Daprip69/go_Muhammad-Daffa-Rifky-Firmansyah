Dalam programming terdapat istilah versionin, yaitu dimana kita melakukan tracking dari perubahan-perubahan dari versi sebelum dan sesudah
Tool dalam versioning ini ada 3, meliputi VCS (Version Control System), RCS  (Revision Control System), dan SRM (Source Code Manager)
Git adalah salah satu tool dari VCS yang paling sering digunakan oleh programmer saat ini
dalam git, setelah melakukan perubahan perlu melakukan commit dengan command git -m "message" atau langsung pada fitur dalam software VCS dan Git Desktop
Kita perlu menghubungkan dahulu akun git dengan software yang digunakan, dan kita perlu mengcopas link repos yang ingin kita lakukan perubahan dengan command git clone -link repo-
setelah itu, kita perlu melakukan pull untuk mengambil data dari server origin, lalu setelah mengubah baru melakukan commit
branch main adalah branch utama, di bawahnya terdapat branch develop dan sub-subnya
command membuat branch adalah git branch "nama branch", lalu switch dengan command git switch "nama branch"
melakukan pull req perlu melalui langkah commit dan dalam branch yang ada, yang nantinya dikonfirmasi pada situs github