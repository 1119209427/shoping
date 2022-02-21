CREATE DATABASE IF NOT EXISTS shopping;
USE shopping;
DROP TABLE IF EXISTS `user`;
CREATE TABLE
(
    `id` INT AUTO_INCREMENT PRIMARY KEY;
    `total_likes` INT NOT NULL,
     `user_name` VARCHAR(15) NOT NULL,
     `password` VARCHAR(20) NOT NULL,
     `email`  VARCHAR(20)  NOT NULL DEFAULT "",
      `phone` VARCHAR(11)  NOT NULL,
      `salt`  VARCHAR(20) NOT NULL,
     `reg_date` DATE       NOT NULL,
      `statement`          VARCHAR(90)  NOT NULL DEFAULT '这个人很懒，什么都没有写',
      `gender` VARCHAR(1)  NOT NULL DEFAULT 男,
      `balance` FLOAT   NOT NULL 0元,
       `cart_id` INT NOT NULL,
)charset="utf8mb4";
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`
(
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `good_id` INT NOT NULL,
    `user_id` INT NOT NULL,
    `value`  VARCHAR(200) NOT NULL,
    `time`  DATE    NOT NULL,
    `likes` INT  NOT NULL,

)charset="utf8mb4";
DROP TABLE IF EXISTS `merchant`
CREATE TABLE  `merchant`
(
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `notice` VARCHAR(200) NOT NULL,
    `merchant_number` INT NOT NULL,
    `merchant_name` VARCHAR (20) NOT NULL,
    `description` VARCHAR(200) NOT NULL,
    `channel`  VARCHAR(50)  NOT NULL,
    `time` DATE NOT NULL,
    `favorable`  FLOAT NOT NULL,
    `volume` INT NOT NULL DEFAULT 0,
    `followers` INT NOT NULL DEFAULT 0,
    `followings` INT NOT NULL DEFAULT 0,

)charset="utf8mb4";

DROP TABLE IF EXISTS `good_like`;
CREATE TABLE `good_like`
(
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `good_id` INT NOT NULL,
    `uid` INT NOT NULL
)charset="utf8mb4";
DROP TABLE IF EXISTS `user_follow`;
CREATE TABLE `user_follow`
(
    `id`  INT AUTO_INCREMENT PRIMARY KEY,
    `follower_uid`  INT NOT NULL,
    `following_uid` INT NOT NULL
) charset="utf8mb4";

CREATE TABLE `good_follow`
(
    `id`  INT AUTO_INCREMENT PRIMARY KEY,
    `follower_uid`  INT NOT NULL,
    `following_good_id` INT NOT NULL
) charset="utf8mb4";
DROP TABLE IF EXISTS `good_label`;

CREATE TABLE `good_label`
(
    `id`          INT AUTO_INCREMENT PRIMARY KEY,
    `good_id`          INT         NOT NULL,
    `good_label` VARCHAR(19) NOT NULL
) charset="utf8mb4";
DROP TABLE IF EXISTS `good`;

CREATE TABLE `good`
(
    `id`          INT AUTO_INCREMENT PRIMARY KEY,
    `title`       VARCHAR(80)  NOT NULL,
    `channel`     VARCHAR(4)   NOT NULL,
    `description` VARCHAR(250) NOT NULL,
    `good_url`   VARCHAR(120) NOT NULL,
    `cover_url`   VARCHAR(120) NOT NULL,
    `merchant_number`  INT          NOT NULL,
    `time`        TIMESTAMP    NOT NULL,
    `views`       INT          NOT NULl DEFAULT 0,
    `likes`       INT          NOT NULL DEFAULT 0,
    `shares`      INT          NOT NULL DEFAULT 0
) charset="utf8mb4";
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`{
    `uid` INT AUTO_INCREMENT PRIMARY KEY;
    `type` VARCHAR(10) NOT NULL,
     `good_id` INT NOT NULL,
}


