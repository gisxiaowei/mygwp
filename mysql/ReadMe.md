# 创建数据库
``` sql
create database mygwp;
use mygwp;
create table movie (
   id int not null auto_increment,
   name varchar(100) not null,
   publish_time datetime,
   primary key (id)
) character set = utf8;
 insert into movie (name, publish_time) values("中国人", "2017-01-02 13:14:15");
 select * from movie;
```