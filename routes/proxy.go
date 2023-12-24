package routes

import (
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Proxy(c *fiber.Ctx) error {

	url := c.Query("url")

	if len(url) == 0 {
		return c.JSON(fiber.Map{"message": "invalid url", "status": "fail"})
	}

	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return c.JSON(fiber.Map{"message": "invalid url", "status": "fail"})
	}

	res, err := httpClient.Do(req)

	if err != nil {
		return c.JSON(fiber.Map{"message": "invalid url", "status": "fail"})
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "something went wrong!", "status": "fail"})
	}

	contentType := res.Header.Get("Content-Type")

	c.Set("Content-Type", contentType)
	return c.Send(body)
}
