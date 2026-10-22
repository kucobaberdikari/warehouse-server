# Analisis Gap Fitur Warehouse Management

Tanggal pemeriksaan ulang: 2026-10-20  
Dasar pemeriksaan: route di `main.go`, handler, service, repository, entity, dan dokumen API yang tersedia. Status di bawah mencerminkan kode saat pemeriksaan, bukan hanya deskripsi pada dokumen lama.

## Ringkasan

Aplikasi memiliki route dan implementasi dasar untuk autentikasi, user, supplier, kategori, produk, mutasi stok, dan stock adjustment. Mutasi dan adjustment saat ini sama-sama memperbarui saldo serta menyimpan ledger dalam transaksi database. Handler juga mengambil aktor kedua operasi dari `currentUser`, bukan dari identitas yang dikirim client.

Validasi ulang `go test ./... -v` pada 2026-10-20 lulus. Temuan lama tentang build yang gagal karena signature `SearchProduct` tidak lagi berlaku: interface dan implementasi sama-sama menerima empat argumen. Namun, urutan parameter yang dikirim service berbeda dari urutan yang dideklarasikan repository, sehingga filter kategori dan SKU berpotensi tertukar.

Prioritas terdekat adalah menutup celah otorisasi user dan validasi claim JWT, memperbaiki query/pencarian yang rusak, serta menguji konsistensi stok. Setelah fondasi tersebut, low-stock dan dashboard memberi nilai operasional paling cepat; PO dan penerimaan barang menjadi tahap berikutnya. Multi-gudang sebaiknya menunggu kebutuhan lokasi yang jelas karena berdampak pada model stok dan transaksi.

## Fitur yang Sudah Tersedia

| Area             | Status berdasarkan kode                                                                                                                                                                                      |
| ---------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Auth             | Register dan login tersedia; token dipakai untuk route terlindungi.                                                                                                                                          |
| User             | Register, login, cek email, detail, dan update tersedia. Update berdasarkan ID belum membatasi target ke pemilik akun/admin. Endpoint upload avatar dan daftar user tidak terdaftar di `main.go`.            |
| Role             | Field `Role` tersedia; registrasi menetapkan role default `staff`; middleware `RequireRole` digunakan untuk sebagian route create/update/delete. Belum ada API aman untuk menetapkan atau mengubah role.     |
| Supplier         | Route create, list, detail, update, dan delete terdaftar. Implementasi list/delete repository perlu diperbaiki (lihat temuan).                                                                               |
| Category         | Route list, detail, create, dan update terdaftar. Route delete belum tersedia.                                                                                                                               |
| Product          | Route create, list/search/paginasi, detail, filter per supplier, update, dan delete terdaftar. Pencarian gabungan berpotensi memakai urutan SKU/kategori yang tertukar dan menggabungkan filter dengan `OR`. |
| Mutation         | Create, list, dan detail tersedia. Create mengubah stok dan menyimpan ledger secara transaksional; aktor diambil dari user terautentikasi. Filter dan pagination belum tersedia.                             |
| Stock adjustment | Create, list, detail, dan route riwayat per produk tersedia. Create mengubah stok dan menyimpan ledger secara transaksional; aktor diambil dari user terautentikasi. Handler riwayat per produk masih stub.  |

## Temuan Prioritas

### Status Temuan Sebelumnya

- Build blocker `SearchProduct`: **sudah tidak berlaku**. `go test ./... -v` lulus, dan signature interface/implementasi kini sama-sama empat argumen.
- Mutasi stok: **sudah diterapkan**. `mutation/repository.go` memperbarui stok dan membuat ledger dalam satu transaksi serta menolak outbound melebihi stok.
- Stock adjustment: **sebagian sudah diterapkan**. Repository memperbarui stok dan menyimpan ledger secara transaksional, memvalidasi produk dan stok decrease. Riwayat per produk masih stub pada handler.
- Identitas pelaku: **sudah diterapkan di handler**. Handler mutasi dan adjustment mengambil aktor dari `currentUser`; input adjustment mengecualikan `CreatedBy` dari JSON.
- Temuan yang masih berlaku: update user belum membatasi target; claim JWT di-cast tanpa validasi; query supplier FindAll/delete salah; urutan/semantik filter pencarian produk berisiko salah; pagination list mutation/adjustment belum ada.

### P0 - Otorisasi dan keamanan user

- `middleware/RoleMiddleware.go` sudah ada dan dipasang pada beberapa route admin, termasuk update produk/supplier. Ini berarti RBAC dasar sudah tersedia; yang belum ada adalah pengelolaan role dan kebijakan yang konsisten untuk semua operasi sensitif.
- Route `POST /user/update/:id` menerima ID target dari path. Handler/service tidak memastikan bahwa ID tersebut sama dengan user yang login atau bahwa pemanggil adalah admin. User ber-token dapat mengubah profil user lain.
- Registrasi publik menetapkan semua akun baru menjadi `staff`, tetapi belum ada alur provisioning admin atau endpoint admin untuk mengubah role.
- Auth middleware langsung melakukan type assertion pada `claim["user_id"]` menjadi `float64`. Claim hilang atau bertipe tidak sesuai dapat menyebabkan panic. Validasi claim perlu dilakukan sebelum konversi.
- Rekomendasi: batasi edit profil ke pemilik akun; izinkan admin mengedit user lain; buat pengelolaan role admin-only; uji route dengan token staff, admin, token invalid, dan tanpa token.

### P1 - Repository dan endpoint CRUD

- `supplier/repository.go`: `FindAll` menggunakan `r.db.Find("Supplier")`, bukan pointer slice model; query delete memakai kondisi `id +?`, bukan kondisi ID yang valid.
- `category/repository.go` memang memiliki `DeleteCategory`, tetapi category service tidak mengeksposnya, tidak ada handler delete, dan tidak ada route delete.
- `product/service.go` mengirim argumen `SearchProduct` dengan urutan name, supplier, category, SKU, sedangkan deklarasi repository mengharapkan name, supplier, SKU, category. Filter pencarian juga menggabungkan kriteria berbeda dengan `OR`, sehingga kombinasi filter dapat menghasilkan cakupan lebih luas dari yang diminta.
- Product dan supplier tidak memiliki field `DeletedAt`; delete yang ada merupakan hard delete, bukan soft delete. Terapkan soft delete hanya bila riwayat/relasi harus dipertahankan, dan tentukan aturan untuk data yang masih direferensikan.

### P1 - Validasi API dan dokumentasi

- `API_TEST_CASES.md` dan `WAREHOUSE_API_BLUEPRINT.md` tidak sepenuhnya cocok dengan route saat ini. Contohnya, route produk sekarang `PUT /product/:id`, stock adjustment memakai `/stock-adjustments` tanpa prefix `/inventory`, dan route daftar produk adalah `/products`.
- `API_TEST_CASES.md` menyatakan `GET /product/supplier/:supplier_id` membaca query parameter dan gagal memfilter; handler aktual sudah memakai `ShouldBindUri`, sehingga catatan itu stale. Sebaliknya, test case menyebut endpoint riwayat stock adjustment per produk seolah berfungsi, padahal handler aktual masih stub.
- Beberapa list sudah memiliki pagination (produk dan supplier), tetapi kategori, mutation, dan stock adjustment belum. List mutation/adjustment dapat membesar tanpa batas.
- Envelope/status response belum konsisten sepenuhnya: handler stock adjustment memakai struktur `status/data` tersendiri, sementara banyak handler lain memakai `helper.APIResponse`.
- Rekomendasi: jadikan route aktual sebagai acuan dokumentasi, tandai endpoint yang belum berfungsi, dan tambahkan test integrasi untuk route serta role/ownership.

## Checklist Fitur dan Perbaikan Berdasarkan Prioritas

### P0 - Keamanan dan integritas data

- [x] Batasi `POST /user/update/:id` agar staff hanya dapat mengubah profil sendiri; admin boleh mengubah user lain.
- [x] Validasi keberadaan dan tipe claim `user_id` di auth middleware sebelum konversi, lalu kembalikan 401 tanpa panic untuk token malformed.
- [x] Tentukan dan terapkan role yang boleh membuat mutation dan stock adjustment; saat ini kedua route hanya memerlukan autentikasi.
- [x] Tambahkan test repository/integrasi untuk memastikan mutation dan adjustment atomik, menolak stok negatif, dan tidak meninggalkan ledger bila update gagal.
- [ ] Implementasikan handler `GET /products/:product_id/stock-adjustments` dengan validasi ID, pemanggilan service, dan response konsisten.
- [ ] Putuskan apakah ledger perlu menyimpan `balance_before` dan `balance_after`; saat ini catatan tidak menyimpan snapshot saldo.

### P1 - Perbaikan alur inti dan visibilitas stok

- [ ] Perbaiki `supplier.FindAll` dan query `DeleteSupplier`; verifikasi list dan delete dengan database test.
- [ ] Selaraskan urutan parameter pencarian produk, lalu tentukan semantik AND/OR untuk filter nama, SKU, kategori, dan supplier; tambahkan test kombinasi filter.
- [ ] Tambahkan endpoint low-stock untuk produk dengan `stock <= min_stock`.
- [ ] Tambahkan dashboard ringkas: jumlah produk, stok rendah, kuantitas/nilai stok bila data harga tersedia, serta ringkasan mutation.
- [ ] Tambahkan pagination dan filter yang diperlukan untuk list mutation dan stock adjustment.
- [ ] Lengkapi delete category pada service, handler, route, dan aturan ketika kategori masih dipakai produk.
- [ ] Selaraskan `API_TEST_CASES.md` dan `WAREHOUSE_API_BLUEPRINT.md` dengan route aktif, response aktual, dan endpoint yang masih stub.
- [ ] Tambahkan test HTTP untuk auth, RBAC, ownership, dan route yang tercatat pada blueprint.

### P2 - Alur pengadaan dan administrasi

- [ ] Tambahkan purchase order: supplier, item, kuantitas, harga, status, dan perubahan status yang tervalidasi.
- [ ] Tambahkan penerimaan barang penuh/parsial yang menambah stok secara transaksional dan terhubung ke PO serta ledger.
- [ ] Tambahkan pengelolaan role admin-only dan prosedur provisioning admin yang tidak membuka eskalasi hak melalui registrasi publik.
- [ ] Tambahkan change password dan reset password dengan verifikasi identitas/token yang sesuai.
- [ ] Tambahkan audit log untuk perubahan master data dan administrasi di luar ledger stok.
- [ ] Tambahkan ekspor laporan setelah format dan kebutuhan laporan disepakati.
- [ ] Tetapkan kebijakan retensi riwayat dan penghapusan untuk produk/supplier yang masih direferensikan; saat ini delete merupakan hard delete.

### P3 - Ekspansi operasional (sesuai kebutuhan bisnis)

- [ ] Rancang multi-warehouse/lokasi dan saldo per lokasi sebelum transfer antar gudang.
- [ ] Tambahkan stock transfer dengan dispatch, penerimaan, dan status in-transit setelah model multi-warehouse tersedia.
- [ ] Evaluasi barcode/QR, batch/lot, retur, dan notifikasi berdasarkan proses operasional yang disepakati.
- [ ] Tambahkan refresh/revoke token bila strategi sesi membutuhkan logout yang mencabut token atau sesi berumur panjang.

Checklist ini membedakan fitur yang benar-benar belum tersedia dari perbaikan atas endpoint yang sudah ada tetapi memiliki celah atau implementasi belum lengkap. Prioritas P0/P1 disusun berdasarkan risiko keamanan, akurasi stok, dan manfaat operasional; P2/P3 membutuhkan keputusan proses bisnis atau fondasi data tambahan.

## Rekomendasi Urutan Pengerjaan

1. **Tutup P0**: ownership user, validasi claim JWT, kebijakan role untuk perubahan stok, test transaksi, dan endpoint riwayat adjustment.
2. **Tutup P1**: query supplier, filter produk, low-stock/dashboard, pagination, category delete, dan dokumentasi/test HTTP.
3. **Bangun P2**: role administration, password recovery, audit, PO, goods receipt, dan ekspor sesuai kebutuhan.
4. **Rencanakan P3**: multi-warehouse mendahului stock transfer; tambahkan fitur lanjutan sesuai proses bisnis.

## Catatan Cakupan

Laporan ini adalah audit statis atas kode dan dokumen yang ada, ditambah `go test ./... -v`; test saat ini lulus tetapi hanya satu test mutation yang tersedia dan belum membuktikan perilaku terhadap database produksi. Jalankan test integrasi dengan fixture sebelum menutup temuan, terutama untuk query supplier, pencarian produk, transaksi stok, dan akses berbasis role.
