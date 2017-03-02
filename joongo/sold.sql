-- 클리앙 중고장터
create table clien (
	num int(10) not null primary key,
	category varchar(10),
	title varchar(255),
	writer varchar(50),
	date DATETIME,
	hit int(10)
);

insert into clien values (글번호, 카테고리, 제목, 작성자, 요일, .)
insert into clien values (0,'기타','제목','작성자','2000-10-10,0');

-- 쿨앤조이 중고장터
create table coolenjoy (
	num int(10) not null primary key auto_increment,
	category varchar(10), -- 카테고리
	title varchar(200), -- 제목
	writer varchar(50), -- 작성자
	price varchar(10), -- 천만원단위까지만
	date DATETIME,
	hit int(10)
);

create table joongo (
	sitenum varchar(100) primary key,
	category varchar(10),
	title varchar(255),
	writer varchar(50),
	price varchar(10),
	date DATETIME,
	hit int(5)
);




delete from joongo where sitenum in ( select * from ( select sitenum from joongo order by date desc limit 0,1) dummyname );