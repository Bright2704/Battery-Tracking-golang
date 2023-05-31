package service

import "api/Battery-Tracking-go/entity"

type RelatedNumberService interface {
	CreateRelatedNumber(*entity.RelatedNumber) error
	GetRelatedNumber(*string) (*entity.RelatedNumber, error)
	GetAll() ([]*entity.RelatedNumber, error)
	UpdateRelatedNumber(*entity.RelatedNumber) error
	DeleteRelatedNumber(*string) error
}