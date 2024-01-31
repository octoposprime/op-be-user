package infrastructure

import (
	"context"

	"github.com/google/uuid"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	tgorm "github.com/octoposprime/op-be-shared/tool/gorm"
	me "github.com/octoposprime/op-be-user/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-user/internal/domain/model/object"
	map_repo "github.com/octoposprime/op-be-user/pkg/infrastructure/mapper/repository"
)

type DbAdapter struct {
	*tgorm.GormClient
	Log func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error)
}

func NewDbAdapter(dbClient *tgorm.GormClient) DbAdapter {
	adapter := DbAdapter{
		dbClient,
		Log,
	}

	err := dbClient.DbClient.AutoMigrate(&map_repo.User{})
	if err != nil {
		panic(err)
	}
	err = dbClient.DbClient.AutoMigrate(&map_repo.UserPassword{})
	if err != nil {
		panic(err)
	}

	return adapter
}

// SetLogger sets logging call-back function
func (a *DbAdapter) SetLogger(LoggerFunc func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error)) {
	a.Log = LoggerFunc
}

// GetUsersByFilter returns the users that match the given filter.
func (a DbAdapter) GetUsersByFilter(ctx context.Context, userFilter me.UserFilter) (me.Users, error) {
	var usersDbMapper map_repo.Users
	var filter map_repo.User
	qry := a.DbClient
	if userFilter.Id.String() != "" && userFilter.Id != (uuid.UUID{}) {
		filter.ID = userFilter.Id
	}
	if userFilter.UserName != "" {
		filter.UserName = userFilter.UserName
	}
	if userFilter.Email != "" {
		filter.Email = userFilter.Email
	}
	if userFilter.UserType != 0 {
		filter.UserType = int(userFilter.UserType)
	}
	if userFilter.UserStatus != 0 {
		filter.UserStatus = int(userFilter.UserStatus)
	}
	if len(userFilter.Tags) > 0 {
		filter.Tags = userFilter.Tags
	}
	if userFilter.FirstName != "" {
		filter.FirstName = userFilter.FirstName
	}
	if userFilter.LastName != "" {
		filter.LastName = userFilter.LastName
	}
	if !userFilter.CreatedAtFrom.IsZero() && !userFilter.CreatedAtTo.IsZero() {
		qry = qry.Where("created_at between ? and ?", userFilter.CreatedAtFrom, userFilter.CreatedAtTo)
	}
	if !userFilter.UpdatedAtFrom.IsZero() && !userFilter.UpdatedAtTo.IsZero() {
		qry = qry.Where("updated_at between ? and ?", userFilter.UpdatedAtFrom, userFilter.UpdatedAtTo)
	}
	if userFilter.SearchText != "" {
		qry = qry.Where(
			qry.Where("UPPER(user_name) LIKE UPPER(?)", "%"+userFilter.SearchText+"%").
				Or("UPPER(email) LIKE UPPER(?)", "%"+userFilter.SearchText+"%").
				Or("UPPER(array_to_string(tags, ',')) LIKE UPPER(?)", "%"+userFilter.SearchText+"%"),
		)
	}
	qry = qry.Where(filter)
	var totalRows int64
	result := qry.Model(&map_repo.User{}).Where(filter).Count(&totalRows)
	if result.Error != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "GetUsersByFilter", userId, result.Error.Error()))
		totalRows = 0
	}
	if userFilter.Limit != 0 {
		qry = qry.Limit(userFilter.Limit)
	}
	if userFilter.Offset != 0 {
		qry = qry.Offset(userFilter.Offset)
	}
	if userFilter.SortType != "" && userFilter.SortField != 0 {
		sortStr := map_repo.UserSortMap[userFilter.SortField]
		if userFilter.SortType == "desc" || userFilter.SortType == "DESC" {
			sortStr += " desc"
		} else {
			sortStr += " asc"
		}
		qry = qry.Order(sortStr)
	}
	result = qry.Find(&usersDbMapper)
	if result.Error != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "GetUsersByFilter", userId, result.Error.Error()))
		return me.Users{}, result.Error
	}
	return me.Users{
		Users:     usersDbMapper.ToEntities(),
		TotalRows: totalRows,
	}, nil
}

// SaveUser insert a new user or update the existing one in the database.
func (a DbAdapter) SaveUser(ctx context.Context, user me.User) (me.User, error) {
	userDbMapper := map_repo.NewUserFromEntity(&user)
	qry := a.DbClient
	if user.Id.String() != "" && user.Id != (uuid.UUID{}) {
		qry = qry.Omit("created_at")
	}
	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	if userDbMapper.ID != (uuid.UUID{}) {
		userDbMapper.UpdatedBy, _ = uuid.Parse(userId)
	} else {
		userDbMapper.CreatedBy, _ = uuid.Parse(userId)
	}
	result := qry.Save(&userDbMapper)
	if result.Error != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "SaveUser", userId, result.Error.Error()))
		return me.User{}, result.Error
	}
	return *userDbMapper.ToEntity(), nil
}

// DeleteUser soft-deletes the given user in the database.
func (a DbAdapter) DeleteUser(ctx context.Context, user me.User) (me.User, error) {
	userDbMapper := map_repo.NewUserFromEntity(&user)
	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	userDbMapper.DeletedBy, _ = uuid.Parse(userId)
	result := a.DbClient.Delete(&userDbMapper)
	if result.Error != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "DeleteUser", userId, result.Error.Error()))
		return me.User{}, result.Error
	}
	return *userDbMapper.ToEntity(), nil
}

// GetUserPasswordByUserId returns active password of the given user.
func (a DbAdapter) GetUserPasswordByUserId(ctx context.Context, userId uuid.UUID) (me.UserPassword, error) {
	var userPasswordsDbMapper *map_repo.UserPasswords
	var filter map_repo.UserPassword
	filter.UserID = userId
	filter.PasswordStatus = int(mo.PasswordStatusACTIVE)
	result := a.DbClient.Where(filter).Find(&userPasswordsDbMapper)
	if result.Error != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "GetUsersByFilter", userId, result.Error.Error()))
		return me.UserPassword{}, result.Error
	}
	if len(*userPasswordsDbMapper) > 0 {
		resultDatas := userPasswordsDbMapper.ToEntities()
		return resultDatas[len(resultDatas)-1], nil
	}
	err := mo.ErrorUserPasswordNotFound
	return me.UserPassword{}, err
}

// ChangePassword changes the given user password in the database.
func (a DbAdapter) ChangePassword(ctx context.Context, userPassword me.UserPassword) (me.UserPassword, error) {
	userPasswordDbMapper := map_repo.NewUserPasswordFromEntity(&userPassword)
	qry := a.DbClient
	if userPassword.Id.String() != "" && userPassword.Id != (uuid.UUID{}) {
		qry = qry.Omit("created_at")
	}
	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	if userPasswordDbMapper.ID != (uuid.UUID{}) {
		userPasswordDbMapper.UpdatedBy, _ = uuid.Parse(userId)
	} else {
		userPasswordDbMapper.CreatedBy, _ = uuid.Parse(userId)
	}
	result := qry.Save(&userPasswordDbMapper)
	if result.Error != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "ChangePassword", userId, result.Error.Error()))
		return me.UserPassword{}, result.Error
	}
	return *userPasswordDbMapper.ToEntity(), nil
}
