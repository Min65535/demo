package biz

import "k8s.io/api/apps/v1beta2"

type ToolCommand interface {
	ImagePull(image string) error
	ImagePush(image string) error
	GetDeployment(namespace, name string) (*v1beta2.Deployment, error)
	UpdateDeploymentImage(namespace, name, imageNew string) error
}

type ToolUseCase struct {
	tool ToolCommand
}

func NewToolUseCase(tool ToolCommand) *ToolUseCase {
	return &ToolUseCase{tool: tool}
}

func (t ToolUseCase) ImagePull(image string) error {
	panic("implement me")
}

func (t ToolUseCase) ImagePush(image string) error {
	panic("implement me")
}

func (t ToolUseCase) GetDeployment(namespace, name string) (*v1beta2.Deployment, error) {
	panic("implement me")
}

func (t ToolUseCase) UpdateDeploymentImage(namespace, name, imageNew string) error {
	panic("implement me")
}
