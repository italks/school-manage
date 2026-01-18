-- 创建数据库
CREATE DATABASE IF NOT EXISTS school_manage DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE school_manage;

-- 允许远程连接 (确保 root 用户可以从任意主机连接)
-- 注意: 这种操作在生产环境中存在安全风险，建议限制特定 IP
-- MySQL 8.0+ 语法
-- CREATE USER IF NOT EXISTS 'root'@'%' IDENTIFIED BY 'root';
-- GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;
-- FLUSH PRIVILEGES;

-- 1. 用户表
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password_hash` longtext NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `role` enum('admin','teacher','student') DEFAULT 'student',
  `status` enum('active','inactive') DEFAULT 'active',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_users_username` (`username`),
  UNIQUE KEY `idx_users_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 2. 教室表
CREATE TABLE IF NOT EXISTS `classrooms` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `capacity` bigint(20) DEFAULT '30',
  `equipment` text,
  `status` enum('available','maintenance','closed') DEFAULT 'available',
  `created_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 3. 排课表
CREATE TABLE IF NOT EXISTS `schedules` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `teacher_id` bigint(20) unsigned NOT NULL,
  `classroom_id` bigint(20) unsigned NOT NULL,
  `course_name` varchar(100) NOT NULL,
  `start_time` datetime(3) NOT NULL,
  `end_time` datetime(3) NOT NULL,
  `max_students` bigint(20) DEFAULT '20',
  `status` enum('scheduled','ongoing','completed','cancelled') DEFAULT 'scheduled',
  `created_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_schedules_teacher_id` (`teacher_id`),
  KEY `idx_schedules_classroom_id` (`classroom_id`),
  KEY `idx_schedules_start_time` (`start_time`),
  KEY `idx_schedules_end_time` (`end_time`),
  CONSTRAINT `fk_schedules_classroom` FOREIGN KEY (`classroom_id`) REFERENCES `classrooms` (`id`),
  CONSTRAINT `fk_schedules_teacher` FOREIGN KEY (`teacher_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 4. 预约表
CREATE TABLE IF NOT EXISTS `bookings` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `student_id` bigint(20) unsigned NOT NULL,
  `schedule_id` bigint(20) unsigned NOT NULL,
  `booking_time` datetime(3) DEFAULT NULL,
  `status` enum('booked','attended','absent','cancelled') DEFAULT 'booked',
  `cost` decimal(10,2) DEFAULT '0.00',
  `checkin_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_bookings_student_id` (`student_id`),
  KEY `idx_bookings_schedule_id` (`schedule_id`),
  CONSTRAINT `fk_bookings_schedule` FOREIGN KEY (`schedule_id`) REFERENCES `schedules` (`id`),
  CONSTRAINT `fk_bookings_student` FOREIGN KEY (`student_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 5. 学生钱包表
CREATE TABLE IF NOT EXISTS `student_wallets` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `student_id` bigint(20) unsigned NOT NULL,
  `balance` decimal(10,2) DEFAULT '0.00',
  `total_recharge` decimal(10,2) DEFAULT '0.00',
  `total_consume` decimal(10,2) DEFAULT '0.00',
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_student_wallets_student_id` (`student_id`),
  CONSTRAINT `fk_student_wallets_student` FOREIGN KEY (`student_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ==========================================
-- 测试数据 (密码统一为: 123456)
-- Hash: $2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy (Example hash for 123456)
-- 注意：实际生产环境请重新生成Hash
-- ==========================================

-- 1. 用户数据
INSERT INTO `users` (`username`, `password_hash`, `email`, `phone`, `role`, `status`, `created_at`, `updated_at`) VALUES
('admin', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'admin@school.com', '13800138000', 'admin', 'active', NOW(), NOW()),
('teacher_wang', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'wang@school.com', '13900139000', 'teacher', 'active', NOW(), NOW()),
('teacher_li', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'li@school.com', '13900139001', 'teacher', 'active', NOW(), NOW()),
('student_xiaoming', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'xiaoming@school.com', '13700137000', 'student', 'active', NOW(), NOW()),
('student_xiaohong', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'xiaohong@school.com', '13700137001', 'student', 'active', NOW(), NOW());

-- 2. 教室数据
INSERT INTO `classrooms` (`name`, `capacity`, `equipment`, `status`, `created_at`) VALUES
('101 钢琴教室', 5, '雅马哈钢琴, 谱架', 'available', NOW()),
('102 舞蹈教室', 20, '音响, 镜面墙, 把杆', 'available', NOW()),
('201 美术教室', 15, '画架, 投影仪', 'available', NOW()),
('202 理论教室', 30, '黑板, 投影仪, 桌椅', 'maintenance', NOW());

-- 3. 排课数据 (假设 teacher_id 2=Wang, 3=Li; classroom_id 1=101, 2=102)
-- 插入几条未来的课程
INSERT INTO `schedules` (`teacher_id`, `classroom_id`, `course_name`, `start_time`, `end_time`, `max_students`, `status`, `created_at`) VALUES
(2, 1, '钢琴基础一对一', DATE_ADD(NOW(), INTERVAL 1 DAY), DATE_ADD(NOW(), INTERVAL '1 1' DAY_HOUR), 1, 'scheduled', NOW()),
(2, 1, '钢琴进阶', DATE_ADD(NOW(), INTERVAL 1 DAY), DATE_ADD(NOW(), INTERVAL '1 2' DAY_HOUR), 1, 'scheduled', NOW()),
(3, 2, '少儿舞蹈初级', DATE_ADD(NOW(), INTERVAL 2 DAY), DATE_ADD(NOW(), INTERVAL '2 2' DAY_HOUR), 15, 'scheduled', NOW()),
(3, 2, '成人瑜伽', DATE_ADD(NOW(), INTERVAL 3 DAY), DATE_ADD(NOW(), INTERVAL '3 2' DAY_HOUR), 12, 'scheduled', NOW());

-- 4. 预约数据 (假设 student_id 4=Xiaoming)
INSERT INTO `bookings` (`student_id`, `schedule_id`, `booking_time`, `status`, `cost`) VALUES
(4, 1, NOW(), 'booked', 200.00);

-- 5. 钱包数据
INSERT INTO `student_wallets` (`student_id`, `balance`, `total_recharge`, `total_consume`, `updated_at`) VALUES
(4, 800.00, 1000.00, 200.00, NOW()),
(5, 0.00, 0.00, 0.00, NOW());
