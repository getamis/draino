package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"go.uber.org/zap"
)

type ASGManager struct {
	Session *session.Session
	Logger  *zap.Logger
}

func NewASGManager(logger *zap.Logger) *ASGManager {
	sess, err := session.NewSession(&aws.Config{})
	if err != nil {
		logger.Fatal("Error creating AWS session:", zap.Error(err))
	}
	logger.Info("AWS session created")
	return &ASGManager{Session: sess, Logger: logger}
}

func (m *ASGManager) SetASGInstanceUnhealthy(instanceID string) error {
	svc := autoscaling.New(m.Session)

	input := &autoscaling.SetInstanceHealthInput{
		HealthStatus:             aws.String("Unhealthy"),
		InstanceId:               aws.String(instanceID),
		ShouldRespectGracePeriod: aws.Bool(true),
	}

	_, err := svc.SetInstanceHealth(input)
	if err != nil {
		m.Logger.Fatal("Error setting instance health:", zap.Error(err))
		return err
	}
	m.Logger.Info("Set instance health to unhealthy", zap.String("instanceID", instanceID))
	return nil
}
