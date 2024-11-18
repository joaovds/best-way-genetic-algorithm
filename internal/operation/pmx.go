package operation

import (
	"errors"
	"math/rand"
	"time"

	"github.com/joaovds/best-way-genetic-algorithm/internal/core"
)

type PMX struct {
	Parent1, Parent2 *core.Chromosome
	StartPoint       int
	EndPoint         int
	Size             int
}

func NewPMX(parent1, parent2 *core.Chromosome) *PMX {
	chromosomeSize := len(parent1.Genes)
	if len(parent1.Genes) != len(parent2.Genes) {
		panic(errors.New("The chromosomes must contain the same number of genes"))
	}

	startPoint, endPoint := getCrossoverPoints(chromosomeSize)
	return &PMX{parent1, parent2, startPoint, endPoint, chromosomeSize}
}

func (pmx *PMX) Run() [2]*core.Chromosome {
	child1Genes := make([]*core.Gene, pmx.Size)
	child2Genes := make([]*core.Gene, pmx.Size)
	for i := pmx.StartPoint; i <= pmx.EndPoint; i++ {
		child1Genes[i] = pmx.Parent2.Genes[i]
		child2Genes[i] = pmx.Parent1.Genes[i]
	}

	fillRemainingGenes(child1Genes, pmx.Parent1.Genes, pmx.StartPoint, pmx.EndPoint)
	fillRemainingGenes(child2Genes, pmx.Parent2.Genes, pmx.StartPoint, pmx.EndPoint)

	children := [2]*core.Chromosome{
		core.NewChromosome(core.NewGene(pmx.Parent1.StartingPoint.GetID(), pmx.Parent1.StartingPoint.Address), child1Genes),
		core.NewChromosome(core.NewGene(pmx.Parent1.StartingPoint.GetID(), pmx.Parent1.StartingPoint.Address), child2Genes),
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
