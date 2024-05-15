package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;type:serial;column:id"`
	Name     string `gorm:"type:varchar(255);column:name"`
	Category string `gorm:"type:varchar(255);column:category"`
	Price    int    `gorm:"type:int;column:price"`
}

type Penduduk struct {
	ID     uint     `gorm:"primaryKey;autoIncrement;type:serial;column:id"`
	Alamat []Alamat `gorm:"many2many:penduduk_alamat"`
}

type Alamat struct {
	ID            uint   `gorm:"primaryKey;autoIncrement;type:serial;column:id"`
	AlamatLengkap string `gorm:"column:alamat_lengkap"`
}

func (Product) TableName() string {
	return "products"
}

func main() {
	connURI := "postgresql://postgres:password@localhost:5432/database?sslmode=disable"
	db, err := gorm.Open(postgres.Open(connURI), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Printf("Gagal menghubungkan database %v\n", err)
		os.Exit(1)
	}

	sqlDB, _ := db.DB()

	defer sqlDB.Close()

	fmt.Println("Database berhasil dihubungkan")

	db.AutoMigrate(&Product{}, &Alamat{}, &Penduduk{})

	fmt.Println("Table berhasil dibuat")

	// produk := Product{Name: "Kertas A4", Category: "Kertas", Price: 45000}

	// result := db.Create(&produk)

	// if result.Error != nil {
	// 	fmt.Printf("Gagal menambahkan data : %v\n", result.Error)
	// 	os.Exit(1)
	// }

	// fmt.Println("Data berhasil ditambahkan")

	// product := Product{ID: 2}

	// result := db.First(&product)

	// if result.Error != nil {
	// 	fmt.Printf("Data tidak ditemukan %v\n", result.Error)
	// }

	// fmt.Println(product)

	// var productSlice []Product

	// result := db.Where(map[string]interface{}{"id": 1}).Find(&productSlice)

	// if result.Error != nil {
	// 	fmt.Printf("Gagal menampilkan product %v\n", result.Error)
	// }

	// fmt.Println(productSlice)

	// result := db.Model(&Product{ID: 1}).Updates(&Product{Name: "Pensil"})

	// if result.Error != nil {
	// 	fmt.Printf("Product gagal di update %v\n", result.Error)
	// }

	// fmt.Println("Product berhasil di update")

	// result := db.Delete(Product{ID: 1})
	// if result.Error != nil {
	// 	fmt.Printf("Hapus data gagal %v\n", result.Error)
	// }

	// fmt.Println("Data berhasil di hapus")

	// db.Transaction(func(tx *gorm.DB) error {
	// 	if result := tx.Delete(&Product{ID: 1}); result.Error != nil {
	// 		fmt.Printf("Transaction gagal %v\n", result.Error)
	// 		return result.Error
	// 	} else {
	// 		fmt.Println("Transaction berhasil")
	// 		return nil
	// 	}
	// })

	//add data penduduk

	// penduduk := Penduduk{
	// 	Alamat: []Alamat{
	// 		{
	// 			AlamatLengkap: "Kota Jakarta",
	// 		},
	// 		{
	// 			AlamatLengkap: "Banjarbaru",
	// 		},
	// 	},
	// }

	// db.Create(&penduduk)

	var penduduk Penduduk

	db.Preload("Alamat").First(&penduduk, 1)

	fmt.Println(penduduk)
}
