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
	n1In := 0
	n0In := 0
	n1All := 0
	for _, el := range cluster {
		// перебираем каждую деталь из этого кластера
		for _, el1 := range el.DetailArr {
			// перебор машин и вектора
			for k, el2 := range el1.Vector {
				if el2 == true {
					n1All += 1 * len(el1.Names)
				}
				if el2 == el.MachineArr[k] {
					if el2 {
						n1In += 1 * len(el1.Names)
					}
					continue
				}
				if el2 == false {
					n0In += 1 * len(el1.Names)
				}

			}
		}
	}
	return float64(n1In) / float64(n0In+n1All)
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

func findCountMachineInCluster(machineArr []bool) int {
	count := 0
	for _, machine := range machineArr {
		if machine {
			count += 1
		}
	}
	return count
}

func ShowAns(clusters []Cluster, maxDetail, maxCluster int) {
	detailAns := make([]int, maxDetail+1)
	machineAns := make([]int, maxCluster+1)
	num := 1
	for _, el := range clusters {
		for i := 0; i < len(el.DetailArr); i++ {
			for _, name := range el.DetailArr[i].Names {
				detailAns[name] = num
			}
		}
		for i, machine := range el.MachineArr {
			if machine {
				machineAns[i+1] = num
			}
		}
		num += 1
	}
	for _, el := range machineAns[1:]{
		fmt.Print(el," ")
	}
	fmt.Println()
	for _, el := range detailAns[1:]{
		fmt.Print(el," ")
	}
	fmt.Println()
}


func ThrowManyMachine(clusters []Cluster) {
	if len(clusters) < 2 {
		//fmt.Println("It unreal ti do it now!")
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
	if len(clusters) < 2 {
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

func SplitClusters(clusters []Cluster, maxCountCluster int) []Cluster {
	if len(clusters) >= maxCountCluster {
		return clusters
	}
	machinePos := -1
	isFoundDetail := false
	var newDetails []detailsPack.Detals
	for i := 0; i < len(clusters); i++ {
		if machinePos == -1 && findCountMachineInCluster(clusters[i].MachineArr) >= 2 {
			for j, machine := range clusters[i].MachineArr {
				if machine {
					machinePos = j
					clusters[i].MachineArr[j] = false
					break
				}
			}
		}
		if isFoundDetail == false && len(clusters[i].DetailArr) >= 2 {
			newDetails = append(newDetails, clusters[i].DetailArr[len(clusters[i].DetailArr)-1])
			clusters[i].DetailArr = clusters[i].DetailArr[:len(clusters[i].DetailArr)-1]
			isFoundDetail = true
		}
		if isFoundDetail == true && machinePos >= 0 {
			break
		}
	}
	newMachineArr := make([]bool, len(clusters[0].MachineArr))
	newMachineArr[machinePos] = true
	newCluster := Cluster{DetailArr: newDetails, MachineArr: newMachineArr}
	clusters = append(clusters, newCluster)
	return clusters
}

func MakeChange(clusters []Cluster){
	chance := rand.Intn(1000)
	if chance <= 250 {
		TransferOneDetail(clusters)
	} else if chance <= 500 {
		TransferOneMachine(clusters)
	} else if chance <= 750 {
		SwapMachine(clusters)
	} else {
		SwapDetails(clusters)
	}
}

func Perturbation(bestCluster []Cluster)[]Cluster{
	for k := 0; k < 10; k++ {
		if k % 2 == 0 {
			ThrowManyMachine(bestCluster)
		} else {
			ThrowManyDetails(bestCluster)
		}
	}
	return bestCluster
}

func Shaking(bestAnsClusterArr [][]Cluster, bestCluster , bestAnsCluster []Cluster, maxSizeClusters int) []Cluster{
	chance := rand.Intn(1000)
	if chance < 500 {
		if chance < 250 || len(bestAnsClusterArr[len(bestCluster)][0].DetailArr) < 1{
			bestCluster = CopyCluster(bestAnsCluster)
		} else {
			bestCluster = CopyCluster(bestAnsClusterArr[len(bestCluster)])
		}
	} else {
		count := rand.Intn(3)
		for i := 0; i < count; i++ {
			if chance < 750 {
				bestCluster = JoinClusters(bestCluster)
			} else {
				bestCluster = SplitClusters(bestCluster, maxSizeClusters)
			}
		}
	}
	return bestCluster
}

func RememberBestAns(bestAnsClusterArr [][]Cluster, bestCluster, bestAnsCluster []Cluster, detailSize,
	machineSize int)([][]Cluster, []Cluster){
	if FindCosts(bestCluster) > FindCosts(bestAnsClusterArr[len(bestCluster)]) + 0.00000000000001 ||
		len(bestAnsClusterArr[len(bestCluster)][0].DetailArr) < 1{
		bestAnsClusterArr[len(bestCluster)] = CopyCluster(bestCluster)
	}
	if FindCosts(bestCluster) > FindCosts(bestAnsCluster)+0.00000000000001 {
		fmt.Println(FindCosts(bestCluster), "--------------", len(bestCluster))
		ShowAns(bestCluster, detailSize, machineSize)
		bestAnsCluster = CopyCluster(bestCluster)
	}
	return bestAnsClusterArr, bestAnsCluster
}

func MakeBestAnsClusters(maxSizeClusters int, clusters []Cluster) [][]Cluster{
	bestAnsClusterArr := make([][]Cluster, maxSizeClusters + 1)
	for i := 1; i <= maxSizeClusters; i++{
		bestAnsClusterArr[i] = make([]Cluster, i)
	}
	bestAnsClusterArr[maxSizeClusters] = CopyCluster(clusters)
	return bestAnsClusterArr
}

func CheckBestCluster(clusters, bestCluster []Cluster, k int) ([]Cluster, int){
	if FindCosts(clusters) > FindCosts(bestCluster)+0.00000000000001 {
		bestCluster = CopyCluster(clusters)
		k = 0
	}
	return bestCluster, k
}
