package endpoint

type DockerInfo struct {
	Architecture      string `json:"Architecture"`
	BridgeNfIP6tables bool   `json:"BridgeNfIp6tables"`
	BridgeNfIptables  bool   `json:"BridgeNfIptables"`
	CPUSet            bool   `json:"CPUSet"`
	CPUShares         bool   `json:"CPUShares"`
	CgroupDriver      string `json:"CgroupDriver"`
	ClusterAdvertise  string `json:"ClusterAdvertise"`
	ClusterStore      string `json:"ClusterStore"`
	ContainerdCommit  struct {
		Expected string `json:"Expected"`
		ID       string `json:"ID"`
	} `json:"ContainerdCommit"`
	Containers         int64       `json:"Containers"`
	ContainersPaused   int64       `json:"ContainersPaused"`
	ContainersRunning  int64       `json:"ContainersRunning"`
	ContainersStopped  int64       `json:"ContainersStopped"`
	CPUCfsPeriod       bool        `json:"CpuCfsPeriod"`
	CPUCfsQuota        bool        `json:"CpuCfsQuota"`
	Debug              bool        `json:"Debug"`
	DefaultRuntime     string      `json:"DefaultRuntime"`
	DockerRootDir      string      `json:"DockerRootDir"`
	Driver             string      `json:"Driver"`
	DriverStatus       [][]string  `json:"DriverStatus"`
	ExperimentalBuild  bool        `json:"ExperimentalBuild"`
	GenericResources   interface{} `json:"GenericResources"`
	HTTPProxy          string      `json:"HttpProxy"`
	HTTPSProxy         string      `json:"HttpsProxy"`
	ID                 string      `json:"ID"`
	IPv4Forwarding     bool        `json:"IPv4Forwarding"`
	Images             int64       `json:"Images"`
	IndexServerAddress string      `json:"IndexServerAddress"`
	InitBinary         string      `json:"InitBinary"`
	InitCommit         struct {
		Expected string `json:"Expected"`
		ID       string `json:"ID"`
	} `json:"InitCommit"`
	Isolation          string        `json:"Isolation"`
	KernelMemory       bool          `json:"KernelMemory"`
	KernelVersion      string        `json:"KernelVersion"`
	Labels             []interface{} `json:"Labels"`
	LiveRestoreEnabled bool          `json:"LiveRestoreEnabled"`
	LoggingDriver      string        `json:"LoggingDriver"`
	MemTotal           int64         `json:"MemTotal"`
	MemoryLimit        bool          `json:"MemoryLimit"`
	Ncpu               int64         `json:"NCPU"`
	NEventsListener    int64         `json:"NEventsListener"`
	NFd                int64         `json:"NFd"`
	NGoroutines        int64         `json:"NGoroutines"`
	Name               string        `json:"Name"`
	NoProxy            string        `json:"NoProxy"`
	OSType             string        `json:"OSType"`
	OomKillDisable     bool          `json:"OomKillDisable"`
	OperatingSystem    string        `json:"OperatingSystem"`
	Plugins            struct {
		Authorization interface{} `json:"Authorization"`
		Log           []string    `json:"Log"`
		Network       []string    `json:"Network"`
		Volume        []string    `json:"Volume"`
	} `json:"Plugins"`
	ProductLicense string `json:"ProductLicense"`
	RegistryConfig struct {
		AllowNondistributableArtifactsCIDRs     []interface{} `json:"AllowNondistributableArtifactsCIDRs"`
		AllowNondistributableArtifactsHostnames []interface{} `json:"AllowNondistributableArtifactsHostnames"`
		IndexConfigs                            struct {
			One0_161_30_207_11036 struct {
				Mirrors  []interface{} `json:"Mirrors"`
				Name     string        `json:"Name"`
				Official bool          `json:"Official"`
				Secure   bool          `json:"Secure"`
			} `json:"10.161.30.207:11036"`
			Docker_io struct {
				Mirrors  []string `json:"Mirrors"`
				Name     string   `json:"Name"`
				Official bool     `json:"Official"`
				Secure   bool     `json:"Secure"`
			} `json:"docker.io"`
		} `json:"IndexConfigs"`
		InsecureRegistryCIDRs []string `json:"InsecureRegistryCIDRs"`
		Mirrors               []string `json:"Mirrors"`
	} `json:"RegistryConfig"`
	RuncCommit struct {
		Expected string `json:"Expected"`
		ID       string `json:"ID"`
	} `json:"RuncCommit"`
	Runtimes struct {
		Runc struct {
			Path string `json:"path"`
		} `json:"runc"`
	} `json:"Runtimes"`
	SecurityOptions []string `json:"SecurityOptions"`
	ServerVersion   string   `json:"ServerVersion"`
	SwapLimit       bool     `json:"SwapLimit"`
	Swarm           struct {
		ControlAvailable bool        `json:"ControlAvailable"`
		Error            string      `json:"Error"`
		LocalNodeState   string      `json:"LocalNodeState"`
		NodeAddr         string      `json:"NodeAddr"`
		NodeID           string      `json:"NodeID"`
		RemoteManagers   interface{} `json:"RemoteManagers"`
	} `json:"Swarm"`
	SystemStatus interface{} `json:"SystemStatus"`
	SystemTime   string      `json:"SystemTime"`
	Warnings     []string    `json:"Warnings"`
}

type ImageList struct {
	Containers  int64    `json:"Containers"`
	Created     int64    `json:"Created"`
	ID          string   `json:"Id"`
	ParentID    string   `json:"ParentId"`
	RepoDigests []string `json:"RepoDigests"`
	RepoTags    []string `json:"RepoTags"`
	SharedSize  int64    `json:"SharedSize"`
	Size        int64    `json:"Size"`
	VirtualSize int64    `json:"VirtualSize"`
}

type ImageInspect struct {
	Architecture string `json:"Architecture"`
	Author       string `json:"Author"`
	Comment      string `json:"Comment"`
	Config       struct {
		ArgsEscaped  bool        `json:"ArgsEscaped"`
		AttachStderr bool        `json:"AttachStderr"`
		AttachStdin  bool        `json:"AttachStdin"`
		AttachStdout bool        `json:"AttachStdout"`
		Cmd          interface{} `json:"Cmd"`
		Domainname   string      `json:"Domainname"`
		Entrypoint   []string    `json:"Entrypoint"`
		Env          []string    `json:"Env"`
		Hostname     string      `json:"Hostname"`
		Image        string      `json:"Image"`
		Labels       interface{} `json:"Labels"`
		OnBuild      interface{} `json:"OnBuild"`
		OpenStdin    bool        `json:"OpenStdin"`
		StdinOnce    bool        `json:"StdinOnce"`
		Tty          bool        `json:"Tty"`
		User         string      `json:"User"`
		Volumes      interface{} `json:"Volumes"`
		WorkingDir   string      `json:"WorkingDir"`
	} `json:"Config"`
	Container       string `json:"Container"`
	ContainerConfig struct {
		ArgsEscaped  bool        `json:"ArgsEscaped"`
		AttachStderr bool        `json:"AttachStderr"`
		AttachStdin  bool        `json:"AttachStdin"`
		AttachStdout bool        `json:"AttachStdout"`
		Cmd          []string    `json:"Cmd"`
		Domainname   string      `json:"Domainname"`
		Entrypoint   []string    `json:"Entrypoint"`
		Env          []string    `json:"Env"`
		Hostname     string      `json:"Hostname"`
		Image        string      `json:"Image"`
		Labels       struct{}    `json:"Labels"`
		OnBuild      interface{} `json:"OnBuild"`
		OpenStdin    bool        `json:"OpenStdin"`
		StdinOnce    bool        `json:"StdinOnce"`
		Tty          bool        `json:"Tty"`
		User         string      `json:"User"`
		Volumes      interface{} `json:"Volumes"`
		WorkingDir   string      `json:"WorkingDir"`
	} `json:"ContainerConfig"`
	Created       string `json:"Created"`
	DockerVersion string `json:"DockerVersion"`
	GraphDriver   struct {
		Data struct {
			LowerDir  string `json:"LowerDir"`
			MergedDir string `json:"MergedDir"`
			UpperDir  string `json:"UpperDir"`
			WorkDir   string `json:"WorkDir"`
		} `json:"Data"`
		Name string `json:"Name"`
	} `json:"GraphDriver"`
	ID       string `json:"Id"`
	Metadata struct {
		LastTagTime string `json:"LastTagTime"`
	} `json:"Metadata"`
	Os          string        `json:"Os"`
	Parent      string        `json:"Parent"`
	RepoDigests []interface{} `json:"RepoDigests"`
	RepoTags    []string      `json:"RepoTags"`
	RootFS      struct {
		Layers []string `json:"Layers"`
		Type   string   `json:"Type"`
	} `json:"RootFS"`
	Size        int64 `json:"Size"`
	VirtualSize int64 `json:"VirtualSize"`
}
