package inotify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"wzinc/dify"
)

// Dataset 结构体定义
type Dataset struct {
	DatasetID      string   `json:"datasetID"`
	LastUpdateTime string   `json:"lastUpdateTime"`
	Paths          []string `json:"paths"`
}

type FolderStatusResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    map[string]Dataset `json:"data"`
}

// PathsComparison 结构体定义
type PathsComparison struct {
	AddPaths    []string `json:"addPaths"`
	DeletePaths []string `json:"deletePaths"`
}

var (
	LastDatasetMap     map[string]Dataset
	DatasetMap         map[string]Dataset
	PathsComparisonMap map[string]PathsComparison
	PathToDatasetMap   map[string]PathComparison
	mutex              sync.Mutex
)

func InitializePathToDatasetMap(watchDir string, datasetID string) {
	PathToDatasetMap = make(map[string]PathComparison)

	pathComparison := PathComparison{
		Base:   []string{datasetID},
		Add:    []string{datasetID},
		Delete: []string{},
		Op:     "add",
	}

	PathToDatasetMap[watchDir] = pathComparison
}

// 更新 DatasetMap
func UpdateDatasetMap() error {
	// 发起 HTTP 请求
	resp, err := CallGetDatasetFolderStatus()
	if err != nil {
		return fmt.Errorf("无法获取数据: %s", err.Error())
	}
	fmt.Println("resp=", string(resp))

	// 解析 JSON 数据
	var data FolderStatusResponse
	err = json.NewDecoder(bytes.NewReader(resp)).Decode(&data)
	if err != nil {
		return fmt.Errorf("解析 JSON 数据失败: %s", err.Error())
	}

	// 使用互斥锁更新全局变量 DatasetMap 和 LastDatasetMap
	mutex.Lock()
	defer mutex.Unlock()

	// 保存当前的 DatasetMap 到 LastDatasetMap
	LastDatasetMap = make(map[string]Dataset)
	for datasetID, dataset := range DatasetMap {
		LastDatasetMap[datasetID] = dataset
	}

	// 更新 DatasetMap
	DatasetMap = make(map[string]Dataset)
	for datasetID, newDataset := range data.Data {
		var cleanedPaths []string
		for _, path := range newDataset.Paths {
			if !strings.HasPrefix(path, "/data") {
				// cleanedPaths = append(cleanedPaths, path)
			} else {
				cleanedPath := strings.TrimPrefix(path, "/data")
				cleanedPaths = append(cleanedPaths, cleanedPath)
			}
		}
		newDataset.Paths = cleanedPaths
		DatasetMap[datasetID] = newDataset
	}

	// 单独处理默认知识库ID作为的键 "A" 和默认WatchDir作为的路径 "B"
	if _, exists := DatasetMap[dify.DatasetId]; !exists {
		// 如果键 "A" 不存在于 DatasetMap 中，则创建一个新的 Dataset
		datasetA := Dataset{
			DatasetID:      dify.DatasetId,
			LastUpdateTime: "",
			Paths:          []string{},
		}
		DatasetMap[dify.DatasetId] = datasetA
	}

	if datasetA, exists := DatasetMap[dify.DatasetId]; exists {
		// 检查路径 "B" 是否已经存在于 Dataset "A" 的 Paths 中
		pathB := WatchDir // 替换为路径 "B" 的实际值
		found := false
		for _, path := range datasetA.Paths {
			if path == pathB {
				found = true
				break
			}
		}

		// 如果路径 "B" 不存在于 Dataset "A" 的 Paths 中，则将其添加进去
		if !found {
			datasetA.Paths = append(datasetA.Paths, pathB)
			DatasetMap[dify.DatasetId] = datasetA
		}
	}
	return nil
}

// 更新 PathsComparisonMap
func UpdatePathsComparison() {
	mutex.Lock()
	defer mutex.Unlock()

	// 获取 DatasetMap 和 LastDatasetMap 的键的并集
	keys := make(map[string]bool)
	for datasetID := range DatasetMap {
		keys[datasetID] = true
	}
	for datasetID := range LastDatasetMap {
		keys[datasetID] = true
	}

	// 更新 PathsComparisonMap
	PathsComparisonMap = make(map[string]PathsComparison)
	for datasetID := range keys {
		// 获取当前和之前的 Dataset
		currentDataset, currentExists := DatasetMap[datasetID]
		previousDataset, previousExists := LastDatasetMap[datasetID]

		// 初始化 PathsComparison
		pathsComparison := PathsComparison{
			AddPaths:    make([]string, 0),
			DeletePaths: make([]string, 0),
		}

		// 处理新增和删除的路径
		if currentExists && previousExists {
			pathsComparison.AddPaths = getAddedPaths(currentDataset.Paths, previousDataset.Paths)
			pathsComparison.DeletePaths = getDeletedPaths(currentDataset.Paths, previousDataset.Paths)
		} else if currentExists {
			pathsComparison.AddPaths = currentDataset.Paths
			pathsComparison.DeletePaths = []string{} // 空列表表示没有删除路径
		} else if previousExists {
			pathsComparison.AddPaths = []string{} // 空列表表示没有新增路径
			pathsComparison.DeletePaths = previousDataset.Paths
		}

		// 更新 PathsComparisonMap
		PathsComparisonMap[datasetID] = pathsComparison
	}
}

// 获取新增的路径
func getAddedPaths(currentPaths, previousPaths []string) []string {
	addedPaths := make([]string, 0)

	for _, currentPath := range currentPaths {
		// 如果当前路径不存在于之前的路径中，则将其添加到新增路径中
		if !contains(previousPaths, currentPath) {
			addedPaths = append(addedPaths, currentPath)
		}
	}

	return addedPaths
}

// 获取删除的路径
func getDeletedPaths(currentPaths, previousPaths []string) []string {
	deletedPaths := make([]string, 0)

	for _, previousPath := range previousPaths {
		// 如果之前的路径不存在于当前路径中，则将其添加到删除路径中
		if !contains(currentPaths, previousPath) {
			deletedPaths = append(deletedPaths, previousPath)
		}
	}

	return deletedPaths
}

// 检查列表中是否包含指定的路径
func contains(paths []string, targetPath string) bool {
	for _, path := range paths {
		if path == targetPath {
			return true
		}
	}
	return false
}

// PathComparison 结构体定义
type PathComparison struct {
	Base   []string `json:"base"`
	Add    []string `json:"add"`
	Delete []string `json:"delete"`
	Op     string   `json:"op"`
}

// 更新 PathToDatasetMap
func UpdatePathToDatasetMap() {
	mutex.Lock()         // 加锁，确保并发安全
	defer mutex.Unlock() // 解锁

	PathToDatasetMap = make(map[string]PathComparison) // 创建一个新的空的 PathToDatasetMap

	// 获取所有路径的并集
	lastAllPaths := make(map[string]bool)
	currentAllPaths := make(map[string]bool)
	allPaths := make(map[string]string) // 用于存储所有路径的状态
	for _, dataset := range LastDatasetMap {
		for _, path := range dataset.Paths {
			lastAllPaths[path] = true
			allPaths[path] = "delete"
		}
	}

	for _, dataset := range DatasetMap {
		for _, path := range dataset.Paths {
			currentAllPaths[path] = true
			allPaths[path] = "add"
		}
	}

	for path := range allPaths {
		if lastAllPaths[path] && currentAllPaths[path] {
			allPaths[path] = "keep"
		}
	}

	// 遍历每个路径
	for path := range allPaths {
		// 初始化 PathComparison
		pathComparison := PathComparison{
			Base:   make([]string, 0), // 初始化 Base 切片为空
			Add:    make([]string, 0), // 初始化 Add 切片为空
			Delete: make([]string, 0), // 初始化 Delete 切片为空
			Op:     allPaths[path],
		}

		// 检查路径在 DatasetMap 中的数据集
		currentDatasetIDs := getDatasetIDsForPath(path) // 获取路径在 DatasetMap 中拥有该路径的数据集ID列表
		if len(currentDatasetIDs) > 0 {
			pathComparison.Base = append(pathComparison.Base, currentDatasetIDs...) // 将数据集ID列表添加到 Base 切片中
		}

		// 遍历 PathsComparisonMap 的每个键
		for key := range PathsComparisonMap {
			// 获取 AddPaths 和 DeletePaths
			addPaths := PathsComparisonMap[key].AddPaths
			deletePaths := PathsComparisonMap[key].DeletePaths

			// 检查路径在 AddPaths 中的存在并更新 PathComparison 的 Add 字段
			if containsPath(addPaths, path) {
				pathComparison.Add = append(pathComparison.Add, key)
			}

			// 检查路径在 DeletePaths 中的存在并更新 PathComparison 的 Delete 字段
			if containsPath(deletePaths, path) {
				pathComparison.Delete = append(pathComparison.Delete, key)
			}
		}

		// 更新 PathToDatasetMap
		PathToDatasetMap[path] = pathComparison // 将路径和对应的 PathComparison 结构体添加到 PathToDatasetMap 中
	}
}

// 获取拥有指定路径的数据集ID列表
func getDatasetIDsForPath(path string) []string {
	datasetIDs := make([]string, 0)
	for datasetID, dataset := range DatasetMap {
		if contains(dataset.Paths, path) {
			datasetIDs = append(datasetIDs, datasetID)
		}
	}
	return datasetIDs
}

// 辅助函数：检查路径切片中是否包含指定的路径
func containsPath(paths []string, path string) bool {
	for _, p := range paths {
		if p == path {
			return true
		}
	}
	return false
}

func CalibPathToDatasetMap() {
	mutex.Lock()
	defer mutex.Unlock()

	// 检查是否存在 inotify.WatchDir 的键
	if pathComparison, exists := PathToDatasetMap[WatchDir]; exists {
		// 检查 Base 中是否包含 dify.DatasetId，如果不包含，则添加到 Base 中
		if !containsDatasetID(pathComparison.Base, dify.DatasetId) {
			pathComparison.Base = append(pathComparison.Base, dify.DatasetId)
		}

		// 检查 Add 中是否包含 dify.DatasetId，如果包含，则移除
		pathComparison.Add = removeDatasetID(pathComparison.Add, dify.DatasetId)

		// 检查 Delete 中是否包含 dify.DatasetId，如果包含，则移除
		pathComparison.Delete = removeDatasetID(pathComparison.Delete, dify.DatasetId)
	} else {
		// 创建新的 PathComparison
		pathComparison := PathComparison{
			Base:   []string{dify.DatasetId},
			Add:    []string{},
			Delete: []string{},
			Op:     "keep",
		}

		// 将 pathComparison 添加到 PathToDatasetMap
		PathToDatasetMap[WatchDir] = pathComparison
	}
}

// 辅助函数：检查数据集ID切片中是否包含指定的数据集ID
func containsDatasetID(datasetIDs []string, datasetID string) bool {
	for _, id := range datasetIDs {
		if id == datasetID {
			return true
		}
	}
	return false
}

// 辅助函数：从数据集ID切片中移除指定的数据集ID
func removeDatasetID(datasetIDs []string, datasetID string) []string {
	result := make([]string, 0)
	for _, id := range datasetIDs {
		if id != datasetID {
			result = append(result, id)
		}
	}
	return result
}

// 根据 datasetID 获取 Dataset 数据
func GetDatasetByID(datasetID string) (Dataset, bool) {
	// 使用互斥锁访问全局变量 datasetMap
	mutex.Lock()
	defer mutex.Unlock()

	dataset, ok := DatasetMap[datasetID]
	return dataset, ok
}

func FindBaseFromPath(filename string) []string {
	var result []string

	for {
		// 遍历每个路径和对应的数据集
		for path, dataset := range PathToDatasetMap {
			// 检查给定的 filename 是否包含当前路径
			if strings.Contains(filename, path) {
				result = dataset.Base
				break
			}
		}

		// 如果找到匹配的路径，终止循环
		if result != nil {
			break
		}

		// 获取 filename 的上一层目录
		dir := filepath.Dir(filename)
		// 如果 dir 与 filename 相同或者为根目录，终止循环
		if dir == filename || dir == "." {
			break
		}

		// 更新 filename 为上一层目录的路径，继续循环
		filename = dir
	}

	// 如果无法找到匹配的路径，将 result 设置为空列表
	if result == nil {
		result = []string{}
	}

	return result
}

// 获取数据集文件夹的状态
func GetAndUpdateDatasetFolderStatus() error {
	fmt.Println("Update Dataset Map")
	err := UpdateDatasetMap()
	if err != nil {
		fmt.Println("Update Dataset Map error: ", err)
		return err
	}

	fmt.Println("Update Paths Comparison")
	UpdatePathsComparison()

	fmt.Println("Update Path to Dataset Map")
	UpdatePathToDatasetMap()

	fmt.Println("Calib Path To Dataset Map")
	CalibPathToDatasetMap()

	fmt.Println("LastDatasetMap:\n", LastDatasetMap)
	fmt.Println("DatasetMap:\n", DatasetMap)
	fmt.Println("PathsComparisonMap:\n", PathsComparisonMap)
	fmt.Println("PathToDatasetMap:\n", PathToDatasetMap)

	fmt.Println("Watch Path")
	WatchPath(PathToDatasetMap)

	fmt.Println("Op 4 functional done!")
	return nil
}

// unused for now
func UpdateDatasetFolderPathsHandler(w http.ResponseWriter, r *http.Request) {
	err := GetAndUpdateDatasetFolderStatus()
	if err != nil {
		// 处理错误
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回响应
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
