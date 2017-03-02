
-- coolenjoy 중고장터
create table coolenjoy (
	num int(10) not null primary key,
	category varchar(10),
	title varchar(255), -- 제목 title
	writer varchar(50), -- 작성자 dc:creator
	date DATETIME, -- 날짜 dc:date
	-- hit int(10)
);

insert into clien values (글번호, 카테고리, 제목, 작성자, 요일, .)
insert into clien values (0,'기타','제목','작성자','2000-10-10,0');
