-- 用户表
create table t_sso_user_info
(
   id                   int not null auto_increment,
   user_id              varchar(36) auto_increment,
   user_name            varchar(36) not null,
   status               varchar(36) NOT NULL DEFAULT 'active',
   is_delete            varchar(36) DEFAULT NULL,
   sys_insert_datetime  datetime not null default CURRENT_TIMESTAMP,
   sys_upd_datetime     datetime not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   PRIMARY KEY (`user_id`),
   UNIQUE KEY `id` (`id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8;

alter table t_sso_user_info comment '用户信息';

-- 角色表
create table t_sso_role_info
(
    id                   int not null auto_increment,
    role_id              varchar(36) not null,
    role_name            varchar(36) not null,
    is_delete            varchar(36) DEFAULT NULL,
    sys_insert_datetime  datetime not null default CURRENT_TIMESTAMP,
    sys_upd_datetime     datetime not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`role_id`),
    UNIQUE KEY `id` (`id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8;

alter table t_sso_role_info comment '角色信息';

-- 用户角色中间表
create table t_sso_user_role
(
    user_id varchar(36) not null,
    role_id varchar(36) not null,
    sys_insert_datetime  datetime not null default CURRENT_TIMESTAMP,
    sys_upd_datetime     datetime not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`,`role_id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8;

alter table t_sso_user_role comment '用户角色中间表';

-- 菜单表
create table t_sso_menu_info
(
    id                   int not null auto_increment,
    menu_id              varchar(36) not null,
    menu_name            varchar(36) not null,
    uri                  varchar(36) default null,
    pid                  varchar(36) not null,
    is_delete            varchar(36) DEFAULT NULL,
    sys_insert_datetime  datetime not null default CURRENT_TIMESTAMP,
    sys_upd_datetime     datetime not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    primary key (`menu_id`),
    UNIQUE KEY `id` (`id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8;

alter table t_sso_user_role comment '菜单信息';

-- 角色菜单中间表
create table t_sso_role_menu
(
    role_id varchar(36) not null,
    menu_id varchar(36) not null,
    sys_insert_datetime  datetime not null default CURRENT_TIMESTAMP,
    sys_upd_datetime     datetime not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`role_id`,`menu_id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8;

alter table t_sso_user_role comment '角色菜单中间表';