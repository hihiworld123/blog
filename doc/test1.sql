/*
 Navicat Premium Dump SQL

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50726 (5.7.26)
 Source Host           : localhost:3306
 Source Schema         : test1

 Target Server Type    : MySQL
 Target Server Version : 50726 (5.7.26)
 File Encoding         : 65001

 Date: 14/11/2025 21:01:27
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `content` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `post_id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_comments_post_id`(`post_id`) USING BTREE,
  INDEX `idx_comments_user_id`(`user_id`) USING BTREE,
  CONSTRAINT `fk_posts_comments` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES (1, '好文章，写得好12321', 2, 7, '2025-11-14 20:38:42.401');
INSERT INTO `comments` VALUES (3, '好文章，写得好12321', 4, 7, '2025-11-14 20:45:44.309');
INSERT INTO `comments` VALUES (4, '好文章，写得好12321', 5, 7, '2025-11-14 20:45:49.075');
INSERT INTO `comments` VALUES (5, '好文章，写得好12321', 6, 7, '2025-11-14 20:45:55.325');
INSERT INTO `comments` VALUES (6, '好文章，写得好12321，21321321321', 2, 7, '2025-11-14 20:46:16.659');

-- ----------------------------
-- Table structure for posts
-- ----------------------------
DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `content` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_posts_user_id`(`user_id`) USING BTREE,
  CONSTRAINT `fk_users_posts` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of posts
-- ----------------------------
INSERT INTO `posts` VALUES (2, '学习1', '学习，学习123', 7, '2025-11-14 16:34:12.912', '2025-11-14 20:14:08.158');
INSERT INTO `posts` VALUES (4, '21321312321321好好学习，天天向上123', '21321312321321好好学习，天天向上123，好好学习，天天向上123, 好好学习，天天向上；好好学习，天天向上；好好学习，天天向上；好好学习，天天向上；好好学习，天天向上；好好学习，天天向上；好好学习，天天向上；好好学习，天天向上，好好学习，天天向上123', 7, '2025-11-14 20:23:07.173', '2025-11-14 20:23:07.173');
INSERT INTO `posts` VALUES (5, '12321321好好学习，天天向上123', '3421321321321, 好好学习，天天向上；好好学习，天天向上；好好学习，天天向上；好好学习，天天向上；好好学习，天天向上；好好学习，天天向上；好好学习，天天向上；好好学习，天天向上，好好学习，天天向上123', 7, '2025-11-14 20:23:23.570', '2025-11-14 20:23:23.570');
INSERT INTO `posts` VALUES (6, 'gorm 更新值', 'gorm 更新值, 好好学习，天天向上；好好学习，天天向上；好好学习，天天向上；好好学习，天天向上；好好学习，天天向上；好好学习，天天向上；好好学习，天天向上；好好学习，天天向上，好好学习，天天向上123', 7, '2025-11-14 20:23:56.074', '2025-11-14 20:23:56.074');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_users_username`(`username`) USING BTREE,
  UNIQUE INDEX `uni_users_email`(`email`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'test1', '$2a$10$WlKYsMJ2MLUIoyUsCSk/s.81QwWGV2Qn326ZLVhbMDOvC/qkRYFBW', 'abc123@gmail.com');
INSERT INTO `users` VALUES (3, 'test2', '$2a$10$HA8UZCsqAEZ2JMzU2NzrYOPorFrdnEoe00enTSRim1WYYMXeqz7Xa', 'abc@gmail.com');
INSERT INTO `users` VALUES (7, 'test3', '$2a$10$II278wIBk4y.8Rz1HH8AMubzrRSP9vFmwDd91LO9bw60HMouxtiCi', 'abcd1@gmail.com');

SET FOREIGN_KEY_CHECKS = 1;
