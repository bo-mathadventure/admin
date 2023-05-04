package workadventure

import (
	"context"
	"github.com/bo-mathadventure/admin/config"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/bo-mathadventure/admin/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func NewTextureHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Get("/list", getWokaList(ctx, db))
}

type TextureCollectionTexture struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	Position int    `json:"position"`
}
type TextureCollection struct {
	Name     string                     `json:"name"`
	Position int                        `json:"position"`
	Textures []TextureCollectionTexture `json:"textures"`
}

type TextureResponse struct {
	Collections []TextureCollection `json:"collections"`
	Required    bool                `json:"required,omitempty"`
}

func getWokaList(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		allTextures, err := db.Textures.Query().All(ctx)
		if err != nil {
			return handler.HandleInvalidLogin(c)
		}
		availableTexturesList := map[string]TextureResponse{}
		for _, texture := range allTextures {
			if _, ok := availableTexturesList[texture.Layer]; !ok {
				availableTexturesList[texture.Layer] = TextureResponse{
					Collections: []TextureCollection{
						{
							Name:     "default",
							Position: 0,
							Textures: []TextureCollectionTexture{},
						},
					},
					Required: utils.Contains([]string{"body", "eyes", "accessory"}, texture.Layer),
				}
			}

			availableTexturesList[texture.Layer].Collections[0].Textures = append(availableTexturesList[texture.Layer].Collections[0].Textures, TextureCollectionTexture{
				ID:       texture.Texture,
				Name:     texture.Texture,
				URL:      strings.ReplaceAll(texture.URL, "%FRONTEND_URL%", config.GetConfig().FrontendURL),
				Position: len(availableTexturesList[texture.Layer].Collections[0].Textures),
			})
		}

		return c.JSON(availableTexturesList)
	}
}
