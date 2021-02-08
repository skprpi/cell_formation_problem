package cluster

import (
	detailsPack "factory/details"
	"fmt"
	"math/rand"
)

type Cluster struct {
	DetailArr  []detailsPack.Detals
	MachineArr []bool
}

//func ValidateCluster(clusters []Cluster) []Cluster{
//	for {
//		count := 0
//		for _, el := range clusters {
//			if len(el.DetailArr) <= 0 {
//				clusters = JoinClusters(clusters, false)
//				count += 1
//				fmt.Println("Validated !!!!")
//				break
//			}
//		}
//		if count <= 0{
//			return clusters
//		}
//	}
//}

func CreateClusters(detailArr []detailsPack.Detals) []Cluster {
	min := len(detailArr[0].Vector)
	if len(detailArr) < min {
		min = len(detailArr)
	}
	clusters := make([]Cluster, min)
	for i := 0; i < len(clusters); i++ {
		clusters[i].MachineArr = make([]bool, len(detailArr[0].Vector))
	}
	for i, detail := range detailArr {
		// кончились машины
		if i >= len(detailArr[0].Vector) {
			ind := i % len(detailArr[0].Vector)
			clusters[ind].DetailArr = append(clusters[ind].DetailArr, detail)
			continue
		}
		clusters[i].DetailArr = append(clusters[i].DetailArr, detail)
		clusters[i].MachineArr[i] = true
	}
	for i := len(detailArr); i < len(detailArr[0].Vector); i++ {
		clusters[i%len(detailArr)].MachineArr[i] = true
	}
	return clusters
}

func FindCosts(cluster []Cluster) float64 {
	// перебираем каждый кластер
	n1_in := 0
	n0_in := 0
	n1_all := 0
	for _, el := range cluster {
		// перебираем каждую деталь из этого кластера
		for _, el1 := range el.DetailArr {
			// перебор машин и вектора
			for k, el2 := range el1.Vector {
				if el2 == true {
					n1_all += 1 * len(el1.Names)
				}
				if el2 == el.MachineArr[k] {
					if el2 {
						n1_in += 1 * len(el1.Names)
					}
					continue
				}
				if el2 == false {
					n0_in += 1 * len(el1.Names)
				}

			}
		}
	}
	return float64(n1_in) / float64(n0_in+n1_all)
}

func CopyCluster(cluster []Cluster) []Cluster {
	newCluster := make([]Cluster, len(cluster))
	for i, el := range cluster {
		newCluster[i].DetailArr = make([]detailsPack.Detals, len(el.DetailArr))
		newCluster[i].MachineArr = make([]bool, len(el.MachineArr))
		copy(newCluster[i].DetailArr, el.DetailArr)
		copy(newCluster[i].MachineArr, el.MachineArr)
	}
	return newCluster
}

func checkProfitOfSwapDetail(detail1, detail2 detailsPack.Detals, machine1, machine2 []bool) bool {
	current1 := findCountIntersection(detail1.Vector, machine1) * len(detail1.Names)
	current2 := findCountIntersection(detail2.Vector, machine2) * len(detail2.Names)
	next1 := findCountIntersection(detail1.Vector, machine2) * len(detail1.Names)
	next2 := findCountIntersection(detail2.Vector, machine1) * len(detail2.Names)
	if current1+current2 < next1+next2 {
		return true
	}
	return false
}

func SwapDetail(cluster []Cluster, safeMode bool) {
	indexCluster1 := rand.Intn(len(cluster))
	indexCluster2 := rand.Intn(len(cluster))
	counter := 0
	for {
		if indexCluster1 != indexCluster2 {
			break
		}
		indexCluster1 = rand.Intn(len(cluster))
		indexCluster2 = rand.Intn(len(cluster))
		counter += 1
		if counter >= 200 {
			fmt.Println("It seems only one cluster", len(cluster))
			return
		}

	}
	if len(cluster[indexCluster1].DetailArr) < 1 || len(cluster[indexCluster2].DetailArr) < 1 {
		fmt.Println(cluster)
		fmt.Println("Error!")
		return
	}
	indexDetalInCluster1 := rand.Intn(len(cluster[indexCluster1].DetailArr))
	indexDetalInCluster2 := rand.Intn(len(cluster[indexCluster2].DetailArr))

	//fmt.Println(cluster[indexCluster1].DetailArr[indexDetalInCluster1])
	if safeMode {
		count := 0
		for {
			detail1 := cluster[indexCluster1].DetailArr[indexDetalInCluster1]
			detail2 := cluster[indexCluster2].DetailArr[indexDetalInCluster2]
			machine1 := cluster[indexCluster1].MachineArr
			machine2 := cluster[indexCluster2].MachineArr
			if checkProfitOfSwapDetail(detail1, detail2, machine1, machine2) {
				break
			}
			if count >= 100 {
				return
			}
			indexDetalInCluster1 = rand.Intn(len(cluster[indexCluster1].DetailArr))
			indexDetalInCluster2 = rand.Intn(len(cluster[indexCluster2].DetailArr))
			count += 1
		}
	}
	cluster[indexCluster1].DetailArr[indexDetalInCluster1], cluster[indexCluster2].DetailArr[indexDetalInCluster2] =
		cluster[indexCluster2].DetailArr[indexDetalInCluster2], cluster[indexCluster1].DetailArr[indexDetalInCluster1]
}

func findCountIntersection(vector, machineVector []bool) int {
	countIntersections := 0
	for i, v := range vector {
		if v == machineVector[i] {
			countIntersections += 1
		}
	}
	return countIntersections
}

func checkTransfer(detail detailsPack.Detals, currentClusterMachine, nextClusterMachine []bool) bool {
	currentCountIntersections := findCountIntersection(detail.Vector, currentClusterMachine) * len(detail.Names)
	nextCountIntersections := findCountIntersection(detail.Vector, nextClusterMachine) * len(detail.Names)
	if nextCountIntersections >= currentCountIntersections {
		return true
	}
	return false
}

func SafeTransferDetail(cluster []Cluster, safeMode bool) {
	counter := 0
	indexCluster1, indexCluster2 := 0, 0
	for {
		if indexCluster1 != indexCluster2 && len(cluster[indexCluster1].DetailArr) > 1 {
			break
		}
		indexCluster1 = rand.Intn(len(cluster))
		indexCluster2 = rand.Intn(len(cluster))
		counter += 1
		if counter >= 200 {
			//fmt.Println("It seems only one cluster", len(cluster))
			return
		}
	}

	indexDetailInCluster1 := rand.Intn(len(cluster[indexCluster1].DetailArr))
	detail := cluster[indexCluster1].DetailArr[indexDetailInCluster1]
	if safeMode {
		count := 0
		for {
			if checkTransfer(detail, cluster[indexCluster1].MachineArr, cluster[indexCluster2].MachineArr) {
				break
			}
			indexDetailInCluster1 = rand.Intn(len(cluster[indexCluster1].DetailArr))
			if count >= 50 {
				return
			}
			count += 1
		}
	}
	cluster[indexCluster2].DetailArr = append(cluster[indexCluster2].DetailArr, cluster[indexCluster1].DetailArr[indexDetailInCluster1])
	cluster[indexCluster1].DetailArr = append(cluster[indexCluster1].DetailArr[:indexDetailInCluster1], cluster[indexCluster1].DetailArr[indexDetailInCluster1+1:]...)
	//fmt.Println(detail)

}

func findProfitOfMachine(machinePos int, detailArr []detailsPack.Detals) int {
	count := 0
	for _, detail := range detailArr {
		if detail.Vector[machinePos] == true {
			count += 1 * len(detail.Names)
		}
	}
	return count
}

func findCountMachineInCluster(machineArr []bool) int {
	count := 0
	for _, machine := range machineArr {
		if machine {
			count += 1
		}
	}
	return count
}

func FindBestPlaceForMachine(clusters []Cluster) {
	// было 100 / 20 -особо разницы не видно (при 100 есть задержки)
	for t := 0; t < 30; t++ {
		isBreak := false
		for k := 0; k < len(clusters); k++ {
			// чередую кластеры - k номер текущего кластера для перебора
			//ShowAnswer(clusters)
			if findCountMachineInCluster(clusters[k].MachineArr) < 2 {
				continue
			}
			for pos, machine := range clusters[k].MachineArr {
				// находим лучшую позицию для машины machine

				maxProfit := -1
				maxProfitIndexCluster := -1
				if machine {
					for i, cluster := range clusters {
						// профит - нахожу для 1 машины в каждом кластере
						//if i == k {
						//	continue
						//}
						profit := findProfitOfMachine(pos, cluster.DetailArr)
						if profit > maxProfit {
							maxProfit = profit
							maxProfitIndexCluster = i
						}
						randNum := rand.Intn(100)
						if randNum < 50 && profit == maxProfit {
							maxProfit = profit
							maxProfitIndexCluster = i
						}
					}
				}
				if maxProfit >= 0 && maxProfitIndexCluster >= 0 && maxProfitIndexCluster != k {
					clusters[maxProfitIndexCluster].MachineArr[pos], clusters[k].MachineArr[pos] =
						clusters[k].MachineArr[pos], clusters[maxProfitIndexCluster].MachineArr[pos]
					isBreak = true
					break
				}
			}
			if isBreak {
				break
			}
		}
	}
}

func FindBestPlaceForDetail(clusters []Cluster) {
	for i := 0; i < 30; i++ {
		for j := 0; j < len(clusters); j++ {
			if len(clusters[j].DetailArr) >= 2 {
				FindBestPlaceForAllDetailCorrect(clusters, j)
			}
		}
	}
}

func ShowAnswer(clusters []Cluster) {
	fmt.Println("detail / machine")
	for _, cluster := range clusters {
		for _, detail := range cluster.DetailArr {
			for _, name := range detail.Names {
				fmt.Print(name, " ")
			}
		}
		fmt.Print("  -  ")
		for i, machine := range cluster.MachineArr {
			if machine {
				fmt.Print(i+1, " ")
			}
		}
		fmt.Println()
	}
}

func HugeShaking(clusters []Cluster) []Cluster {
	newCluster := CopyCluster(clusters)
	indexCluster1 := rand.Intn(len(newCluster))
	indexCluster2 := rand.Intn(len(newCluster))
	count := 0
	for {
		if indexCluster1 != indexCluster2 {
			break
		}
		indexCluster1 = rand.Intn(len(newCluster))
		indexCluster2 = rand.Intn(len(newCluster))
		if count >= 200 {
			break
		}
		count += 1
	}
	rNum := rand.Intn(100)
	if rNum < 50 {
		newCluster[indexCluster1].DetailArr, newCluster[indexCluster2].DetailArr =
			newCluster[indexCluster2].DetailArr, newCluster[indexCluster1].DetailArr
	} else {
		newCluster[indexCluster1].MachineArr, newCluster[indexCluster2].MachineArr =
			newCluster[indexCluster2].MachineArr, newCluster[indexCluster1].MachineArr
	}
	for i := 0; i < 500; i++ {
		SwapTrueMachine(newCluster)
		SafeTransferDetail(newCluster, false)
		SwapDetail(newCluster, false)
	}
	return newCluster
}

func FindBestPlaceForAllDetailCorrect(clusters []Cluster, clusterIndex int) {
	// 2 варианта есть - перетаскавать в любом случае или же оставлять если это лучшее место
	if len(clusters[clusterIndex].DetailArr) < 2 {
		fmt.Println("Errorsssss!")
		fmt.Println(clusters)
		return
	}
	count := 0
	size := len(clusters[clusterIndex].DetailArr)
	for {
		if len(clusters[clusterIndex].DetailArr) <= 1 || count >= size*3 {
			break
		}
		detailIndex := rand.Intn(len(clusters[clusterIndex].DetailArr))
		detail := clusters[clusterIndex].DetailArr[detailIndex]
		maxCount := -1
		nextClusterIndex := -1
		for i, cluster := range clusters {
			for j, currentDetail := range cluster.DetailArr {
				if j == detailIndex {
					continue
				}
				countInter := findCountIntersection(detail.Vector, currentDetail.Vector)
				if countInter > maxCount {
					maxCount = countInter
					nextClusterIndex = i
				}
			}

		}
		//if len(clusters[0].DetailArr[0].Vector) - maxCount >  countDiff{
		//	return
		//}
		if nextClusterIndex != clusterIndex && nextClusterIndex != -1 {
			clusters[nextClusterIndex].DetailArr = append(clusters[nextClusterIndex].DetailArr, clusters[clusterIndex].DetailArr[detailIndex])
			clusters[clusterIndex].DetailArr = append(clusters[clusterIndex].DetailArr[:detailIndex], clusters[clusterIndex].DetailArr[detailIndex+1:]...)
		}
		count += 1
	}
}

func SwapTrueMachine(clusters []Cluster) {
	counter := 0
	machineIndex1, machineIndex2 := 0, 0
	indexCluster1, indexCluster2 := 0, 0
	for {
		indexCluster1 = rand.Intn(len(clusters))
		indexCluster2 = rand.Intn(len(clusters))

		machineIndex1 = rand.Intn(len(clusters[indexCluster1].MachineArr))
		machineIndex2 = rand.Intn(len(clusters[indexCluster2].MachineArr))

		if indexCluster1 != indexCluster2 && clusters[indexCluster1].MachineArr[machineIndex1] &&
			clusters[indexCluster2].MachineArr[machineIndex2] && machineIndex1 != machineIndex2 {
			break
		}
		if counter >= 400 {
			return
		}
		counter += 1
	}
	clusters[indexCluster1].MachineArr[machineIndex1], clusters[indexCluster2].MachineArr[machineIndex1] =
		clusters[indexCluster2].MachineArr[machineIndex1], clusters[indexCluster1].MachineArr[machineIndex1]

	clusters[indexCluster1].MachineArr[machineIndex2], clusters[indexCluster2].MachineArr[machineIndex2] =
		clusters[indexCluster2].MachineArr[machineIndex2], clusters[indexCluster1].MachineArr[machineIndex2]

}

func ThrowManyMachine(clusters []Cluster) {
	if len(clusters) < 4 {
		fmt.Println("It unreal ti do it now!")
		return
	}
	counter := 0
	clusterIndex1, clusterIndex2 := 0, 0
	countMachine := 0
	for {

		clusterIndex1 = rand.Intn(len(clusters))
		clusterIndex2 = rand.Intn(len(clusters))
		countMachine = findCountMachineInCluster(clusters[clusterIndex1].MachineArr)
		if countMachine >= 3 {
			break
		}
		if counter >= 400 {
			return
		}
		counter += 1
	}
	count := 0
	for i, machine := range clusters[clusterIndex1].MachineArr {
		if countMachine-count <= 2 {
			return
		}

		if machine {
			chance := rand.Intn(100)
			if chance < 70 {
				clusters[clusterIndex1].MachineArr[i] = false
				clusters[clusterIndex2].MachineArr[i] = true
				count += 1
			}
		}
	}
}

func ThrowManyDetails(clusters []Cluster) {
	if len(clusters) < 4 {
		fmt.Println("It unreal ti do it now!")
		return
	}
	counter := 0
	clusterIndex1, clusterIndex2 := 0, 0
	countDetail := 0
	for {

		clusterIndex1 = rand.Intn(len(clusters))
		clusterIndex2 = rand.Intn(len(clusters))
		countDetail = len(clusters[clusterIndex1].DetailArr)
		if countDetail >= 3 {
			break
		}
		if counter >= 400 {
			return
		}
		counter += 1
	}
	count := 0
	size := len(clusters[clusterIndex1].DetailArr)
	var newArr []detailsPack.Detals
	for j := 0; j < size/2; j++ {
		// цикл в цикле т к изменяется clusters
		for i, _ := range clusters[clusterIndex1].DetailArr {
			if countDetail-count <= 2 {
				return
			}
			chance := rand.Intn(100)
			lenOfPrimaryDetailArr := len(clusters[clusterIndex1].DetailArr)
			if chance < 50 {
				// меняем выбранный элемент с последним, чтобы облегчить перестановку
				clusters[clusterIndex1].DetailArr[i], clusters[clusterIndex1].DetailArr[lenOfPrimaryDetailArr-1] =
					clusters[clusterIndex1].DetailArr[lenOfPrimaryDetailArr-1], clusters[clusterIndex1].DetailArr[i]
				// перестановка
				newArr = append(newArr, clusters[clusterIndex1].DetailArr[lenOfPrimaryDetailArr-1])
				clusters[clusterIndex1].DetailArr = clusters[clusterIndex1].DetailArr[:lenOfPrimaryDetailArr-1]
				break
			}
		}
	}
	for _, el := range newArr {
		clusters[clusterIndex2].DetailArr = append(clusters[clusterIndex2].DetailArr, el)
	}

}

func findProfitAndFineIndex(detailVector1 []bool, detailVector2 []bool) float64 {
	countProfit := 0
	countFine := 0
	for i, detailV1 := range detailVector1 {
		if detailV1 == detailVector2[i] {
			countProfit += 1
			continue
		}
		countFine += 1
	}
	return float64(countProfit) / float64(countFine)
}

func findBestPlaceForOneDetail(clusters []Cluster, clusterIndex, detailIndex int) bool {
	// не находим лучшее положение если в кластере только 1 деталь или ошибка с индексами
	lenOfPrimaryDetailArr := len(clusters[clusterIndex].DetailArr)
	if len(clusters) <= 1 || lenOfPrimaryDetailArr <= detailIndex || lenOfPrimaryDetailArr < 2 {
		fmt.Println("It seems too little clusters")
		return false
	}
	maxProfitIndex := -1.0
	maxProfitClusterPosition := 0
	for i, cluster := range clusters {
		for detIndex, detail := range cluster.DetailArr {
			// не учитываем профит от данной детали
			if detIndex == detailIndex && clusterIndex == i {
				continue
			}
			// профит / штраф - находим наибольший
			currentProfit := findProfitAndFineIndex(detail.Vector, clusters[clusterIndex].DetailArr[detailIndex].Vector)
			if currentProfit > maxProfitIndex {
				maxProfitIndex = currentProfit
				maxProfitClusterPosition = i
			}
		}
	}
	if maxProfitClusterPosition == clusterIndex {
		return false
	}
	// меняем выбранный элемент с последним, чтобы облегчить перестановку
	clusters[clusterIndex].DetailArr[detailIndex], clusters[clusterIndex].DetailArr[lenOfPrimaryDetailArr-1] =
		clusters[clusterIndex].DetailArr[lenOfPrimaryDetailArr-1], clusters[clusterIndex].DetailArr[detailIndex]
	// перестановка
	clusters[maxProfitClusterPosition].DetailArr = append(clusters[maxProfitClusterPosition].DetailArr, clusters[clusterIndex].DetailArr[lenOfPrimaryDetailArr-1])
	clusters[clusterIndex].DetailArr = clusters[clusterIndex].DetailArr[:lenOfPrimaryDetailArr-1]
	return true
}

func FindBestPlaceForAllDetails(clusters []Cluster) {
	// не более 20 иначе перебор оптимизаии ( проверено) !!!! уменьшил и стало намного лучше
	for i := 0; i <= 19; i++ {
		clusterIndex := rand.Intn(len(clusters))
		if len(clusters[clusterIndex].DetailArr) < 2 {
			continue
		}
		detailIndex := rand.Intn(len(clusters[clusterIndex].DetailArr))
		findBestPlaceForOneDetail(clusters, clusterIndex, detailIndex)
	}
}

func JoinClusters(clusters []Cluster) []Cluster {
	if len(clusters) <= 1 {
		return clusters
	}
	for i, machine := range clusters[0].MachineArr {
		if machine {
			clusters[1].MachineArr[i] = true
		}
	}
	for _, detail := range clusters[0].DetailArr {
		clusters[1].DetailArr = append(clusters[1].DetailArr, detail)
	}
	return clusters[1:]
}

func TransferOneDetail(clusters []Cluster) {
	if len(clusters) < 2 {
		return
	}
	clusterIndex1, clusterIndex2 := 0, 0
	count := 0
	for {
		count += 1
		clusterIndex1 = rand.Intn(len(clusters))
		clusterIndex2 = rand.Intn(len(clusters))
		if count >= 200 {
			return
		}
		if len(clusters[clusterIndex1].DetailArr) <= 1 {
			continue
		}
		if clusterIndex1 != clusterIndex2 {
			break
		}
	}
	detailIndex := rand.Intn(len(clusters[clusterIndex1].DetailArr))
	// swap for more comfortable changes
	size := len(clusters[clusterIndex1].DetailArr)
	clusters[clusterIndex1].DetailArr[detailIndex], clusters[clusterIndex1].DetailArr[size-1] =
		clusters[clusterIndex1].DetailArr[size-1], clusters[clusterIndex1].DetailArr[detailIndex]

	clusters[clusterIndex2].DetailArr =
		append(clusters[clusterIndex2].DetailArr, clusters[clusterIndex1].DetailArr[size-1])
	clusters[clusterIndex1].DetailArr = clusters[clusterIndex1].DetailArr[:size-1]
}

func SwapDetails(clusters []Cluster) {
	if len(clusters) < 2 {
		return
	}
	clusterIndex1, clusterIndex2 := 0, 0
	count := 0
	for {
		count += 1
		clusterIndex1 = rand.Intn(len(clusters))
		clusterIndex2 = rand.Intn(len(clusters))
		if count >= 200 {
			return
		}
		if len(clusters[clusterIndex1].DetailArr) <= 1 {
			continue
		}
		if clusterIndex1 != clusterIndex2 {
			break
		}
	}
	detailIndex1 := rand.Intn(len(clusters[clusterIndex1].DetailArr))
	detailIndex2 := rand.Intn(len(clusters[clusterIndex2].DetailArr))

	// swap 1
	size1 := len(clusters[clusterIndex1].DetailArr)
	clusters[clusterIndex1].DetailArr[detailIndex1], clusters[clusterIndex1].DetailArr[size1-1] =
		clusters[clusterIndex1].DetailArr[size1-1], clusters[clusterIndex1].DetailArr[detailIndex1]

	size2 := len(clusters[clusterIndex2].DetailArr)
	clusters[clusterIndex2].DetailArr[detailIndex2], clusters[clusterIndex2].DetailArr[size2-1] =
		clusters[clusterIndex2].DetailArr[size2-1], clusters[clusterIndex2].DetailArr[detailIndex2]

	det1 := clusters[clusterIndex1].DetailArr[size1-1]
	det2 := clusters[clusterIndex2].DetailArr[size2-1]

	clusters[clusterIndex1].DetailArr = clusters[clusterIndex1].DetailArr[:size1-1]
	clusters[clusterIndex2].DetailArr = clusters[clusterIndex2].DetailArr[:size2-1]

	clusters[clusterIndex1].DetailArr = append(clusters[clusterIndex1].DetailArr, det2)
	clusters[clusterIndex2].DetailArr = append(clusters[clusterIndex2].DetailArr, det1)
}

func TransferOneMachine(clusters []Cluster) {
	if len(clusters) < 2 {
		return
	}
	clusterIndex1, clusterIndex2 := 0, 0
	count := 0
	machineIndex := 0
	for {
		count += 1
		clusterIndex1 = rand.Intn(len(clusters))
		clusterIndex2 = rand.Intn(len(clusters))
		machineIndex = rand.Intn(len(clusters[clusterIndex1].MachineArr))
		if count >= 400 {
			return
		}
		if findCountMachineInCluster(clusters[clusterIndex1].MachineArr) <= 1 {
			continue
		}
		if clusterIndex1 != clusterIndex2 && clusters[clusterIndex1].MachineArr[machineIndex] {
			break
		}
	}

	// swap
	clusters[clusterIndex1].MachineArr[machineIndex], clusters[clusterIndex2].MachineArr[machineIndex] =
		clusters[clusterIndex2].MachineArr[machineIndex], clusters[clusterIndex1].MachineArr[machineIndex]
}

func SwapMachine(clusters []Cluster) {
	if len(clusters) < 2 {
		return
	}
	clusterIndex1, clusterIndex2 := 0, 0
	count := 0
	machineIndex1, machineIndex2 := 0, 0
	for {
		count += 1
		clusterIndex1 = rand.Intn(len(clusters))
		clusterIndex2 = rand.Intn(len(clusters))
		machineIndex1 = rand.Intn(len(clusters[clusterIndex1].MachineArr))
		machineIndex2 = rand.Intn(len(clusters[clusterIndex2].MachineArr))
		if count >= 400 {
			return
		}
		if findCountMachineInCluster(clusters[clusterIndex1].MachineArr) <= 1 ||
			findCountMachineInCluster(clusters[clusterIndex2].MachineArr) <= 1 {
			continue
		}
		if clusterIndex1 != clusterIndex2 && clusters[clusterIndex1].MachineArr[machineIndex1] &&
			clusters[clusterIndex2].MachineArr[machineIndex2] {
			break
		}
	}
	// swap
	clusters[clusterIndex1].MachineArr[machineIndex1], clusters[clusterIndex2].MachineArr[machineIndex1] =
		clusters[clusterIndex2].MachineArr[machineIndex1], clusters[clusterIndex1].MachineArr[machineIndex1]

	clusters[clusterIndex1].MachineArr[machineIndex2], clusters[clusterIndex2].MachineArr[machineIndex2] =
		clusters[clusterIndex2].MachineArr[machineIndex2], clusters[clusterIndex1].MachineArr[machineIndex2]
}
