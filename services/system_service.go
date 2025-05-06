package services

import (
	"fmt"
	"time"

	"github.com/naufal/simba-qr-app/models"
	"github.com/naufal/simba-qr-app/repository"
	"github.com/naufal/simba-qr-app/services/requests"
	"github.com/naufal/simba-qr-app/utils"
)

func GetSystemByID(id string) (models.System, error) {
	return repository.GetSystemByID(id)
}

func CreateSystem(input requests.PostSystemRequest) (models.System, error) {
	data := input.Keygen + ":" + input.Identifier
	encrypted, err := utils.EncryptAES(data)

	if err != nil {
		fmt.Println("Encryption failed:", err)
		return models.System{}, err
	}

	expired := time.Now().Unix() + 300

	inputModel := models.System{
		Key:     encrypted,
		Expired: expired,
	}

	return repository.CreateSystem(inputModel)
}

func DeleteSystem(id string) error {
	return repository.DeleteSystem(id)
}
