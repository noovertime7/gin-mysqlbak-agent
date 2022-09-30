alter table es_bak_history
    add status int default 0 null comment '备份状态 1成功，0失败';

