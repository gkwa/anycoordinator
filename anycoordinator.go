package anycoordinator

import (
	"fmt"
	"log/slog"
)

func Main() int {
	slog.Debug("anycoordinator", "test", true)

	main()
	return 0
}

// Abstract Factory Interface
type ReportFactory interface {
	CreateIncomeStatement() IncomeStatement
	CreateBalanceSheet() BalanceSheet
}

// Concrete Factory A
type ConcreteFactoryA struct{}

func (c ConcreteFactoryA) CreateIncomeStatement() IncomeStatement {
	return IncomeStatementA{}
}

func (c ConcreteFactoryA) CreateBalanceSheet() BalanceSheet {
	return BalanceSheetA{}
}

// Concrete Factory B
type ConcreteFactoryB struct{}

func (c ConcreteFactoryB) CreateIncomeStatement() IncomeStatement {
	return IncomeStatementB{}
}

func (c ConcreteFactoryB) CreateBalanceSheet() BalanceSheet {
	return BalanceSheetB{}
}

// Abstract Product: Report
type Report interface {
	Generate()
}

// Concrete Products: IncomeStatement and BalanceSheet
type IncomeStatement interface {
	Report
	GenerateIncomeStatement()
}

type BalanceSheet interface {
	Report
	GenerateBalanceSheet()
}

// Concrete Income Statement A
type IncomeStatementA struct{}

func (i IncomeStatementA) Generate() {
	fmt.Println("Generating Income Statement A")
}

func (i IncomeStatementA) GenerateIncomeStatement() {
	fmt.Println("Generating detailed income statement A")
}

// Concrete Income Statement B
type IncomeStatementB struct{}

func (i IncomeStatementB) Generate() {
	fmt.Println("Generating Income Statement B")
}

func (i IncomeStatementB) GenerateIncomeStatement() {
	fmt.Println("Generating summarized income statement B")
}

// Concrete Balance Sheet A
type BalanceSheetA struct{}

func (b BalanceSheetA) Generate() {
	fmt.Println("Generating Balance Sheet A")
}

func (b BalanceSheetA) GenerateBalanceSheet() {
	fmt.Println("Generating detailed balance sheet A")
}

// Concrete Balance Sheet B
type BalanceSheetB struct{}

func (b BalanceSheetB) Generate() {
	fmt.Println("Generating Balance Sheet B")
}

func (b BalanceSheetB) GenerateBalanceSheet() {
	fmt.Println("Generating summarized balance sheet B")
}

func main() {
	factoryA := ConcreteFactoryA{}
	factoryB := ConcreteFactoryB{}

	incomeStatementA := factoryA.CreateIncomeStatement()
	balanceSheetA := factoryA.CreateBalanceSheet()

	incomeStatementB := factoryB.CreateIncomeStatement()
	balanceSheetB := factoryB.CreateBalanceSheet()

	fmt.Println("Factory A:")
	incomeStatementA.Generate()
	incomeStatementA.GenerateIncomeStatement()
	balanceSheetA.Generate()
	balanceSheetA.GenerateBalanceSheet()

	fmt.Println("Factory B:")
	incomeStatementB.Generate()
	incomeStatementB.GenerateIncomeStatement()
	balanceSheetB.Generate()
	balanceSheetB.GenerateBalanceSheet()
}
