package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const uploadFolder="./uploads"

func main(){
	if _,err:=os.Stat(uploadFolder);os.IsNotExist(err){
		os.Mkdir(uploadFolder,os.ModePerm)
	}

	app:=fiber.New()

	app.Post("/upload",uploadFile)
	app.Get("/files/:filename",downloadFile)
	app.Get("/files",listFiles)
	app.Delete("/files/:filename",deleteFile)

	log.Fatal(app.Listen(":8080"))
}

func uploadFile(c *fiber.Ctx)error{
	file,err:=c.FormFile("file")
	if err!=nil{
		return c.Status(400).JSON(fiber.Map{"error":"File upload failed"})

	}

	filepath:=filepath.Join(uploadFolder,file.Filename)
	err=c.SaveFile(file,filepath)
	if err!=nil{
		
	return c.Status(500).JSON(fiber.Map{"error":"Filed to save the file"})

	}

	return c.JSON(fiber.Map{"message": "File uploaded successfully","url":fmt.Sprintf("/files/%s",file.Filename)})

}



func downloadFile(c *fiber.Ctx) error {
	filename := c.Params("filename")
	filePath := filepath.Join(uploadFolder, filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.Status(404).JSON(fiber.Map{"error": "File not found"})
	}

	return c.SendFile(filePath)
}


func listFiles(c *fiber.Ctx) error {
	files, err := ioutil.ReadDir(uploadFolder)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to list files"})
	}
	var fileList []string

	for _, file := range files {
		fileList = append(fileList, file.Name())
	}
	return c.JSON(fiber.Map{"files": fileList})

}

func deleteFile(c *fiber.Ctx) error {
	filename := c.Params("filename")
	filePath := filepath.Join(uploadFolder, filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.Status(404).JSON(fiber.Map{"error": "File not found"})
	}

	err := os.Remove(filePath)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete file"})
	}

	return c.JSON(fiber.Map{"message": "File deleted successfully"})
}

