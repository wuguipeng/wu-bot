create database wu-bot;

create table stores
(
    id             int auto_increment comment '文件唯一ID'
        primary key,
    user_id        int          not null comment '该用户或机器人的唯一标识',
    file_id        varchar(128) not null comment '文件唯一ID',
    file_unique_id varchar(64)  not null,
    file_size      bigint       not null,
    file_name      varchar(255) null,
    mime_type      varchar(32)  null,
    duration       int          null,
    width          int          null,
    height         int          null,
    title          varchar(255) null,
    performer      varchar(255) null,
    file_type      int          not null comment '1 视频，2音频，3 文档，4 图片',
    local_path     varchar(255) not null comment '本地路径',
    bak_local_path varchar(255) not null comment '备份路径',
    create_time    datetime     not null comment '日期'
);

-- auto-generated definition
create table users
(
    id                          int         not null comment '该用户或机器人的唯一标识'
        primary key,
    is_bot                      tinyint(1)  not null comment '标识该用户是否是机器人，True如果是机器人',
    first_name                  varchar(32) not null comment '用户或者机器人的first_name',
    last_name                   varchar(32) not null comment '可选。用户或者机器人的last_name',
    user_name                   varchar(32) not null comment '可选。用户或者机器人的username',
    language_code               varchar(32) not null comment '可选。用户语言的IETF语言标签',
    can_join_groups             varchar(32) not null comment '可选。返回True如果该机器人可以被邀请加入群组，只在getMe方法返回',
    can_read_all_group_messages varchar(32) not null comment '可选。返回True如果该机器人禁用了隐私模式，只在getMe方法返回',
    supports_inline_queries     varchar(32) not null comment '可选。返回True，如果这个自持内联查询，只在getMe方法返回'
);


