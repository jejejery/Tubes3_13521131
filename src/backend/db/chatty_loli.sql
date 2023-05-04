-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 04, 2023 at 05:30 AM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.0.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `chatty_loli`
--

-- --------------------------------------------------------

--
-- Table structure for table `input_users`
--

CREATE TABLE `input_users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `session` bigint(20) DEFAULT NULL,
  `input_text` longtext DEFAULT NULL,
  `algorithm` tinyint(1) DEFAULT NULL,
  `answer` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `input_users`
--

INSERT INTO `input_users` (`id`, `session`, `input_text`, `algorithm`, `answer`) VALUES
(40, 1683170870939, 'Tambahkan pertanyaan film kesukaan nobita? dengan jawaban doraemon', 1, 'Pertanyaan berhasil ditambahkan ke database!'),
(41, 1683170870939, '(2+8)^2', 1, '100.00'),
(42, 1683170870939, '4^2^2', 1, '256.00'),
(43, 1683170870939, 'Siapa penemu lampu?', 1, 'Thomas Alfa Edison dong'),
(44, 1683170954459, 'Siapa dosen saat ini?', 1, 'Pertanyaan tidak ditemukan di database.\nApakah maksud anda: \n1. Siapa dosen stima saat ini?\n2. apa dia suka aku?\n3. Siapa penemu lampu?\n'),
(45, 1683170954459, '8/2', 1, '4.00'),
(46, 1683170954459, 'Hapus pertanyaan dafklsad ', 1, 'Pertanyaan yang ingin dihapus tidak tersedia di database!');

-- --------------------------------------------------------

--
-- Table structure for table `qn_as`
--

CREATE TABLE `qn_as` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `question` longtext DEFAULT NULL,
  `answer` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `qn_as`
--

INSERT INTO `qn_as` (`id`, `question`, `answer`) VALUES
(1, 'Matkul wajib terseru sem 4', 'STIMA DONGGG!!'),
(2, 'Apa tubes terseru stima?', 'Tubes 3 inii, yaitu tubes bikin chatgptt hurrayy!'),
(3, 'Apa tucil terseru stima?', 'Tucil 3 kemarinnn soalnya saya jadi tak tersesat lagi di itb hehe'),
(4, 'Siapa asisten yang paling ganteng?', 'Kak Saul :)'),
(5, 'Apa ibukota dari Jawa Barat?', 'Bandung'),
(6, 'Siapa dosen stima saat ini?', 'Pak Rinaldi'),
(7, 'Bagaimana cara membuat chatgpt?', 'Ya cari tau sendiri dongg, masa nanya saya'),
(8, 'Mengapa kuda tidak bisa terbang?', 'Karena gak punya sayap'),
(14, 'Siapa penemu lampu?', 'Thomas Alfa Edison dong'),
(16, 'hai kamu', 'iya '),
(24, 'apa dia suka aku?', 'iya'),
(26, 'aku suka kamu', 'oh'),
(28, 'kuda makan apa?', 'rumput'),
(29, 'apa itu curut', 'bebek'),
(30, 'film kesukaan nobita?', 'doraemon');

-- --------------------------------------------------------

--
-- Table structure for table `sessions`
--

CREATE TABLE `sessions` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `session` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `sessions`
--

INSERT INTO `sessions` (`id`, `session`) VALUES
(39, 1683170870939),
(40, 1683170954459);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `input_users`
--
ALTER TABLE `input_users`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `qn_as`
--
ALTER TABLE `qn_as`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `sessions`
--
ALTER TABLE `sessions`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `input_users`
--
ALTER TABLE `input_users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=47;

--
-- AUTO_INCREMENT for table `qn_as`
--
ALTER TABLE `qn_as`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=31;

--
-- AUTO_INCREMENT for table `sessions`
--
ALTER TABLE `sessions`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=41;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
