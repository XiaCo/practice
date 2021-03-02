package data_structures

func GetBitMatrixMemCostByte(b *BitMatrix) uint32 {
	return uint32(b.width) * GetBitArrayMemCostByte(b.matrix[0])
}

type BitMatrix struct {
	length uint64
	width  uint64
	matrix []*BitArray
}

func (b *BitMatrix) SetBit1(x, y uint64) bool {
	if x > b.length || y > b.width {
		return false
	}

	return b.matrix[y].SetBit1(x)
}

func (b *BitMatrix) SetBit0(x, y uint64) bool {
	if x > b.length || y > b.width {
		return false
	}

	return b.matrix[y].SetBit0(x)
}

func (b *BitMatrix) GetBit(x, y uint64) bool {
	if x > b.length || y > b.width {
		return false
	}

	return b.matrix[y].GetBit(x)
}

func NewBitMatrix(length, width uint64) *BitMatrix {
	bm := BitMatrix{length: length, width: width}
	bm.matrix = make([]*BitArray, width)
	for i := uint64(0); i < width; i++ {
		bm.matrix[i] = NewBitArray(length)
	}

	return &bm
}
