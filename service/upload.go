package service

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"restapi/model"
	"time"

	"github.com/gocql/gocql"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func (s *CommonService) HandleSendFeedback(c echo.Context) error {
	ctx := c.Request().Context()
	logger := zerolog.Ctx(ctx)

	logger.Debug().Msg("handle send feedback")

	fmt.Println("1")

	// file upload
	attachment, _ := upload(c)

	feedback := &model.Feedback{
		Id:         gocql.TimeUUID(),
		SenderId:   gocql.UUID{},
		Subject:    c.FormValue("subject"),
		Message:    c.FormValue("message"),
		Device:     c.FormValue("device"),
		Attachment: attachment,
		CreatedAt:  time.Now().UTC(),
	}

	fmt.Println("feedback", feedback)

	return errors.WithStack(c.NoContent(http.StatusOK))
}

func (s *CommonService) HandleFileUpload1(c echo.Context) error {
	ctx := c.Request().Context()
	logger := zerolog.Ctx(ctx)

	logger.Debug().Msg("handled file upload 1")

	attachment, _ := upload(c)

	logger.Debug().Msgf("the file has been uploaded: %s", attachment)

	return errors.WithStack(c.NoContent(http.StatusOK))
}

func (s *CommonService) HandleFileUpload2(c echo.Context) error {
	ctx := c.Request().Context()
	logger := zerolog.Ctx(ctx)

	logger.Debug().Msg("handled file upload 2")

	attachment, _ := upload(c)

	logger.Debug().Msgf("the file has been uploaded: %s", attachment)

	return errors.WithStack(c.NoContent(http.StatusOK))
}

func fileUpload(c echo.Context) (string, error) {

	/*
		file, err := c.FormFile("attachment")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := os.Create(file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully.</p>", file.Filename))
	*/

	ctx := c.Request().Context()
	logger := zerolog.Ctx(ctx)

	// file upload
	var attachment string = ""

	fileHeader, err := c.FormFile("attachment")
	if err != nil {
		logger.Error().Err(err).Msg("failed to get file")
		return attachment, errors.WithStack(err)
	}

	if fileHeader != nil {
		//f, _ := fileHeader.Open()
		logger.Debug().Msg(fileHeader.Filename)
		attachment = fileHeader.Filename
	}

	return attachment, nil
}

func upload(c echo.Context) (string, error) {
	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("email")

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return "", err
	}
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	//return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, name, email))
	fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, name, email)
	return "", nil
}
