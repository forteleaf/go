CREATE USER 'packt'@'%' IDENTIFIED BY 'packt1234';

CREATE DATABASE IF NOT EXISTS packtpub;

CREATE USER 'packt'@'localhost' IDENTIFIED BY 'packt1234';

grant all privileges on packtpub.* to 'packt'@'%';
grant all privileges on packtpub.* to 'packt'@'localhost';


CREATE TABLE mybook(
	num int(5), -- 넘버링 할 수 있게 해볼까앙
	thumbnail VARCHAR(255),
	name VARCHAR(255),
	pdf VARCHAR(200),
	epub VARCHAR(200),
	mobi VARCHAR(200),
	codefiles VARCHAR(200),
);

alter table mybook MODIFY COLUMN num int(5) auto_increment
GRANT RELOAD ON *.* TO 'packt'@'localhost';


-- thubnail 의 길이가 140 이상인 값을 추출
select * from (
	select char_length(thumbnail) a from packtpub.mybook
) a where a > 140;

