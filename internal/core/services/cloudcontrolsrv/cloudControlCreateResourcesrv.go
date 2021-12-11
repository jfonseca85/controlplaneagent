//Cria resources usando o aws sdk go v2 - service cloudcontrol
package cloudcontrolsrv

import (
	"context"
	"log"

	"github.com/jfonseca85/controlplaneagent/internal/core/domain"
	"github.com/jfonseca85/controlplaneagent/internal/types"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

type Service struct {
	controlsdk *cloudcontrol.Client
}

func New(controlsdk *cloudcontrol.Client) Service {
	return Service{
		controlsdk: controlsdk,
	}
}

func (srv *Service) Create(model domain.CloudControlModel) (*domain.ProgressEvent, error) {
	//Metodo que destinado para criar os recursos usando o cloudcontrol
	log.Printf("Chamando o CreateResource>>> ")
	output, err := srv.controlsdk.CreateResource(context.TODO(), toCreateResourceInput(model))
	progressEvent := toCreateResourceProgressEvent(output)
	if err != nil {
		return nil, err
	}
	return progressEvent, nil
}

func toCreateResourceInput(model domain.CloudControlModel) *cloudcontrol.CreateResourceInput {
	result := cloudcontrol.CreateResourceInput{
		TypeName:      &(model.TypeName),
		DesiredState:  &(model.DesiredState),
		ClientToken:   &(model.ClientToken),
		RoleArn:       &(model.RoleArn),
		TypeVersionId: &(model.TypeVersionId),
	}
	return &result
}

func toCreateResourceProgressEvent(model *cloudcontrol.CreateResourceOutput) *domain.ProgressEvent {
	result := domain.ProgressEvent{
		ErrorCode:       types.HandlerErrorCode(model.ProgressEvent.ErrorCode),
		EventTime:       model.ProgressEvent.EventTime,
		Identifier:      model.ProgressEvent.Identifier,
		Operation:       types.Operation(model.ProgressEvent.Operation),
		OperationStatus: types.OperationStatus(model.ProgressEvent.Operation),
		RequestToken:    model.ProgressEvent.RequestToken,
		ResourceModel:   model.ProgressEvent.ResourceModel,
		RetryAfter:      model.ProgressEvent.RetryAfter,
		StatusMessage:   model.ProgressEvent.StatusMessage,
		TypeName:        model.ProgressEvent.TypeName,
	}
	return &result
}
