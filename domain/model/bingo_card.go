package model

import (
	"github.com/samber/lo"
	"math/rand"
)

type BingoRow struct {
	bingoNumber uint
	isMarked    bool
}

type BingoCard struct {
	bingoSeed int64
	card      [5][5]BingoRow
}

func NewBingoCard(bingoSeed int64) *BingoCard {
	return &BingoCard{
		bingoSeed: bingoSeed,
		card:      GenerateCard(bingoSeed),
	}
}

func RebuildBingoCard(bingoSeed int64, bingoNumbers []uint) *BingoCard {
	bc := &BingoCard{
		bingoSeed: bingoSeed,
		card:      GenerateCard(bingoSeed),
	}
	bc.UpdateBingoCard(bingoNumbers)

	return bc
}

func GenerateCard(bingoSeed int64) [5][5]BingoRow {
	var card [5][5]BingoRow
	bingoNums := lo.RangeFrom(1, 75)
	random := rand.New(rand.NewSource(bingoSeed))

	for i := 0; i < 5; i++ {
		bingoNumsCopy := lo.Slice(bingoNums, i*15, (i+1)*15)
		for j := 0; j < 5; j++ {
			idx := random.Intn(len(bingoNumsCopy))
			card[i][j] = BingoRow{
				bingoNumber: uint(bingoNumsCopy[idx]),
				isMarked:    false,
			}
			bingoNumsCopy = append(bingoNumsCopy[:idx], bingoNumsCopy[idx+1:]...)
		}
	}
	// 中心部だけは最初からマーク済みにする
	card[2][2].isMarked = true

	return card
}

func isBingoLine(bingoLine []BingoRow) bool {
	return lo.EveryBy(bingoLine, func(bingoRow BingoRow) bool {
		return bingoRow.isMarked
	})
}

func (bc *BingoCard) checkVerticalBingo() bool {
	for i := 0; i < 5; i++ {
		if isBingoLine(bc.card[i][:]) {
			return true
		}
	}
	return false
}

func (bc *BingoCard) checkHorizontalBingo() bool {
	for i := 0; i < 5; i++ {
		var bingoLine []BingoRow
		for j := 0; j < 5; j++ {
			bingoLine = append(bingoLine, bc.card[j][i])
		}
		if isBingoLine(bingoLine) {
			return true
		}
	}
	return false
}

func (bc *BingoCard) checkDiagonalBingo() bool {
	var diagonalBingoLineIdxArr [2][]int
	lo.ForEach(lo.Range(len(bc.card)), func(_ int, i int) {
		diagonalBingoLineIdxArr[0] = append(diagonalBingoLineIdxArr[0], i)
		diagonalBingoLineIdxArr[1] = append(diagonalBingoLineIdxArr[1], len(bc.card)-i-1)
	})
	for i, _ := range diagonalBingoLineIdxArr {
		diagonalLine := lo.Map(diagonalBingoLineIdxArr[i], func(idx1 int, idx2 int) BingoRow {
			return bc.card[idx1][idx2]
		})
		if isBingoLine(diagonalLine) {
			return true
		}
	}

	return false
}

func (bc *BingoCard) CheckBingo() bool {
	return bc.checkVerticalBingo() || bc.checkHorizontalBingo() || bc.checkDiagonalBingo()
}

func (bc *BingoCard) UpdateBingoCard(bingoNumbers []uint) {
	for i, v := range bc.card {
		for j, _ := range v {
			if lo.Contains(bingoNumbers, bc.card[i][j].bingoNumber) {
				bc.card[i][j].isMarked = true
			}
		}
	}
}
