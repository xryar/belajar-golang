package dao

type Number interface {
	int64 | float64
}

type UserModel[T int | float64] struct {
	Name   string
	Scores []T
}

func Sum[V int | float32 | float64](numbers []V) V {
	var total V
	for _, e := range numbers {
		total += e
	}
	return total
}

func SumNumbers1(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers2[k comparable, V int64 | float64](m map[k]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers3[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func (m *UserModel[int]) setScoresA(scores []int) {
	m.Scores = scores
}

func (m *UserModel[float64]) setScoresB(scores []float64) {
	m.Scores = scores
}

func (m *UserModel[float64]) SetScoresB(scores []float64) {
	m.Scores = scores
}
