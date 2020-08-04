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

 Date: 04/08/2020 23:00:58
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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_login_logs
-- ----------------------------
INSERT INTO `hm_login_logs` VALUES (1, 'admin', 1596548392, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (2, 'admin', 1596552286, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (3, 'admin', 1596552313, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);
INSERT INTO `hm_login_logs` VALUES (4, 'admin', 1596552651, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36', 1);

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
