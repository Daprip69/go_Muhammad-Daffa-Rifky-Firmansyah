database relationship meliputi
one to one = satu entitas berhub dengan satu entitas lain, contoh 1 user hanya memiliki 1 field foto profil
one to many = satu entitas dapat memiliki hub dengan banyak entitas atau field, contoh 1 user dapat memposting banyak postingan
many to many = banyak entitas bisa berhub dengan satu entitas, dan satu entitas tsb dapat berhub dengan banyak entitas, contoh 1 user dapat difollow banyak orang sekaligus dapat memfollow banyak orang

DDL STATEMENT
membuat database = CREATE DATABASE -namanya-
menggunakan database atau skema = USE -nama database-
membuat tabel = CREATE TABLE -nama- .... (isi macam strukturnya, seperti tipe data, primary.foreign/null/dsb)
menghapus tabel = DROP TABLE -nama- ....
merename tabel = RENAME TABLE -nama- ....
primary key = kata kunci, unique, cont dalam coding = column1 -nama kolom- PRIMARY KEY
foreign key = gabungan
modifikasi table skema = ALTER TABLE -name- < duplikat
ADD COLUMN -nama- < tambah kolom, bisa langsung tapi
tipe data dapat berupa num (int) atau huruf (varchar, text) dan date (timestamps)

DML
==insert==
INSERT INTO -nama tabel tujuan- (struktur tabel atau field)
==select==
SELECT * FROM -nama tabel asal- 
bisa + WHERE -field name- = -valuenya-
SETELAH STRUKTUR WHERE BISA JUGA + LIKE "valuenya" atau BETWEEN -value1- AND -value2- atau Like "value" OR -nama tabel atau field- BETWEEN -value1- AND -value2-
==UPDATE, DELETE== SAMA, MIRIP
=ORDER By== mengurutkan
....struktur... ORDER BY -field- DESC/ASC -pilih-
==limit== membatasi data yg dicari
.....struktur.... LIMIT -valuenya-