create user alta@localhost identified by 'root';

grant all privileges on *.* to  alta@localhost with grant option;

create user alta@'%' identified by 'root';

grant all privileges on *.* to alta@'%' with grant option;