package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.61

import (
	"context"
	"craftnet/config"
	"craftnet/graph/model"
	"craftnet/internal/db"
	"craftnet/internal/util"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.User, error) {
	db := db.GetDB()

	var existingUser model.User
	if err := db.QueryRow(ctx, "SELECT id FROM accounts WHERE username = $1", input.Username).Scan(&existingUser.ID); err == nil {
		util.GetLogger().LogErrorWithMsgAndError("username "+input.Username+" already exists", err, false)
		return nil, errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		util.GetLogger().LogErrorWithMsgAndError("failed to hash password", err, false)
		return nil, errors.New("failed to hash password")
	}

	var newUser model.User
	query := `
        INSERT INTO accounts (username, password_hash, created_at) 
        VALUES ($1, $2, $3)
		RETURNING id, username
    `
	if err := db.QueryRow(ctx, query, input.Username, string(hashedPassword), util.TimeNowJSTZone()).Scan(&newUser.ID, &newUser.Username); err != nil {
		util.GetLogger().LogErrorWithMsgAndError("failed to create user", err, false)
		return nil, errors.New("failed to create user")
	}

	util.GetLogger().LogInfo(fmt.Sprintf("Register new account succesfully. Id: %s, username: %s", newUser.ID, newUser.Username))
	return &newUser, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthPayload, error) {
	db := db.GetDB()

	var account model.Account
	if err := db.QueryRow(ctx, "SELECT id, username, password_hash FROM accounts WHERE username = $1", input.Username).Scan(&account.Username, &account.PasswordHash); err != nil {
		util.GetLogger().LogErrorWithMsgAndError("invalid username or password", err, false)
		return nil, errors.New("invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.PasswordHash), []byte(input.Password)); err != nil {
		util.GetLogger().LogErrorWithMsgAndError("invalid username or password", err, false)
		return nil, errors.New("invalid username or password")
	}

	token, err := util.GenerateJWT(account.Username, config.GetJwtSecret())

	if err != nil {
		util.GetLogger().LogErrorWithMsgAndError("failed to generate token", err, false)
		return nil, errors.New("failed to generate token")
	}

	return &model.AuthPayload{
		AccessToken: token,
		User:        account.Username,
	}, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	db := db.GetDB()

	rows, err := db.Query(ctx, "SELECT id, username FROM accounts")

	if err != nil {
		util.GetLogger().LogErrorWithMsgAndError("failed to get users", err, false)
		return nil, errors.New("failed to get users")
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			util.GetLogger().LogErrorWithMsgAndError("failed to parse user", err, false)
			return nil, errors.New("failed to parse user")
		}
		users = append(users, &user)
	}

	return users, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
