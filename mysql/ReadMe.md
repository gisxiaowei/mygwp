# 创建数据库
``` sql
create database mygwp;
use mygwp;
create table movie (
   id int not null auto_increment,
   name varchar(100) not null,
   publish_time datetime,
   primary key (id)
);
 insert into movie (name, publish_time) values("asd", "2017-01-02 13:14:15");
```