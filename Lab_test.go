package testlab

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `valid:"required~name cannot be blank"`
	Email string `gorm:"uniqueIndex"valid:"email~กรุณากรอกอีเมลให้ถูกต้อง"`
}

// test name cannot be blank
func TestUserValidate(t *testing.T) {

	g := NewGomegaWithT(t)

	n := User{

		Name:  "",
		Email: "wanatsanan@gmail.com",
	}

	ok, err := govalidator.ValidateStruct(n)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("name cannot be blank"))

}

// test email ผิด
func TestEmailValidate(t *testing.T) {

	g := NewGomegaWithT(t)

	n := User{
		Name:  "wanat",
		Email: "wanatsanan",
	}

	ok, err := govalidator.ValidateStruct(n)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("กรุณากรอกอีเมลให้ถูกต้อง"))
}

// test name cannot be blank and email ผิด
func TestValidate(t *testing.T) {

	g := NewGomegaWithT(t)

	n := User{
		Name:  "",
		Email: "wanatsanan",
	}

	ok, err := govalidator.ValidateStruct(n)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("name cannot be blank;กรุณากรอกอีเมลให้ถูกต้อง"))

}

func TestTrueValidate(t *testing.T) {

	g := NewGomegaWithT(t)

	n := User{
		Name:  "wanat",
		Email: "wanatsanan@gmail.com",
	}

	ok, err := govalidator.ValidateStruct(n)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
	// g.Expect(err.Error()).To(Equal("name cannot be blank;กรุณากรอกอีเมลให้ถูกต้อง"))

}

