package get

import (
	"net/http"

	"github.com/AlejandroPintosAlcarazo/asteroid.API/configs"
	"github.com/AlejandroPintosAlcarazo/asteroid.API/handlers"
	"github.com/AlejandroPintosAlcarazo/asteroid.API/models"
	"github.com/AlejandroPintosAlcarazo/asteroid.API/responses"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAsteroidByID(c echo.Context) error {
	ctx, cancel, AsteroidCollection, err := handlers.SetupDBContext()
	if err != nil {
		return handlers.InternalServerErrorResponse(c, err)
	}
	defer cancel()

	asteroidID := c.Param("id")
	var objID primitive.ObjectID

	// Intentar convertir a ObjectID
	objID, err = primitive.ObjectIDFromHex(asteroidID)
	if err != nil {
		// Si falla, buscar en la colección de mapeo
		MappingCollection := configs.GetMappingCollection(configs.ConnectDB(), "asteroid_mappings")
		var mapping bson.M
		err := MappingCollection.FindOne(ctx, bson.M{"name": asteroidID}).Decode(&mapping)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return handlers.AsteroidExistsErrorJSON(c, "Asteroid does not exist")
			}
			return handlers.InternalServerErrorResponse(c, err)
		}
		objID = mapping["id"].(primitive.ObjectID)
	}

	var asteroid models.Asteroid
	err = AsteroidCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&asteroid)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return handlers.AsteroidExistsErrorJSON(c, "Asteroid does not exist")
		}
		return handlers.InternalServerErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, responses.AsteroidResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &echo.Map{"data": asteroid},
	})
}
