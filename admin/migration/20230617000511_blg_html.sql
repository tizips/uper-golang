-- +goose Up
-- +goose StatementBegin
create table `blg_html`
(
    `id`         int unsigned not null auto_increment,
    `type`       char(3)      not null default '' comment '类型：cat=栏目；art=文章',
    `other_id`   varchar(64)  not null default '' comment 'ID',
    `content`    text comment '内容',
    `text`       text comment '文本',
    `created_at` timestamp    not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp    not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp             default null,
    primary key (`id`),
    key (`other_id`)
) auto_increment = 1000
  collate = utf8mb4_unicode_ci comment ='博客-HTML 表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `blg_html`;
-- +goose StatementEnd
