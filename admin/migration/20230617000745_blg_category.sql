-- +goose Up
-- +goose StatementBegin
create table `blg_category`
(
    `id`         varchar(64)      not null,
    `parent_id`  varchar(64)               default null comment '父级ID',
    `type`       varchar(10)               default '' comment '类型：parent=父级；page=单页；list=列表',
    `name`       varchar(120)     not null default '' comment '名称',
    `picture`    varchar(255)     not null default '' comment '图片',
    `order`      tinyint unsigned not null default 50 comment '序号',
    `is_comment` tinyint unsigned not null default 0 comment '开启评论：1=是；2=否',
    `is_enable`  tinyint unsigned not null default 0 comment '是否启用：1=是；2=否',
    `created_at` timestamp        not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp        not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp                 default null,
    primary key (`id`),
    key (`parent_id`)
) default collate utf8mb4_unicode_ci comment ='博客-栏目表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `blg_category`;
-- +goose StatementEnd
