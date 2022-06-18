package app_handler

// type GetFilesArgs struct {
// 	Ids []string `json:"ids"`
// }

// func (a GetFilesArgs) Validate() ([]interface{}, error) {
// 	rules := []*validation.FieldRules{
// 		validation.Field(&a.Ids, validation.Required),
// 	}

// 	error := validation.ValidateStruct(&a, rules...)
// 	data := util.ErroObjToArray(error)

// 	return data, error
// }

// type GetFile struct {
// 	ID  string `json:"id"`
// 	Url string `json:"url"`
// }

// func (pkg AppHandler) UploadHandler(c *gin.Context) {
// 	user, _ := auth.GetUserFromContext(c)

// 	multipart, err := c.Request.MultipartReader()
// 	if err != nil {
// 		log.Println("Failed to create MultipartReader", err)
// 		error := util.ResponseError{
// 			Message: err.Error(),
// 		}
// 		error.Send(c, http.StatusUnauthorized)
// 		return
// 	}

// 	// var files []mediatypes.File
// 	var getFiles []GetFile

// 	for {
// 		mimePart, err := multipart.NextPart()
// 		if err == io.EOF {
// 			log.Printf("multipart.NextPart: %v", err)
// 			break
// 		}
// 		if err != nil {
// 			log.Printf("Error reading multipart section: %v", err)
// 			break
// 		}
// 		_, params, err := mime.ParseMediaType(mimePart.Header.Get("Content-Disposition"))
// 		if err != nil {
// 			log.Printf("Invalid Content-Disposition: %v", err)
// 		}

// 		b, err := ioutil.ReadAll(mimePart)
// 		if err != nil {
// 			log.Printf("Error reading multipart section: %v", err)
// 			break
// 		}

// 		namespace, _ := uuid.Parse("b9cfdb9d-f741-4e1f-89ae-fac6b2a5d740")
// 		checkSum := uuid.NewSHA1(namespace, b)
// 		// fmt.Println(checkSum)

// 		fileRecord, err := media_util.GetFileByCheckSum(checkSum.String(), pkg.GormDB.DB)
// 		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
// 			log.Printf("db error: %v", err)
// 			continue
// 		}

// 		if fileRecord != nil {
// 			// files = append(files, *fileRecord)
// 			getFiles = append(getFiles, GetFile{
// 				ID:  fileRecord.ID,
// 				Url: media_util.BuildUrl(*fileRecord),
// 			})
// 			log.Println("Record already exists with same checksum")
// 			continue
// 		}

// 		fileName := params["filename"]
// 		// dst := filepath.Join("/public", media_util.GetYearMontPath())
// 		dst := media_util.GetYearMontPath()
// 		media_util.CreateDir(dst)
// 		ext := strings.ToLower(filepath.Ext(fileName))
// 		newFileName := media_util.GenerateUniqueFileName(media_util.FileNameWithoutExtSliceNotation(fileName), ext)
// 		path := filepath.Join(dst, newFileName)
// 		fullPath := filepath.Join(media_util.GetStorageDir(), path)

// 		f, err := os.Create(fullPath)
// 		if err != nil {
// 			log.Printf("os.Create: %v", err)
// 			continue
// 		}
// 		f.Write(b)
// 		f.Close()

// 		fileEntity := &media_type.File{
// 			UserID:       user.ID,
// 			CheckSum:     checkSum.String(),
// 			FileName:     newFileName,
// 			OriginalName: fileName,
// 			Mime:         mimePart.Header.Get("Content-Type"),
// 			Size:         uint(len(b)), // in bytes
// 			Ext:          ext[1:],
// 			Destination:  dst,
// 			Path:         path,
// 		}

// 		err = pkg.GormDB.DB.Create(fileEntity).Error
// 		// fmt.Println(err)

// 		if err != nil {
// 			log.Printf("os.Create: %v", err)
// 			continue
// 		}

// 		// fmt.Println(fileEntity)
// 		getFiles = append(getFiles, GetFile{
// 			ID:  fileEntity.ID,
// 			Url: media_util.BuildUrl(*fileEntity),
// 		})

// 	}

// 	// fmt.Println(files)
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": getFiles,
// 	})
// }
