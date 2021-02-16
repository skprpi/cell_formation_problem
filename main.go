package main

import (
	clusterPack "factory/claster"
	dataPack "factory/data"
	detailPack "factory/details"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Hello")
	detailArr, detailSize, machineSize := dataPack.ReadDataAndFormPrimaryArr()
	detailArr = detailPack.FormDetailWithoutRepetition(detailArr)
	clusters := clusterPack.CreateClusters(detailArr)

	//clusters = clusterPack.ValidateCluster(clusters)

	fmt.Println(len(clusters), clusters)
	maxSizeClusters := len(clusters)

	//for {
	//	clusters = clusterPack.JoinClusters(clusters)
	//	if len(clusters) == 17{
	//		break
	//	}
	//}

	bestAnsCluster := clusterPack.CopyCluster(clusters)

	bestAnsClusterArr := make([][]clusterPack.Cluster, maxSizeClusters+1)
	for i := 1; i <= maxSizeClusters; i++ {
		bestAnsClusterArr[i] = make([]clusterPack.Cluster, i)
	}
	bestAnsClusterArr[maxSizeClusters] = clusterPack.CopyCluster(clusters)

	bestCluster := clusterPack.CopyCluster(clusters)

	for {
		for ind := 0; ind < 5; ind++ {
			index := 0
			for k := 0; k < 1000; k++ {
				if len(clusters) <= 1 {
					break
				}
				chanse := rand.Intn(1000)
				if chanse <= 400 {
					clusterPack.TransferOneDetail(clusters)
				} else if chanse <= 800 {
					clusterPack.TransferOneMachine(clusters)
				} else if chanse <= 900 {
					clusterPack.SwapMachine(clusters)
				} else {
					clusterPack.SwapDetails(clusters)
				}

				if clusterPack.FindCosts(clusters) > clusterPack.FindCosts(bestCluster)+0.00000000000001 {
					bestCluster = clusterPack.CopyCluster(clusters)
					k = 0
				}
				clusters = clusterPack.CopyCluster(bestCluster)

				index += 1
			}
			if clusterPack.FindCosts(bestCluster) > clusterPack.FindCosts(bestAnsClusterArr[len(bestCluster)])+0.00000000000001 || len(bestAnsClusterArr[len(bestCluster)][0].DetailArr) < 1 {
				bestAnsClusterArr[len(bestCluster)] = clusterPack.CopyCluster(bestCluster)
			}
			if clusterPack.FindCosts(bestCluster) > clusterPack.FindCosts(bestAnsCluster)+0.00000000000001 {
				fmt.Println(clusterPack.FindCosts(bestCluster), "--------------", len(bestCluster))
				clusterPack.ShowAnswer(bestCluster)
				clusterPack.WriteAnswerInFile(bestCluster, detailSize, machineSize)
				bestAnsCluster = clusterPack.CopyCluster(bestCluster)
			}
			for k := 0; k < 2; k++ {
				if k%2 == 0 {
					clusterPack.ThrowManyMachine(bestCluster)
				} else {
					clusterPack.ThrowManyDetails(bestCluster)
				}
			}
			//if len(bestAnsClusterArr[len(bestCluster)][0].DetailArr) >= 1 {
			//	bestCluster = clusterPack.CopyCluster(bestAnsClusterArr[len(bestCluster)])
			//}
			clusters = clusterPack.CopyCluster(bestCluster)
		}
		//fmt.Println("len ---------------- ", len(bestCluster), clusterPack.FindCosts(bestAnsClusterArr[len(bestCluster)]))
		chance := rand.Intn(1000)
		if chance < 500 {
			if chance < 250 || len(bestAnsClusterArr[len(bestCluster)][0].DetailArr) < 1 {
				bestCluster = clusterPack.CopyCluster(bestAnsCluster)
			} else {
				bestCluster = clusterPack.CopyCluster(bestAnsClusterArr[len(bestCluster)])
			}
		} else {
			if chance < 750 {
				bestCluster = clusterPack.SplitClusters(bestCluster, maxSizeClusters)
			} else {
				bestCluster = clusterPack.JoinClusters(bestCluster)
			}
		}
		clusters = clusterPack.CopyCluster(bestCluster)
		//fmt.Println(count)

	}
}
