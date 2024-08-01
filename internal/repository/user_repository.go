package repository

// import (
// 	"Go-TiketPemesanan/internal/domain"
// 	"errors"
// 	"sync"
// )

// // interface yang mengabungkan beberapa interface kecil
// type UserRepositoryInterface interface {
// 	UserSaver
// 	UserFindById
// 	UserUpdater
// 	UserDeleter
// 	UserLister
// 	UpdateBalance
// }

// // interface yang digunakan untuk menyimpan penguna, dengan menerima pointer ke domain.user dan mengembalikan domain.user dan error
// type UserSaver interface {
// 	UserSaver(user *domain.User) (domain.User, error)
// }

// // interface yang digunakan untuk menerima id user, dengan mengemablikan domain.user dan error
// type UserFindById interface {
// 	UserFindById(id int) (domain.User, error)
// }

// // interface yang digunakan untuk memperbaruhi penguna, dengan menerima pointer ke domain.user dan mengembalikan domain.user dan error
// type UserUpdater interface {
// 	UpdateUser(user *domain.User) (domain.User, error)
// }

// // interface yang di gunakan untuk menghapus user dengan menerima id user dan mengembalikan domain.user dan error
// type UserDeleter interface {
// 	DeleteUser(id int) (domain.User, error)
// }

// // interface yang di gunakan untuk melihat semua user, dan mengembalikan daftar semua penguna dalam bentuk slice dan error
// type UserLister interface {
// 	GetAllUser() ([]domain.User, error)
// }

// // interface yang digunakan untuk mengupdate saldo, dengan menerima id user dan saldo baru, dan mengembalikan domain.user dan error
// type UpdateBalance interface {
// 	UpdateBalance(id int, newBalance float64) (domain.User, error)
// }

// // struktur data berfungsi untuk menyimpan data penguna dalam bentuk map dengan key int dan value sebagai user
// type UserRepository struct {
// 	users map[int]domain.User
// 	mu    sync.Mutex
// }

// /*
// fungsi ini mengembalikan sebuah object yang mengimplementasi userrepositroyinterface dan
// membuat object baru userrepository dan mengembalikan alamat memori dari object tersebut
// */
// func NewUserRepository() UserRepositoryInterface {
// 	return &UserRepository{
// 		users: map[int]domain.User{},
// 		mu:    sync.Mutex{},
// 	}
// }

// // implementasi metode yang di diguanakn untuk mendapat kan id user dan mengembalikan kesalahan jika id tidak di temukan
// // *userrepository  tipe penerima ke userrepository pointer ini memungkin kita mengubah nilai dalam userrepository
// func (repo *UserRepository) UserFindById(id int) (domain.User, error) {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	usr, exist := repo.users[id]
// 	if !exist {
// 		return domain.User{}, errors.New("user not found")
// 	}
// 	return usr, nil
// }

// // implementasi metode yang digunakan untuk menghapus user , dengan menerima id dan mengembalikan engembalikan domain.user dan error
// func (repo *UserRepository) DeleteUser(id int) (domain.User, error) {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	deletedUser, exist := repo.users[id]
// 	if !exist {
// 		return domain.User{}, errors.New("user not found")
// 	}

// 	delete(repo.users, id)
// 	return deletedUser, nil
// }

// // implementasi metode yang digunakan untuk mendapatkan semua user.
// func (repo *UserRepository) GetAllUser() ([]domain.User, error) {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	users := []domain.User{}
// 	for _, user := range repo.users {
// 		users = append(users, user)
// 	}
// 	return users, nil
// }

// // implementasi sebuah metode yang menyimpan penguna baru. dan mengatur id secara otomatis dan mengembalikan kesalahan
// func (repo *UserRepository) UserSaver(user *domain.User) (domain.User, error) {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	if _, exist := repo.users[user.ID]; exist {
// 		return *user, errors.New("user already exist")
// 	}

// 	if repo.users == nil || len(repo.users) == 0 {
// 		user.ID = 1
// 	} else {
// 		user.ID = repo.users[len(repo.users)].ID + 1
// 	}

// 	repo.users[user.ID] = *user
// 	return *user, nil
// }

// // implementasi sebuah metode yang memperbaruhi penguna, dan mengembalikan domain.user dan error
// func (repo *UserRepository) UpdateUser(user *domain.User) (domain.User, error) {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	if _, exist := repo.users[user.ID]; exist {
// 		repo.users[user.ID] = *user
// 		return *user, nil
// 	}
// 	return *user, errors.New("user not found")
// }

// // UpdateBalance implements UserRepositoryInterface.
// func (repo *UserRepository) UpdateBalance(id int, newBalance float64)  (domain.User,error) {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	user, exist := repo.users[id]
// 	if !exist {
// 		return domain.User{}, errors.New("user not found")
// 	}

// 	user.Balance = newBalance
// 	repo.users[id] = user
// 	return user, nil

// }