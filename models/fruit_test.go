package models_test

import (
	"testing"
	
	"github.com/relax-space/go-api/models"
	"github.com/pangpanglabs/goutils/test"
)

func TestFruitCRUD(t *testing.T) {
	t.Run("GetAll", func(t *testing.T) {
		count, fruits, err := models.Fruit{}.GetAll(ctx,[]string{"id"},[]string{"asc"},0, 1)
		test.Ok(t, err)
		test.Assert(t,count>=1,"result must be greater than 1")
		v :=fruits[0]
		test.Equals(t, int64(1), v.Id)
		test.Equals(t, "apple", v.Code)
		test.Equals(t, "apple", v.Name)
		test.Equals(t, int64(11), v.Price)
		test.Equals(t, "10001", v.StoreCode)
	})
	t.Run("Get", func(t *testing.T) {
		has, v, err := models.Fruit{}.GetById(ctx, 1)
		test.Ok(t, err)
		test.Equals(t, true, has)
		test.Equals(t, int64(1), v.Id)
		test.Equals(t, "apple", v.Code)
		test.Equals(t, "apple", v.Name)
		test.Equals(t, int64(11), v.Price)
		test.Equals(t, "10001", v.StoreCode)
	})
	t.Run("Create", func(t *testing.T) {
		code := "apple"
		f := &models.Fruit{
			Code: code,
		}
		affectedRow, err := f.Create(ctx)
		test.Ok(t, err)
		test.Equals(t, int64(1), affectedRow)
		test.Assert(t, f.Id != int64(0), "create failure")
	})

	t.Run("Update", func(t *testing.T) {
		var price int64 = 10
		f := &models.Fruit{
			Price: price,
		}
		affectedRow, err := f.Update(ctx, 1)
		test.Ok(t, err)
		test.Equals(t, int64(1), affectedRow)
	})
	t.Run("Delete", func(t *testing.T) {
		affectedRow, err := models.Fruit{}.Delete(ctx, 1)
		test.Ok(t, err)
		test.Equals(t, int64(1), affectedRow)
	})
	rollback()
}
