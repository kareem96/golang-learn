package golanggorm

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB  {
	dialect := mysql.Open("root:developer@tcp(localhost:3306)/belajar_golang_gorm?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		PrepareStmt: true, // save query to memory
	})
	if err != nil{
		panic(err)
	}
	sqlDB, _ := db.DB()
	
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)
	return db
}

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestExecuteSQL(t *testing.T) {
	err := db.Exec("insert into sample(id, name) values (?, ?)", "1", "kareem").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values (?, ?)", "2", "abdul").Error
	assert.Nil(t, err)
	
	err = db.Exec("insert into sample(id, name) values (?, ?)", "3", "melayu").Error
	assert.Nil(t, err)
	
}

type Sample struct {
	Id string
	Name string
}
func TestRawSQL(t *testing.T) {
	var sample Sample
	err := db.Raw("select id, name from sample where id = ?", 1).Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, "1", sample.Id)

	var samples []Sample
	err = db.Raw("select id, name from sample").Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(samples))

}

func TestRowSQL(t *testing.T) {
	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample
	for rows.Next(){
		var id string
		var name string
		err := rows.Scan(&id, &name)
		assert.Nil(t, err)

		samples = append(samples, Sample{
			Id: id,
			Name: name,
		})
	}
	assert.Equal(t, 3, len(samples))

}
func TestScanRow(t *testing.T) {
	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample
	for rows.Next(){
		err := db.ScanRows(rows, &samples)
		assert.Nil(t, err)
	}

	assert.Equal(t, 3, len(samples))
}


func TestCreateUser(t *testing.T) {
	user := User{
		ID: "1",
		Password: "rahasia",
		Name: Name{
			FirstName: "Abdul",
			MiddleName: "Karim",
			LastName: "Melayu",
		},
		Infromation: "ini colum ignore",
	}
	response := db.Create(&user)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(1), response.RowsAffected)
}
func TestBatchInsert(t *testing.T) {
	var users []User
	for i := 2; i < 10; i++ {
		users = append(users, User{
			ID: strconv.Itoa(i),
			Password: "rahasia",
			Name: Name{
				FirstName: "User " + strconv.Itoa(i),
			},
			Infromation: "ini colum ignore",
		})
	}
	
	response := db.Create(&users)
	assert.Nil(t, response.Error)
	assert.Equal(t, 8, int(response.RowsAffected))
}
func TestTransactionSuccess(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{ID: "10", Password: "rahasia", Name: Name{FirstName: "User 10"}}).Error
		if err != nil{
			return err
		}
		err = tx.Create(&User{ID: "11", Password: "rahasia", Name: Name{FirstName: "User 11"}}).Error
		if err != nil{
			return err
		}
		err = tx.Create(&User{ID: "12", Password: "rahasia", Name: Name{FirstName: "User 12"}}).Error
		if err != nil{
			return err
		}
		return nil
	})
	assert.Nil(t, err)
}
func TestTransactionError(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		
		err := tx.Create(&User{ID: "11", Password: "rahasia", Name: Name{FirstName: "User 11"}}).Error
		if err != nil{
			return err
		}
		err = tx.Create(&User{ID: "13", Password: "rahasia", Name: Name{FirstName: "User 12"}}).Error
		if err != nil{
			return err
		}
		return nil
	})
	assert.NotNil(t, err)
}


func TestManualTransactionSuccess(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{ID: "13", Password: "rahasia", Name: Name{FirstName: "User 13"}}).Error
	assert.Nil(t, err)

	err = tx.Create(&User{ID: "14", Password: "rahasia", Name: Name{FirstName: "User 14"}}).Error
	assert.Nil(t, err)

	if err == nil {
		tx.Commit()
	}
}
func TestManualTransactionError(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{ID: "15", Password: "rahasia", Name: Name{FirstName: "User 15"}}).Error
	assert.Nil(t, err)

	err = tx.Create(&User{ID: "14", Password: "rahasia", Name: Name{FirstName: "User 14"}}).Error
	assert.Nil(t, err)

	if err == nil {
		tx.Commit()
	}
}
func TestQuerySingleObject(t *testing.T) {
	user := User{}
	err := db.First(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "1", user.ID)

	user = User{}
	err = db.Last(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "9", user.ID)
}
func TestQueryInlineCondition(t *testing.T) {
	user := User{}
	err := db.Take(&user, "id = ?", "5").Error
	assert.Nil(t, err)
	assert.Equal(t, "5", user.ID)
	assert.Equal(t, "User 5", user.Name.FirstName)
}
func TestQueryAllObjects(t *testing.T) {
	users := []User{}
	err := db.Find(&users, "id in ?", []string{"1","2","3","4"}).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))
}
func TestQueryCondition(t *testing.T) {
	users := []User{}
	err := db.Where("first_name like ?", "%User%").Where("password = ?", "rahasia").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 13, len(users))
}
func TestOrOperator(t *testing.T) {
	users := []User{}
	err := db.Where("first_name like ?", "%User%").Or("password = ?", "rahasia").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}
func TestNotOperator(t *testing.T) {
	users := []User{}
	err := db.Not("first_name like ?", "%User%").Where("password = ?", "rahasia").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}
func TestSelectFields(t *testing.T) {
	users := []User{}
	err := db.Select("id","first_name").Find(&users).Error
	assert.Nil(t, err)
	for _, user := range users {
		assert.NotNil(t, user.ID)
		assert.NotEqual(t, "", user.Name.FirstName)
	}

	assert.Equal(t, 14, len(users))
}
func TestStructCondition(t *testing.T) {
	usersCondition := User{
		Name: Name{
			FirstName: "User 5",
			LastName: "", // tidak bisa karena default value
		},
	}
	var users []User
	err := db.Where(usersCondition).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}
func TestMapCondition(t *testing.T) {
	usersCondition := map[string]interface{}{
		"middle_name": "",
	}
	var users []User
	err := db.Where(usersCondition).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 13, len(users))
}
func TestOrderLimitOffset(t *testing.T) {
	var users []User
	err := db.Order("id asc, first_name desc").Limit(5).Offset(5).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 5, len(users))
	assert.Equal(t, "14", users[0].ID)
}

type UserResponse struct {
	ID string
	FirstName string
	LastName string
}

func TestQueryNoModel(t *testing.T)  {
	var users []UserResponse
	err := db.Model(&User{}).Select("id", "first_name", "last_name").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 14, len(users))
	fmt.Println(users)
}
func TestUpdate(t *testing.T)  {
	user := User{}
	err := db.Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)
	
	user.Name.FirstName = "KDev"
	user.Name.MiddleName = "KDev"
	user.Name.LastName = "Developer"
	user.Password = "Developer123"
	
	err = db.Save(&user).Error
	assert.Nil(t, err)
	fmt.Println(user)
}
func TestSelectedColumnsUpdate(t *testing.T)  {
	err := db.Model(&User{}).Where("id = ?", "1").Updates(map[string]interface{}{
		"middle_name":"",
		"last_name":"K",
	}).Error
	assert.Nil(t, err)

	err = db.Model(&User{}).Where("id = ?", "1").Update("password", "diubahlagi").Error
	assert.Nil(t, err)

	err = db.Where("id = ?", "1").Updates(User{
		Name: Name{
			FirstName: "Abdul",
			LastName: "Kareem",
		},
	}).Error
	assert.Nil(t, err)
}

func TestAutoIncrement(t *testing.T)  {
	for i := 0; i < 10; i++ {
		userLog := UserLog{
			UserID: "1",
			Action: "Test Action",
		}
		err := db.Create(&userLog).Error
		assert.Nil(t, err)
		assert.NotEqual(t, 0, userLog.ID)
		fmt.Println(userLog.ID)
	}
}
func TestSaveOrUpdate(t *testing.T)  {
	userLog := UserLog{
		UserID: "1",
		Action: "Text Action",
	}

	err := db.Save(&userLog).Error //insert
	assert.Nil(t, err)

	userLog.UserID = "2"
	err = db.Save(&userLog).Error //update
	assert.Nil(t, err)

}

func TestSaveOrUpdateNonAutoIncrement(t *testing.T)  {
	user := User{
		ID: "101",
		Name: Name{
			FirstName: "User 101",
		},
	}

	err := db.Save(&user).Error //create
	assert.Nil(t, err)

	user.ID = "User 101 update"
	err = db.Save(&user).Error //update
	assert.Nil(t, err)

}
func TestConflict(t *testing.T)  {
	user := User{
		ID: "88",
		Name: Name{
			FirstName: "User 88",
		},
	}

	err := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&user).Error
	assert.Nil(t, err)

}
func TestDelete(t *testing.T)  {
	user := User{}

	err := db.Take(&user, "id = ?", "100").Delete(&user).Error
	assert.Nil(t, err)

	err = db.Delete(&User{}, "id = ?", "101").Error
	assert.Nil(t, err)
	
	err = db.Where("id = ?", "1").Delete(&User{}).Error
	assert.Nil(t, err)

}
func TestSoftDelete(t *testing.T)  {
	todo := Todo{
		UserID: "1",
		Title: "Todo 1",
		Description: "Isi Todo 1",
	}

	err := db.Create(&todo).Error
	assert.Nil(t, err)

	err = db.Delete(&todo).Error
	assert.Nil(t, err)
	assert.NotNil(t, todo.DeletedAt)

	var todos []Todo
	err = db.Find(&todos).Error
	assert.Nil(t, err)
	assert.Equal(t, 0, len(todos))
}
func TestUnscopped(t *testing.T)  {
	todo := Todo{}

	err := db.Unscoped().First(&todo, "id = ?", "1").Error
	assert.Nil(t, err)

	err = db.Unscoped().Delete(&todo).Error
	assert.Nil(t, err)

	var todos []Todo
	err = db.Unscoped().Find(&todos).Error
	assert.Nil(t, err)
}
func TestLock(t *testing.T)  {
	err := db.Transaction(func(tx *gorm.DB) error {
		var user User
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&user, "id = ?", "2").Error
		if err != nil {
			return err
		}

		user.Name.FirstName = "Joko"
		user.Name.LastName = "Morro"
		err = tx.Save(&user).Error
		if err != nil{
			return err
		}

		return nil
	})
	assert.Nil(t, err)
}
func TestCreateWallet(t *testing.T)  {
	wallet := Wallet{
		ID: "2",
		UserId: "2",
		Balance: 10000,
	}
	err := db.Create(&wallet).Error
	assert.Nil(t, err)
}
func TestRetrieveRelation(t *testing.T)  {
	var user User
	err := db.Model(&User{}).Preload("Wallet").Take(&user, "id = ?", "2").Error
	assert.Nil(t, err)
	
	assert.Equal(t, "2", user.ID)
	assert.Equal(t, "2", user.Wallet.ID)
}
func TestRetrieveRelationJoin(t *testing.T)  {
	var user User
	err := db.Model(&User{}).Joins("Wallet").Find(&user, "users.id = ?", "2").Error
	assert.Nil(t, err)
	
	assert.Equal(t, "2", user.ID)
	assert.Equal(t, "2", user.Wallet.ID)
}
func TestAutoCreateUpdate(t *testing.T)  {
	user := User{
		ID: "1",
		Password: "rahasia",
		Name: Name{
			FirstName: "User 1",
		},
		Wallet: Wallet{
			ID: "1",
			UserId: "1",
			Balance: 100000,
		},
	}
	err := db.Create(&user).Error
	assert.Nil(t, err)
}
func TestSkipAutoCreateUpdate(t *testing.T)  {
	user := User{
		ID: "20",
		Password: "rahasia",
		Name: Name{
			FirstName: "User 20",
		},
		Wallet: Wallet{
			ID: "20",
			UserId: "20",
			Balance: 100000,
		},
	}
	err := db.Omit(clause.Associations).Create(&user).Error
	assert.Nil(t, err)
}
func TestUserAndAdresses(t *testing.T)  {
	user := User{
		ID: "71",
		Password: "rahasia",
		Name: Name{
			FirstName: "User 71",
		},
		Wallet: Wallet{
			ID: "71",
			UserId: "71",
			Balance: 100000,
		},
		Addresses: []Address{
			{
				UserId: "71",
				Address: "Jalan A",
			},
			{
				UserId: "71",
				Address: "Jalan B",
			},
		},
	}
	err := db.Create(&user).Error
	assert.Nil(t, err)
}

func TestPreloadJoinOneToMany(t *testing.T) {
	var userPreload []User
	err := db.Model(&User{}).Preload("Addresses").Joins("Wallet").Find(&userPreload).Error
	assert.Nil(t, err)
}
func TestTakePreloadJoinOneToMany(t *testing.T) {
	var userPreload []User
	err := db.Model(&User{}).Preload("Addresses").Joins("Wallet").Take(&userPreload, "users.id = ?","71").Error
	assert.Nil(t, err)
}
func TestBelongsTo(t *testing.T) {
	fmt.Println("Preload")
	var addresses []Address
	err := db.Model(&Address{}).Preload("User").Find(&addresses).Error
	assert.Nil(t, err)
	
	fmt.Println("Joins")
	addresses = []Address{}
	err = db.Model(&Address{}).Joins("User").Find(&addresses).Error
	assert.Nil(t, err)
}

func TestBelongsToOneToOne(t *testing.T) {
	fmt.Println("Preload")
	var wallets []Wallet
	err := db.Model(&Wallet{}).Preload("User").Find(&wallets).Error
	assert.Nil(t, err)
	
	fmt.Println("Joins")
	wallets = []Wallet{}
	err = db.Model(&Wallet{}).Joins("User").Find(&wallets).Error
	assert.Nil(t, err)
}
func TestCreateManyToMany(t *testing.T) {
	product :=  Product{
		ID: "P003",
		Name: "Contoh Product",
		Price: 100000,
	}
	err := db.Create(&product).Error
	assert.Nil(t, err)

	err = db.Table("user_like_product").Create(map[string]interface{}{
		"user_id": "4",
		"product_id": "P003",
	}).Error
	assert.Nil(t, err)
	
	err = db.Table("user_like_product").Create(map[string]interface{}{
		"user_id": "5",
		"product_id": "P003",
	}).Error
	assert.Nil(t, err)
}

func TestPreloadManyToMany(t *testing.T) {
	var product Product
	err := db.Preload("LikeByUsers").Take(&product, "id = ?", "P002").Error
	assert.Nil(t, err)
	assert.Equal(t, 2, len(product.LikeByUsers))
}

func TestPreloadManyToManyUser(t *testing.T) {
	var user User
	err := db.Preload("LikeProducts").Take(&user, "id = ?", "4").Error
	assert.Nil(t, err)
	assert.Equal(t, 2,len(user.LikeProducts))
}

func TestAssociationFind(t *testing.T) {
	var product Product
	err := db.First(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	var users []User
	err = db.Model(&product).Where("first_name LIKE ?", "User%").Association("LikeByUsers").Find(&users)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}
func TestAssociationAdd(t *testing.T) {
	var user User
	err := db.First(&user, "id = ?", "4").Error
	assert.Nil(t, err)

	var product Product
	err = db.First(&product, "id = ?", "P002").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikeByUsers").Append(&user)
	assert.Nil(t, err)
}
func TestAssociationReplace(t *testing.T) {

	err := db.Transaction(func(tx *gorm.DB) error {
		var user User
		err := db.Take(&user, "id = ?", "5").Error
		assert.Nil(t, err)
	
		wallet := Wallet{
			ID: "04",
			UserId: user.ID,
			Balance: 10000,
		}
	
		err = db.Model(&user).Association("Wallet").Replace(&wallet)
		return err
	}).Error()
	assert.Nil(t, err)
}
func TestAssociationDelete(t *testing.T) {

	var user User
	err := db.First(&user, "id = ?", "5").Error
	assert.Nil(t, err)

	var product Product
	err = db.First(&product, "id = ?", "P002").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikeByUsers").Append(&user)
	assert.Nil(t, err)
}
func TestAssociationClear(t *testing.T) {
	var product Product
	err := db.First(&product, "id = ?", "P002").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikeByUsers").Clear()
	assert.Nil(t, err)
}
func TestPreloadingWithCondition(t *testing.T) {
	var user User
	err := db.Preload("Wallet", "balance > ?", 100).First(&user, "id = ?", "1").Error
	assert.Nil(t, err)
	fmt.Println(user)
}
func TestNestedPreloadingWithCondition(t *testing.T) {
	var wallet Wallet
	err := db.Preload("User.Addresses",).Find(&wallet, "id = ?", "71").Error
	assert.Nil(t, err)
	fmt.Println(wallet)
	fmt.Println(wallet.User)
	fmt.Println(wallet.User.Addresses)
}
func TestPreloadAll(t *testing.T) {
	var user User
	err := db.Preload(clause.Associations).Take(&user, "id = ?", "71").Error
	assert.Nil(t, err)
	fmt.Println(user)
}
func TestJoinQuery(t *testing.T) {
	var users []User
	err := db.Joins("join wallets on wallets.user_id = users.id").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 6, len(users))

	users = []User{}
	err = db.Joins("Wallet").Find(&users).Error //left join
	assert.Nil(t, err)
	assert.Equal(t, 20, len(users))


	fmt.Println(users)
}
func TestJoinQueryCondition(t *testing.T) {
	var users []User
	err := db.Joins("join wallets on wallets.user_id = users.id AND wallets.balance > ?", 1000).Find(&users).Error
	assert.Nil(t, err)
	// assert.Equal(t, 6, len(users))

	users = []User{}
	err = db.Joins("Wallet").Where("Wallet.balance > ?", 1000).Find(&users).Error // alias menggunakan nama field
	assert.Nil(t, err)
	// assert.Equal(t, 20, len(users))


	fmt.Println(users)
}


func TestAgregationCount(t *testing.T) {
	var count int64
	err := db.Model(&User{}).Joins("Wallet").Where("Wallet.balance > ?", 500).Count(&count).Error
	assert.Nil(t, err)
	assert.Equal(t, int64(6), count)

	fmt.Println(count)
}

type AggregationResult struct {
	TotalBalance int64
	MinBalance int64
	MaxBalance int64
	AvgBalance float64
}

func TestAggregation(t *testing.T) {
	var result AggregationResult
	err := db.Model(&Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance", "max(balance) as max_balance, avg(balance) as avg_balance").Take(&result).Error
	assert.Nil(t, err)

	assert.Equal(t, int64(240000), result.TotalBalance)
	assert.Equal(t, int64(10000), result.MinBalance)
	assert.Equal(t, int64(100000), result.MaxBalance)
	assert.Equal(t, float64(40000), result.AvgBalance)
}
func TestAggregationGroupByAndHaving(t *testing.T) {
	var result []AggregationResult
	err := db.Model(&Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance", "max(balance) as max_balance, avg(balance) as avg_balance").
	Joins("User").Group("User.id").Having("sum(balance) > ?", 10000).Find(&result).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(result))

	fmt.Println(result)
}

func TestContext(t *testing.T) {
	ctx := context.Background()
	var users []User
	err := db.WithContext(ctx).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 19, len(users))

	fmt.Println(users)
}

func BrokeWalletBalance(db *gorm.DB) *gorm.DB  {
	return db.Where("balance = ?", 0)
}
func SultanWalletBalance(db *gorm.DB) *gorm.DB  {
	return db.Where("balance > ?", 10000)
}

func TestScopes(t *testing.T) {
	var wallets []Wallet
	err := db.Scopes(BrokeWalletBalance).Find(&wallets).Error
	assert.Nil(t, err)

	wallets = []Wallet{}
	err = db.Scopes(SultanWalletBalance).Find(&wallets).Error
	assert.Nil(t, err)

	fmt.Println(wallets)
}

func TestMigrator(t *testing.T) {
	err := db.Migrator().AutoMigrate(&GuestBook{})
	assert.Nil(t, err)
}

func TestUserHook(t *testing.T) {
	user := User{
		Password: "rahasia",
		Name: Name{
			FirstName: "User 103",
		},
	}
	err := db.Create(&user).Error
	assert.Nil(t, err)
	assert.NotEqual(t, "", user.ID)

	fmt.Println(user.ID)
}