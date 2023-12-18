package service

import (
	"feature/internal/model"
	"feature/internal/repository"
)

type PatientService interface {
	GetPatientById(id int64) (*model.Patient, error)
}

func NewPatientService(service *Service, patientRepository repository.PatientRepository) PatientService {
	return &patientService{
		Service:        service,
		patientRepository: patientRepository,
	}
}

type patientService struct {
	*Service
	patientRepository repository.PatientRepository
}

func (s *patientService) GetPatientById(id int64) (*model.Patient, error) {
	return s.patientRepository.FirstById(id)
}
