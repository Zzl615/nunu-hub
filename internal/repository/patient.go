package repository

import (
	"feature/internal/model"
)

type PatientRepository interface {
	FirstById(id int64) (*model.Patient, error)
}

func NewPatientRepository(repository *Repository) PatientRepository {
	return &patientRepository{
		Repository: repository,
	}
}

type patientRepository struct {
	*Repository
}

func (r *patientRepository) FirstById(id int64) (*model.Patient, error) {
	var patient model.Patient
	// TODO: query db
	return &patient, nil
}
