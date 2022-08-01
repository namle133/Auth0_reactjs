package service

import (
	"CMS/domain"
	"fmt"

	"github.com/jinzhu/gorm"
)

type UserService struct {
	Db *gorm.DB
}

func (us *UserService) GetNewsByTitle(title string) ([]domain.News, error) {
	if title == "" {
		return []domain.News{}, fmt.Errorf("title cannot be empty")
	}
	var n []domain.News
	err := us.Db.Where("title = ?", title).Find(&n).Error
	if err != nil {
		return []domain.News{}, err
	}
	return n, nil
}

func (us *UserService) GetNews() ([]domain.News, error) {
	var n []domain.News
	us.Db.LogMode(true)
	err := us.Db.Debug().Order("id asc").Find(&n).Error
	if err != nil {
		return []domain.News{}, err
	}
	return n, nil
}

func CheckUsersEmpty(u domain.Users) error {
	if len(u.Email) == 0 {
		return fmt.Errorf("%s is empty", "email")
	} else if len(u.Mobile) == 0 {
		return fmt.Errorf("%s is empty", "mobile")
	} else if len(u.Username) == 0 {
		return fmt.Errorf("%s is empty", "username")
	} else if len(u.Referrer) == 0 {
		return fmt.Errorf("%s is empty", "referrer")
	}
	return nil
}

func (us UserService) AddUsers(u domain.Users) error {
	err := CheckUsersEmpty(u)
	if err != nil {
		return err
	}
	err = us.Db.Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}
func CheckNewsEmpty(n domain.News) error {
	if len(n.Title) == 0 {
		return fmt.Errorf("%s is empty", "title")
	} else if len(n.TitleContent) == 0 {
		return fmt.Errorf("%s is empty", "title_content")
	} else if len(n.Content) == 0 {
		return fmt.Errorf("%s is empty", "content")
	} else if len(n.CreatorContent) == 0 {
		return fmt.Errorf("%s is empty", "creatorContent")
	}
	return nil
}
func (us UserService) AddContents(n domain.News) error {
	err := CheckNewsEmpty(n)
	if err != nil {
		return err
	}
	err = us.Db.Create(&n).Error
	if err != nil {
		return err
	}
	return nil
}

func (us UserService) UpdateNewsLikes(id int, likes int) error {
	var n domain.News
	err := us.Db.Model(&n).Where("id=?", id).Update("likes", likes).Error
	if err != nil {
		return err
	}
	return nil

}

func (us UserService) UpdateContent(ct domain.News) error {
	err := CheckNewsEmpty(ct)
	if err != nil {
		return err
	}
	var n domain.News
	err = us.Db.Model(&n).Where("title=? AND title_content=? AND creator_content=?", ct.Title, ct.TitleContent, ct.CreatorContent).Update("content", ct.Content).Error
	if err != nil {
		return err
	}
	return nil
}

func CheckContentsEmpty(dc domain.DeleteContent) error {
	if len(dc.Title) == 0 {
		return fmt.Errorf("%s is empty", "title")
	} else if len(dc.Title_content) == 0 {
		return fmt.Errorf("%s is empty", "title_content")
	} else if len(dc.Username) == 0 {
		return fmt.Errorf("%s is empty", "username")
	}
	return nil
}

func (us UserService) DeleteContent(ct domain.DeleteContent) error {
	err := CheckContentsEmpty(ct)
	if err != nil {
		return err
	}
	var n domain.News
	err = us.Db.Where("title=? AND title_content=? AND creator_content=?", ct.Title, ct.Title_content, ct.Username).Delete(&n).Error
	if err != nil {
		return err
	}
	return nil
}
