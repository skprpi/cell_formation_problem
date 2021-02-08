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

	for {
		clusters = clusterPack.JoinClusters(clusters)
		if len(clusters) == 10 {
			break
		}
	}
	bestAnsCluster := clusterPack.CopyCluster(clusters)
	bestCluster := clusterPack.CopyCluster(clusters)
	count := 0
	for {
		index := 0
		start := time.Now()
		for {
			if len(clusters) <= 1 {
				break
			}

			if index%4 == 0 {
				clusterPack.TransferOneDetail(clusters)
			} else if index%4 == 1 {
				clusterPack.TransferOneMachine(clusters)
			} else if index%4 == 2 {
				clusterPack.SwapMachine(clusters)
			} else {
				clusterPack.SwapDetails(clusters)
			}

			if clusterPack.FindCosts(clusters) > clusterPack.FindCosts(bestCluster)+0.00000000000001 {
				bestCluster = clusterPack.CopyCluster(clusters)
				start = time.Now()
			}

			clusters = clusterPack.CopyCluster(bestCluster)

			if time.Since(start) >= time.Second*1 {
				break
			}

			index += 1
		}
		if clusterPack.FindCosts(bestCluster) > clusterPack.FindCosts(bestAnsCluster)+0.00000000000001 {
			fmt.Println(clusterPack.FindCosts(bestCluster))
			clusterPack.ShowAnswer(bestCluster)
			bestAnsCluster = clusterPack.CopyCluster(bestCluster)
		}
		if count%2 == 0 {
			bestCluster = clusterPack.CopyCluster(bestAnsCluster)
		} else {
			for i := 0; i < 50; i++ {
				clusterPack.ThrowManyDetails(bestCluster)
				clusterPack.ThrowManyMachine(bestCluster)
			}
		}
		clusters = clusterPack.CopyCluster(bestCluster)
		//fmt.Println(count)
		count += 1
	}
	fmt.Println(len(bestAnsCluster))
	fmt.Println(clusterPack.FindCosts(bestAnsCluster))
	fmt.Println(bestAnsCluster)
}
