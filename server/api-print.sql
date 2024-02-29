-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Feb 09, 2024 at 03:51 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `api-print`
--

-- --------------------------------------------------------

--
-- Table structure for table `auth_user_tokens`
--

CREATE TABLE `auth_user_tokens` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `access_token` varchar(350) DEFAULT NULL,
  `refesh_token` varchar(350) DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `auth_user_tokens`
--

INSERT INTO `auth_user_tokens` (`id`, `access_token`, `refesh_token`, `user_id`) VALUES
(2, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc0ODQ1OTEsImlzcyI6IjIifQ.WWY6Yt7ScGguQL34B0SYEYqAWvRbWSaX-S7AUGye8FM', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc1NzA5NjEsImlzcyI6IjIifQ.ktez0dEMkFfn2hPH0DCii8Bu1F3XemsKX8IrvB278Bo', 2),
(3, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc0ODQ2NDMsImlzcyI6IjMifQ.hj2RL4MlAI4KV9FID8idjhAbvRXmGHeWWDZL8fivELc', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc1NzEwMTMsImlzcyI6IjMifQ.fSvv6lD-ek6OE0CuWU9YKIIQp-VCLrNIhcysvyxJPeg', 3);

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `role` varchar(80) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id`, `role`) VALUES
(1, 'Admin'),
(2, 'User');

-- --------------------------------------------------------

--
-- Table structure for table `saldo_users`
--

CREATE TABLE `saldo_users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `saldo_user` varchar(350) DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `saldo_users`
--

INSERT INTO `saldo_users` (`id`, `saldo_user`, `user_id`) VALUES
(2, '10000', 2);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `username` varchar(80) DEFAULT NULL,
  `email` varchar(80) DEFAULT NULL,
  `password` longblob NOT NULL,
  `id_role` bigint(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `email`, `password`, `id_role`) VALUES
(2, 'nieBota25', 'dennieBotak25@gmail.com', 0x2432612431302454585844794a7152356f3548636d6a5451784735524f53656e4f714471424c39646459762e32556c7667764d6a4b6c66434732772e, 1),
(3, 'ekopass25', 'ekopass25@gmail.com', 0x24326124313024544c71575a375a326a5746436a45756948517053712e3058726539424f6a4d35476e44762f566536775058754f3861554f4e726547, 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `auth_user_tokens`
--
ALTER TABLE `auth_user_tokens`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_auth_user_tokens_users` (`user_id`);

--
-- Indexes for table `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `saldo_users`
--
ALTER TABLE `saldo_users`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_saldo_users_users` (`user_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD UNIQUE KEY `email` (`email`),
  ADD KEY `fk_users_role` (`id_role`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `auth_user_tokens`
--
ALTER TABLE `auth_user_tokens`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `roles`
--
ALTER TABLE `roles`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `saldo_users`
--
ALTER TABLE `saldo_users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `auth_user_tokens`
--
ALTER TABLE `auth_user_tokens`
  ADD CONSTRAINT `fk_auth_user_tokens_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `saldo_users`
--
ALTER TABLE `saldo_users`
  ADD CONSTRAINT `fk_saldo_users_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `users`
--
ALTER TABLE `users`
  ADD CONSTRAINT `fk_users_role` FOREIGN KEY (`id_role`) REFERENCES `roles` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
