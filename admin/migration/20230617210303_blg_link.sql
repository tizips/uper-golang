-- +goose Up
-- +goose StatementBegin
create table `blg_link`
(
    `id`         int unsigned     not null auto_increment,
    `name`       varchar(20)      not null default '' comment '名称',
    `url`        varchar(64)      not null default '' comment '链接',
    `logo`       varchar(120)     not null default '' comment 'LOGO',
    `email`      varchar(64)      not null default '' comment '邮箱',
    `position`   varchar(10)      not null default '' comment '位置：all=所有；bottom=底部；other=其他',
    `order`      tinyint unsigned not null default 0 comment '序号（正序）',
    `is_enable`  tinyint unsigned not null default 0 comment '是否启用：1=是；2=否',
    `created_at` timestamp        not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp        not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp                 default NULL,
    primary key (`id`)
) collate = utf8mb4_unicode_ci comment ='博客-友链表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `blg_link`;
-- +goose StatementEnd
