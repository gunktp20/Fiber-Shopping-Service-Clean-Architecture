package repository

import (
	userDto "shopping-service-be/modules/user/userDto"
	userEntity "shopping-service-be/modules/user/userEntity"

	"github.com/jmoiron/sqlx"
)

type userPostgresRepository struct {
	db *sqlx.DB
}

func NewUserPostgresRepository(db *sqlx.DB) UserRepositoryService {
	return &userPostgresRepository{db}
}

func (r *userPostgresRepository) IsUniqueUser(email string) bool {
	return true
}

func (r *userPostgresRepository) CreateOneUser(createUserReq *userDto.CreateUserReq) (*userDto.CreateUserRes, error) {
	// ? sqlx query row x is not support named parameters (:name , :description , :price ) that why we use $ instead
	var query string
	var args []interface{}

	query = `
			INSERT INTO users (email,password) 
			VALUES ($1, $2) 
			RETURNING id , email
		`
	args = []interface{}{createUserReq.Email, createUserReq.Password}

	user := new(userEntity.User)

	err := r.db.QueryRowx(query, args...).StructScan(user)
	if err != nil {
		return nil, err
	}

	return &userDto.CreateUserRes{
		Email: user.Email,
	}, nil

}

func (r *userPostgresRepository) GetOneUserByEmail(email string) (*userEntity.User, error) {

	var user userEntity.User

	query := `SELECT id, email, password FROM users WHERE email = $1`

	err := r.db.Get(&user, query, email)
	if err != nil {
		return &userEntity.User{}, err
	}

	return &userEntity.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
