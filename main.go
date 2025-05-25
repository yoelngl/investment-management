package main

import (
	"fmt"
	"managementInvestment/internal/service"
	"managementInvestment/models"
	"managementInvestment/pkg/input"
	"strconv"
	"strings"
)

const MaxInvest = 100

var defaultStocks = []models.Stock{
	{Symbol: "BBCA", Name: "Bank Central Asia", Price: 900000},
	{Symbol: "TLKM", Name: "Telekomunikasi Indonesia", Price: 3500},
	{Symbol: "BMRI", Name: "Bank Mandiri", Price: 7800},
	{Symbol: "UNVR", Name: "Unilever Indonesia", Price: 42000},
	{Symbol: "ASII", Name: "Astra International", Price: 6700},
}

func main() {
	wallet := &models.CustomerWallet{}
	stockService := service.NewStockService(defaultStocks)
	investmentService := service.NewInvestmentService(MaxInvest, stockService, wallet)

	initializeWallet(wallet)

	for {
		fmt.Println("\n=== Management Investasi Sederhana ===")
		fmt.Println("1. Tampilkan List Saham")
		fmt.Println("2. Update Harga Saham")
		fmt.Println("3. Beli Saham")
		fmt.Println("4. Jual Saham")
		fmt.Println("5. Tampilkan Portfolio")
		fmt.Println("6. Urutkan Portfolio")
		fmt.Println("7. Cari Portfolio")
		fmt.Println("8. Keluar")

		choice := input.GetInput("Masukkan Pilihan anda (1-8): ")

		switch choice {
		case "1":
			displayStocks(stockService)
		case "2":
			updateStockPrice(stockService)
		case "3":
			buyStock(investmentService, stockService)
		case "4":
			sellStock(investmentService)
		case "5":
			displayInvestments(investmentService)
		case "6":
			sortInvestments(investmentService)
		case "7":
			searchInvestment(investmentService)
		case "8":
			fmt.Println("Terimakasih sudah menggunakan Management Investasi Sederhana")
			return
		default:
			fmt.Println("Pilihan tidak sesuai!")
		}
	}
}

func initializeWallet(wallet *models.CustomerWallet) {
	for {
		balanceInput := input.GetInput("Masukkan saldo awal anda: ")
		balance, err := strconv.Atoi(balanceInput)
		if err == nil && balance >= 0 {
			wallet.Deposit(balance)
			fmt.Printf("Init saldo adalah %d IDR.\n", balance)
			break
		}
		fmt.Println("Input tidak sesuai format.")
	}
}

func displayStocks(stockService *service.StockService) {
	fmt.Println("\n--- List saham ---")
	fmt.Printf("%-6s %-25s %-12s\n", "Symbol", "Nama", "Harga (IDR)")
	for _, stock := range stockService.Stocks {
		fmt.Printf("%-6s %-25s %-12d\n", stock.Symbol, stock.Name, stock.Price)
	}
}

func updateStockPrice(stockService *service.StockService) {
	displayStocks(stockService)
	symbol := strings.ToUpper(input.GetInput("Masukkan Symbol saham: "))
	priceStr := input.GetInput("Masukkan Harga baru dalam Rupiah: ")
	price, err := strconv.Atoi(priceStr)
	if err != nil || price <= 0 {
		fmt.Println("Harga tidak valid.")
		return
	}
	if !stockService.UpdatePrice(symbol, price) {
		fmt.Println("Emiten tidak ditemukan.")
		return
	}
	fmt.Println("Harga berhasil di update.")
}

func buyStock(investmentService *service.InvestmentService, stockService *service.StockService) {
	displayStocks(stockService)
	symbol := strings.ToUpper(input.GetInput("Masukkan Symbol saham: "))
	qtyStr := input.GetInput("Jumlah lot saham: ")
	qty, err := strconv.Atoi(qtyStr)
	if err != nil || qty <= 0 {
		fmt.Println("Input tidak valid.")
		return
	}

	if err := investmentService.BuyStock(symbol, qty); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Saham berhasil di beli!")
}

func sellStock(investmentService *service.InvestmentService) {
	symbol := strings.ToUpper(input.GetInput("Masukkan Symbol saham: "))
	qtyStr := input.GetInput("Jumlah lot saham yang akan dijual: ")
	qty, err := strconv.Atoi(qtyStr)
	if err != nil || qty <= 0 {
		fmt.Println("Input tidak valid.")
		return
	}

	if err := investmentService.SellStock(symbol, qty); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Saham berhasil dijual!")
}

func displayInvestments(investmentService *service.InvestmentService) {
	if investmentService.GetCount() == 0 {
		fmt.Println("\nTidak ada investasi.")
		return
	}

	fmt.Println("\n--- Portfolio Anda ---")
	fmt.Printf("%-6s %-10s %-15s %-15s %-15s\n", "Symbol", "Qty", "Avg Buy Price", "Current Price", "Profit/Loss (IDR)")

	investments := investmentService.GetInvestments()
	for _, inv := range investments {
		currentPrice := investmentService.StockService.GetPrice(inv.Symbol)
		profitLoss := inv.ProfitLoss(currentPrice)
		fmt.Printf("%-6s %-10d %-15d %-15d %-15d\n",
			inv.Symbol, inv.Quantity, inv.BuyPrice, currentPrice, profitLoss)
	}

	fmt.Printf("\nSaldo Wallet: %d IDR\n", investmentService.GetWallet().Balance)
}

func sortInvestments(investmentService *service.InvestmentService) {
	fmt.Println("Pilih algoritma pengurutan:")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	algChoice := input.GetInput("Masukkan pilihan (1 atau 2): ")

	fmt.Println("Pilih urutan:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	orderChoice := input.GetInput("Masukkan pilihan (1 atau 2): ")

	ascending := orderChoice != "2"
	useSelection := algChoice == "1"

	investmentService.SortInvestments(ascending, useSelection)
	fmt.Println("Portfolio berhasil diurutkan.")
}

func searchInvestment(investmentService *service.InvestmentService) {
	symbol := strings.ToUpper(input.GetInput("Masukkan Symbol saham yang dicari: "))
	investments := investmentService.GetInvestments()

	found := false
	for _, inv := range investments {
		if inv.Symbol == symbol {
			currentPrice := investmentService.StockService.GetPrice(symbol)
			fmt.Printf("Ditemukan investasi: %s\n", inv.Symbol)
			fmt.Printf("Jumlah: %d\n", inv.Quantity)
			fmt.Printf("Harga Beli Rata-rata: %d IDR\n", inv.BuyPrice)
			fmt.Printf("Nilai Saat Ini: %d IDR\n", inv.Value(currentPrice))
			fmt.Printf("Profit/Loss: %d IDR\n", inv.ProfitLoss(currentPrice))
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Investasi tidak ditemukan.")
	}
}
