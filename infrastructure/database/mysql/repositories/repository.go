package repositories

import (
	"GoMetricsExporter/app/entities/mysql/BabyPetControl/model"
	"GoMetricsExporter/app/entities/mysql/BabyPetControl/table"
	"GoMetricsExporter/infrastructure/database/mysql/config"
	"fmt"

	"github.com/go-jet/jet/v2/mysql"
)

// ControlRepository é a implementação do gateway para acessar dados
type ControlRepository struct{}

func (x ControlRepository) GetMatrizes() ([]model.Matrices, error) {
	db := config.DB
	if db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	// Define a tabela usando o schema importado
	matrizesTable := table.Matrices

	// Define a consulta SQL
	stmt := mysql.SELECT(
		matrizesTable.ID,
		matrizesTable.Name,
		matrizesTable.Color,
		matrizesTable.Age,
	).
		FROM(matrizesTable)

	var matrizesList []model.Matrices
	err := stmt.Query(db, &matrizesList)
	if err != nil {
		return nil, err
	}

	return matrizesList, nil
}

func (x ControlRepository) GetPadreadores() ([]model.Matrices, error) {
	db := config.DB
	if db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	// Define a tabela usando o schema importado
	matrizesTable := table.Matrices

	// Define a consulta SQL
	stmt := mysql.SELECT(
		matrizesTable.ID,
		matrizesTable.Name,
		matrizesTable.Color,
		matrizesTable.Age,
	).
		FROM(matrizesTable)

	var matrizesList []model.Matrices
	err := stmt.Query(db, &matrizesList)
	if err != nil {
		return nil, err
	}

	return matrizesList, nil
}
