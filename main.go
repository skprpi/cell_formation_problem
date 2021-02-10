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
	detailArr := dataPack.ReadDataAndFormPrimaryArr()
	detailArr = detailPack.FormDetailWithoutRepetition(detailArr)
	clusters := clusterPack.CreateClusters(detailArr)

	//clusters = clusterPack.ValidateCluster(clusters)

	fmt.Println(len(clusters), clusters)
	maxSizeClusters := len(clusters)
	//for {
	//	clusters = clusterPack.JoinClusters(clusters)
	//	if len(clusters) == 14 {
	//		break
	//	}
	//}
	bestAnsCluster := clusterPack.CopyCluster(clusters)
	bestCluster := clusterPack.CopyCluster(clusters)
	count := 0
	count2 := 0
	for {
		index := 0
		for k := 0; k < 2000; k++ {
			if len(clusters) <= 1 {
				break
			}
			chanse := rand.Intn(1000)
			if chanse <= 250 {
				clusterPack.TransferOneDetail(clusters)
			} else if chanse <= 500 {
				clusterPack.TransferOneMachine(clusters)
			} else if chanse <= 750 {
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
		if clusterPack.FindCosts(bestCluster) > clusterPack.FindCosts(bestAnsCluster)+0.00000000000001 {
			fmt.Println(clusterPack.FindCosts(bestCluster), "--------------", len(bestCluster))
			clusterPack.ShowAnswer(bestCluster)
			bestAnsCluster = clusterPack.CopyCluster(bestCluster)
		}
		chance := rand.Intn(1000)
		if chance < 500 {
			bestCluster = clusterPack.CopyCluster(bestAnsCluster)
			if count%2 == 0 {
				clusterPack.ThrowManyMachine(bestCluster)
			} else {
				clusterPack.ThrowManyDetails(bestCluster)
			}
			count2 += 1
		} else {
			if count%2 == 0 {
				bestCluster = clusterPack.JoinClusters(bestCluster)
			} else {
				bestCluster = clusterPack.SplitClusters(bestCluster, maxSizeClusters)
			}
			count += 1
		}
		clusters = clusterPack.CopyCluster(bestCluster)
		//fmt.Println(count)

	}
	fmt.Println(len(bestAnsCluster))
	fmt.Println(clusterPack.FindCosts(bestAnsCluster))
	fmt.Println(bestAnsCluster)
}
