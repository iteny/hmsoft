/*
 Navicat Premium Data Transfer

 Source Server         : hmsoft
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : localhost:3306
 Source Schema         : hmsoft

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 01/09/2020 22:42:43
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for hm_login_logs
-- ----------------------------
DROP TABLE IF EXISTS `hm_login_logs`;
CREATE TABLE `hm_login_logs`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `time` int NOT NULL DEFAULT 0 COMMENT '登录时间',
  `ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '登录ip',
  `useragent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '浏览器信息',
  `uid` bigint NOT NULL COMMENT '登录用户ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 40 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_login_logs
-- ----------------------------
INSERT INTO `hm_login_logs` VALUES (1, 'admin', 1596548392, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (2, 'admin', 1596552286, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (3, 'admin', 1596552313, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (4, 'admin', 1596552651, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (5, 'admin', 1596721787, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (6, 'admin', 1596721793, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (7, 'admin', 1596721842, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (8, 'admin', 1596721843, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (9, 'admin', 1596721851, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (10, 'admin', 1596721964, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (11, 'admin', 1596722105, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (12, 'admin', 1596722528, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (13, 'admin', 1596722717, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (14, 'admin', 1596723308, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (15, 'admin', 1596723622, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (16, 'admin', 1597155036, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (17, 'admin', 1597156266, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (18, 'admin', 1597156294, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (19, 'admin', 1597156357, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (20, 'admin', 1597156403, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (21, 'admin', 1597156557, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (22, 'admin', 1598188570, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (23, 'admin', 1598357542, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (24, 'admin', 1598358787, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (25, 'admin', 1598360054, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (26, 'admin', 1598360273, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (27, 'admin', 1598360373, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (28, 'admin', 1598363374, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (29, 'admin', 1598364858, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (30, 'admin', 1598365003, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (31, 'admin', 1598365248, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (32, 'admin', 1598365304, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (33, 'admin', 1598537912, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (34, 'admin', 1598705302, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (35, 'admin', 1598796464, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (36, 'admin', 1598964271, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (37, 'admin', 1598967135, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (38, 'admin', 1598967998, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (39, 'admin', 1598968425, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36', 1);

-- ----------------------------
-- Table structure for hm_menus
-- ----------------------------
DROP TABLE IF EXISTS `hm_menus`;
CREATE TABLE `hm_menus`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` int NOT NULL COMMENT '父级菜单id',
  `name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `sort` int NOT NULL COMMENT '排序值',
  `path` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '跳转路由',
  `icon` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '图标',
  `isshow` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否显示',
  `level` tinyint(1) NOT NULL DEFAULT 1 COMMENT '菜单等级',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_menus
-- ----------------------------
INSERT INTO `hm_menus` VALUES (1, 0, '用户管理', 1, '/user', 'icon-users', 1, 0);
INSERT INTO `hm_menus` VALUES (2, 1, '用户列表', 1, '/userlist', 'icon-users', 1, 0);
INSERT INTO `hm_menus` VALUES (3, 1, '角色列表', 2, '/role', 'icon-users', 1, 0);
INSERT INTO `hm_menus` VALUES (4, 1, '权限列表', 3, '/permission', 'icon-users', 1, 0);

-- ----------------------------
-- Table structure for hm_users
-- ----------------------------
DROP TABLE IF EXISTS `hm_users`;
CREATE TABLE `hm_users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '登录密码',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `status` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '用户状态;0:禁用,1:正常,2:未验证',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE COMMENT '唯一',
  UNIQUE INDEX `nickname`(`nickname`) USING BTREE COMMENT '唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_users
-- ----------------------------
INSERT INTO `hm_users` VALUES (1, 'admin', '90b9aa7e25f80cf4f64e990b78a9fc5ebd6cecad', '', '', 1);

SET FOREIGN_KEY_CHECKS = 1;
