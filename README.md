# openwrt-search

Basit bir Go programı ile [OpenWRT Table of Hardware](https://openwrt.org/toh.json) verisini kullanarak cihazların marka, model ve desteklediği sürüm bilgilerini listeler.

## 📋 Özellikler

- OpenWRT JSON verisini çeker (`toh.json`)
- "24." ile başlayan sürüm bilgilerini filtreler
- Terminalde tablo halinde gösterir

## 🚀 Kurulum

### Gereksinimler

- Go (1.21 veya üstü önerilir)
- İnternet bağlantısı (veri çekmek için)

### Kurulum ve Çalıştırma

```bash
git clone https://github.com/murat-akpinar/openwrt-search.git
cd openwrt-search
go run openwrt_search.go
```

Alternatif olarak binary derlemek için:

```bash
go build -o openwrt-search
./openwrt-search
```

## 🗂 Örnek Çıktı

```
Marka      Model           Sürüm
------     -----           -----
TP-Link    Archer C7       24.05.0
Netgear    WNDR3700        24.02.1
...
```

