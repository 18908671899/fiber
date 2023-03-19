package models

type Test struct {
	Model
	Name string `query:"name" json:"name" xml:"name" form:"name"`
}

func GetTest(pageNum int, pageSize int, maps interface{}) (tests []Test) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tests)

	return
}

func AddTest(test *Test) bool {
	db.Create(&test)

	return true
}

func EditTest(id int, data interface{}) bool {
	db.Model(&Test{}).Where("id = ?", id).Updates(data)
	return true
}

func DeleteTest(id int) bool {
	db.Where("id = ?", id).Delete(Test{})
	return true
}

func ExistTestById(id int) bool {
	var test Test
	db.Select("id").Where("id = ?").First(&test)

	return test.ID > 0
}
