func UpdateBook(c *gin.Context) {
	var book database.Book
	id := c.Param("id")
	err := c.ShouldBind(&book)
   
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
  
	var updateBook database.Book
	res := database.DB.Model(&updateBook).Where("id = ?", id).Updates(book)
  
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "book not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
	return
 }
 