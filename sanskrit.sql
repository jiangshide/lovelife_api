/*
 Navicat MySQL Data Transfer

 Source Server         : localhost
 Source Server Version : 50712
 Source Host           : localhost
 Source Database       : backstage

 Target Server Version : 50712
 File Encoding         : utf-8

 Date: 01/19/2018 23:57:42 PM
*/

-- SET NAMES utf8mb4;
-- SET FOREIGN_KEY_CHECKS = 0;

DROP DATABASE IF EXISTS `sanskrit`;
-- CREATE  DATABASE `sanskrit`;
-- CREATE DATABASE IF NOT EXISTS sanskrit default charset utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS sanskrit;


USE `sanskrit`;

-- ---------------------------
--  Table structure for `zd_uc_user`
-- ----------------------------
-- SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `zd_uc_user`;
-- SET FOREIGN_KEY_CHECKS=1;
CREATE TABLE `zd_uc_user` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` char(60) NOT NULL DEFAULT '' COMMENT '密码',
  `role_ids` varchar(255) NOT NULL DEFAULT '0' COMMENT '角色id字符串，如：2,3,4',
  `salt` char(10) NOT NULL DEFAULT '' COMMENT '密码盐',
  `ip` char(15) NOT NULL DEFAULT '' COMMENT '登录IP',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态:1~审核中,2~审核通过,3移到回忆箱,-1~审核拒绝,-2~禁言,-3~关闭/折叠,-4~被举报',
  `reason` text NOT NULL COMMENT '原由',
  `online`           TINYINT(1)       NOT NULL DEFAULT '3' COMMENT '在线状态:1~在线:-1~离线',
  `source` INT NOT NULL DEFAULT '0' COMMENT '注册方式:0~android,1~ios,2~web,3～小程序',
  `create_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  -- UNIQUE KEY `idx_user_name` (`name`)
-- ) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';
) ENGINE=InnoDB COMMENT='管理员表';


DROP TABLE IF EXISTS `zd_user_weixin`; #微信登录
CREATE TABLE `zd_user_weixin` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `uid`           BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '用户ID',
  `openid`           VARCHAR(100)  NOT NULL DEFAULT '' COMMENT '普通用户的标识，对当前开发者帐号唯一',
  `nick`          VARCHAR(50)      NOT NULL DEFAULT ''  COMMENT '昵称',
  `sex`       Int      NOT NULL DEFAULT '0' COMMENT '性别',
  `language` VARCHAR(10)             NOT NULL   COMMENT '' COMMENT '语言',
  `city` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '城市',
  `province` VARCHAR(80) NOT NULL DEFAULT '' COMMENT '省',
  `country`           VARCHAR(20) NOT NULL DEFAULT '' COMMENT '国家代码',
  `icon` text NOT NULL COMMENT '用户头像',
  `privilege` text NOT NULL COMMENT '用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）',
  `unionid` text NOT NULL COMMENT '微信:用户统一标识。针对一个微信开放平台帐号下的应用，同一用户的 unionid 是唯一的',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_weixin_openid` (`openid`)
  -- KEY `zd_user_invite_id` (`from_id`),    
  -- CONSTRAINT `zd_user_invite_ibfk_1` FOREIGN KEY (`from_id`) REFERENCES `zd_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
)
  ENGINE = InnoDB
  -- DEFAULT CHARSET = utf8mb4
  COMMENT ='微信登录';  

DROP TABLE IF EXISTS `zd_user_pay`; #支付订单
CREATE TABLE `zd_user_pay` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `uid`           BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '用户ID',
  `source` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '渠道来源:0~wechat,1~支付宝,2~qq钱包...',
  `appid`           VARCHAR(50)  NOT NULL DEFAULT '' COMMENT '开放平台审核通过的应用APPID',
  `mchid`          VARCHAR(50)      NOT NULL DEFAULT ''  COMMENT '支付分配的商户号',
  `apikey` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '支付平台密钥',
  `package_value` VARCHAR(30) NOT NULL DEFAULT '' COMMENT 'Sign=WXPay',
  `nonce_str`       VARCHAR(100)             NOT NULL DEFAULT '' COMMENT '随机字符串，不长于32位',
  `sign` VARCHAR(100)             NOT NULL DEFAULT '' COMMENT '签名，详见签名生成算法注意：签名方式一定要与统一下单接口使用的一致',
  `prepay_id` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '返回的支付交易会话ID',
  `date` BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '时间戳',         
  `trade_type` VARCHAR(20)             NOT NULL DEFAULT '' COMMENT '交易类型',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态:1~成功,-1失败',
  `reason` text NOT NULL COMMENT '原由',    
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
  -- UNIQUE KEY `idx_user_pay_prepay_id` (`prepay_id`)
  -- KEY `zd_user_invite_id` (`from_id`),    
  -- CONSTRAINT `zd_user_invite_ibfk_1` FOREIGN KEY (`from_id`) REFERENCES `zd_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
)
  ENGINE = InnoDB
  -- DEFAULT CHARSET = utf8mb4
  COMMENT ='支付订单';  

-- ----------------------------
--  Records of `zd_uc_admin`
-- ----------------------------
-- BEGIN;
-- INSERT INTO `zd_uc_user` VALUES ('1','admin', 'c1875edcd37820e1346f0e3be812dd41', '0', 'e5Ps','[', '1','','0', '0');
-- COMMIT;

-- ----------------------------
--  Table structure for `zd_uc_auth`
-- ----------------------------
DROP TABLE IF EXISTS `zd_uc_auth`;
CREATE TABLE `zd_uc_auth` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `pid` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '上级ID，0为顶级',
  `auth_name` varchar(64) NOT NULL DEFAULT '0' COMMENT '权限名称',
  `auth_url` varchar(255) NOT NULL DEFAULT '0' COMMENT 'URL地址',
  `sort` int(1) unsigned NOT NULL DEFAULT '999' COMMENT '排序，越小越前',
  `icon` varchar(255) NOT NULL,
  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否显示，0-隐藏，1-显示',
  `user_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '操作者ID',
  `create_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态，1-正常，0-删除',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
-- ) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8mb4 COMMENT='权限因子';
) ENGINE=InnoDB COMMENT='权限因子';


-- ----------------------------
--  Records of `zd_uc_auth`
-- ----------------------------
-- BEGIN;
-- INSERT INTO `zd_uc_auth` VALUES ('1', '0', '所有权限', '/', '1', '', '0', '1', '1', '1', '1', '1505620970', '1505620970'), ('2', '1', '权限管理', '/', '999', 'fa-id-card', '1', '1', '0', '1', '1', '0', '1505622360'), ('3', '2', '管理员', '/admin/list', '1', 'fa-user-o', '1', '1', '1', '1', '1', '1505621186', '1505621186'), ('4', '2', '角色管理', '/role/list', '2', 'fa-user-circle-o', '1', '1', '0', '1', '1', '0', '1505621852'), ('5', '3', '新增', '/admin/add', '1', '', '0', '1', '0', '1', '1', '0', '1505621685'), ('6', '3', '修改', '/admin/edit', '2', '', '0', '1', '0', '1', '1', '0', '1505621697'), ('7', '3', '删除', '/admin/ajaxdel', '3', '', '0', '1', '1', '1', '1', '1505621756', '1505621756'), ('8', '4', '新增', '/role/add', '1', '', '1', '1', '0', '1', '1', '0', '1505698716'), ('9', '4', '修改', '/role/edit', '2', '', '0', '1', '1', '1', '1', '1505621912', '1505621912'), ('10', '4', '删除', '/role/ajaxdel', '3', '', '0', '1', '1', '1', '1', '1505621951', '1505621951'), ('11', '2', '权限因子', '/auth/list', '3', 'fa-list', '1', '1', '1', '1', '1', '1505621986', '1505621986'), ('12', '11', '新增', '/auth/add', '1', '', '0', '1', '1', '1', '1', '1505622009', '1505622009'), ('13', '11', '修改', '/auth/edit', '2', '', '0', '1', '1', '1', '1', '1505622047', '1505622047'), ('14', '11', '删除', '/auth/ajaxdel', '3', '', '0', '1', '1', '1', '1', '1505622111', '1505622111'), ('15', '1', '个人中心', 'profile/edit', '1001', 'fa-user-circle-o', '1', '1', '0', '1', '1', '0', '1506001114'),('16', '15', '资料修改', '/user/edit', '1', 'fa-edit', '1', '1', '0', '1', '1', '0', '1506057468');
-- COMMIT;

-- ----------------------------
--  Table structure for `zd_uc_role`
-- ----------------------------
DROP TABLE IF EXISTS `zd_uc_role`;
CREATE TABLE `zd_uc_role` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_name` varchar(32) NOT NULL DEFAULT '0' COMMENT '角色名称',
  `detail` varchar(255) NOT NULL DEFAULT '0' COMMENT '备注',
  `create_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改这ID',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态:1-正常，0-删除',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`)
-- ) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='角色表';
) ENGINE=InnoDB COMMENT='角色表';


-- ----------------------------
--  Records of `zd_uc_role`
-- ----------------------------
-- BEGIN;
-- INSERT INTO `zd_uc_role` VALUES ('2', '系统管理员', '系统管理员', '0', '0', '1', '1506124114', '1506124114');
-- COMMIT;

-- ----------------------------
--  Table structure for `zd_uc_role_auth`
-- ----------------------------
DROP TABLE IF EXISTS `zd_uc_role_auth`;
CREATE TABLE `zd_uc_role_auth` (
  `role_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '角色ID',
  `auth_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '权限ID',
  PRIMARY KEY (`role_id`,`auth_id`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限和角色关系表';
) ENGINE=InnoDB COMMENT='权限和角色关系表';


-- ----------------------------
--  Records of `zd_uc_role_auth`
-- ----------------------------
-- BEGIN;
-- INSERT INTO `zd_uc_role_auth` VALUES ('2', '0'), ('2', '1'), ('2', '15'), ('2', '20'), ('2', '21'), ('2', '22'), ('2', '23'), ('2', '24');
-- COMMIT;

DROP TABLE IF EXISTS `zd_user_profile`; #用户信息
CREATE TABLE `zd_user_profile` (
  `id`           BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT 'ID',
  `union_id` VARCHAR(80)      NOT NULL DEFAULT ''  COMMENT '唯一凭证:针对第三登录',
  `online`           TINYINT(1)       NOT NULL DEFAULT '-1' COMMENT '在线状态:1~在线,2~音乐,3~电影,4~游戏,5~学习,6~工作,-1~离线',
  `online_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '在线状态资源ID',
  `online_name`        text    NOT NULL  COMMENT '在线状态资源名称',
  `icon`          text             NOT NULL   COMMENT '用户头像',
  `level`         TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '级别:0~普通...',
  `score`         INT              NOT NULL DEFAULT '0' COMMENT '积分:',
  `vip`           INT              NOT NULL DEFAULT '0' COMMENT 'vip:会员及级别',
  `certification` BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '实名认证ID', 
  `nick`          VARCHAR(50)      NOT NULL DEFAULT ''  COMMENT '昵称',
  `full_name`          VARCHAR(50)      NOT NULL DEFAULT ''  COMMENT '用户真实名',
  `phone`         VARCHAR(20)      NOT NULL DEFAULT '0' COMMENT '手机号码',
  `intro`         VARCHAR(500)      NOT NULL DEFAULT ''  COMMENT '简介',
  `birthday`      VARCHAR(20)      NOT NULL DEFAULT ''  COMMENT '生日',
  `sex`           TINYINT(1)       NOT NULL DEFAULT '0' COMMENT '性别:0:保密,1:男,2:女',
  `age`           INT   NOT NULL DEFAULT '0' COMMENT '年龄',
  `nation` VARCHAR(30) NOT NULL DEFAULT '' COMMENT '民族',
  `zodiac`      VARCHAR(10) NOT NULL DEFAULT '' COMMENT '生肖',
  `constellation` VARCHAR(10)      NOT NULL DEFAULT ''  COMMENT '星座',
  `hobby` VARCHAR(120) NOT NULL DEFAULT '' COMMENT '爱好',
  `email`         VARCHAR(50)      NOT NULL DEFAULT ''  COMMENT '邮箱',
  `weixin`        VARCHAR(30)      NOT NULL DEFAULT ''  COMMENT '微信',
  `qq`            VARCHAR(20)      NOT NULL DEFAULT ''  COMMENT 'qq',
  `weibo`         VARCHAR(30)      NOT NULL DEFAULT ''  COMMENT '微博',
  `alipay`         VARCHAR(50)      NOT NULL DEFAULT ''  COMMENT '支付宝',
  `latitude`          DOUBLE(10,6)      NOT NULL DEFAULT '0.0'        COMMENT '经度',
  `longitude`         DOUBLE(10,6)      NOT NULL DEFAULT '0.0' COMMENT '纬度',
  `location_type`      VARCHAR(20)      NOT NULL DEFAULT '' COMMENT '定位类型',
  `country`           VARCHAR(20) NOT NULL DEFAULT '' COMMENT '国家代码',
  `province` VARCHAR(80) NOT NULL DEFAULT '' COMMENT '省', 
  `city` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '城市',
  `ad_code`        VARCHAR(10)      NOT NULL DEFAULT ''  COMMENT '区域码',
  `channel_id`           BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '默认公共频道id',
  `im_sign` text             NOT NULL   COMMENT 'im鉴权',
  `praise_notice`         TINYINT(1)       NOT NULL DEFAULT '0' COMMENT '点赞通知:0~打开,-1~关闭',
  `follow_notice`         TINYINT(1)       NOT NULL DEFAULT '0' COMMENT '关注通知:0~打开,-1~关闭',
  `browe_home_notice`         TINYINT(1)       NOT NULL DEFAULT '0' COMMENT '看主页通知:0~打开,-1~打开',
  `browe_channel_notice`         TINYINT(1)       NOT NULL DEFAULT '0' COMMENT '看频道通知:0~打开,-1~打开',
  `browe_blog_notice`         TINYINT(1)       NOT NULL DEFAULT '0' COMMENT '看动态通知:0~打开,-1~打开',
  `create_channel_notice`         TINYINT(1)       NOT NULL DEFAULT '0' COMMENT '创建频道通知:0~打开,-1~关闭',
  `create_blog_notice`         TINYINT(1)       NOT NULL DEFAULT '0' COMMENT '发布动态通知:0~打开,-1~关闭',
  `notice_type`         VARCHAR(10) NOT NULL DEFAULT '' COMMENT '通知类型:可选范围为 -1～7 ，对应 Notification.DEFAULT_ALL = -1 或者 Notification.DEFAULT_SOUND = 1, Notification.DEFAULT_VIBRATE = 2, Notification.DEFAULT_LIGHTS = 4 的任意 “or” 组合。默认按照 -1 处理',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `zd_user_profile_full_name` (`full_name`),    
  CONSTRAINT `zd_user_profile_ibfk_1` FOREIGN KEY (`id`) REFERENCES `zd_uc_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
)
  ENGINE = InnoDB
  -- DEFAULT CHARSET = utf8mb4
  COMMENT ='用户信息表';

DROP TABLE IF EXISTS `zd_user_nation`; #所属民族
CREATE TABLE `zd_user_nation` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '名称',
  `icon`          text             NOT NULL   COMMENT '图标',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
  -- KEY `zd_user_invite_id` (`from_id`),    
  -- CONSTRAINT `zd_user_invite_ibfk_1` FOREIGN KEY (`from_id`) REFERENCES `zd_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
)
  ENGINE = InnoDB
  -- DEFAULT CHARSET = utf8mb4
  COMMENT ='所属民族'; 

DROP TABLE IF EXISTS `zd_user_frequency`; #用户修改频率
CREATE TABLE `zd_user_frequency` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `uid`           BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '用户ID',
  `type` INT NOT NULL DEFAULT '0' COMMENT '类型:1~头像,2~昵称,3~性别,4~生日,5~城市,6~手机号,7~爱好...',
  `number`           INT   NOT NULL DEFAULT '0' COMMENT '次数',
  `frequency` BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '更改频率(秒)',
  `time` BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '创建时间',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
  -- KEY `zd_user_invite_id` (`from_id`),    
  -- CONSTRAINT `zd_user_invite_ibfk_1` FOREIGN KEY (`from_id`) REFERENCES `zd_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
)
  ENGINE = InnoDB
  -- DEFAULT CHARSET = utf8mb4
  COMMENT ='用户修改频率';  

DROP TABLE IF EXISTS `zd_user_certification`; #实名认证
CREATE TABLE `zd_user_certification` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `uid`           BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '用户ID',
  `name`          VARCHAR(50)      NOT NULL DEFAULT ''  COMMENT '真实姓名',
  `id_card`       VARCHAR(20)      NOT NULL DEFAULT '' COMMENT '身份证号码',
  `id_card_pic_front` text             NOT NULL   COMMENT '' COMMENT '手持身份证前照',
  `id_card_pic_behind` text             NOT NULL   COMMENT '' COMMENT '手持身份证后照',
  `url` text             NOT NULL   COMMENT '视频/语音',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态:1~审核中,2~审核通过,-1~审核拒绝',
  `reason` text NOT NULL COMMENT '原由',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_certification_uid` (`uid`)
  -- KEY `zd_user_invite_id` (`from_id`),    
  -- CONSTRAINT `zd_user_invite_ibfk_1` FOREIGN KEY (`from_id`) REFERENCES `zd_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
)
  ENGINE = InnoDB
  -- DEFAULT CHARSET = utf8mb4
  COMMENT ='实名认证';  

DROP TABLE IF EXISTS `zd_user_remarks`; #备注名
CREATE TABLE `zd_user_remarks` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `uid`           BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '用户ID',
  `from_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '来自用户ID',
  `name`          VARCHAR(50)      NOT NULL DEFAULT ''  COMMENT '备注名称',
  `url` text             NOT NULL   COMMENT '图片/视频/语音',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
  -- KEY `zd_user_invite_id` (`from_id`),    
  -- CONSTRAINT `zd_user_invite_ibfk_1` FOREIGN KEY (`from_id`) REFERENCES `zd_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
)
  ENGINE = InnoDB
  -- DEFAULT CHARSET = utf8mb4
  COMMENT ='备注名';  

DROP TABLE IF EXISTS `zd_user_active`; #活跃度
CREATE TABLE `zd_user_active` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `uid`           BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '用户ID',
  `source` INT NOT NULL DEFAULT '0' COMMENT '来源:1.开始页面启动1活跃度(间隔需要6分钟),2.账号登录2活跃度(间隔需要30分钟),3.邀请朋友成功2活跃度,4.发布视频2个活跃度,5.发布音乐1活跃度,6.发布图片1活跃度,7.发布文章1活跃度,8.创建一个频道1活跃度',
  `content_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '内容来源',#内容id
  `num`          INT NOT NULL DEFAULT '0'  COMMENT '活跃值',#1000活跃度=1梵币=1人民币
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态:0~正常,-1~失效,-2~异常',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
  -- KEY `zd_user_invite_id` (`from_id`),    
  -- CONSTRAINT `zd_user_invite_ibfk_1` FOREIGN KEY (`from_id`) REFERENCES `zd_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
)
  ENGINE = InnoDB
  -- DEFAULT CHARSET = utf8mb4
  COMMENT ='活跃度';  


DROP TABLE IF EXISTS `zd_user_course`; #历程
CREATE TABLE `zd_user_course` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `uid`           BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '用户ID',
  `name`          VARCHAR(50)      NOT NULL DEFAULT ''  COMMENT '名称', 
  `des` VARCHAR(500) NOT NULL DEFAULT '' COMMENT '描述',
  `cover` text NOT NULL COMMENT '封面',
  `url` text             NOT NULL   COMMENT '图片/视频/语音',
  `source` INT NOT NULL DEFAULT '0' COMMENT '来源:0~未知,1~注册,2~登录,3～创建频道,4~发布动态',
  `source_id`           BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '来源ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
  -- KEY `zd_user_invite_id` (`from_id`),    
  -- CONSTRAINT `zd_user_invite_ibfk_1` FOREIGN KEY (`from_id`) REFERENCES `zd_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
)
  ENGINE = InnoDB
  -- DEFAULT CHARSET = utf8mb4
  COMMENT ='历程';  


DROP TABLE IF EXISTS `zd_user_invite`; #邀请码
CREATE TABLE `zd_user_invite` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `uid`           BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '用户ID',
  `from_id`           BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '内容来源ID', 
  `times` INT              NOT NULL DEFAULT '0' COMMENT '邀请次数',
  `time`       INT NOT NULL DEFAULT '0' COMMENT '到期时间',
  `status`  TINYINT(1)       NOT NULL DEFAULT '0' COMMENT '状态:0～无效,1～有效,-1:过期',
  `especially` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '有效状态:1~永久,0～取消永久',
  `reason` text NOT NULL COMMENT '原由',
  `code` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '邀请码',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
  -- KEY `zd_user_invite_id` (`from_id`),    
  -- CONSTRAINT `zd_user_invite_ibfk_1` FOREIGN KEY (`from_id`) REFERENCES `zd_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
)
  ENGINE = InnoDB
  -- DEFAULT CHARSET = utf8mb4
  COMMENT ='邀请码';  

DROP TABLE IF EXISTS `zd_user_friend`; #zd:user_firend
CREATE TABLE `zd_user_friend` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `uid`     BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
  `from_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '来自用户ID',
  `status`  TINYINT(1)       NOT NULL DEFAULT '0' COMMENT '状态,0:取消,1~好友,-1:拉黑,-2:删除好友',
  `reason` text NOT NULL COMMENT '原由',
  `url` text NOT NULL COMMENT '快照',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
  -- KEY `zd_user_friend_id` (`from_id`),    
  -- CONSTRAINT `zd_user_friend_ibfk_1` FOREIGN KEY (`from_id`) REFERENCES `zd_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
)
  ENGINE = InnoDB
  -- DEFAULT CHARSET = utf8mb4
  COMMENT ='好友管理';  

-- SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `zd_user_top`;#zd_user_top
CREATE TABLE `zd_user_top`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `from_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内容来源ID:UID',
    `status` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '状态:-1~取消置顶,0~未置顶,1~置顶,2~超级置顶',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_user_top_id` (`from_id`),    
    -- CONSTRAINT `zd_user_top_ibfk_1` FOREIGN KEY (`from_id`) REFERENCES `zd_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='用户置顶';

DROP TABLE IF EXISTS `zd_user_follow`;#zd_user_follow
CREATE TABLE `zd_user_follow`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `from_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内容来源ID:UID',
    `status` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '状态:-1~取消关注,0~推荐关注,1~已关注,2~特别关注',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_user_follow_id` (`from_id`),    
    -- CONSTRAINT `zd_user_follow_ibfk_1` FOREIGN KEY (`from_id`) REFERENCES `zd_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='用户关注';

-- SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `zd_user_recommend`;#zd_user_recommend
CREATE TABLE `zd_user_recommend`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `user_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内容来源ID',
    `status` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '状态:-1~取消推荐,0~未推荐,1~推荐,2～特别推荐',
    `from_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '来自用户ID',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_user_recommend_id` (`from_id`),    
    -- CONSTRAINT `zd_user_recommend_ibfk_1` FOREIGN KEY (`from_id`) REFERENCES `zd_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='用户推荐';

DROP TABLE IF EXISTS `zd_uc_position`; #位置信息
CREATE TABLE `zd_uc_position` (
  `id`            BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `uid`           BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '用户ID',
  `from_id`           BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '来自ID',
  `event`     BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '事件:REG=1,LOGIN=2,CHANNEL=3,BLOG=4',
  `err`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '错误描述',
  `ip` char(15) NOT NULL DEFAULT '' COMMENT '登录IP',
  `latitude`          DOUBLE(10,6)      NOT NULL DEFAULT '0.0'        COMMENT '经度',
  `longitude`         DOUBLE(10,6)      NOT NULL DEFAULT '0.0' COMMENT '纬度',
  `location_type`      VARCHAR(20)      NOT NULL DEFAULT '' COMMENT '定位类型',
  `accuracy`           VARCHAR(20)      NOT NULL DEFAULT '' COMMENT '精度',
  `provider`          VARCHAR(30)      NOT NULL DEFAULT ''  COMMENT '提供者',
  `speed`          VARCHAR(20)      NOT NULL DEFAULT ''  COMMENT '速度',
  `bearing`         VARCHAR(20)      NOT NULL DEFAULT '' COMMENT '角度',
  `satellites`         VARCHAR(20)      NOT NULL DEFAULT ''  COMMENT '星数',
  `country`      VARCHAR(20)      NOT NULL DEFAULT ''  COMMENT '国家',  
  `province`         VARCHAR(30)      NOT NULL DEFAULT ''  COMMENT '省',
  `city`           VARCHAR(30)      NOT NULL DEFAULT '' COMMENT '市',
  `district` VARCHAR(30)      NOT NULL DEFAULT ''  COMMENT '区',
  `city_code`         VARCHAR(10)      NOT NULL DEFAULT ''  COMMENT '城市编码',
  `ad_code`        VARCHAR(10)      NOT NULL DEFAULT ''  COMMENT '区域码',
  `address`            VARCHAR(100)      NOT NULL DEFAULT ''  COMMENT '地址',
  `poi_name`         VARCHAR(80)      NOT NULL DEFAULT ''  COMMENT '兴趣点',
  `network_type`         VARCHAR(10)      NOT NULL DEFAULT ''  COMMENT '网络类型',
  `g_p_s_status`       VARCHAR(20)      NOT NULL DEFAULT '' COMMENT 'GPS状态',
  `g_p_s_satellites` VARCHAR(20)      NOT NULL DEFAULT '' COMMENT 'GPS星数',
  `time_zone` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '时区',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) 
)
  ENGINE = InnoDB
  -- DEFAULT CHARSET = utf8mb4
  COMMENT ='位置信息';

DROP TABLE IF EXISTS `zd_uc_user_device_info`; #设备信息
CREATE TABLE `zd_uc_user_device_info` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `device_id`     BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '设备ID',
  `uid`     BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
  `uuid`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '唯一设备标志管理表',
  `position_id`     BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '位置ID',
  `net_operator`      VARCHAR(50  )      NOT NULL DEFAULT '' COMMENT '获取运营商',
  `net_name`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '获取联网方式',
  `net_speed`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '网络速度',
  `meid`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '获取meid',
  `imei1`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '获取imei1',
  `imei2`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '获取imei2',
  `inumeric`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '获取双卡手机的imei',
  `total_mem`      BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '总共内存',
  `threshold`      BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内存阀值',
  `avail_mem`      BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '可用内存',
  `mac`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT 'mac地址',
  `board`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '主板',
  `brand`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '品牌',
  `device`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '设备参数',
  `display`      VARCHAR(100)      NOT NULL DEFAULT '' COMMENT '显示屏参数',
  `fingerprint`      VARCHAR(100)      NOT NULL DEFAULT '' COMMENT '唯一编号',
  `serial`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '硬件序列号',
  `manufacturer`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '硬件制造商',
  `model`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '机型',
  `hardware`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '硬件名',
  `product`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '手机产品名',
  `type`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT 'Builder类型',
  `host`      VARCHAR(100)      NOT NULL DEFAULT '' COMMENT 'HOST值',
  `user`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT 'User名',
  `time`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '编译时间',
  `os_version`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT 'os版本号',
  `os_name`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT 'os名称',
  `os_arch`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT 'os架构',
  `sdk_version`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '当前sdk版本',
  `app_name`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '应用名称',
  `pkg`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '包名',
  `version_code`      INT              NOT NULL DEFAULT '0' COMMENT '应用code',
  `version_name`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '应用版本号',
  `os`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '操作系统',
  `resolution`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '分辨率800x400',
  `time_zone`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '时区',
  `battery`      INT              NOT NULL DEFAULT '0' COMMENT '电量',
  `elapsed_realtime`      BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '开机时长',
  `language`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '语言',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  -- DEFAULT CHARSET = utf8mb4
  COMMENT ='设备信息管理表';

DROP TABLE IF EXISTS `zd_uc_user_device`; #唯一设备标志管理表
CREATE TABLE `zd_uc_user_device` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `uuid`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '唯一设备标志管理表',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_uc_user_device_uuid` (`uuid`)
)
  ENGINE = InnoDB
  -- DEFAULT CHARSET = utf8mb4
  COMMENT ='唯一设备标志管理表';

-- DROP TABLE IF EXISTS `zd_uc_user_brand`; #银行卡账户
-- CREATE TABLE `zd_uc_user_brand` (
--   `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
--   `user_id`     BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
--   `icon` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '银行logo',
--   `name`        VARCHAR(50)     NOT NULL  DEFAULT '' COMMENT '银行卡名称',
--   `code`        VARCHAR(20)              NOT NULL  DEFAULT '0' COMMENT '银行卡卡号',
--   `create_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '创建者ID',
--   `update_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '修改者ID',
--   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--   `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
--   PRIMARY KEY (`id`),
--   UNIQUE KEY `idx_user_brand_name` (`name`)
-- )
--   ENGINE = InnoDB
--   DEFAULT CHARSET = utf8mb4
--   COMMENT ='银行卡账户管理表';

-- DROP TABLE IF EXISTS `zd_uc_user_address`; #用户邮寄地址
-- CREATE TABLE `zd_uc_user_address` (
--   `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
--   `user_id`     BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
--   `name`        VARCHAR(200)     NOT NULL  DEFAULT ''  COMMENT '邮件地址',
--   `code`        INT              NOT NULL  DEFAULT '0' COMMENT '邮政编码',
--   `create_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '创建者ID',
--   `update_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '修改者ID',
--   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--   `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
--   PRIMARY KEY (`id`),
--   UNIQUE KEY `idx_user_address_name` (`name`)
-- )
--   ENGINE = InnoDB
--   DEFAULT CHARSET = utf8mb4
--   COMMENT ='用户地址管理表';

DROP TABLE IF EXISTS `zd_app`;#App更新管理
CREATE TABLE `zd_app` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `uid`     BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
  `name`        VARCHAR(50)     NOT NULL  DEFAULT '' COMMENT '应用名称',
  `pkg`        VARCHAR(100)              NOT NULL  DEFAULT '' COMMENT '应用包名',
  `platform` VARCHAR(30)       NOT NULL DEFAULT '' COMMENT '平台',
  `channel` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '渠道名称',
  `version` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '版本',
  `code`      INT              NOT NULL DEFAULT '0' COMMENT '版本code',
  `env` TINYINT(1)       NOT NULL DEFAULT '0'  COMMENT '环境:2~production,1~pre,0~qa',
  `build` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '打包状态:1:打包中,2:打包成功,3:打包失败,4,无',
  `cmd` text NOT NULL COMMENT '打包脚本',
  `log` MEDIUMTEXT NOT NULL COMMENT '打包日志详情',
  `reason` text NOT NULL COMMENT '原由',
  `duration`      INT              NOT NULL DEFAULT '0' COMMENT '提示安装间隔时间',
  `times`      INT              NOT NULL DEFAULT '0' COMMENT '提示安装间隔时间次数',
  `des` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '更新内容提示',
  `url` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '下载地址',
  `qr` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '二维码:可扫描下载',
  `size`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '文件尺寸',
  `status` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '更新状态:-2~删除,-1~不可用,0~普通提示更新,1~提示强制更新,2～后台自动下载后更新(非静默更新),3~静默更新',
  `count`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '更新次数',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
  -- UNIQUE KEY `idx_app_version` (`version`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='App管理';

DROP TABLE IF EXISTS `zd_app_count`;#App更新次数
CREATE TABLE `zd_app_count` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `app_id`     BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '应用ID',
  `status` TINYINT(1)       NOT NULL DEFAULT '0' COMMENT '更新状态:0~请求更新,2~下载更新文件,3~安装文件,4~安装成功,-1~请求更新失败,-2~下载更新文件失败,-3~安装文件失败',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_app_id` (`app_id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='App更新次数';

-- DROP TABLE IF EXISTS `zd_app_stop`;#App停服管理
-- CREATE TABLE `zd_app_stop` (
--   `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
--   `name`        VARCHAR(50)     NOT NULL  DEFAULT '' COMMENT '应用名称',
--   `channel` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '渠道名称',
--   `version` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '版本',
--   `url` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '下载地址',
--   `platform` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '平台:0~无限制, 1~android, 2~ios, 3~web',
--   `count`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '更新次数',
--   `create_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '创建者ID',
--   `update_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '修改者ID',
--   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--   `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
--   PRIMARY KEY (`id`)
-- )
--   ENGINE = InnoDB
--   DEFAULT CHARSET = utf8mb4
--   COMMENT ='App停服管理';

DROP TABLE IF EXISTS `zd_app_ads`;#App广告管理
CREATE TABLE `zd_app_ads` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `uid`     BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
  `name`        VARCHAR(50)     NOT NULL  DEFAULT '' COMMENT '广告名称',
  `type` TINYINT(4)       NOT NULL DEFAULT '2' COMMENT '类型:1~开屏,2~banner',
  `platform` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '平台:0~无限制, 1~android, 2~ios, 3~web',
  `price` INT NOT NULL DEFAULT '0' COMMENT '价格',
  `cover` text NOT NULL COMMENT '封面',
  `url` VARCHAR(500) NOT NULL DEFAULT '' COMMENT '内容地址',
  `status` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '状态:1:正常,2:下架,3:异常',
  `count`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '展示次数',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_app_ads` (`url`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='App广告管理';


-- DROP TABLE IF EXISTS `zd_app_web_banner`;#Web~Banner
-- CREATE TABLE `zd_app_web_banner`(
--   `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
--   `name`        VARCHAR(50)     NOT NULL  DEFAULT '' COMMENT '应用名称',
--   `platform` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '平台:0~无限制, 1~android, 2~ios, 3~web',
--   `channel` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '渠道名称',
--   `version` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '版本',
--   `content` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '更新内容提示',
--   `url` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '下载地址',
--   `status` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '更新状态:0~普通提示更新,1~提示强制更新,2～后台自动下载后更新(非静默更新),3~静默更新',
--   `count`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '更新次数',
--   `create_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '创建者ID',
--   `update_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '修改者ID',
--   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--   `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
--   PRIMARY KEY (`id`)
-- )
--   ENGINE = InnoDB
--   DEFAULT CHARSET = utf8mb4
--   COMMENT ='App更新管理';


  -- SET FOREIGN_KEY_CHECKS=0;
  DROP TABLE IF EXISTS `zd_channel`;#zd~channel
  -- SET FOREIGN_KEY_CHECKS=1;
  CREATE TABLE `zd_channel`(
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
  `channel_type_id`   INT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '频道类型ID',
  `channel_nature_id`   INT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '频道所属ID',
  `cover` text NOT NULL COMMENT '封面',
  `blog_cover` text NOT NULL COMMENT '封面:最后一个动态的封面',
  `name`        VARCHAR(50)     NOT NULL  DEFAULT '' COMMENT '频道名称',
  `des` VARCHAR(200)     NOT NULL  DEFAULT '' COMMENT '频道描述',
  `status` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '状态:0~未审核,1~审核中,2~审核通过,-1~移到回忆箱,-2~审核拒绝,-3～禁言，-4~关闭/折叠,-5~被投诉',
  `official` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '官方推荐:-1~取消推荐,0~未推荐,1~推荐,2~特别推荐',
  `reason`        VARCHAR(200)     NOT NULL  DEFAULT '' COMMENT '审核拒绝/禁言/关闭原因',
  `latitude`          DOUBLE(10,6)      NOT NULL DEFAULT '0.0'        COMMENT '经度',
  `longitude`         DOUBLE(10,6)      NOT NULL DEFAULT '0.0' COMMENT '纬度',
  `location_type`      VARCHAR(20)      NOT NULL DEFAULT '' COMMENT '定位类型',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_channel_name` (`name`)
)
  ENGINE = InnoDB 
  -- DEFAULT CHARSET = utf8mb4
  COMMENT ='频道管理';

 DROP TABLE IF EXISTS `zd_channel_nature`;#zd_channel_nature
    CREATE TABLE `zd_channel_nature`(
    `id`          INT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `name`        VARCHAR(20)     NOT NULL  DEFAULT '' COMMENT '频道所属名称',
    `des`        VARCHAR(200)     NOT NULL  DEFAULT '' COMMENT '频道所属描述',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_channel_nature_name` (`name`)
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='频道所属';

 DROP TABLE IF EXISTS `zd_channel_nature_user`;#zd_channel_nature_user
    CREATE TABLE `zd_channel_nature_user`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `channel_nature_id`   INT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '频道所属ID',
    `num`        INT              NOT NULL DEFAULT '0' COMMENT '频道所属可创建条数',
    `time`       INT NOT NULL DEFAULT '0' COMMENT '频道所属可创建的时间周期',
    `note`        VARCHAR(20)     NOT NULL  DEFAULT '' COMMENT '标注',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_channel_nature_user_id` (`channel_nature_id`),    
    -- CONSTRAINT `zd_channel_nature_user_ibfk_1` FOREIGN KEY (`channel_nature_id`) REFERENCES `zd_channel_nature` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='用户频道所属';

 DROP TABLE IF EXISTS `zd_channel_type`;#zd_channel_type
 CREATE TABLE `zd_channel_type`(
    `id`          INT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `name`        VARCHAR(20)     NOT NULL  DEFAULT '' COMMENT '频道类型名称',
    `des`        VARCHAR(200)     NOT NULL  DEFAULT '' COMMENT '频道类型描述',
    `status` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '状态:1~显示,-1~隐藏',
    `sequence` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '显示顺序',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_channel_type_name` (`name`)
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='频道类型';

DROP TABLE IF EXISTS `zd_channel_type_user`;#zd_channel_type_user
CREATE TABLE `zd_channel_type_user`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `channel_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '频道ID',
    `channel_type_id`   INT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '频道类型ID',
    `selected`        INT              NOT NULL DEFAULT '0' COMMENT '选中状态:1~选中,0~未选中',
    `num`        INT              NOT NULL DEFAULT '0' COMMENT '排序顺序',
    `note`        VARCHAR(20)     NOT NULL  DEFAULT '' COMMENT '标注',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_channel_type_user_id` (`channel_id`),    
    -- CONSTRAINT `zd_channel_type_user_ibfk_1` FOREIGN KEY (`channel_id`) REFERENCES `zd_channel_type` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='用户频道类型';  
   
-- SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `zd_channel_top`;#zd_channel_top
CREATE TABLE `zd_channel_top`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `channel_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '频道ID',
    `status` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '状态:-1~取消置顶,0~未置顶,1~置顶,2~超级置顶',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_channel_top_id` (`channel_id`),    
    -- CONSTRAINT `zd_channel_top_ibfk_1` FOREIGN KEY (`channel_id`) REFERENCES `zd_channel` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='频道置顶';

-- SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `zd_channel_recommend`;#zd_channel_recommend
CREATE TABLE `zd_channel_recommend`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `channel_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '动态ID',
    `status` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '状态:-1～取消推荐,0~未推荐,1~推荐,2~特别推荐',
    `from_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '来自用户ID',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_channel_recommend_id` (`channel_id`),    
    -- CONSTRAINT `zd_channel_recommend_ibfk_1` FOREIGN KEY (`channel_id`) REFERENCES `zd_channel` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='频道推荐';

  DROP TABLE IF EXISTS `zd_format`;#zd~format
  CREATE TABLE `zd_format`(
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
  `name`        VARCHAR(50)     NOT NULL  DEFAULT '' COMMENT '内容类型名称',
  `des` VARCHAR(200)     NOT NULL  DEFAULT '' COMMENT '内容类型简介',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_format_name` (`name`)
)
  ENGINE = InnoDB
  -- DEFAULT CHARSET = utf8mb4
  COMMENT ='内容格式';

-- SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `zd_blog`;#zd~blog
-- SET FOREIGN_KEY_CHECKS=1;
CREATE TABLE `zd_blog`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '创建者ID',
    `channel_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '频道ID',
    `position_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '位置ID',
    `content` text NOT NULL COMMENT '内容',
    `title`        VARCHAR(50)     NOT NULL  DEFAULT '' COMMENT '动态名称',
    `des` VARCHAR(200)     NOT NULL  DEFAULT '' COMMENT '动态描述',
    `latitude`          DOUBLE(10,6)      NOT NULL DEFAULT '0.0'        COMMENT '经度',
    `longitude`         DOUBLE(10,6)      NOT NULL DEFAULT '0.0' COMMENT '纬度',
    `location_type`      VARCHAR(20)      NOT NULL DEFAULT '' COMMENT '定位类型',
    `country`      VARCHAR(20)      NOT NULL DEFAULT ''  COMMENT '国家',
    `city` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '所在城市',
    `position` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '位置',
    `address` text NOT NULL COMMENT '详细位置',
    `city_code`         VARCHAR(10)      NOT NULL DEFAULT ''  COMMENT '城市编码',
    `ad_code`        VARCHAR(10)      NOT NULL DEFAULT ''  COMMENT '区域码',
    `time_zone` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '时区',
    `tag` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '标签',
    `status` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '状态:0~未审核,1~审核中,2~审核通过,-1~移到回忆箱,-2~审核拒绝,-3～禁言，-4~关闭/折叠,-5~被投诉',
    `reason`        VARCHAR(200)     NOT NULL  DEFAULT '' COMMENT '原由',
    `official` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '官方推荐:-1~取消推荐,0~未推荐,1~推荐,2~特别推荐',
    `cover` text NOT NULL COMMENT '封面',
    `dynamic_cover` text NOT NULL COMMENT '动态封面',
    `url` text NOT NULL COMMENT '内容Url', 
    `name`        text    NOT NULL  COMMENT '文件名称',
    `sufix` text NOT NULL COMMENT '文件名后缀',
    `format`      TINYINT(4)              NOT NULL DEFAULT '0' COMMENT '内容格式:0:图片,1:音频,2:视频,3:文档,4:web,5:VR',
    `duration`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内容时长',
    `width`      INT              NOT NULL DEFAULT '0' COMMENT '内容宽',
    `height`      INT              NOT NULL DEFAULT '0' COMMENT '内容高',
    `size`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内容尺寸',
    `rotate`      INT              NOT NULL DEFAULT '0' COMMENT '角度旋转',
    `bitrate`      INT              NOT NULL DEFAULT '0' COMMENT '采样率', 
    `sample_rate`      INT              NOT NULL DEFAULT '0' COMMENT '频率', 
    `level`      INT              NOT NULL DEFAULT '0' COMMENT '质量:0~标准,-1~差,1~好', 
    `mode`      INT              NOT NULL DEFAULT '0' COMMENT '模式', 
    `wave` text NOT NULL COMMENT '频谱',
    `lrc_zh` text NOT NULL COMMENT '歌词~中文',
    `lrc_es` text NOT NULL COMMENT '歌词~英文',
    `source`      INT              NOT NULL DEFAULT '0' COMMENT '创作类型:0~原创,1~其它', 
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='动态';

-- SET FOREIGN_KEY_CHECKS=0;  
DROP TABLE IF EXISTS `zd_blog_style`;#zd_blog_style
CREATE TABLE `zd_blog_style`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `blog_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '动态ID',
    `position` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '位置:左上-51,上中-49,上右-53,右中-21,右下-85,下中-81,下左-83,左中-19',
    `size` INT NOT NULL DEFAULT '0' COMMENT '文字大小',
    `txt_color` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '文字颜色',
    `bg_color` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '背景颜色',
    `scroll` TINYINT(1)       NOT NULL DEFAULT '0' COMMENT '是否滚动:1~滚动,0~不滚动',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_blog_praise_id` (`blog_id`),    
    -- CONSTRAINT `zd_blog_praise_ibfk_1` FOREIGN KEY (`blog_id`) REFERENCES `zd_blog` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='动态样式';

-- SET FOREIGN_KEY_CHECKS=0;  
DROP TABLE IF EXISTS `zd_blog_praise`;#zd_blog_praise
CREATE TABLE `zd_blog_praise`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `blog_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '动态ID',
    `status` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '状态:-1~取消,1~点赞,2~超级赞',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_blog_praise_id` (`blog_id`),    
    -- CONSTRAINT `zd_blog_praise_ibfk_1` FOREIGN KEY (`blog_id`) REFERENCES `zd_blog` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='动态点赞';

-- SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `zd_blog_recommend`;#zd_blog_recommend
CREATE TABLE `zd_blog_recommend`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `blog_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '动态ID',
    `status` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '状态:-1～取消推荐,0~未被推荐,~1推荐,2~特别推荐',
    `from_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '来自用户ID',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_blog_recommend_id` (`blog_id`),    
    -- CONSTRAINT `zd_blog_recommend_ibfk_1` FOREIGN KEY (`blog_id`) REFERENCES `zd_blog` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='动态推荐';

-- SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `zd_blog_top`;#zd_blog_top
CREATE TABLE `zd_blog_top`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `blog_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '动态ID',
    `status` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '状态:-1~取消置顶,0~未置顶,1~置顶,2~超级置顶',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_top_id` (`blog_id`),    
    -- CONSTRAINT `zd_blog_top_ibfk_1` FOREIGN KEY (`blog_id`) REFERENCES `zd_blog` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='动态置顶';

DROP TABLE IF EXISTS `zd_blog_view`;#zd_blog_view
    CREATE TABLE `zd_blog_view`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `blog_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内容来源ID',
    `num`      INT              NOT NULL DEFAULT '0' COMMENT '内容详情位置:来自哪张图片',
    `nums`      INT              NOT NULL DEFAULT '0' COMMENT '浏览次数',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_blog_view_id` (`blog_id`),    
    -- CONSTRAINT `zd_blog_view_ibfk_1` FOREIGN KEY (`blog_id`) REFERENCES `zd_blog` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='动态阅览';

-- SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `zd_blog_share`;#zd_blog_share
CREATE TABLE `zd_blog_share`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `blog_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内容来源ID',
    `status` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '状态:1~成功,0～失败',
    `from_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '来自用户ID',
    `station` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '类型:0~站内,1~微信好友,2～微信朋友圈,3~微信收藏',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_share_id` (`content_id`),    
    -- CONSTRAINT `zd_share_ibfk_1` FOREIGN KEY (`content_id`) REFERENCES `zd_share` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='动态分享';    

-- SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `zd_blog_collection`;#zd_blog_collection
CREATE TABLE `zd_blog_collection`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `blog_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内容来源ID',
    `status` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '状态:-1~取消收藏,0～未收藏,1~收藏,2~珍藏',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',    
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='动态珍藏';

-- SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `zd_praise`;#zd_praise
CREATE TABLE `zd_praise`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `content_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内容来源ID',
    `status` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '状态:-1~取消,1~点赞,2~超级赞',
    `source`      INT              NOT NULL DEFAULT '0' COMMENT '内容类型:1~用户,2~频道,3~动态,4~评论',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_praise_id` (`content_id`),    
    -- CONSTRAINT `zd_praise_ibfk_1` FOREIGN KEY (`content_id`) REFERENCES `zd_praise` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='点赞|投票(用户|频道)';

-- SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `zd_comment`;#zd_comment
-- SET FOREIGN_KEY_CHECKS=1;
    CREATE TABLE `zd_comment`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `content_id` BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '动态ID',
    `content` text NOT NULL COMMENT '内容',
    `at` text NOT NULL COMMENT '@内容:content:id',
    `topic` text NOT NULL COMMENT '频道(话题)内容:content:id',
    `urls` text NOT NULL COMMENT '内容Url/音频波文件数组地址',
    `status` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '状态:1~正常评论,2~优秀评论,-1~异常评论,-4~关闭/折叠,-5~被投诉',
    `reason`        VARCHAR(200)     NOT NULL  DEFAULT '' COMMENT '原由',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_comment_id` (`content_id`),    
    -- CONSTRAINT `zd_comment_ibfk_1` FOREIGN KEY (`content_id`) REFERENCES `zd_comment` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='评论';

-- SET FOREIGN_KEY_CHECKS=0;  
DROP TABLE IF EXISTS `zd_comment_praise`;#zd_comment_praise
CREATE TABLE `zd_comment_praise`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `comment_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '评论ID',
    `status` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '状态:-1~取消,1~点赞,2~超级赞',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_blog_praise_id` (`blog_id`),    
    -- CONSTRAINT `zd_blog_praise_ibfk_1` FOREIGN KEY (`blog_id`) REFERENCES `zd_blog` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB 
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='评论点赞';

-- SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `zd_comment_uid`;#zd_comment_uid
-- SET FOREIGN_KEY_CHECKS=1;
    CREATE TABLE `zd_comment_uid`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `comment_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '评论ID',
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `from_uid` BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '来自用户ID',
    `content_id` BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '动态ID',
    `content` text NOT NULL COMMENT '内容',
    `at` text NOT NULL COMMENT '@内容:content:id',
    `topic` text NOT NULL COMMENT '频道(话题)内容:content:id',
    `urls` text NOT NULL COMMENT '内容Url/音频波文件数组地址',
    `reply` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '回复:1~是,0~不是',
    `status` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '状态:1~正常评论,2~优秀评论,-1~异常评论,-4~关闭/折叠,-5~被投诉',
    `reason`        VARCHAR(200)     NOT NULL  DEFAULT '' COMMENT '原由',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_comment_id` (`content_id`),    
    -- CONSTRAINT `zd_comment_ibfk_1` FOREIGN KEY (`content_id`) REFERENCES `zd_comment_id` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='针对评论ID的评论';

-- SET FOREIGN_KEY_CHECKS=0;  
DROP TABLE IF EXISTS `zd_comment_uid_praise`;#zd_comment_uid_praise
CREATE TABLE `zd_comment_uid_praise`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `comment_uid_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '评论用户ID',
    `status` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '状态:-1~取消,1~点赞,2~超级赞',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_blog_praise_id` (`blog_id`),    
    -- CONSTRAINT `zd_blog_praise_ibfk_1` FOREIGN KEY (`blog_id`) REFERENCES `zd_blog` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB 
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='评论点赞';

-- SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `zd_top`;#zd_top
CREATE TABLE `zd_top`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `content_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内容来源ID',
    `status` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '状态:1~置顶,0~取消',
    `especially` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '特别状态:1~超级顶,0～取消超级顶',
    `type`      INT              NOT NULL DEFAULT '0' COMMENT '内容类型:1~用户,2~频道,3~动态,4~评论',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_top_id` (`content_id`),    
    -- CONSTRAINT `zd_top_ibfk_1` FOREIGN KEY (`content_id`) REFERENCES `zd_top` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='置顶';

-- SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `zd_share`;#zd_share
CREATE TABLE `zd_share`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `content_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内容来源ID',
    `status` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '状态:1~成功,0～失败',
    `type`      INT              NOT NULL DEFAULT '0' COMMENT '内容类型:1~用户,2~频道,3~动态,4~评论',
    `from_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '来自用户ID',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_share_id` (`content_id`),    
    -- CONSTRAINT `zd_share_ibfk_1` FOREIGN KEY (`content_id`) REFERENCES `zd_share` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='分享';

-- SET FOREIGN_KEY_CHECKS=0;
  DROP TABLE IF EXISTS `zd_at`;#zd~at
  -- SET FOREIGN_KEY_CHECKS=1;
      CREATE TABLE `zd_at`(
      `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
      `uid` BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '用户ID', 
      `content_id` BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '内容来源ID',
      `source`      INT              NOT NULL DEFAULT '0' COMMENT '内容类型:1~用户,2~频道,3~动态,4~评论',
      `from_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '来自用户ID',
      `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
      `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
      PRIMARY KEY (`id`)
      -- KEY `zd_at_id` (`content_id`),    
      -- CONSTRAINT `zd_at_ibfk_1` FOREIGN KEY (`content_id`) REFERENCES `zd_at` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
    )
      ENGINE = InnoDB
      -- DEFAULT CHARSET = utf8mb4
      COMMENT ='@xxx';

-- SET FOREIGN_KEY_CHECKS=0;
    DROP TABLE IF EXISTS `zd_file`;#zd~file
    SET FOREIGN_KEY_CHECKS=1;
    CREATE TABLE `zd_file`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `content_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内容来源ID',
    `classify` TINYINT(4)              NOT NULL DEFAULT '0' COMMENT '内容分类:1~用户,2~频道,3~动态,4~评论',
    `url` text NOT NULL COMMENT '内容Url', 
    `cover` text NOT NULL COMMENT '文件封面',
    `dynamic_cover` text NOT NULL COMMENT '动态封面',
    `name`        text     NOT NULL  COMMENT '文件名称',
    `sufix` text NOT NULL COMMENT '文件名后缀',
    `format`      TINYINT(4)              NOT NULL DEFAULT '0' COMMENT '内容格式:0:图片,1:音频,2:视频,3:文档,4:web,5:VR',
    `duration`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内容时长',
    `width`      INT              NOT NULL DEFAULT '0' COMMENT '内容宽',
    `height`      INT              NOT NULL DEFAULT '0' COMMENT '内容高',
    `size`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内容尺寸',
    `rotate`      INT              NOT NULL DEFAULT '0' COMMENT '角度旋转',
    `bitrate`      INT              NOT NULL DEFAULT '0' COMMENT '采样率', 
    `sample_rate`      INT              NOT NULL DEFAULT '0' COMMENT '频率', 
    `level`      INT              NOT NULL DEFAULT '0' COMMENT '质量:0~标准', 
    `mode`      INT              NOT NULL DEFAULT '0' COMMENT '模式', 
    `wave` text NOT NULL COMMENT '频谱',
    `lrc_zh` text NOT NULL COMMENT '字幕~中文',
    `lrc_es` text NOT NULL COMMENT '字母~英文',
    `source`      INT              NOT NULL DEFAULT '0' COMMENT '创作类型:0~原创,1~其它', 
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    -- KEY `zd_blog_file_id` (`blog_id`),    
    -- CONSTRAINT `zd_blog_file_ibfk_1` FOREIGN KEY (`blog_id`) REFERENCES `zd_blog` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='文件';

DROP TABLE IF EXISTS `zd_view`;#zd_view
    CREATE TABLE `zd_view`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `content_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内容来源ID',
    `num`      INT              NOT NULL DEFAULT '0' COMMENT '内容详情位置:来自哪张图片',
    `nums`      INT              NOT NULL DEFAULT '0' COMMENT '浏览次数',
    `type`      INT              NOT NULL DEFAULT '0' COMMENT '内容类型:1~用户,2~频道,3~动态,4~评论,5~音乐',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='阅览';

DROP TABLE IF EXISTS `zd_history`;#zd_history
    CREATE TABLE `zd_history`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `content_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内容来源ID',
    `num`      INT              NOT NULL DEFAULT '0' COMMENT '内容详情位置:来自哪张图片',
    `source`      INT              NOT NULL DEFAULT '0' COMMENT '内容类型:1~用户,2~频道,3~动态~图片,4~动态~音乐,5~动态~视频',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='历史';

DROP TABLE IF EXISTS `zd_word`;#zd_word
    CREATE TABLE `zd_word`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `content_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内容来源ID',
    `name`        VARCHAR(50)     NOT NULL  DEFAULT '' COMMENT '关键字名称',
    `source`      INT              NOT NULL DEFAULT '0' COMMENT '内容类型:1~用户,2~频道,3~动态,4~评论',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='关键字';

DROP TABLE IF EXISTS `zd_delete`;#zd_word
    CREATE TABLE `zd_delete`(
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `content_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内容来源ID',
    `status`        VARCHAR(50)     NOT NULL  DEFAULT '' COMMENT '状态:1~可以删除,0~不能删除',
    `type`      INT              NOT NULL DEFAULT '0' COMMENT '内容类型:1~用户,2~频道,3~动态,4~评论',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
  )
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='关键字';

-- SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `zd_report`;#zd_report
CREATE TABLE `zd_report`( 
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `content_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内容来源ID',
    `type`  INT              NOT NULL DEFAULT '0' COMMENT '投诉类型:',
    `source`  INT              NOT NULL DEFAULT '0' COMMENT '内容来源类型:1~用户,2~频道,3~动态,4~评论,5~用户评论的评论',
    `status` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '状态:1~未处理,2～成功,-1~失败',
    `reason`        VARCHAR(200)     NOT NULL  DEFAULT '' COMMENT '原由',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
  ) 
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='举报/投诉';   

  DROP TABLE IF EXISTS `zd_report_type`;#zd_report_type
  CREATE TABLE `zd_report_type`(
      `id`          INT UNSIGNED  NOT NULL  AUTO_INCREMENT,
      `name`        VARCHAR(20)     NOT NULL  DEFAULT '' COMMENT '投诉类型名称',
      `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
      `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
      PRIMARY KEY (`id`),
      UNIQUE KEY `idx_report_type_name` (`name`)
    ) 
      ENGINE = InnoDB
      -- DEFAULT CHARSET = utf8mb4
      COMMENT ='投诉类型';

-- SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `zd_feedback`;#zd_feedback
CREATE TABLE `zd_feedback`( 
    `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
    `uid`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
    `content_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '内容来源ID',
    `content` text NOT NULL COMMENT '描述',
    `url` text NOT NULL COMMENT '文件',
    `contact`        VARCHAR(30)     NOT NULL  DEFAULT '' COMMENT '联系方式:微信,QQ,邮箱或者联系电话',
    `status` TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '状态:0~未处理,1～成功,-1~失败',
    `source`  TINYINT(4)              NOT NULL DEFAULT '0' COMMENT '内容来源类型:0~app,1~音乐',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
  ) 
    ENGINE = InnoDB
    -- DEFAULT CHARSET = utf8mb4
    COMMENT ='反馈';  


-- SET FOREIGN_KEY_CHECKS = 1;
