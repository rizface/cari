# 🕶️ CARI
simple terminal app for search and watch movie or anime

## 🖥️ INSTALLATION
```shell
git clone  https://github.com/rizface/cari.git
cd /path/to/cari
go build -o cari
```

## 💻 USAGE
- :film_projector: Film
```shell
./cari film -t cars
```
- :japanese_ogre: Anime
```shell
./cari anime -t naruto
```

- :framed_picture: Gambar
  - download image from page 1 to page 30
  ```shell
    ./cari gambar -c anime -d /dir/to/save -p 30
  ```
  
  - download image from page 1
  ```shell
    ./cari gambar -c anime -d /dir/to/save
  ```
