-- +goose Up
-- +goose StatementBegin
create table `blg_article`
(
    `id`          varchar(64)      not null,
    `category_id` varchar(64)      not null default '' comment '栏目ID',
    `user_id`     varchar(64)      not null default '' comment '作者ID',
    `name`        varchar(120)     not null default '' comment '名称',
    `picture`     varchar(255)     not null default '' comment '图片',
    `source`      varchar(32)      not null default '' comment '转载标题',
    `url`         varchar(120)     not null default '' comment '转载链接',
    `summary`     varchar(255)     not null default '' comment '简介',
    `is_comment`  tinyint unsigned not null default 0 comment '开启评论：1=是；2=否',
    `is_enable`   tinyint unsigned not null default 0 comment '是否启用：1=是；2=否',
    `created_at`  timestamp        not null default CURRENT_TIMESTAMP,
    `updated_at`  timestamp        not null default CURRENT_TIMESTAMP,
    `deleted_at`  timestamp                 default null,
    primary key (`id`),
    key (`category_id`),
    key (`user_id`)
) collate = utf8mb4_unicode_ci comment ='博客-文章表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `blg_article`;
-- +goose StatementEnd
