package operation

import (
	"errors"
	"math/rand"
	"time"

	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
)

type PMX struct {
	StartPoint, EndPoint int
}

func NewPMX() *PMX {
	return &PMX{}
}

func (pmx *PMX) Run(parent1, parent2 *core.Chromosome) [2]*core.Chromosome {
	size := len(parent1.Genes)
	if len(parent1.Genes) != len(parent2.Genes) {
		panic(errors.New("The chromosomes must contain the same number of genes"))
	}

	if pmx.StartPoint == 0 && pmx.EndPoint == 0 {
		pmx.StartPoint, pmx.EndPoint = getCrossoverPoints(size)
	}

	child1Genes := make([]*core.Gene, size)
	child2Genes := make([]*core.Gene, size)
	for i := pmx.StartPoint; i <= pmx.EndPoint; i++ {
		child1Genes[i] = parent2.Genes[i]
		child2Genes[i] = parent1.Genes[i]
	}

	fillRemainingGenes(child1Genes, parent1.Genes, pmx.StartPoint, pmx.EndPoint)
	fillRemainingGenes(child2Genes, parent2.Genes, pmx.StartPoint, pmx.EndPoint)

	children := [2]*core.Chromosome{
		core.NewChromosome(core.NewGene(parent1.StartingPoint.GetID(), parent1.StartingPoint.Address), child1Genes),
		core.NewChromosome(core.NewGene(parent1.StartingPoint.GetID(), parent1.StartingPoint.Address), child2Genes),
	}
	return children
}

func getCrossoverPoints(size int) (startPoint, endPoint int) {
	randSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randSource)
	startPoint = rnd.Intn(size)
	endPoint = rnd.Intn(size - 1)
	if endPoint >= startPoint {
		endPoint++
	} else {
		startPoint, endPoint = endPoint, startPoint
	}
	return
}

func fillRemainingGenes(childGenes, parentGenes []*core.Gene, startPoint, endPoint int) {
	chromosomeSize := len(parentGenes)
	for i := range chromosomeSize {
		if i >= startPoint && i <= endPoint {
			continue
		}

		gene := parentGenes[i]
		for contains(childGenes, gene) {
			gene = parentGenes[indexOfGene(childGenes, gene)]
		}
		childGenes[i] = core.NewGene(gene.GetID(), gene.Address)
	}
}

func contains(genes []*core.Gene, gene *core.Gene) bool {
	for _, g := range genes {
		if g == nil {
			continue
		}

		if g.GetID() == gene.GetID() {
			return true
		}
	}
	return false
}

func indexOfGene(genes []*core.Gene, gene *core.Gene) int {
	for i, g := range genes {
		if g == nil {
			continue
		}
		if g.GetID() == gene.GetID() {
			return i
		}
	}
	return -1
}
