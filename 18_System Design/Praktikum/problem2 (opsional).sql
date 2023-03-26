--Terdapat sebuah query pada SQL yaitu `SELECT * FROM users;`--

--Dengan tujuan yang sama, tuliskan dalam bentuk perintah:--

--1. Redis--
--2. Neo4j--
--3. Cassandra--

--1. Redis--
GET * FROM users
-- 2. Neo4j--
MATCH (n:User) RETURN n
-- 3. Cassandra--
SELECT * FROM users_table