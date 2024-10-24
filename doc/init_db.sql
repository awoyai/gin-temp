create table tb_user (
    id bigint unsigned auto_increment primary key,
    created_at datetime(3) null,
    updated_at datetime(3) null,
    nike_name varchar(255) null comment '名字',
    username varchar(255) null comment '用户名',
    email varchar(255) null comment '邮箱',
    department varchar(255) null comment '部门',
    enable tinyint(1) null comment '是否启用',
    user_type tinyint(1) null comment '用户类型',
    roles varchar(10000) null comment '权限',
    index idx_username(username)
);

create table tb_menu (
    id bigint auto_increment primary key,
    created_at datetime(3) null,
    updated_at datetime(3) null,
    name varchar(191) null,
    pid bigint unsigned null,
    path varchar(191) null,
    component varchar(191) null,
    redirect varchar(191) null,
    icon varchar(191) null,
    sort bigint null,
    menu_type bigint null,
    `key` varchar(191) null,
    api_list longblob null,
    index idx_pid (pid),
    index idx_name (`name`)
);

create table tb_casbin (
    id bigint unsigned auto_increment primary key,
    ptype varchar(100) null,
    v0 varchar(100) null,
    v1 varchar(100) null,
    v2 varchar(100) null,
    v3 varchar(100) null,
    v4 varchar(100) null,
    v5 varchar(100) null,
    constraint unique_index unique (ptype, v0, v1, v2, v3, v4, v5)
);

create table tb_roles (
    id bigint unsigned auto_increment primary key,
    created_at datetime(3) null,
    updated_at datetime(3) null,
    name varchar(191) null,
    code varchar(191) null,
    auths varchar(10000) null,
    description varchar(191) null,
    index idx_name (`name`)
);