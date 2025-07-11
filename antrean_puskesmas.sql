-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jul 11, 2025 at 06:59 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `antrean_puskesmas`
--

-- --------------------------------------------------------

--
-- Table structure for table `antreans`
--

CREATE TABLE `antreans` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `pasien_id` bigint(20) UNSIGNED DEFAULT NULL,
  `nama_lengkap` varchar(100) NOT NULL,
  `nik` varchar(20) NOT NULL,
  `no_hp` varchar(20) NOT NULL,
  `layanan` varchar(100) NOT NULL,
  `tanggal_kunjungan` datetime(3) DEFAULT NULL,
  `nomor_antrean` varchar(10) NOT NULL,
  `status` enum('menunggu','dipanggil','selesai') DEFAULT 'menunggu',
  `created_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `antreans`
--

INSERT INTO `antreans` (`id`, `pasien_id`, `nama_lengkap`, `nik`, `no_hp`, `layanan`, `tanggal_kunjungan`, `nomor_antrean`, `status`, `created_at`) VALUES
(26, NULL, 'Putri Aisyah', '3312345678910111', '081234567890', 'Umum', '2025-07-11 00:00:00.000', 'A001', 'selesai', '2025-07-11 11:41:26.309'),
(27, NULL, 'Febia Riyana', '3312345678910112', '081234567891', 'Gigi dan Mulut', '2025-07-11 00:00:00.000', 'A002', 'selesai', '2025-07-11 11:43:03.574'),
(28, NULL, 'Lusia Dheanatalie', '3312345678910113', '081234567892', 'Kandungan', '2025-07-11 00:00:00.000', 'A003', 'selesai', '2025-07-11 11:44:00.837'),
(30, 3, 'Nida Hanifah', '3312345678910111', '081234567890', 'Umum', '2025-07-11 00:00:00.000', 'A004', 'selesai', '2025-07-11 16:20:25.494');

-- --------------------------------------------------------

--
-- Table structure for table `pasiens`
--

CREATE TABLE `pasiens` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `nama_lengkap` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `pasiens`
--

INSERT INTO `pasiens` (`id`, `nama_lengkap`, `email`, `password`, `created_at`) VALUES
(3, 'Nida Hanifah', 'nidahanifah@gmail.com', '$2a$10$X3xyCbtwx5aVbatWbgqyCeHrSAwIjTifO47cUxsuXROl6ZYK0M4Eu', '2025-07-11 16:16:55.990');

-- --------------------------------------------------------

--
-- Table structure for table `petugas`
--

CREATE TABLE `petugas` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `petugas`
--

INSERT INTO `petugas` (`id`, `username`, `password`, `created_at`) VALUES
(1, 'petugas', '$2a$10$ZDjyHJFbFqHFRTe/fYiq4OK6yQCyBBDW6zqwKaMtzVSvqZ85uGOkm', '2025-07-10 07:07:56.968');

-- --------------------------------------------------------

--
-- Table structure for table `riwayat_bulanans`
--

CREATE TABLE `riwayat_bulanans` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `tahun` bigint(20) NOT NULL,
  `bulan` bigint(20) NOT NULL,
  `jumlah_pasien` bigint(20) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `riwayat_bulanans`
--

INSERT INTO `riwayat_bulanans` (`id`, `tahun`, `bulan`, `jumlah_pasien`, `created_at`) VALUES
(10, 2025, 7, 4, '2025-07-11 16:40:46.294');

-- --------------------------------------------------------

--
-- Table structure for table `riwayat_harians`
--

CREATE TABLE `riwayat_harians` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `tanggal` datetime(3) NOT NULL,
  `nomor_antrean` varchar(10) NOT NULL,
  `nama_pasien` varchar(100) NOT NULL,
  `layanan` varchar(100) NOT NULL,
  `status` enum('selesai') DEFAULT 'selesai',
  `created_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `riwayat_harians`
--

INSERT INTO `riwayat_harians` (`id`, `tanggal`, `nomor_antrean`, `nama_pasien`, `layanan`, `status`, `created_at`) VALUES
(33, '2025-07-11 00:00:00.000', 'A001', 'Putri Aisyah', 'Umum', 'selesai', '2025-07-11 16:40:46.228'),
(34, '2025-07-11 00:00:00.000', 'A002', 'Febia Riyana', 'Gigi dan Mulut', 'selesai', '2025-07-11 16:40:46.243'),
(35, '2025-07-11 00:00:00.000', 'A003', 'Lusia Dheanatalie', 'Kandungan', 'selesai', '2025-07-11 16:40:46.256'),
(36, '2025-07-11 00:00:00.000', 'A004', 'Nida Hanifah', 'Umum', 'selesai', '2025-07-11 16:40:46.266');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `antreans`
--
ALTER TABLE `antreans`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_pasiens_antrean` (`pasien_id`);

--
-- Indexes for table `pasiens`
--
ALTER TABLE `pasiens`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_pasiens_email` (`email`);

--
-- Indexes for table `petugas`
--
ALTER TABLE `petugas`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_petugas_username` (`username`);

--
-- Indexes for table `riwayat_bulanans`
--
ALTER TABLE `riwayat_bulanans`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `riwayat_harians`
--
ALTER TABLE `riwayat_harians`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `antreans`
--
ALTER TABLE `antreans`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=32;

--
-- AUTO_INCREMENT for table `pasiens`
--
ALTER TABLE `pasiens`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `petugas`
--
ALTER TABLE `petugas`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `riwayat_bulanans`
--
ALTER TABLE `riwayat_bulanans`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `riwayat_harians`
--
ALTER TABLE `riwayat_harians`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=37;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `antreans`
--
ALTER TABLE `antreans`
  ADD CONSTRAINT `fk_pasiens_antrean` FOREIGN KEY (`pasien_id`) REFERENCES `pasiens` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
