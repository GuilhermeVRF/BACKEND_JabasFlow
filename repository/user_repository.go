package repository

import (
	"database/sql"
	"jabas-flow/models"
	_"github.com/go-sql-driver/mysql"
)

type UserRepository struct {
	databaseConnection *sql.DB
}

func NewUserRepository(databaseConnection *sql.DB) UserRepository {
	return UserRepository{
		databaseConnection: databaseConnection,
	}
}

func (userRepository *UserRepository) GetUsers() ([]models.User, error) {
	usersQuery := "SELECT * FROM User"
	userResults, err := userRepository.databaseConnection.Query(usersQuery)

	if err != nil{
		return []models.User{}, err
	}

	var users []models.User
	var user models.User

	for userResults.Next(){
		err := userResults.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Profile)
		if err != nil {
			return []models.User{}, err
		}

		users = append(users, user)
	}

	userResults.Close()
	return users, nil
}

func (UserRepository *UserRepository) GetUser(id int) (models.User, error) {
	userQuery := "SELECT * FROM User WHERE id = ?"
	userResult := UserRepository.databaseConnection.QueryRow(userQuery, id)

	var user models.User
	err := userResult.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Profile)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (userRepository* UserRepository) InsertUser(user models.User) (int64, error) {
	insertQuery := "INSERT INTO User (name, email, password, profile) VALUES (?, ?, ?, ?)"
	
	var userId int64
	userInsertResult, err := userRepository.databaseConnection.Exec(insertQuery, user.Name, user.Email, user.Password, user.Profile)
    if err != nil {
        return 0, err
    }
	
	userId, err = userInsertResult.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (userRepository *UserRepository) UpdateUser(user models.User) (models.User, error) {
	updateQuery := "UPDATE User SET name = ?, email = ?, password = ?, profile = ? WHERE id = ?"
	_, err := userRepository.databaseConnection.Exec(updateQuery, user.Name, user.Email, user.Password, user.Profile, user.ID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}