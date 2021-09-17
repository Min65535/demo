package dao

import (
	"context"
	"demo/dfm-test/inter/publisher/biz"
	"fmt"
	"github.com/dipperin/go-ms-toolkit/json"
	dc "github.com/docker/cli/cli/command"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"k8s.io/api/apps/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ToolClient struct {
	docCli client.APIClient
	kbsCli kubernetes.Interface
}


func NewToolClient() biz.ToolCommand {
	// docker := exec.Command("docker")
	// kub := exec.Command("kubectl")
	cl, err := dc.NewDockerCli()
	if err != nil {
		panic("error NewDockerCli")
	}

	return &ToolClient{docCli: cl.Client()}
}

func (t *ToolClient) ImagePull(image string) error {
	resp, err := t.docCli.ImagePull(context.Background(), image, types.ImagePullOptions{})
	if err != nil {
		return errors.Wrapf(err, "failed to pull %s", image)
	}
	defer resp.Close()
	return nil
}

func (t *ToolClient) ImagePush(image string) error {
	resp, err := t.docCli.ImagePush(context.Background(), image, types.ImagePushOptions{})
	if err != nil {
		return errors.Wrapf(err, "failed to pull %s", image)
	}
	defer resp.Close()
	return nil
}

func (t *ToolClient) GetDeployment(namespace, name string) (*v1beta2.Deployment, error) {
	dp, err := t.kbsCli.AppsV1beta2().Deployments(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	fmt.Println("deployment:", json.StringifyJson(dp))
	return dp, nil
}

func (t *ToolClient) UpdateDeploymentImage(namespace, name, imageNew string) error {
	dp, err := t.GetDeployment(namespace, name)
	if err != nil {
		return err
	}
	dp.Spec.Template.Spec.Containers[0].Image = imageNew
	fmt.Println("change deployment:", json.StringifyJson(dp))
	ndp, err := t.kbsCli.AppsV1beta2().Deployments(namespace).Update(context.Background(), dp, v1.UpdateOptions{})
	if err != nil {
		return err
	}
	fmt.Println("new deployment:", json.StringifyJson(ndp))
	return nil
}

