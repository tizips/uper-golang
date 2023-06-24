-- +goose Up
-- +goose StatementBegin
create table `sys_role_bind_permission`
(
    `id`         int unsigned not null auto_increment,
    `role_id`    int unsigned not null default 0 comment '角色ID',
    `permission` varchar(64)  not null default '' comment '权限',
    `created_at` timestamp    not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp             default null,
    primary key (`id`),
    key (`role_id`)
) auto_increment = 10000
  default collate = utf8mb4_unicode_ci comment '系统-角色绑定权限表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `sys_role_bind_permission`;
-- +goose StatementEnd
