-- +goose Up
-- +goose StatementBegin
create table `blg_seo`
(
    `id`          int unsigned not null auto_increment,
    `type`        char(3)      not null default '' comment '类型：cat=栏目；art=文章',
    `other_id`    varchar(64)  not null default '' comment 'ID',
    `title`       varchar(255) not null default '' comment '标题',
    `keyword`     varchar(255) not null default '' comment '关键词',
    `description` varchar(255) not null default '' comment '描述',
    `created_at`  timestamp    not null default CURRENT_TIMESTAMP,
    `updated_at`  timestamp    not null default CURRENT_TIMESTAMP,
    `deleted_at`  timestamp             default null,
    primary key (`id`),
    key (`other_id`)
) auto_increment = 1000
  collate = utf8mb4_unicode_ci comment ='博客-SEO 表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `blg_seo`;
-- +goose StatementEnd
