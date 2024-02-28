CREATE DATABASE IF NOT EXISTS  `songs` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE songs;
-- songs.catalogo_fuente definition

CREATE TABLE IF NOT EXISTS `catalogo_fuente` (
  `id_catalogo_fuente` int NOT NULL AUTO_INCREMENT,
  `nombre_fuente` text NOT NULL,
  `url_api` text,
  PRIMARY KEY (`id_catalogo_fuente`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- songs.transaccion_songs definition

CREATE TABLE IF NOT EXISTS `transaccion_songs` (
  `id_transaccion_song` int NOT NULL AUTO_INCREMENT,
  `id_track` text NOT NULL,
  `nombre_song` text NOT NULL,
  `artist` text NOT NULL,
  `duration` text,
  `album` text,
  `url_art_work` text NOT NULL,
  `price` double DEFAULT NULL,
  `origin` text,
  `fuente` int DEFAULT NULL,
  `fecha` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_transaccion_song`),
  KEY `transaccion_song_catalogo_fuente_FK` (`fuente`),
  CONSTRAINT `transaccion_song_catalogo_fuente_FK` FOREIGN KEY (`fuente`) REFERENCES `catalogo_fuente` (`id_catalogo_fuente`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT IGNORE INTO songs.catalogo_fuente
(id_catalogo_fuente, nombre_fuente, url_api)
VALUES(1, 'Apple', 'https://itunes.apple.com/search?term=jack+johnson');
INSERT IGNORE INTO songs.catalogo_fuente
(id_catalogo_fuente, nombre_fuente, url_api)
VALUES(2, 'CharLyrics', 'http://api.chartlyrics.com/apiv1.asmx/SearchLyric?artist=string&song=string');
