package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"ttdocker/network"
	"ttdocker/cgroups"
	"ttdocker/cgroups/subsystems"
	"ttdocker/container"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)




func Run(tty bool, comArray []string, res *subsystems.ResourceConfig, volume string, containerName string, imageName string, envSlice []string, nw string, portmapping []string){

	containerID := randStringBytes(10)
	if containerName == "" {

		containerName = containerID
	}

	//将环境变量传递给 process
	parent, writePipe := container.NewParentProcess(tty, volume, containerName, imageName, envSlice)
	if parent == nil {

		log.Errorf("new parent process error")
		return
	}

	//start 调用前面创建好的command 命令
	// start 以非阻塞方式运行， run 为阻塞，等待命令结束
	//首先会clone 出一个namspace 隔离的进程, 然后在子进程中,调用/proc/self/exe  调用自己, 发送init 参数
	if err := parent.Start(); err != nil {

		log.Error(err)
	}

	//记录容器信息
	//fmt.Println("开始记录容器信息")
	//fmt.Println("comarray = ", comArray)
	containerName, err := recordContainerInfo(parent.Process.Pid, comArray, containerName, containerID, volume)
	if err != nil {

		log.Errorf("recode container info error %v", err)
		return
	}

	// use mydocker-cgroup as cgroup name
	//创建 cgroup manager ，并通过调用 set 和 apply 设置资源限制并使限制在容器上生效
	cgroupsManager := cgroups.NewCgroupManager(containerID)
	defer cgroupsManager.Destroy()
	//设置资源限制
	cgroupsManager.Set(res)
	//将容器进程加入到各个 subsystem 挂载对应的 cgroup
	cgroupsManager.Apply(parent.Process.Pid)

	if nw != "" {

		//config container network
		network.Init()
		containerInfo := &container.ContainerInfo{

			Id: containerID,
			Pid: strconv.Itoa(parent.Process.Pid),
			Name: containerName,
			PortMapping: portmapping,
		}

		if err := network.Connect(nw, containerInfo); err != nil {

			log.Errorf("error connect network %v", err)
			return
		}
	}

	//对容器设置完限制之后，初始化容器
	//发送用户命令
/*
	mntURL := "/root/mnt"
	rootURL := "/root/"
*/
	sendInitCommand(comArray, writePipe)
	//　阻塞在这
	if tty {

		parent.Wait()
		deleteContainerInfo(containerName)
		container.DeleteWorkSpace(volume,containerName)
	}

	/*
	mntURL := "/root/mnt"
	rootURL := "/root/"
	//退出前删除对应的目录
	os.Exit(0)
*/
}

func sendInitCommand(comArray []string, writePipe *os.File){

	command := strings.Join(comArray, " ")
	log.Infof("command all is %s", command)

	writePipe.WriteString(command)
	writePipe.Close()
}



//记录容器信息,将容器的信息持久化到磁盘中
func recordContainerInfo (containerPID int, commandArray []string, containerName string, id string, volume string) (string, error){
	//首先生成　10 为数字的容器ID
	/*fmt.Println("开始获取随机数")
	id := randStringBytes(10)
	fmt.Println("成功获取随机数")*/
	//以当前时间为容器创建时间
	createTime := time.Now().Format("2020-08-28 13:08:00")
	command := strings.Join(commandArray, "")


	//如果用户不指定容器名，　那么就以容器ID　当做容器名
	/*if containerName == "" {
		containerName = id
	}*/

	//生成容器信息的结构体实例

	containerInfo := &container.ContainerInfo{

		Id: id,
		Pid: strconv.Itoa(containerPID),
		Command: command,
		CreatedTime: createTime,
		Status: container.RUNNING,
		Name: containerName,
		Volume: volume,
	}

	//将容器信息对象 json 序列化成字符串
	jsonBytes, err := json.Marshal(containerInfo)
	if err != nil {

		log.Error("Record container info error %v", err)
		return "",err
	}

	jsonStr := string(jsonBytes)

	//拼凑一下存储容器信息的路径
	dirUrl := fmt.Sprintf(container.DefaultInfoLocation, containerName)

	//如果改路径不存在，级联创建

	//fmt.Println("开始创建目录")
	if err := os.MkdirAll(dirUrl, 0622); err != nil {

		log.Errorf("mkdir error %s error %v", dirUrl, err)
		return "",err
	}

	//fmt.Println("成功创建目录")
	fileName := dirUrl + "/" + container.ConfigName
	//创建最终的配置文件 -- config.json 文件
	//fmt.Println("开始创建文件")
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {

		log.Errorf("create file %s error %v", fileName, err)
		return "", err
	}

	//将json化之后的数据写到文件中
	if _, err := file.WriteString(jsonStr); err != nil {

		log.Errorf("file write string error %v", err)
		return "", err
	}

	return containerName, nil
}

//ID 生成器
func randStringBytes(n int) string {

	letterBytes := "1234567890"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)

	for i := range b {

		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

func deleteContainerInfo(containerId string){

	dirUrl := fmt.Sprintf(container.DefaultInfoLocation, containerId)
	if err := os.RemoveAll(dirUrl); err != nil {

		log.Errorf("remove dir %s error %v", dirUrl, err)
	}
}
