package main

import (
	"fmt"
)

const MaxInvest = 100

type Stock struct {
	Symbol string
	Name   string
	Price  int
}

type Investment struct {
	Symbol   string
	Quantity int
	BuyPrice int
}

func InvestmentValue(i Investment, currentPrice int) int {
	return i.Quantity * currentPrice
}

func InvestmentProfitLoss(i Investment, currentPrice int) int {
	var currentValue int = InvestmentValue(i, currentPrice)
	var buyValue int = i.Quantity * i.BuyPrice
	return currentValue - buyValue
}

type CustomerWallet struct {
	Balance int
}

func WalletHasSufficientBalance(w CustomerWallet, amount int) bool {
	return w.Balance >= amount
}

func WalletWithdraw(w CustomerWallet, amount int) (CustomerWallet, error) {
	if amount < 0 {
		return w, fmt.Errorf("jumlah penarikan harus positif")
	}
	if w.Balance < amount {
		return w, fmt.Errorf("saldo tidak mencukupi")
	}
	w.Balance -= amount
	return w, nil
}

func WalletDeposit(w CustomerWallet, amount int) (CustomerWallet, error) {
	if amount < 0 {
		return w, fmt.Errorf("jumlah deposit harus positif")
	}
	w.Balance += amount
	return w, nil
}

func SelectionSort(investments []Investment, count int, getPriceFunc func(string) int, ascending bool) {
	var i, j, selectedIndex int
	for i = 0; i < count-1; i++ {
		selectedIndex = i
		for j = i + 1; j < count; j++ {
			var currentValue int = InvestmentValue(investments[j], getPriceFunc(investments[j].Symbol))
			var selectedValue int = InvestmentValue(investments[selectedIndex], getPriceFunc(investments[selectedIndex].Symbol))
			if ascending {
				if currentValue < selectedValue {
					selectedIndex = j
				}
			} else {
				if currentValue > selectedValue {
					selectedIndex = j
				}
			}
		}
		if selectedIndex != i {
			investments[i], investments[selectedIndex] = investments[selectedIndex], investments[i]
		}
	}
}

func InsertionSort(investments []Investment, count int, getPriceFunc func(string) int, ascending bool) {
	var i, j int
	var key Investment
	for i = 1; i < count; i++ {
		key = investments[i]
		j = i - 1
		for j >= 0 {
			var currentValue int = InvestmentValue(investments[j], getPriceFunc(investments[j].Symbol))
			var keyValue int = InvestmentValue(key, getPriceFunc(key.Symbol))
			if ascending {
				if currentValue > keyValue {
					investments[j+1] = investments[j]
				} else {
					break
				}
			} else {
				if currentValue < keyValue {
					investments[j+1] = investments[j]
				} else {
					break
				}
			}
			j--
		}
		investments[j+1] = key
	}
}

func SortBySymbol(investments []Investment, ascending bool) {
	var n int = len(investments)
	var i, j int
	for i = 0; i < n-1; i++ {
		for j = 0; j < n-1-i; j++ {
			var shouldSwap bool = false
			if ascending {
				if investments[j].Symbol > investments[j+1].Symbol {
					shouldSwap = true
				}
			} else {
				if investments[j].Symbol < investments[j+1].Symbol {
					shouldSwap = true
				}
			}
			if shouldSwap {
				investments[j], investments[j+1] = investments[j+1], investments[j]
			}
		}
	}
}

func SequentialSearch(symbol string, investments []Investment) int {
	var i int
	var investment Investment
	for i, investment = range investments {
		if investment.Symbol == symbol {
			return i
		}
	}
	return -1
}

func BinarySearchStock(symbol string, stockLists []Stock) int {
	var low int = 0
	var high int = len(stockLists) - 1
	for low <= high {
		var mid int = (low + high) / 2
		if stockLists[mid].Symbol == symbol {
			return mid
		} else if stockLists[mid].Symbol < symbol {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func BinarySearchInvestments(symbol string, investments []Investment) int {
	var low int = 0
	var high int = len(investments) - 1
	for low <= high {
		var mid int = (low + high) / 2
		if investments[mid].Symbol == symbol {
			return mid
		} else if investments[mid].Symbol < symbol {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func getStringInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}

func getIntInput(prompt string) (int, error) {
	var input int
	fmt.Print(prompt)
	var err error
	_, err = fmt.Scanln(&input)
	return input, err
}

func manualToUpper(s string) string {
	var r []rune = []rune(s)
	var i int
	var c rune
	for i, c = range r {
		if c >= 'a' && c <= 'z' {
			r[i] = c - ('a' - 'A')
		}
	}
	return string(r)
}

type StockService struct {
	Stocks []Stock
}

func NewStockService(stocks []Stock) StockService {
	var newStocks []Stock = make([]Stock, len(stocks))
	copy(newStocks, stocks)
	return StockService{
		Stocks: newStocks,
	}
}

func StockServiceGetPrice(s StockService, symbol string) int {
	var idx int = BinarySearchStock(symbol, s.Stocks)
	if idx == -1 {
		return 0
	}
	return s.Stocks[idx].Price
}

func StockServiceUpdatePrice(s StockService, symbol string, newPrice int) (StockService, bool) {
	fmt.Printf("Memperbarui harga untuk %s\n", symbol)
	var idx int = BinarySearchStock(symbol, s.Stocks)
	if idx == -1 {
		return s, false
	}
	s.Stocks[idx].Price = newPrice
	return s, true
}

type InvestmentService struct {
	Investments    []Investment
	Count          int
	MaxInvestments int
	StockService   StockService
	Wallet         CustomerWallet
}

func NewInvestmentService(maxInvestments int, stockService StockService, wallet CustomerWallet) InvestmentService {
	var investments []Investment = make([]Investment, 0, maxInvestments)
	return InvestmentService{
		Investments:    investments,
		Count:          0,
		MaxInvestments: maxInvestments,
		StockService:   stockService,
		Wallet:         wallet,
	}
}

func InvestmentServiceGetInvestments(s InvestmentService) []Investment {
	var investmentsCopy []Investment = make([]Investment, len(s.Investments))
	copy(investmentsCopy, s.Investments)
	return investmentsCopy
}

func InvestmentServiceGetCount(s InvestmentService) int {
	return s.Count
}

func InvestmentServiceGetWallet(s InvestmentService) CustomerWallet {
	return s.Wallet
}

func InvestmentServiceBuyStock(s InvestmentService, symbol string, quantity int) (InvestmentService, error) {
	var stockPrice int = StockServiceGetPrice(s.StockService, symbol)
	if stockPrice == 0 {
		return s, fmt.Errorf("saham tidak ditemukan")
	}

	var totalPrice int = stockPrice * quantity
	if !WalletHasSufficientBalance(s.Wallet, totalPrice) {
		return s, fmt.Errorf("saldo tidak mencukupi")
	}

	var invIdx int = SequentialSearch(symbol, s.Investments)
	if invIdx == -1 {
		if s.Count >= s.MaxInvestments {
			return s, fmt.Errorf("daftar investasi penuh")
		}
		var newInvestment Investment = Investment{Symbol: symbol, Quantity: quantity, BuyPrice: stockPrice}
		s.Investments = append(s.Investments, newInvestment)
		s.Count++
	} else {
		var existingInvestment Investment = s.Investments[invIdx]
		var totalQty int = existingInvestment.Quantity + quantity
		var newAvgPrice int
		if totalQty > 0 {
			newAvgPrice = ((existingInvestment.BuyPrice * existingInvestment.Quantity) + (stockPrice * quantity)) / totalQty
		} else {
			newAvgPrice = stockPrice
		}
		s.Investments[invIdx].Quantity = totalQty
		s.Investments[invIdx].BuyPrice = newAvgPrice
	}

	var updatedWallet CustomerWallet
	var errWithdraw error
	updatedWallet, errWithdraw = WalletWithdraw(s.Wallet, totalPrice)
	if errWithdraw != nil {
		return s, errWithdraw
	}
	s.Wallet = updatedWallet
	return s, nil
}

func InvestmentServiceSellStock(s InvestmentService, symbol string, quantity int) (InvestmentService, error) {
	var invIdx int = SequentialSearch(symbol, s.Investments)
	if invIdx == -1 {
		return s, fmt.Errorf("investasi tidak ditemukan")
	}

	if quantity > s.Investments[invIdx].Quantity {
		return s, fmt.Errorf("saham tidak mencukupi")
	}

	var currentPrice int = StockServiceGetPrice(s.StockService, symbol)
	var totalSellPrice int = currentPrice * quantity

	var updatedWallet CustomerWallet
	var errDeposit error
	updatedWallet, errDeposit = WalletDeposit(s.Wallet, totalSellPrice)
	if errDeposit != nil {
		return s, errDeposit
	}
	s.Wallet = updatedWallet

	s.Investments[invIdx].Quantity -= quantity

	if s.Investments[invIdx].Quantity == 0 {
		var newInvestments []Investment = make([]Investment, 0, len(s.Investments)-1)
		var i int
		for i = 0; i < len(s.Investments); i++ {
			if i != invIdx {
				newInvestments = append(newInvestments, s.Investments[i])
			}
		}
		s.Investments = newInvestments
		s.Count--
	}
	return s, nil
}

func InvestmentServiceSortInvestments(s InvestmentService, ascending bool, useSelection bool) InvestmentService {
	if InvestmentServiceGetCount(s) > 0 {
		var getPriceAdapter func(string) int
		getPriceAdapter = func(symbol string) int {
			return StockServiceGetPrice(s.StockService, symbol)
		}

		if useSelection {
			SelectionSort(s.Investments, InvestmentServiceGetCount(s), getPriceAdapter, ascending)
		} else {
			InsertionSort(s.Investments, InvestmentServiceGetCount(s), getPriceAdapter, ascending)
		}
	}
	return s
}

var defaultStocks = []Stock{
	{Symbol: "ASII", Name: "Astra International", Price: 6700},
	{Symbol: "BBCA", Name: "Bank Central Asia", Price: 900000},
	{Symbol: "BMRI", Name: "Bank Mandiri", Price: 7800},
	{Symbol: "TLKM", Name: "Telekomunikasi Indonesia", Price: 3500},
	{Symbol: "UNVR", Name: "Unilever Indonesia", Price: 42000},
}

func initializeWallet() CustomerWallet {
	var wallet CustomerWallet
	for {
		var balance int
		var err error
		balance, err = getIntInput("Masukkan saldo awal anda: ")
		if err == nil && balance >= 0 {
			var tempWallet CustomerWallet = wallet
			var depositErr error
			tempWallet, depositErr = WalletDeposit(tempWallet, balance)
			if depositErr == nil {
				wallet = tempWallet
				fmt.Printf("Init saldo adalah %d IDR.\n", wallet.Balance)
				break
			} else {
				fmt.Printf("Kesalahan saat deposit: %v\n", depositErr)
			}
		}
		fmt.Println("Input tidak sesuai format atau terjadi kesalahan input.")
	}
	return wallet
}

func displayStocks(stockService StockService) {
	fmt.Println("\n--- List saham ---")
	fmt.Printf("%-6s %-25s %-12s\n", "Symbol", "Nama", "Harga (IDR)")
	var stock Stock
	for _, stock = range stockService.Stocks {
		fmt.Printf("%-6s %-25s %-12d\n", stock.Symbol, stock.Name, stock.Price)
	}
}

func updateStockPrice(stockService StockService) StockService {
	displayStocks(stockService)
	var symbolRaw string = getStringInput("Masukkan Symbol saham: ")
	var symbol string = manualToUpper(symbolRaw)

	var price int
	var err error
	price, err = getIntInput("Masukkan Harga baru dalam Rupiah: ")
	if err != nil || price <= 0 {
		fmt.Println("Harga tidak valid atau terjadi kesalahan input.")
		return stockService
	}

	var success bool
	var updatedStockService StockService
	updatedStockService, success = StockServiceUpdatePrice(stockService, symbol, price)
	if !success {
		fmt.Println("Emiten tidak ditemukan.")
		return stockService
	}
	fmt.Println("Harga berhasil di update.")
	return updatedStockService
}

func buyStock(investmentService InvestmentService, stockService StockService) InvestmentService {
	displayStocks(stockService)
	var symbolRaw string = getStringInput("Masukkan Symbol saham: ")
	var symbol string = manualToUpper(symbolRaw)

	var qty int
	var err error
	qty, err = getIntInput("Jumlah lot saham: ")
	if err != nil || qty <= 0 {
		fmt.Println("Input tidak valid atau terjadi kesalahan input.")
		return investmentService
	}

	var updatedInvestmentService InvestmentService
	var errBuy error
	updatedInvestmentService, errBuy = InvestmentServiceBuyStock(investmentService, symbol, qty)
	if errBuy != nil {
		fmt.Printf("Kesalahan: %v\n", errBuy)
		return investmentService
	}
	fmt.Println("Saham berhasil di beli!")
	return updatedInvestmentService
}

func sellStock(investmentService InvestmentService) InvestmentService {
	var investmentsToDisplay []Investment = InvestmentServiceGetInvestments(investmentService)
	if len(investmentsToDisplay) == 0 {
		fmt.Println("\nTidak ada investasi untuk dijual.")
		return investmentService
	}
	fmt.Println("\n--- Portfolio Anda (untuk dijual) ---")
	fmt.Printf("%-6s %-10s\n", "Symbol", "Quantity")
	var invDisplay Investment
	for _, invDisplay = range investmentsToDisplay {
		fmt.Printf("%-6s %-10d\n", invDisplay.Symbol, invDisplay.Quantity)
	}

	var symbolRaw string = getStringInput("Masukkan Symbol saham yang akan dijual: ")
	var symbol string = manualToUpper(symbolRaw)

	var qty int
	var err error
	qty, err = getIntInput("Jumlah lot saham yang akan dijual: ")
	if err != nil || qty <= 0 {
		fmt.Println("Input tidak valid atau terjadi kesalahan input.")
		return investmentService
	}

	var updatedInvestmentService InvestmentService
	var errSell error
	updatedInvestmentService, errSell = InvestmentServiceSellStock(investmentService, symbol, qty)
	if errSell != nil {
		fmt.Printf("Kesalahan: %v\n", errSell)
		return investmentService
	}
	fmt.Println("Saham berhasil dijual!")
	return updatedInvestmentService
}

func displayInvestments(investmentService InvestmentService) {
	if InvestmentServiceGetCount(investmentService) == 0 {
		fmt.Println("\nTidak ada investasi.")
		fmt.Printf("\nSaldo Wallet: %d IDR.\n", InvestmentServiceGetWallet(investmentService).Balance)
		return
	}

	fmt.Println("\n--- Portfolio Anda ---")
	fmt.Printf("%-6s %-10s %-15s %-15s %-15s\n", "Symbol", "Quantity", "Avg Buy Price", "Current Price", "Profit/Loss (IDR)")

	var investments []Investment = InvestmentServiceGetInvestments(investmentService)
	var inv Investment
	for _, inv = range investments {
		var currentPrice int = StockServiceGetPrice(investmentService.StockService, inv.Symbol)
		var profitLoss int = InvestmentProfitLoss(inv, currentPrice)
		fmt.Printf("%-6s %-10d %-15d %-15d %-15d\n",
			inv.Symbol, inv.Quantity, inv.BuyPrice, currentPrice, profitLoss)
	}
	fmt.Printf("\nSaldo Wallet: %d IDR.\n", InvestmentServiceGetWallet(investmentService).Balance)
}

func sortInvestments(investmentService InvestmentService) InvestmentService {
	if InvestmentServiceGetCount(investmentService) == 0 {
		fmt.Println("Tidak ada investasi untuk diurutkan.")
		return investmentService
	}
	fmt.Println("Pilih algoritma pengurutan:")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	var algChoiceStr string = getStringInput("Masukkan pilihan (1 atau 2): ")

	fmt.Println("Pilih urutan:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	var orderChoiceStr string = getStringInput("Masukkan pilihan (1 atau 2): ")

	var ascending bool = orderChoiceStr != "2"
	var useSelection bool = algChoiceStr == "1"

	var sortedInvestmentService InvestmentService
	sortedInvestmentService = InvestmentServiceSortInvestments(investmentService, ascending, useSelection)
	fmt.Println("Portfolio berhasil diurutkan.")
	return sortedInvestmentService
}

func searchInvestment(investmentService InvestmentService) {
	if InvestmentServiceGetCount(investmentService) == 0 {
		fmt.Println("\nTidak ada investasi untuk dicari.")
		return
	}
	var symbolRaw string = getStringInput("Masukkan Symbol saham yang dicari: ")
	var symbol string = manualToUpper(symbolRaw)

	fmt.Println("Select search method:")
	fmt.Println("1. Linear Search")
	fmt.Println("2. Binary Search (berdasarkan simbol, portfolio akan disalin dan diurutkan sementara)")
	var searchMethodChoice string = getStringInput("Masukkan pilihan (1 atau 2): ")

	var investments []Investment = InvestmentServiceGetInvestments(investmentService)
	var idx int = -1

	if searchMethodChoice == "1" {
		idx = SequentialSearch(symbol, investments)
	} else if searchMethodChoice == "2" {
		if len(investments) == 0 {
			fmt.Println("Tidak ada investasi untuk dicari dengan binary search.")
			return
		}
		var tempInvestments []Investment = make([]Investment, len(investments))
		copy(tempInvestments, investments)
		SortBySymbol(tempInvestments, true)
		var foundInTempIdx int = BinarySearchInvestments(symbol, tempInvestments)

		if foundInTempIdx != -1 {
			idx = SequentialSearch(symbol, investments)
		}
	} else {
		fmt.Println("Pilihan metode pencarian tidak valid.")
		return
	}

	if idx != -1 {
		var inv Investment = investments[idx]
		var currentPrice int = StockServiceGetPrice(investmentService.StockService, inv.Symbol)
		fmt.Printf("\nDitemukan investasi: %s\n", inv.Symbol)
		fmt.Printf("Jumlah: %d\n", inv.Quantity)
		fmt.Printf("Harga Beli Rata-rata: %d IDR\n", inv.BuyPrice)
		fmt.Printf("Nilai Saat Ini: %d IDR\n", InvestmentValue(inv, currentPrice))
		fmt.Printf("Untung/Rugi: %d IDR\n", InvestmentProfitLoss(inv, currentPrice))
	} else {
		fmt.Println("Investasi tidak ditemukan.")
	}
}

func main() {
	var wallet CustomerWallet = initializeWallet()
	var initialStockService StockService = NewStockService(defaultStocks)
	var investmentService InvestmentService = NewInvestmentService(MaxInvest, initialStockService, wallet)
	var currentStockService StockService = initialStockService

	for {
		fmt.Println("\n=== Management Investasi Sederhana ===")
		fmt.Println("1. Tampilkan List Saham")
		fmt.Println("2. Update Harga Saham")
		fmt.Println("3. Beli Saham")
		fmt.Println("4. Jual Saham")
		fmt.Println("5. Tampilkan Portfolio")
		fmt.Println("6. Urutkan Portfolio (berdasarkan nilai saat ini)")
		fmt.Println("7. Cari Portfolio (berdasarkan simbol)")
		fmt.Println("8. Keluar")

		var choice string = getStringInput("Masukkan Pilihan anda (1-8): ")

		switch choice {
		case "1":
			displayStocks(currentStockService)
		case "2":
			currentStockService = updateStockPrice(currentStockService)
			investmentService.StockService = currentStockService
		case "3":
			investmentService.StockService = currentStockService
			investmentService = buyStock(investmentService, currentStockService)
		case "4":
			investmentService.StockService = currentStockService
			investmentService = sellStock(investmentService)
		case "5":
			investmentService.StockService = currentStockService
			displayInvestments(investmentService)
		case "6":
			investmentService.StockService = currentStockService
			investmentService = sortInvestments(investmentService)
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
