-- 创建用户、数据库、赋权
SET NAMES UTF8;
create database if not exists echodemo default charset utf8 collate utf8_general_ci;
  grant select,update,delete,insert,alter,create,drop on echodemo.* to "echodemo"@"%" identified by "echodemo";
  grant select,update,delete,insert,alter,create,drop on echodemo.* to "echodemo"@"localhost" identified by "echodemo";
USE echodemo;



create table t_user
(
	id bigint auto_increment
		primary key,
	username varchar(20) default '' not null,
	account varchar(20) not null,
	password varchar(120) not null,
	phone_num varchar(20) not null,
	status int default '1' not null,
	createTime timestamp default CURRENT_TIMESTAMP not null,
	constraint t_user_account_uindex
		unique (account)
)
;
