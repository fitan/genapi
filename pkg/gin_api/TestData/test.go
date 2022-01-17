package TestData

import (
	"github.com/fitan/genapi/pkg/gin_api/TestData/nest"
	"k8s.io/api/apps/v1"
	"time"
)

type Form struct {
	Name string `json:"name" ngform:"title=邮箱,format=id-card"`
	Age int `json:"age"`
	Notes []struct{
		Node string `json:"node"`
	} `json:"notes" ngform:"title=笔记"`
}

type K8sDeploy struct {
	Result v1.Deployment `json:"result"`
	Service v1.StatefulSet `json:"service"`
}

type UserResult struct {
	//Code AliaseInt                 `json:"code"`
	Data      map[string]nest.Nest      `json:"data"`
	M         map[nest.Fater]nest.Fater `json:"m"`
	N         nest.Nest                 `json:"n"`
	Err       interface{}               `json:"err"`
	User      []*User                      `json:"user"`
	AliaseInt `json:"aint"`
	UserIncludes
	//this is nest
	nest.Nest
	// this is fater
	nest.Fater
	//UserIncludes
	//ATime time.Time
}

type AliaseInt int

type User struct {
	ID         int `json:"id"`
	Age        int `json:"age"`
	Name       string
	Nest       []nest.Nest          `json:"nest"`
	M          map[string]nest.Nest `json:"m"`
	Fater      nest.Fater           `json:"fater"`
	Time       time.Time            `json:"time"`
	//UserResult *UserResult
}

type UserIncludes struct {
	// Association query Multiple choice:
	// role_binding
	// alert
	Includes []string `form:"includes" json:"includes" binding:"dive,oneof=role_binding alert"`
}

type Test1T []nest.Nest

type Worker struct {
	Metadata  K8sKey          `json:"metadata"`
	Component WorkerComponent `json:"component"`
}

type K8sKey struct {
	Namespace string `json:"namespace,omitempty" uri:"namespace"`
	Name      string `json:"name,omitempty" uri:"name"`
}

type WorkerComponent struct {
	Properties struct {
		Image            string    `json:"image"`
		ImagePullPolicy  *string   `json:"imagePullPolicy,omitempty"`
		ImagePullSecrets *string   `json:"imagePullSecrets,omitempty"`
		Cmd              *[]string `json:"cmd,omitempty"`
		Env              *[]struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"env,omitempty"`
		Cpu    *string `json:"cpu,omitempty"`
		Memory *string `json:"memory,omitempty"`

		Volumes *[]struct {
			Name      string `json:"name"`
			MountPath string `json:"mountPath"`
			Type      string `json:"type"`
		} `json:"volumes,omitempty"`

		LivenessProbe  *HealthProbe `json:"livenessProbe,omitempty"`
		ReadinessProbe *HealthProbe `json:"readinessProbe,omitempty"`
	} `json:"properties"`

	Traits struct {
		Ingress       *TraitIngress  `json:"ingress,omitempty"`
		Labels        *Labels        `json:"labels,omitempty"`
		Annotations   *Annotations   `json:"annotations,omitempty"`
		Sidecar       *Sidecar       `json:"sidecar,omitempty"`
		Expose        *Expose        `json:"expose,omitempty"`
		InitContainer *InitContainer `json:"initContainer,omitempty"`
		ConfigMap     *ConfigMap     `json:"configMap,omitempty"`
		Pvc           *Pvc           `json:"pvc,omitempty"`
		Scaler        *Scaler        `json:"scaler,omitempty"`
		Ports         *MyPorts       `json:"ports,omitempty"`
		MyEnv         *MyEnv         `json:"myEnv,omitempty"`
		Test          *struct{}      `json:"test,omitempty"`
	} `json:"traits,omitempty"`
}

type HealthProbe struct {
	Exec    *[]string `json:"exec"`
	HttpGet *struct {
		Path        string `json:"path"`
		Port        string `json:"port"`
		HttpHeaders *[]struct {
			Name  string
			Value string
		} `json:"httpHeaders"`
	} `json:"httpGet"`
	TcpSocket *struct {
		Port int `json:"port"`
	} `json:"tcpSocket"`
	InitialDelaySeconds int `json:"initialDelaySeconds"`
	PeriodSeconds       int `json:"periodSeconds"`
	TimeoutSeconds      int `json:"timeoutSeconds"`
	SuccessThreshold    int `json:"successThreshold"`
	FailureThreshold    int `json:"failureThreshold"`
}

type TraitIngress struct {
	Domain string         `json:"domain"`
	Http   map[string]int `json:"http"`
}


type Labels map[string]string



type Annotations map[string]string



type Sidecar struct {
	Name    string    `json:"name"`
	Image   string    `json:"image"`
	Cmd     *[]string `json:"cmd"`
	Args    *[]string `json:"args"`
	Volumes *[]struct {
		Name string `json:"name"`
		Path string `json:"path"`
	} `json:"volumes"`
}



type Scaler struct {
	Replicas int `json:"replicas"`
}



type InitContainer struct {
	Name          string    `json:"name"`
	Image         string    `json:"image"`
	Cmd           *[]string `json:"cmd"`
	Args          *[]string `json:"args"`
	MountName     string    `json:"mountName"`
	AppMountPath  string    `json:"appMountPath"`
	InitMountPath string    `json:"initMountPath"`
}



type Expose struct {
	Port []int  `json:"port"`
	Type string `json:"type"`
}



type ConfigMap struct {
	Volumes []struct {
		Name      string             `json:"name"`
		MountPath string             `json:"mountPath"`
		readOnly  bool               `json:"readOnly"`
		Data      *map[string]string `json:"data"`
	} `json:"volumes"`
}



type Pvc struct {
	ClaimName        string   `json:"claimName"`
	VolumeMode       string   `json:"volumeMode"`
	VolumeName       *string  `json:"volumeName"`
	AccessModes      []string `json:"accessModes"`
	StorageClassName *string  `json:"storageClassName"`
	Resources        struct {
		Requests string `json:"requests"`
		Limits   string `json:"limits"`
	} `json:"resources"`

	VolumesToMount []struct {
		Name       string `json:"name"`
		DevicePath string `json:"devicePath"`
		MountPath  string `json:"mountPath"`
	} `json:"volumesToMount"`
}



type MyPorts struct {
	Ports []struct {
		ContainerPort int    `json:"containerPort"`
		Ptotocol      string `json:"ptotocol"`
	} `json:"ports"`
}



type MyEnv struct {
}