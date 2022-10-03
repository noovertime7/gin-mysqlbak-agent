alter table es_bak_history
    add status int default 0 null comment '备份状态 1成功，0失败';

alter table t_host
    add type int default 1 not null comment '主机类型;1:mysql2:elastic';

alter table es_task drop column host;
alter table es_task drop column password;
alter table es_task drop column username;
alter table es_task add host_id int not null;
alter table es_bak_history add bak_time datetime null;




