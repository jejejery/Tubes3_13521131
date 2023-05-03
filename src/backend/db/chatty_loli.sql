-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: May 01, 2023 at 07:41 AM
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
-- Database: `chatty_loli`
--

-- --------------------------------------------------------

--
-- Table structure for table `histories`
--

CREATE TABLE `histories` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `urutan` bigint(20) UNSIGNED NOT NULL,
  `text` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `input_users`
--

CREATE TABLE `input_users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `input_text` longtext DEFAULT NULL,
  `algorithm` tinyint(1) DEFAULT NULL,
  `answer` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `input_users`
--

INSERT INTO `input_users` (`id`, `input_text`, `algorithm`, `answer`) VALUES
(1, 'haijasdfk', 1, ''),
(2, 'marikitacoba', 1, ''),
(3, 'lets try', 1, ''),
(4, 'lets trydsafkd', 1, ''),
(5, 'hai gais', 1, ''),
(6, 'markicob', 1, ''),
(7, 'hai lagi', 1, ''),
(8, '2*5', 1, ''),
(9, '25/08/2023', 1, ''),
(10, 'test', 1, ''),
(11, '2+5', 1, ''),
(12, 'try', 1, ''),
(13, 'masih masuk njr', 1, ''),
(14, '', 0, ''),
(15, '', 0, ''),
(16, '', 0, ''),
(17, '', 0, ''),
(18, '', 0, ''),
(19, '', 0, ''),
(20, '', 0, ''),
(21, 'cape', 0, ''),
(22, 'cape bgttt', 0, ''),
(23, '', 1, ''),
(24, '', 0, ''),
(25, 'ok marki', 1, ''),
(26, '2+5', 1, '7'),
(27, '7+9*3', 1, '34'),
(28, 'tambahkan pertanyaan ini dengan beberapa kata-kata acak jawaban: ini adalah jawaban saya', 1, ''),
(29, 'tambahkan pertanyaan ini dengan beberapa kata-kata acak jawaban: ini adalah jawaban saya', 1, ''),
(30, 'Tambahkan pertanyaan xx dengan jawaban yy', 1, 'adding'),
(31, 'Tambahkan pertanyaan aaaaa dengan jawaban bbbb', 1, 'adding'),
(32, 'hai yorobun', 1, ''),
(33, '2+5*3/2', 1, '9'),
(34, '29/04/2023', 1, 'Sabtu'), 
(35, '31/04/2023', 1, 'Senin'),
(36, '-1/04/2023', 1, '0'),
(37, '32/04/2023', 1, 'Selasa'),
(38, 'tambahkan pertanyaan xxx dengan jawaban yyyy', 1, 'adding'),
(39, 'tambahkan pertanyaan xxx dengan jawaban yyyy', 1, 'adding'),
(40, 'tAmbaHkan Pertanyaan xxx dengan jawaban yyyyY', 1, 'adding'),
(41, 'tAmbaHkan Pertanyaan xxx dengan jawaban yyyyY', 1, 'adding'),
(42, 'tAmbaHkan Pertanyaan xxxdfkaod dengan jawaban yyyyY', 1, 'adding'),
(43, 'tAmbaHkan Pertanyaan xxxdfkaod dengan jawaban yyyyYdaf', 1, 'adding'),
(44, 'tambahkan pertanyaan asdjfhaj dengan jawaban yydfhasd', 1, 'adding'),
(45, 'Tambahan pertanyaan xxx dengan jawaban yyy', 1, ''),
(46, 'Tambahan pertanyaan xxx dengan jawaban yyy', 1, ''),
(47, 'Tambahan pertanyaan xxx dengan jawaban yyy', 1, ''),
(48, 'Tambahan pertanyaan xxx dengan jawaban yyy', 1, ''),
(49, 'Tambahan pertanyaan xxx dengan jawaban yyy', 1, ''),
(50, 'Tambahan pertanyaan xxx dengan jawaban yyy', 1, ''),
(51, 'tambahan pertanyaan xxx dengan jawaban yyy', 1, ''),
(52, 'tambahkan pertanyaan xxx dengan jawaban yyyy', 1, 'adding'),
(53, 'Tambahkan pertanyaan xx dengan jawaban yy', 1, 'adding'),
(54, 'Tambahkan pertanyaan xxfakl;dksa dengan jawaban yydfsa;l', 1, 'adding'),
(55, 'Tambahkan pertanyaan xxx dengan jawaban yyy', 1, 'adding'),
(56, 'Tambahkan pertanyaan xxx dasflks dsafkdsa dengan jawaban yyyda dsflaksdmf', 1, 'adding'),
(57, 'Tambahkan pertanyaan xxx dengan jawaban dsafjkasl', 1, 'adding'),
(58, 'tambahkan pertanyaan sdkjfa dengan jawaban dsfdlkafja', 1, 'adding'),
(59, 'tambahkan pertanyaan sdkjfa dengan jawaban dsfdlkafja', 1, 'adding'),
(60, 'tambahkan pertanyaan sdkjfa dengan jawaban dsfdlkafja', 1, 'adding'),
(61, 'tambahkan pertanyaan sdkjfa dengan jawaban dsfdlkafja', 1, 'adding'),
(62, 'Tambahkan pertanyaan Siapakan dosen stima? dengan jawaban Pak Rin', 1, 'adding'),
(63, 'Tambahkan pertanyaan Ini apa? dengan jawaban Chatty-Loli', 1, 'adding'),
(64, 'Tambahkan pertanyaan xjakldsfj denga jawaban djasfkljsdalk', 1, ''),
(65, 'Tambahkan pertanyaan Ini apa? dengan jawaban Chatty-Loli', 1, 'adding'),
(66, 'Tambahkan pertanyaan Ini apa? dengan jawaban Chatty-Loli\'', 1, 'adding'),
(67, 'Tambahkan pertanyaan Ini apa? dengan jawaban Chatty-Loli\'', 1, 'adding'),
(68, 'Tambahkan pertanyaan Ini apa? dengan jawaban Chatty-Loli\'', 1, 'adding'),
(69, 'Tambahkan pertanyaan Ini apa? dengan jawaban Chatty-Loli\'', 1, 'adding'),
(70, 'Tambahkan pertanyaan Ini apa? dengan jawaban Chatty-Loli\'', 1, 'adding'),
(71, 'tAmbahKan pertanyaan Siapa penemu lampu? dengan jawaban Thomas Alfa Edison dong', 1, 'adding'),
(72, 'Hapus pertanyaan Ini Apa?', 1, 'erasing'),
(73, 'Hapus pertanyaan Ini apa?', 1, 'erasing'),
(74, 'Hapus pertanyaan Siapa penemu lampu?', 1, 'erasing'),
(75, 'Hapus pertanyaan Siapa penemu lampu?', 1, 'erasing'),
(76, 'Hapus pertanyaan Ini apa? ', 1, 'erasing'),
(77, 'Hapus pertanyaan Ini apa?', 1, 'erasing'),
(78, 'Hapus pertanyaan Ini apa?', 1, 'erasing'),
(79, 'Hapus pertanyaan Ini apa?', 1, 'erasing'),
(80, 'Hapus pertanyaan Ini apa?', 1, 'erasing'),
(81, 'Hapus pertanyaan afdsaf?', 1, 'erasing'),
(82, 'Hapus pertanyaan Ini apa?', 1, 'erasing'),
(83, 'Hapus pertanyaan Ini apa?', 1, 'erasing'),
(84, 'Hapus pertanyaan Ini apa?', 1, 'erasing'),
(85, 'Hapus pertanyaan Siapa penemu lampu?', 1, 'erasing'),
(86, 'Hapus pertanyaan Ini apa?', 1, 'erasing'),
(87, 'Hapus pertanyaan Ini apa?', 1, 'erasing'),
(88, 'Hapus pertanyaan Ini apa?', 1, 'erasing'),
(89, 'Hapus pertanyaan Ini apa?', 1, 'erasing'),
(90, 'Hapus pertanyaan Ini apa?', 1, 'erasing'),
(91, 'Hapus pertanyaan Ini apa?', 1, 'erasing'),
(92, 'Hapus pertanyaan Siapa penemu lampu?', 1, 'erasing'),
(93, 'Tambahkan pertanyaan hai dengan jawaban oke', 1, 'adding'),
(94, 'Tambahkan pertanyaan hai dengan jawaban oke', 1, 'adding'),
(95, 'Tambahkan pertanyaan hai dengan jawaban oke', 1, 'adding'),
(96, 'Tambahkan pertanyaan hai dengan jawaban oke', 1, 'adding'),
(97, 'Tambahkan pertanyaan hai dengan jawaban oke', 1, 'adding'),
(98, 'Hapus pertanyaan hai', 1, 'erasing'),
(99, 'Hapus pertanyaan hai', 1, 'erasing'),
(100, 'Tambahkan pertanyaan hai kamu dengan jawaban iya ', 1, 'adding'),
(101, '2+3/5', 1, '2'),
(102, '2+3/5', 1, '2'),
(103, '2+3/5', 1, '2'),
(104, 'Tambahkan pertanyaan xxx dengan jawaban mantap', 1, 'adding'),
(105, 'Tambahkan pertanyaan xxx dengan jawaban mantap', 1, 'adding'),
(106, 'Hapus pertanyaan 1', 1, 'erasing'),
(107, 'Hapus pertanyaan Siapa penemu lampu?', 1, 'erasing'),
(108, 'Hapus pertanyaan siapa penemu lampu?', 1, 'erasing'),
(109, 'Hapus pertanyaan siapa penemu lampu?', 1, 'erasing'),
(110, 'Hapus pertanyaan siapa penemu lampu?', 1, 'erasing'),
(111, 'Hapus pertanyaan Siapa penemu lampu?', 1, 'erasing'),
(112, 'Hapus pertanyaan Siapa penemu lampu?', 1, 'erasing'),
(113, 'Hapus pertanyaan Siapa penemu kuda?', 1, 'erasing'),
(114, 'Hapus pertanyaan Siapa penemu kuda?', 1, 'erasing'),
(115, 'Tambahkan pertanyaan aasfj dengan jawaban asdifjails', 1, 'adding'),
(116, 'Tambahkan pertanyaan aasfj dengan jawaban asdifjails', 1, 'adding'),
(117, 'Tambahkan pertanyaan aasfj dengan jawaban asdifjails', 1, 'adding'),
(118, 'Tambahkan pertanyaan Siapa penemu lamput? dengan jawaban dkajfsda', 1, 'adding'),
(119, 'Tambahkan pertanyaan Siapa penemu lamput? dengan jawaban dkajfsda', 1, 'adding'),
(120, 'Hapus pertanyaan Siapa penemu lamput?', 1, 'erasing'),
(121, 'Hapus pertanyaan Siapa penemu lamput?', 1, 'erasing'),
(122, 'Hapus pertanyaan Siapa penemu lamput?', 1, 'erasing'),
(123, 'Hapus pertanyaan Siapa penemu lamput?', 1, 'erasing'),
(124, 'Hapus pertanyaan Siapa penemu lamput?', 1, 'erasing'),
(125, 'Hapus pertanyaan Siapa penemu lamput?', 1, 'erasing'),
(126, 'Hai yorobun', 1, ''),
(127, 'Hai yorobun', 1, ''),
(128, 'Ok lah', 1, ''),
(129, 'Tambahkan pertanyaan ini dengan jawaban itu', 1, 'adding'),
(130, '1+5', 1, '6'),
(131, 'test', 0, 'Pertanyaan tidak ditemukan di database.\nApakah maksud anda: \n. ini\n. aasfj\n. hai\n'),
(132, 'makan', 0, 'Pertanyaan tidak ditemukan di database.\nApakah maksud anda: \n. hai\n. ini\n. hai kamu\n'),
(133, 'hai', 0, 'oke'),
(134, 'ini', 0, 'itu'),
(135, 'ini', 0, 'itu'),
(136, 'ini', 0, 'itu'),
(137, 'hai', 0, 'oke'),
(138, '2+8/4', 1, '4'),
(139, '02/05/2023', 1, 'Selasa'),
(140, 'ini', 1, 'itu'),
(141, 'ini', 1, 'itu'),
(142, 'hai', 1, 'oke'),
(143, 'Tambahkan pertanyaan aaa dengan jawaban bbb', 1, 'Question is added to database!');

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
(15, 'hai', 'oke'),
(16, 'hai kamu', 'iya '),
(17, 'xxx', 'mantap'),
(18, 'xxx', 'mantap'),
(19, 'aasfj', 'asdifjails'),
(21, 'ini', 'itu'),
(22, 'aaa', 'bbb');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `histories`
--
ALTER TABLE `histories`
  ADD PRIMARY KEY (`id`,`urutan`);

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
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `histories`
--
ALTER TABLE `histories`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `input_users`
--
ALTER TABLE `input_users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=144;

--
-- AUTO_INCREMENT for table `qn_as`
--
ALTER TABLE `qn_as`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=23;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
