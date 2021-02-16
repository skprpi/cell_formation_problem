package main

import (
	clusterPack "factory/claster"
	dataPack "factory/data"
	detailPack "factory/details"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	detailArr, detailSize, machineSize := dataPack.ReadDataAndFormPrimaryArr()
	detailArr = detailPack.FormDetailWithoutRepetition(detailArr)
	clusters := clusterPack.CreateClusters(detailArr)

	maxSizeClusters := len(clusters)
	bestAnsCluster := clusterPack.CopyCluster(clusters)
	bestAnsClusterArr := clusterPack.MakeBestAnsClusters(maxSizeClusters, clusters)

	bestCluster := clusterPack.CopyCluster(clusters)
	for {
		for ind := 0; ind < 3; ind++ {
			for k := 0; k < 5000; k++ {
				if len(clusters) <= 1 {
					break
				}
				clusterPack.MakeChange(clusters)
				bestCluster, k = clusterPack.CheckBestCluster(clusters, bestCluster, k)
				clusters = clusterPack.CopyCluster(bestCluster)
			}
			bestAnsClusterArr, bestAnsCluster = clusterPack.RememberBestAns(
				bestAnsClusterArr, bestCluster, bestAnsCluster, detailSize, machineSize)
			bestCluster = clusterPack.Perturbation(bestCluster)
			clusters = clusterPack.CopyCluster(bestCluster)
		}
		bestCluster = clusterPack.Shaking(bestAnsClusterArr, bestCluster, bestAnsCluster, maxSizeClusters)
		clusters = clusterPack.CopyCluster(bestCluster)
	}
}
