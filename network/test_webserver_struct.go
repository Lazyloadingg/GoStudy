package main

type Person struct {
	Name     string `json:"name" form:"name" `
	Age      int    `json:"age" form:"age" `
	ID       int    `json:"id" form:"id"`
	Phone    int    `json:"phone" form:"phone" db:"phone" binding:"required"`
	Gender   string `json:"gender" form:"gender" db:"gender" binding:"required"`
	IDNumber int    `json:"idNumber" form:"idNumber" db:"id_number" binding:"required"`
	NickName string `json:"nickName" form:"nickName" db:"nick_name" binding:"required"` //添加 binding:"required" 则在使用反射时，必须有这个字段，否则报错  'Person.IDNumber' Error:Field validation for 'IDNumber' failed on the 'required' tag
}
